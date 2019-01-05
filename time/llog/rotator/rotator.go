package rotator

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// nl is a byte slice containing a newline byte.  It is used to avoid creating
// additional allocations when writing newlines to the log file.
var nl = []byte{'\n'}

// 这就是我们的主类
//threshold:阈值
type Rotator struct {
	size      int64
	threshold int64
	maxRolls  int
	filename  string
	out       *os.File
	tee       bool
	wg        sync.WaitGroup
}

func New(filenam string, threshholdKB int64, tee bool, maxRolls int) (*Rotator, error) {
	f, err := os.OpenFile(filenam, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return &Rotator{
		size:      stat.Size(),
		threshold: 1000 * threshholdKB,
		maxRolls:  maxRolls,
		filename:  filenam,
		out:       f,
		tee:       tee,
	}, nil
}

// Rotator关闭.主要是关闭输出文件
func (r *Rotator) Close() error {
	err := r.out.Close()
	r.wg.Wait()
	return err
}

// run 从一个reader中读取行并且按照需要totating logs
// 不能和写同时进行
func (r *Rotator) Run(reader io.Reader) error {
	in := bufio.NewReader(reader)

	if r.size >= r.threshold {
		if err := r.rotate(); err != nil {
			return err
		}
		r.size = 0
	}
	for {
		// isPrefix就是判断缓冲区一行有没有读完
		line, isPrefix, err := in.ReadLine()
		if err != nil {
			return err
		}
		n, err := r.out.Write(line)
		r.size += int64(n)
		if r.tee {
			os.Stdout.Write(line)
		}
		if isPrefix {
			continue
		}
		m, _ := r.out.Write(n1)
		if r.tee {
			os.Stdout.Write(n1)
		}
		r.size += int64(m)
		if r.size >= r.threshold {
			err := r.rotate()
			if err != nil {
				return err
			}
			r.size = 0
		}
	}
}

func (r *Rotator) rotate() error {
	dir := filepath.Dir(r.filename)
	glob := filepath.Join(dir, filepath.Base(r.filename)+".*")
	// Glob returns the names of all files matching pattern or nil
	existing, err := filepath.Glob(glob)
	if err != nil {
		return err
	}
	maxNum := 0
	for _, name := range existing {
		parts := strings.Split(name, ".")
		if len(parts) < 2 {
			continue
		}
		numIdx := len(parts) - 1
		if parts[numIdx] == "gz" {
			numIdx--
		}
		num, err := strconv.Atoi(parts[numIdx])
		if err != nil {
			continue
		}
		if num > maxNum {
			maxNum = num
		}
	}

	err := r.out.Close()
	if err != nil {
		return err
	}
	rtname := fmt.Sprintf("%s.%d", r.filename, maxNum+1)
	err = os.Rename(r.filename, rotname)
	if err != nil {
		return err
	}
	if r.maxRolls > 0 {
		for n := maxNum + 1 - r.maxRolls; ; n-- {
			err := os.Remove(fmt.Sprintf("%s.%d.gz", r.filename, n))
			if err != nil {
				break
			}
		}
	}
	r.out, err = os.OpenFile(r.filenam, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	r.wg.Add(1)
	go func() {
		err := compress(rotname)
		if err == nil {
			os.Remove(rotname)
		}
		r.wg.Done()
	}()
	return nil
}

// 传一个文件名，我就给你压缩
func compress(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	arc, err := os.OpenFile(name+".gz", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	z := gzip.NewWriter(arc)
	if _, err = io.Copy(z, f); err != nil {
		return err
	}
	if err = z.Close(); err != nil {
		return err
	}
	return arc.Close()

}

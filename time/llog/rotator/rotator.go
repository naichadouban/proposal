package rotator

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

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
		if parts[numIdx] == "gz"{
			numIdx--
		}
		num,err := strconv.Atoi(parts[numIdx])
		if err != nil {
			continue
		}
		if num > maxNum {
			maxNum = num
		}
	}
}

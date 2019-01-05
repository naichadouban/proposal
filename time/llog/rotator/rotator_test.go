package rotator

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFilePath(t *testing.T){
	fmt.Println(filepath.Dir("rotator.go"))
}

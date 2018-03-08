package utils_test

import (
	"gitlab.xinghuolive.com/Backend-Go/kangaroo/utils"
	"testing"
	"fmt"
)

func TestExecute(t *testing.T) {
	var execStr = "du -b /home/chenqicong/.bashrc | awk '{print $1}'"
	size, err := utils.ReadFileSize(execStr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(size)
}

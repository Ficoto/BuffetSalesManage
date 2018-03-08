package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ReadLine - 读取每行
func ReadLine(filePath string, handler func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func ReadFileSize(filePath string) (size int64, err error) {
	var execStr = fmt.Sprintf("du -b %s | awk '{print $1}'", filePath)
	var lines []string
	err, lines = Execute(execStr)
	if err != nil {
		return
	}
	if len(lines) == 0 {
		return
	}
	size, err = strconv.ParseInt(lines[0], 10, 64)
	return
}

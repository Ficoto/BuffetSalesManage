package utils

import (
	"bytes"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/Sirupsen/logrus"
)

type CustomTextFormatter struct {
	Formatter logrus.TextFormatter
}

func (f *CustomTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	oriLog, _ := f.Formatter.Format(entry)
	_, file, line, ok := runtime.Caller(5)

	if ok {
		oriLog = oriLog[:len(oriLog)-2]
		_, filename := filepath.Split(file)
		byteLog := [][]byte{oriLog, []byte(" file=\""), []byte(filename), []byte(":"),
							[]byte(strconv.Itoa(line)), []byte("\"\n")}
		oriLog = bytes.Join(byteLog, nil)
	}

	return oriLog, nil
}

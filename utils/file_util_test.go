package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestM3U8File2Object(t *testing.T) {
	filePath := "E:/598c5858d0b04b6ac36df0dd.m3u8"
	type arg struct {
		ObjectKey string  `json:"object_key"`
		Duration  float64 `json:"duration"`
	}

	var (
		args  []arg
		a     arg
		count int
	)
	ReadLine(
		filePath,
		func(i string) {
			if strings.HasPrefix(i, "#EXTINF:") {
				f, err := strconv.ParseFloat(i[8:16], 64)
				if err != nil {
					log.Printf("ERROR(%s): Parse to float64 error!", err.Error())
				}
				a.Duration = f
			}
			if !strings.HasPrefix(i, "#") && i != "\n" {
				a.ObjectKey = i
			}
			if a.Duration != 0 && len(a.ObjectKey) != 0 {
				fmt.Println(count)
				args = append(args, a)
				a = arg{}
				count++
			}
		},
	)
	fmt.Println(args)
}

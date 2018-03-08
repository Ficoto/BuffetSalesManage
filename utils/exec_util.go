package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func Execute(execStr string) (err error, lines []string) {
	log.Printf("INFO(%s): Echo command!", execStr)
	cmd := exec.Command("/bin/sh", "-c", execStr)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Printf("ERROR(%s): Stdout pipe error!", err.Error())
		return
	}

	err = cmd.Start()
	if err != nil {
		log.Printf("ERROR(%s): Command start error!", err.Error())
		return
	}

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
		lines = append(lines, strings.Replace(line, "\n", "", -1))
	}

	err = cmd.Wait()
	if err != nil {
		log.Printf("ERROR(%s): Execute command fail!", err.Error())
		return
	}
	//if err := cmd.Run(); err != nil {
	//	log.Printf("ERROR(%s): Execute command '%s' fail!", err, execStr)
	//	return err
	//}
	return
}

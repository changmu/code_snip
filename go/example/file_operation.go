package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 按行读取文件
func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		line = strings.TrimSpace(line)
		handler(line)
	}
	return nil
}

func main() {
	cnt := 1
	err := ReadLine("./file_operation.go", func(line string) {
		fmt.Printf("line:%d,[%v]\n", cnt, line)
		cnt++
	})
	if err != nil {
		panic(err)
	}
}

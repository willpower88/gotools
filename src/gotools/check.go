package main

import (
	"fmt"
	"os"
	"gotools/tools"
	"bufio"
	"io"
)

func main() {
	fmt.Println("hello")
	f, err := os.OpenFile("/Users/didi/cf-work/driverlog/driverid-grade1.txt", os.O_RDWR, 0666)
	tools.CheckError(err)
	defer f.Close()

	rd := bufio.NewReader(f);
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}
}

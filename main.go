package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Complete the staircase function below.
func staircase(n int32) {
	var arr []string
	num := n
	var s string
	var i int32
	var a int32
	var ii int32
	var iii int32
	i = 0
	for i < num {
		i = i + 1
		a = a + 1
		ii = 0
		for ii < num-a {
			s = s + " "
			ii = ii + 1
		}
		ii = 0
		iii = 0
		for iii < a {
			s = s + "#"
			iii = iii + 1
		}
		iii = 0
		arr = append(arr, s)
		s = ""
	}
	//fmt.Println(arr)
	for _, arrval := range arr {
		fmt.Println(arrval)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

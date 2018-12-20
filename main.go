package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */
	fmt.Println(s)

	if strings.HasSuffix(s, "PM") {
		s = strings.TrimSuffix(s, "PM")
		var hourint int
		ssplitted := strings.Split(s, ":")

		switch hour := ssplitted[0]; hour {
		case "01":
			hourint = 1
		case "02":
			hourint = 2
		case "03":
			hourint = 3
		case "04":
			hourint = 4
		case "05":
			hourint = 5
		case "06":
			hourint = 6
		case "07":
			hourint = 7
		case "08":
			hourint = 8
		case "09":
			hourint = 9
		case "10":
			hourint = 10
		case "11":
			hourint = 11
		case "12":
			ssplitted[0] = "12"
			hourint = 12
		}

		if hourint < 12 {
			ssplitted[0] = strconv.Itoa(hourint + 12)
		}

		s = ssplitted[0] + ":" + ssplitted[1] + ":" + ssplitted[2]
	}

	if strings.HasSuffix(s, "AM") {
		s = strings.TrimSuffix(s, "AM")
		ssplitted := strings.Split(s, ":")
		if strings.Contains(ssplitted[0], "12") {
			ssplitted[0] = "00"
		}
		s = ssplitted[0] + ":" + ssplitted[1] + ":" + ssplitted[2]
	}
	return s
}

func main() {
	s := "12:45:54PM"
	a := timeConversion(s)
	fmt.Println(a)
}

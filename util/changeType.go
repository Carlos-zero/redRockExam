package util

import "strconv"

func StringToInt(str string) int {
	if str=="" {
		return 0
	}
	int,err:=strconv.Atoi(str)
	FailOnErr(err,"string to int failure!")
	return int
}

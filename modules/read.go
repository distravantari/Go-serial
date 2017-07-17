package modules

import (
	"io/ioutil"
	"strconv"
)

func ReadTxtMax() int {

	dat, err := ioutil.ReadFile("max.txt")
	check(err)
	s := string(dat)
	out, _ := strconv.Atoi(s)
	return out
}

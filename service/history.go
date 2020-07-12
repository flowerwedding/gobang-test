package service

import (
	"strconv"
	"strings"
)

func Fenjie(s string) (i int ,j int,comma int) {
	comma = strings.Index(s, ",")
	i,_  = strconv.Atoi(s[1:comma])
	j, _ = strconv.Atoi(s[comma + 1:len(s)-1])
	return i,j,comma
}
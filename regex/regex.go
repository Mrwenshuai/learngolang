package main

import (
	"fmt"
	"regexp"
)

const text  =
	"My email is 115435@qq.com@QQ.COM test aa@cc.com adf daf@bb.com"

func main() {
	resp := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9.]+\\.[a-zA-Z0-9]+")
	//match := resp.FindString(text)
	match := resp.FindAllString(text,-1)
	fmt.Println(match)

}

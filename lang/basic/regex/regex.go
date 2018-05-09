package main

import (
	"fmt"
	"regexp"
)

const text = `My email is chen@qq.com
email2 is wandu@gmail.com
email 3 is chen@163.com.cn`

func main() {
	re := regexp.MustCompile(`(\w+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindString(text)         //匹配第一个邮箱
	match2 := re.FindAllString(text, -1) //匹配所有
	match3 := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	fmt.Println(match2)
	for _, m := range match3 {
		fmt.Println(m)
	}
}

// 创建一个方法，以将字符串从小写字母转换位大写字母
package main

import (
	"fmt"
	"strings"
)

// 基于基本类型创建自定义类型
type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func main() {
	s := upperstring("Learning Go!")
	fmt.Println(s)
	fmt.Println(s.Upper())
}

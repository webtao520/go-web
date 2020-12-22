package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		fmt.Println(strings.Contains("seafood", "foo")) // true
		s := []string{"foo", "bar", "baz"}
		fmt.Println(strings.Join(s, ", ")) // foo, bar, baz

		fmt.Println(strings.Index("lang go", "test")) // -1  如果找不到返回 -1

		fmt.Println("ba" + strings.Repeat("na", 3))

		%q      单引号围绕的字符字面值，由Go语法安全地转义 Printf("%q", 0x4E2D)        '中'
		fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	*/
	// Fields are: ["foo" "bar" "baz"]   Fields are: ["foobar" "baz"]
	fmt.Printf("Fields are: %q", strings.Fields("  foobar  baz   ")) // 去除s字符串的空格符，并且按照空格分割返回slice

}

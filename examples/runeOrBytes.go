package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var sentence string
	fmt.Printf("请输入需要计算的词或句子? ")
	//跳过可能的错误检查
	fmt.Scanf("%s", &sentence)

	fmt.Printf("%s 有%d个字符，共%d个字节！\n",
                    strconv.Quote(sentence),
                    utf8.RuneCountInString(sentence),
                    len(sentence))

	//可以基于此进行验证
	//for i, r := range sentence {
	//	fmt.Printf("sentence[%d]=\t%q\t0x%x\n", i, r, r)
	//}
}
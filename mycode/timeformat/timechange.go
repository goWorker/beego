package main

import (
	"fmt"
	//"log"
	"time"
)

func main() {
	//datetime := "20200428 01:28:35.465"  	    //待转化为时间戳的字符串

	tm2, _ := time.Parse("20060102 03:04:05", "20200428 01:27:12.465")

	fmt.Println(tm2.Unix())

}
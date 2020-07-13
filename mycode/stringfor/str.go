package main

import "fmt"

func main(){
	var path string = "6.0.0"
	var s string = "'"
	path = fmt.Sprintf("%s%s%s",s,path,s)
	fmt.Println(path)
}
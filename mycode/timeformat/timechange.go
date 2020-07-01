package main

import (
	"fmt"
	"time"
)
var timeLayoutStr = "2006-01-02 15:04:05"
func testFormat()  {
	t := time.Now() //当前时间
	t.Unix() //时间戳

	ts := t.Format(timeLayoutStr) //time转string
	fmt.Println(ts)
	st, _ := time.Parse(timeLayoutStr, ts) //string转time
	fmt.Println(st)

	//在go中, 可以格式化一个带前后缀的时间字符串
	prefixTStr := "PREFIX-- 2029-06-30 -TEST- 10:31:12 --SUFFIX" //带前后缀的时间字符串
	preTimeLayoutStr := "PREFIX-- 2006-01-02 -TEST- 15:04:05 --SUFFIX" //需要转换的时间格式, 格式和前后缀需要一致, 这种写法的限制很大, 但一些特殊场景可以用到
	prefixTime, _ := time.Parse(preTimeLayoutStr, prefixTStr)
	fmt.Println(prefixTime)

	//时间加减 time.ParseDuration()
	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	at, _ := time.ParseDuration("2h") //2个小时后的时间, 负数就是之前的时间
	fmt.Println((t.Add(at)).Format(timeLayoutStr))

	//两个时间差
	sub := t.Sub(prefixTime)
	fmt.Println(sub.Seconds()) //秒,  sub.Minutes()分钟,  sub.Hours()小时...

}

func main() {
	//datetime := "20200428 01:28:35.465"  	    //待转化为时间戳的字符串
	//
	//tm2, _ := time.Parse("20060102 03:04:05", "20200428 01:27:12.465")
	//
	//fmt.Println(tm2.Unix())
	//m, _ := time.ParseDuration("200s")
	//fmt.Printf("The movie is %.0f minutes long.", m.Minutes())



	//st, _ := time.Parse(timeLayoutStr, "2020-06-30 11:28:19")
	////fmt.Println(st)
	//
	//sub := time.Now().Sub(st)
	//fmt.Println(sub.Seconds())
	//switch {
	//case 0 < sub.Seconds() && sub.Seconds() < 60:
	//	m, _ := time.ParseDuration(strconv.Itoa(int(sub.Seconds())))
	//	fmt.Printf("The movie is %.0f seconds long.", m.Seconds())
	//case 60 <= sub.Seconds() && sub.Seconds() <= 3600:
	//	m, _ := time.ParseDuration(strconv.Itoa(int(sub.Seconds())))
	//	fmt.Printf("The movie is %.0f minutes long.", m.Minutes())
	//case 3600 < sub.Seconds() && sub.Seconds() <= 14400:
	//	m, _ := time.ParseDuration(strconv.Itoa(int(sub.Seconds())))
	//	fmt.Printf("The movie is %.0f hours long.", m.Hours())
	//case 14400 < sub.Seconds() && sub.Seconds() <= 432000:
	//	m, _ := time.ParseDuration(strconv.Itoa(int(sub.Seconds())))
	//	fmt.Printf("The movie is %.0f day long.", m.Round(sub))
	//}
	testFormat()

}
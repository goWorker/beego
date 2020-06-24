package utils

import (
	"fmt"
	"reflect"
	"time"
)

func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

func ConvertTimeToSeconds(timestring string) int64{
	tm2, _ := time.Parse("20060102 03:04:05", timestring)

	return tm2.Unix()
}


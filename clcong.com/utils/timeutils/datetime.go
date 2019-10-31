package timeutils

import (
	"time"
	//"fmt"
)

func Unix64TimeToUnix32Time(data int64) int32 {

	// value := data >> 32
	return (int32)(data)
}
func GetUnix13NowTime() int64 {
	unixtime := time.Now().UTC().UnixNano()
	return unixtime / 1000000
}

func GetTimeStamp() int32 {
	return int32(GetUnix13NowTime() / 1000)
}

func GetUnixTimeStamp() int64 {
	unixtime := time.Now().UTC().UnixNano()
	return unixtime / 1000000000
}

//字符串转13位时间戳
func StringToTimeStamp(timeStr string) int64 {
	//第一个参数是格式，第二个是要转换的时间字符串
	tm2, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	//fmt.Println("err:", err.Error())
	return tm2.UTC().UnixNano() / 1000000

}



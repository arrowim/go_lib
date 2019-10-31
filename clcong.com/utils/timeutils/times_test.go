package timeutils

import (
	"fmt"
	"testing"
	"time"
)

func TestUnix64TimeToUnix32Time(t *testing.T) {
	var data int64
	data = time.Now().Unix()

	d2 := Unix64TimeToUnix32Time(data)

	d := StringToTimeStamp("2017-12-12 00:00:00")

	fmt.Println(fmt.Sprintf("aa:%d",d))

	fmt.Println(data)
	fmt.Println(d2)
}

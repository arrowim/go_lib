package osutil

import (
	//	"fmt"
	"os"
	//	"runtime/debug"
	//	"time"

	logs "github.com/YoungPioneers/blog4go"
)

func Exit(code int) {
	defer func() {
		if err := recover(); err != nil {
		}
		os.Exit(code)
	}()
	//	time.Sleep(time.Second)

	//	fmt.Println(string (debug.Stack()))
	logs.Close()
	os.Exit(code)
}

package main

import (
	"log"
)

func Test() {
	wr, _ := NewRhapLogger(&LogConfig{
		Filename: "F:\\project\\rhap-logger\\test.log",
		Stdout:   true,
		Prefix:   "",
		Flags:    log.Llongfile,
	})

	info := wr.NewLogInfo()

	//todo reformat output menjadi JSON
	wr.Logger().Print(info)
}

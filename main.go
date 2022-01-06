package main

import (
	"log"

	"github.com/bernylinville/learn-wtf/flags"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // 设置log标准输出格式

	// paras and handle flags
	flags := flags.NewFlags()
	flags.Parse()
}

package loggerinfo

import (
	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	var logpath = build.Default.GOPATH + "/src/chat/logger/info.log"

	fmt.Println(build.Default.GOPATH + "/src/chat/logger/info.log")

	flag.Parse()
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}

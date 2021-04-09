package main

import (
	"flag"
	"fmt"
	"github.com/google/logger"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/driver"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/middleware"
	"io/ioutil"
	"runtime"
	"time"
)

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// args
	flag.Parse()
	// log
	defer logger.Init("LoggerExample", *verbose, false, ioutil.Discard).Close()
	// db
	driver.Connect(10, 100)
	defer driver.Disconnect()
	// middleware
	middleware.Start()
}

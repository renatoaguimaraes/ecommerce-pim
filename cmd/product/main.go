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

	//defer elapsed("page")()
	//file, err := os.Open("/Users/renato/Downloads/Age-sex-by-ethnic-group-grouped-total-responses-census-usually-resident-population-counts-2006-2013-2018-Censuses-RC-TA-SA2-DHB/Data8277.csv")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//
	//fileinfo, err := file.Stat()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//filesize := fileinfo.Size()
	//buffer := make([]byte, filesize)
	//
	//bytesread, err := file.Read(buffer)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("bytes read: ", bytesread)
	// r := csv.NewReader(bytes.NewReader(buffer))
	// arr, _ := r.ReadAll()
	// fmt.Println("bytestream to string: ", len(arr))

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

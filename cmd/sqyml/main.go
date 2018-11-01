package main

import (
	"flag"
	"fmt"
	// "github.com/19700101000000/go-sqyml"
	// "io/ioutil"
	// "log"
	"sync"
)

func fileOpen(wg *sync.WaitGroup, file_name string) {
	defer wg.Done()
	fmt.Println(file_name)
}

func main() {
	flag.Parse()

	args := flag.Args()
	wg := new(sync.WaitGroup)

	for _, arg := range args {
		wg.Add(1)
		go fileOpen(wg, arg)
	}
	wg.Wait()
}

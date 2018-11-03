package main

import (
	"os"
	"sync"
	"testing"
)

func ExampleFileOpen() {
	wg := new(sync.WaitGroup)
	str := "foo"

	wg.Add(1)
	fileOpen(wg, str)
	// Output:
	// foo
}

func TestMain(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"foo", "bar"}
	main()
}

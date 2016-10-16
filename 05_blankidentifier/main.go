package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// This code works, but is a little untidy due to the error checking.

	// res, err := http.Get("https://crocker.io/")
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// page, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("%s", page)

	// This code is much tidier:
	// _ is the blank identifier
	res, _ := http.Get("https://crocker.io")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

// In Go, you're unable to have an unused variable; if you declare a variable,
// but don't use it, then you'll get an error.

// This is only really useful in prototype code; in production code you want to
// check for errors.

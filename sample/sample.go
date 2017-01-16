package main

import (
	"errors"
	"fmt"
	"github.com/Simbory/trycatch"
)

func main() {
	trycatch.TryCatch(func(){
		panic(errors.New("exception message for trycache.TryCache"))
	}, func(ex interface{}) {
		fmt.Println(ex.(error).Error())
	})

	trycatch.TryFinally(func(){
		panic(errors.New("exception message for trycache.TryFinally"))
	}, func(){
		fmt.Println("msg1: finally function is called")
	})

	trycatch.TryFinally(func(){
		fmt.Println("code block for try")
	}, func(){
		fmt.Println("msg2: finally function is called")
	})

	trycatch.TryCatchFinally(func(){
		fmt.Println("code block for try")
		panic(errors.New("exception message for trycatch.TryCatchFinally"))
	}, func(ex interface{}) {
		fmt.Println(ex.(error).Error())
	}, func() {
		fmt.Println("msg2: finally function is called")
	})
}

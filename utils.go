package main

import "fmt"

func testPanic() {
	panic("I'm panicing!")
}

func testAfterPanic() {
	fmt.Printf("I paniced, but I'm ok now...")
}

package main

type greeting string

func (g greeting) Greet(log string) {
	println("hello hotfix:" + log)
}

var Greeter greeting

func HotFixPrint(log string) {
	// println(log)
}

func PrintNoParam() {
	
}
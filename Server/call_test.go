package main

import (
	"plugin"
	"testing"
)

func BenchmarkCallHotFix(b *testing.B) {
	plug, _ := plugin.Open("../HotFix/greet.so")
	symPrint, _ := plug.Lookup("HotFixPrint")
	hPrint := symPrint.(func(string))
	b.ResetTimer()
	for i:= 0; i < b.N; i++ {
		hPrint("test string")
	}
}

func localPrint(log string) {
	// println(log)
}

func BenchmarkCallLocal(b *testing.B) {
	for i := 0; i< b.N; i++ {
		localPrint("test string")
	}
}


func localPrintNoParam() {	
}

func BenchmarkCallHotFixNoParam(b *testing.B) {
	plug, _ := plugin.Open("../HotFix/greet.so")
	symPrint, _ := plug.Lookup("PrintNoParam")
	hPrint := symPrint.(func())
	b.ResetTimer()
	for i:= 0; i < b.N; i++ {
		hPrint()
	}
}

func BenchmarkCallLocalNoParam(b *testing.B) {
	for i := 0; i< b.N; i++ {
		localPrintNoParam()
	}
}
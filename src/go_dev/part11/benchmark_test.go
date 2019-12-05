package main

import "testing"

func BenchMark_Add(b *testing.B) {
	var n int
	for i:=0;i<b.N;i++{
		n++
	}

}

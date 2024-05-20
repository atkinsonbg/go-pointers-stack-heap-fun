package main

import "testing"

func Benchmark_PassByValueReturnValueConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassByValueReturnValueConcatenation("test")
	}
}

func Benchmark_PassByPointerReturnValueConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "test"
		PassByPointerReturnValueConcatenation(&s)
	}
}

func Benchmark_PassByPointerReturnPointerConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "test"
		PassByPointerReturnPointerConcatenation(&s)
	}
}

func Benchmark_PassByValueReturnValueFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassByValueReturnValueFmt("test")
	}
}

func Benchmark_PassByValueReturnValueBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassByValueReturnValueBuilder("test")
	}
}

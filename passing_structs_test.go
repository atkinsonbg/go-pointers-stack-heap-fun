package main

import "testing"

func Benchmark_PassStructByValueReturnByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := SmallStruct{
			ID: 1,
		}
		PassStructByValueReturnByValue(s)
	}
}

func Benchmark_PassStructByPointerReturnByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := SmallStruct{
			ID: 1,
		}
		PassStructByPointerReturnByPointer(&s)
	}
}

package main

import (
	"bytes"
	"io"
	"testing"
)

func writeToBuffer(w io.Writer, msg []byte) {
	w.Write(msg)
}

func BenchmarkWriteToBuffer(b *testing.B) {
	msg := []byte("Foo")
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ { //internal benching logic by Golang
		for i := 0; i < 100; i++ {
			writeToBuffer(buf, msg)
		}
	}
}

func BenchmarkSliceNonAlloc(b *testing.B) {
	n := 100

	b.Run("non alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ints := []int{}
			for i := 0; i < n; i++ {
				ints = append(ints, i)
			}
		}
	})

	b.Run("alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ints := make([]int, n)
			for i := 0; i < n; i++ {
				ints[i] = i
			}
		}
	})
}

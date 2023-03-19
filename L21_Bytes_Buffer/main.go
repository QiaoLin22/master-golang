package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type BytesClosingBuffer struct {
	*bytes.Buffer
	io.Closer
}

func NewBytesClosingBuffer() *BytesClosingBuffer {
	return &BytesClosingBuffer{
		Buffer: new(bytes.Buffer),
	}
}

func (b *BytesClosingBuffer) Close() error {
	fmt.Println("closing buffer...")
	return nil
}

func writeTo(wc io.WriteCloser, msg []byte) error {
	defer wc.Close()
	_, err := wc.Write(msg)
	return err
}

func main() {
	buf := NewBytesClosingBuffer()

	if err := writeTo(buf, []byte("hello world")); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

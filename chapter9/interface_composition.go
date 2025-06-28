package main

import (
	"bytes"
	"fmt"
	"io"
)

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// just like we embed two structs, we can also compose two or more interfaces
type WriterCloser interface {
	Writer
	Closer
}

// Now, whichever type that implements methods Write() and Close(), that is said to
// implement WriterCloser interface

type BufferedWriterCloser struct {
	buffer *bytes.Buffer // internal buffer
}

func MainInterfaceComposition() {

	writer := NewBufferedWriterCloser()
	writer.Write([]byte("Hello, I am Lazinerd. I am learning GO from Scratch"))
	writer.Close()

	// inorder to type converge later, you need to define
	// the variable of type WriterCloser interface
	var wc WriterCloser = NewBufferedWriterCloser()

	// type assertion/conversion
	// sometimes we may have to work with concrete data inside the named types
	// but we can't figure out thaat using some interfaces
	// for that reason: we use type conversion to get the underlying type to work with

	data, ok := wc.(*BufferedWriterCloser) // then type converge here...
	fmt.Println(data, ok)

	if khurapatiData, ok := wc.(io.Reader); ok {
		fmt.Println(khurapatiData)
	} else {
		fmt.Println("Conversion Failed !!")
	}

	// or just call the function that accepts the interface type as an argument
	Public(writer)
	// Public(io.Reader) // cannot type converge the interface because our interface does not have Read() method defined
}

// implements Write methods
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data) // .Write(data) appends the content of data to the buffer, growing the buffer
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8) // creating a memory that can store 8 bits

	// as long as the buffer has more than 8 characters, its going to write it out.
	// but as soon as the buffer has less than 8 characters, its not goint to write it out.
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v) // reads the next len(v) byte of data from the buffer
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v)) // returns the number of bytes written and err if any
		if err != nil {
			return 0, err
		}
	}

	return n, err
}

// now while printing the data from the buffer, the last chunk may have the length less than 8
// so its gonna be missed out by the Write() methods
// so we implement the Close() method that can flush the remaining buffered data

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8) // returns a slice containing the next n bytes from the buffer,
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func Public(wc WriterCloser) {

	fmt.Println(wc)

	// type switches
	// switch v := wc.(type) {
	// case *BufferedWriterCloser:
	// 	{
	// 		fmt.Println(*v.buffer)

	// 	}
	// }

}

// We want to calculate in the form of chunks of data, so in such situations as needing the length of the overly large string,
// we can use the buffer; the buffer allows us to handle any size of the dynamic length, making it flexible.

package main

import (
	"bytes" // importing the bytes package so that buffer can be used
	"fmt"
)

// func main() {
// 	//Creating buffer variable to hold and manage the string data
// 	var strBuffer bytes.Buffer
// 	strBuffer.WriteString("Ranjan")
// 	strBuffer.WriteString("Kumar")
// 	fmt.Println("The string buffer output is", strBuffer.String()) // out: The string buffer output is RanjanKumar
// }

func main() {
	//Creating buffer variable to hold and manage the string data
	var strByte bytes.Buffer
	strByte.Grow(64)                         // Grow(64) method preallocates enough capacity in the buffer to hold 64 bytes
	strBytestrByte := strByte.Bytes()        //Bytes() method returns the current underlying byte slice of the buffer
	strByte.Write([]byte("It is a 64 byte")) // Now buffer's length is 15 (number of bytes in the string), capacity remains 64
	fmt.Printf("%b", strBytestrByte[:strByte.Len()])
	//The %b format specifier in fmt.Printf prints the binary representation of each byte in the slice.
}

package gllrb_test

import (
	"fmt"

	"github.com/levigross/gllrb"
)

func Example_basicUsage() {
	word := []byte("hello")

	// Create a new left leaning red black tree (LLRB)

	bytesLLRB := gllrb.NewLLRB()

	// every item placed into the LLRB needs to be encapsulated within a `Comparer` interface

	bytesLLRB.Put(gllrb.Bytes(word))

	// The library supports the following wrappers
	gllrb.String("foo")
	gllrb.UInt(123)
	gllrb.Int(-123)

	// Get the item as an interface{} so it must be typecast
	myWord := bytesLLRB.Get(gllrb.Bytes(word)).([]byte)

	// Delete the item
	bytesLLRB.Delete(gllrb.Bytes(word))

	// Trying to get an item that doesn't exist will return a nil
	if bytesLLRB.Get(gllrb.Bytes(word)) == nil { // true
		fmt.Println("This will always print", myWord)
	}

}

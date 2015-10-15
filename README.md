# GLLRB
Left Leaning Red Black Tree written in Go

[![Build Status](https://travis-ci.org/levigross/gllrb.svg?branch=master)](https://travis-ci.org/levigross/gllrb) [![GoDoc](https://godoc.org/github.com/levigross/gllrb?status.svg)](https://godoc.org/github.com/levigross/gllrb) [![Coverage Status](https://coveralls.io/repos/levigross/gllrb/badge.svg?branch=master&service=github)](https://coveralls.io/github/levigross/gllrb?branch=master)

License
======

GLLRB is licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for the full license text

Built In Primitives
===================

- `string` -> `String`
- `[]byte` -> `Bytes`
- `uint` -> `UInt`
- `int` -> `Int`

Example
=======

```go

import "github.com/levigross/gllrb"

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
myWord := bytesLLRB.Get(gllrb.Bytes(word)).(gllrb.ByteComparer)

// Delete the item
bytesLLRB.Delete(gllrb.Bytes(word))

// Trying to get an item that doesn't exist will return a nil
if bytesLLRB.Get(gllrb.Bytes(word)) == nil { // true
    fmt.Println("This will always print", myWord)
}


```


Building Tree Primitives
========================

All objects that are inserted into the LLRB must implement the `Comparer` interface

```go
type Comparer interface {
    // Compare should implement a response akin to bytes.Compare
    // The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
    Compare(treeKey Comparer) int
    Value() interface{} // the value that is returned
}
```

Here is an example of the `StringComparer` built into the library

```go
// String is used when a user wishes to insert a string into the LLRB
func String(v string) StringComparer {
	return StringComparer(v)
}

// Value will return the value of our string as an {}interface
func (sc StringComparer) Value() interface{} {
	return sc
}

// Compare compares two strings (and follows the return format of bytes.Compare)
func (sc StringComparer) Compare(a Comparer) int {
	if sc.Value().(StringComparer) > a.Value().(StringComparer) {
		return +1
	}

	if sc.Value().(StringComparer) < a.Value().(StringComparer) {
		return -1
	}

	return 0

}
```

Concurrency
===========

GLLRB does *not* offer any form of concurrency protection.

package gllrb

import "bytes"

// Comparer is an interface that we use to compare LLRB keys
type Comparer interface {
	Compare(treeKey Comparer) int
	Value() interface{}
}

// BytesComparer is a wrapper function around a []byte value that enables us
// to use it with our red black tree
type BytesComparer struct {
	value []byte
}

// Bytes inserts the []byte type into our LLRB
func Bytes(v []byte) *BytesComparer {
	return &BytesComparer{value: v}
}

// Value will return the value of our bytes as an {}interface
func (bc *BytesComparer) Value() interface{} {
	return bc.value
}

// Compare is a shadow method for bytes.Compare – make sure you pass
// BytesComparer or the program will panic!
func (bc *BytesComparer) Compare(a Comparer) int {
	return bytes.Compare(bc.Value().([]byte), a.Value().([]byte))
}

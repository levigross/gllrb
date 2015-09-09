package gllrb

import "bytes"

// Comparer is an interface that we use to compare LLRB keys
type Comparer interface {
	Compare(treeKey Comparer) int
	Value() interface{}
}

// ByteComparer is a wrapper struct around a []byte value that enables us
// to use it with our red black tree
type ByteComparer struct {
	value []byte
}

// UIntComparer is a wrapper struct around an unsigned integer
// that enables us to use it with our red black tree.
// The internal type is uint64
type UIntComparer struct {
	value uint64
}

// IntComparer is a wrapper struct around an integer
// that enables us to use it with our red black tree.
// The internal type is int64
type IntComparer struct {
	value int64
}

// StringComparer is a wrapper struct around a string that enables us
// to use it with our red black tree
type StringComparer struct {
	value string
}

// UInt is used when a user wishes to insert a string into the LLRB
func UInt(v uint) *UIntComparer {
	return &UIntComparer{value: uint64(v)}
}

// Value will return the value of our string as an {}interface
func (ui *UIntComparer) Value() interface{} {
	return ui.value
}

// Compare compares two unsigned integers (and follows the return format of bytes.Compare)
func (ui *UIntComparer) Compare(a Comparer) int {
	if ui.Value().(uint64) > a.Value().(uint64) {
		return +1
	}

	if ui.Value().(uint64) < a.Value().(uint64) {
		return -1
	}

	return 0
}

// Int is used when a user wishes to insert a string into the LLRB
func Int(v int) *IntComparer {
	return &IntComparer{value: int64(v)}
}

// Value will return the value of our string as an {}interface
func (i *IntComparer) Value() interface{} {
	return i.value
}

// Compare compares two unsigned intergers (and follows the return format of bytes.Compare)
func (i *IntComparer) Compare(a Comparer) int {
	if i.Value().(int64) > a.Value().(int64) {
		return +1
	}

	if i.Value().(int64) < a.Value().(int64) {
		return -1
	}

	return 0
}

// String is used when a user wishes to insert a string into the LLRB
func String(v string) *StringComparer {
	return &StringComparer{value: v}
}

// Value will return the value of our string as an {}interface
func (sc *StringComparer) Value() interface{} {
	return sc.value
}

// Compare compares two strings (and follows the return format of bytes.Compare)
func (sc *StringComparer) Compare(a Comparer) int {
	if sc.Value().(string) > a.Value().(string) {
		return +1
	}

	if sc.Value().(string) < a.Value().(string) {
		return -1
	}

	return 0

}

// Bytes inserts the []byte type into our LLRB
func Bytes(v []byte) *ByteComparer {
	return &ByteComparer{value: v}
}

// Value will return the value of our bytes as an {}interface
func (bc *ByteComparer) Value() interface{} {
	return bc.value
}

// Compare is a shadow method for bytes.Compare – make sure you pass
// ByteComparer or the program will panic!
func (bc *ByteComparer) Compare(a Comparer) int {
	return bytes.Compare(bc.Value().([]byte), a.Value().([]byte))
}

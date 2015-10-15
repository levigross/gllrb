package gllrb

import "bytes"

// Comparer is an interface that we use to compare LLRB keys
type Comparer interface {
	Compare(treeKey Comparer) int
	Value() interface{}
}

// ByteComparer is a wrapper struct around a []byte value that enables us
// to use it with our red black tree
type ByteComparer []byte

// UIntComparer is a wrapper struct around an unsigned integer
// that enables us to use it with our red black tree.
// The internal type is uint64
type UIntComparer uint64

// IntComparer is a wrapper struct around an integer
// that enables us to use it with our red black tree.
// The internal type is int64
type IntComparer int64

// StringComparer is a wrapper struct around a string that enables us
// to use it with our red black tree
type StringComparer string

// UInt is used when a user wishes to insert a uint into the LLRB
func UInt(v uint) UIntComparer {
	return UIntComparer(v)
}

// Value will return the value of our uint as an {}interface
func (ui UIntComparer) Value() interface{} {
	return ui
}

// Compare compares two unsigned integers (and follows the return format of bytes.Compare)
func (ui UIntComparer) Compare(a Comparer) int {
	if ui.Value().(UIntComparer) > a.Value().(UIntComparer) {
		return +1
	}

	if ui.Value().(UIntComparer) < a.Value().(UIntComparer) {
		return -1
	}

	return 0
}

// Int is used when a user wishes to insert a int into the LLRB
func Int(v int) IntComparer {
	return IntComparer(v)
}

// Value will return the value of our int as an {}interface
func (i IntComparer) Value() interface{} {
	return i
}

// Compare compares two unsigned intergers (and follows the return format of bytes.Compare)
func (i IntComparer) Compare(a Comparer) int {
	if i.Value().(IntComparer) > a.Value().(IntComparer) {
		return +1
	}

	if i.Value().(IntComparer) < a.Value().(IntComparer) {
		return -1
	}

	return 0
}

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

// Bytes inserts the []byte type into our LLRB
func Bytes(v []byte) ByteComparer {
	return ByteComparer(v)
}

// Value will return the value of our bytes as an {}interface
func (bc ByteComparer) Value() interface{} {
	return bc
}

// Compare is a shadow method for bytes.Compare – make sure you pass
// ByteComparer or the program will panic!
func (bc ByteComparer) Compare(a Comparer) int {
	return bytes.Compare([]byte(bc.Value().(ByteComparer)), []byte(a.Value().(ByteComparer)))
}

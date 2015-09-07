package gllrb

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestLLRBCreate(t *testing.T) {
	tree := NewLLRB()
	if tree.root != nil {
		t.Error("Initial LLRB node (root node) must be nil")
	}
}

func TestLLRBInsert(t *testing.T) {
	dict, err := ioutil.ReadFile("test_data/words")
	if err != nil {
		t.Fatal("Unable to read words dict", err)
	}
	words := bytes.Split(dict, []byte("\n"))
	llrb := NewLLRB()
	for _, word := range words {
		llrb.Put(Bytes(word))
	}

	llrb.Put(Bytes([]byte("while")))

	if uint64(len(words)) != llrb.Size() {
		t.Error("RB tree height not where it needs to be. Is",
			llrb.root.Number, "Should be", len(words))
	}

	if string(llrb.root.Value.([]byte)) != "funnel" {
		t.Error("Root value is not wrong", string(llrb.root.Value.([]byte)))
	}
}

func TestLLRBMax(t *testing.T) {
	dict, err := ioutil.ReadFile("test_data/words")
	if err != nil {
		t.Fatal("Unable to read words dict", err)
	}
	words := bytes.Split(dict, []byte("\n"))
	llrb := NewLLRB()
	if m := llrb.Max(); m != nil {
		t.Error("Blank tree returns non-nil when Max is called")
	}

	for _, word := range words {
		llrb.Put(Bytes(word))
	}

	if string(llrb.Max().([]byte)) != "A" {
		t.Error("Right most element on the tree isn't 'A' it is", string(llrb.Max().([]byte)))
	}
}

func TestLLRBMin(t *testing.T) {
	dict, err := ioutil.ReadFile("test_data/words")
	if err != nil {
		t.Fatal("Unable to read words dict", err)
	}
	words := bytes.Split(dict, []byte("\n"))
	llrb := NewLLRB()

	if m := llrb.Min(); m != nil {
		t.Error("Blank tree returns non-nil when Min is called")
	}

	for _, word := range words {
		llrb.Put(Bytes(word))
	}

	if string(llrb.Min().([]byte)) != "zythum" {
		t.Error("Right most element on the tree isn't 'zythum' it is ", string(llrb.Min().([]byte)))
	}
}

func TestLLRBDelete(t *testing.T) {
	dict, err := ioutil.ReadFile("test_data/words")
	if err != nil {
		t.Fatal("Unable to read words dict", err)
	}
	words := bytes.Split(dict, []byte("\n"))
	llrb := NewLLRB()
	for _, word := range words {
		llrb.Put(Bytes(word))
	}

	llrb.Delete(Bytes([]byte("while")))
	//
	// if sen := llrb.Get(Bytes([]byte("while"))); sen != nil {
	// 	t.Error("Word 'while'  in LLRB")
	// }
}

func TestLLRBGet(t *testing.T) {
	dict, err := ioutil.ReadFile("test_data/words")
	if err != nil {
		t.Fatal("Unable to read words dict", err)
	}
	words := bytes.Split(dict, []byte("\n"))

	llrb := NewLLRB()

	if n := llrb.Get(Bytes([]byte("foo"))); n != nil {
		t.Error("Blank LLRB returns value for get")
	}

	for _, word := range words {
		llrb.Put(Bytes(word))
	}

	if uint64(len(words)) != llrb.root.Number {
		t.Error("RB tree height not where it needs to be. Is",
			llrb.root.Number, "Should be", len(words))
	}

	if string(llrb.root.Value.([]byte)) != "funnel" {
		t.Error("Root value is not wrong", string(llrb.root.Value.([]byte)))
	}

	if sen := llrb.Get(Bytes([]byte("while"))); bytes.Compare(sen.([]byte), []byte("while")) != 0 {
		t.Error("Word 'while' not in LLRB")
	}

	if sen := llrb.Get(Bytes([]byte("fjdaslkfdslka"))); sen != nil {
		t.Error("Somehow we have fjdaslkfdslka in our LLRB ")
	}
}

func BenchmarkLLRBInsert(b *testing.B) {
	dict, err := ioutil.ReadFile("test_data/words")
	if err != nil {
		b.Fatal("Unable to read words dict", err)
	}
	words := bytes.Split(dict, []byte("\n"))
	llrb := NewLLRB()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		llrb.Put(Bytes(words[rand.Intn(len(words))]))
	}
}

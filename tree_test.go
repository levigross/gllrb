package gllrb

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"testing"
)

func WordList() [][]byte {
	dict, err := ioutil.ReadFile("test_data/words")
	if err != nil {
		log.Fatal("Unable to read words dict", err)
	}
	return bytes.Split(dict, []byte("\n"))

}

func TestLLRBCreate(t *testing.T) {
	tree := NewLLRB()
	if tree.root != nil {
		t.Error("Initial LLRB node (root node) must be nil")
	}
}

func TestLLRBInsert(t *testing.T) {
	words := WordList()
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

func TestGoMapInsert(t *testing.T) {
	words := WordList()
	gomap := map[string][]byte{}
	for _, word := range words {
		gomap[string(word)] = word
	}

	if len(words) != len(gomap) {
		t.Error("RB tree height not where it needs to be. Is",
			len(gomap), "Should be", len(words))
	}

}

func TestLLRBMax(t *testing.T) {
	words := WordList()
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
	words := WordList()
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
	words := WordList()
	llrb := NewLLRB()
	for _, word := range words {
		llrb.Put(Bytes(word))
	}

	for _, word := range words {
		llrb.Delete(Bytes(word))
		if sen := llrb.Get(Bytes(word)); sen != nil {
			t.Error("Word", string(word), "in LLRB")
		}
	}
}

func TestLLRBGet(t *testing.T) {
	words := WordList()

	llrb := NewLLRB()

	if n := llrb.Get(Bytes([]byte("foo"))); n != nil {
		t.Error("Blank LLRB returns value for get")
	}

	for _, word := range words {
		llrb.Put(Bytes(word))
	}

	if uint64(len(words)) != llrb.Size() {
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
	words := WordList()
	llrb := NewLLRB()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		llrb.Put(Bytes(words[rand.Intn(len(words))]))
	}
}

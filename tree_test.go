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

	if string(llrb.root.Value.([]byte)) != "pockwood" {
		t.Error("Root value is not 'pockwood", string(llrb.root.Value.([]byte)))
	}
}

func TestLLRBInsertUInt(t *testing.T) {

	llrb := NewLLRB()
	for i := 0; i < 30000; i++ {
		llrb.Put(UInt(uint(i)))
	}

	if uint64(30000) != llrb.Size() {
		t.Error("RB tree height not where it needs to be. Is",
			llrb.root.Number, "Should be", 300000)
	}

	if llrb.root.Value.(uint64) != uint64(16383) {
		t.Error("Root value is not '16383", llrb.root.Value.(uint64))
	}

	if llrb.Max().(uint64) != uint64(29999) {
		t.Error("Right most element on the tree isn't '29999' it is", llrb.Max().(uint64))
	}

	if llrb.Min().(uint64) != uint64(0) {
		t.Error("Left most element on the tree isn't '0' it is ", llrb.Min().(uint64))
	}

	for i := 0; i < 30000; i++ {
		llrb.Delete(UInt(uint(i)))
		if sen := llrb.Get(UInt(uint(i))); sen != nil {
			t.Error("Number", i, "in LLRB")
		}
	}

}

func TestLLRBInsertInt(t *testing.T) {

	llrb := NewLLRB()
	for i := -15000; i < 15000; i++ {
		llrb.Put(Int(int(i)))
	}

	if uint64(30000) != llrb.Size() {
		t.Error("RB tree height not where it needs to be. Is",
			llrb.root.Number, "Should be", 300000)
	}

	if llrb.root.Value.(int64) != int64(1383) {
		t.Error("Root value is not '16383", llrb.root.Value.(int64))
	}

	if llrb.Max().(int64) != int64(14999) {
		t.Error("Right most element on the tree isn't '14999' it is", llrb.Max().(int64))
	}

	if llrb.Min().(int64) != int64(-15000) {
		t.Error("Left most element on the tree isn't '-15000' it is ", llrb.Min().(int64))
	}

	for i := -15000; i < 15000; i++ {
		llrb.Delete(Int(int(i)))
		if sen := llrb.Get(Int(int(i))); sen != nil {
			t.Error("Number", i, "in LLRB")
		}
	}

}

func TestLLRBInsertString(t *testing.T) {
	words := WordList()
	llrb := NewLLRB()
	for _, word := range words {
		llrb.Put(String(string(word)))
	}

	if uint64(len(words)) != llrb.Size() {
		t.Error("RB tree height not where it needs to be. Is",
			llrb.root.Number, "Should be", len(words))
	}

	if llrb.root.Value.(string) != "pockwood" {
		t.Error("Root value is not 'pockwood", llrb.root.Value.(string))
	}

	if llrb.Min().(string) != "A" {
		t.Error("Left most element on the tree isn't 'A' it is", llrb.Min().(string))
	}

	if llrb.Max().(string) != "zythum" {
		t.Error("Right most element on the tree isn't 'zythum' it is ", llrb.Max().(string))
	}

	for _, word := range words {
		llrb.Delete(String(string(word)))
		if sen := llrb.Get(String(string(word))); sen != nil {
			t.Error("Word", string(word), "in LLRB")
		}
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

	if string(llrb.Max().([]byte)) != "zythum" {
		t.Error("Right most element on the tree isn't 'zythum' it is", string(llrb.Max().([]byte)))
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

	if string(llrb.Min().([]byte)) != "A" {
		t.Error("Left most element on the tree isn't 'A' it is ", string(llrb.Min().([]byte)))
	}
}

func TestLLRBDelete(t *testing.T) {
	words := WordList()
	llrb := NewLLRB()

	llrb.Delete(Bytes([]byte("while")))

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

	if string(llrb.root.Value.([]byte)) != "pockwood" {
		t.Error("Root value is not 'pockwood", string(llrb.root.Value.([]byte)))
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

package main

import (
	"fmt"
	"hashtable/hashing"
	"hashtable/tableLP"
	"testing"
)

func baseTesting(t *testing.T, hf hashing.HashFunction) {
	size := 7
	ht := tableLP.New(size, hf)
	data := []int{12, 1, 4, 34, 13, 56, 69}
	for _, key := range data {
		// ht.Insert(key)
		ok := ht.Insert(key)
		if ok {
			// check if same item inserted
			hash := hf(key, size)
			if ht.NthItem(hash).Key() != key {
				t.Errorf("item mismatched - %d", key)
			}
		} else {
			// check for collision
			hash := hf(key, size)
			if ht.NthItem(hash) == nil {
				t.Errorf("failed without any collision - %d", key)
			}
		}
	}
	fmt.Println(ht.NilString())
}

func TestDivision(t *testing.T) {
	baseTesting(t, hashing.DivisionMethod)
}

func TestMultiplication(t *testing.T) {
	baseTesting(t, hashing.MultiplicationMethod)
}

func TestMidSquare(t *testing.T) {
	baseTesting(t, hashing.MidSquareMethod)
}

func TestFolding(t *testing.T) {
	baseTesting(t, hashing.FoldingMethod)
}

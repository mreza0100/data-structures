package hashtable

import (
	"bytes"
	"fmt"

	"data-structures/utils"
)

type pairT struct {
	key   interface{}
	value interface{}
}

type hashTable struct {
	bucket [][]*pairT
	size   int
}

type controller interface {
	Set(key, value interface{}) bool
	Get(key interface{}) interface{}
	Delete(key interface{}) bool
	Loop(func(key, value interface{}) bool)
	Print()
}

func NewHastable(size int) controller {
	bucket := make([][]*pairT, size)

	return &hashTable{
		bucket: bucket,
		size:   size,
	}
}

func (ht *hashTable) hash(key interface{}) (idx int) {
	inByte := utils.ToByte(key)

	for _, b := range inByte {
		idx += int(b)
	}

	return idx % ht.size
}

func (ht *hashTable) Set(key, value interface{}) bool {
	idx := ht.hash(key)

	for _, p := range ht.bucket[idx] {
		if p.key == key {
			p.value = value
			return true
		}
	}

	ht.bucket[idx] = append(ht.bucket[idx], &pairT{key: key, value: value})

	return true
}

func (ht *hashTable) compare(keyBytes []byte, pairKey interface{}) bool {
	pairKeyBytes := utils.ToByte(pairKey)

	return bytes.Equal(keyBytes, pairKeyBytes)
}

func (ht *hashTable) Get(key interface{}) interface{} {
	hash := ht.hash(key)
	keyBytes := utils.ToByte(key)

	for _, p := range ht.bucket[hash] {
		if ht.compare(keyBytes, p.key) {
			return p.value
		}
	}

	return nil
}

func (ht *hashTable) Delete(key interface{}) bool {
	hash := ht.hash(key)
	newBucket := make([]*pairT, 0, len(ht.bucket[hash]))
	keyBytes := utils.ToByte(key)

	for _, p := range ht.bucket[hash] {
		if !ht.compare(keyBytes, p.key) {
			newBucket = append(newBucket, p)
		}
	}
	ht.bucket[hash] = newBucket

	return false
}

func (ht *hashTable) Loop(executor func(key, value interface{}) bool) {
bucket_loop:
	for _, b := range ht.bucket {
		for _, p := range b {
			if !executor(p.key, p.value) {
				break bucket_loop
			}
		}
	}
	fmt.Printf("\n")
}

func (ht *hashTable) Print() {
	for _, b := range ht.bucket {
		fmt.Printf("\n--")
		for _, p := range b {
			fmt.Printf("%+v", p)
		}
	}
}

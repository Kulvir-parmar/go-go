package main

import (
	"fmt"

	"github.com/twmb/murmur3"
)

var hasher = murmur3.New32()

/*
Murmur Hash function is used to hash the values
https://en.wikipedia.org/wiki/MurmurHash

hash(key) -> returns a number

insertion index = number % size
filter[index] = true
*/
func mHash(key string, size int32) int32 {
	hasher.Write([]byte(key))
	index := hasher.Sum32() % uint32(size)
	hasher.Reset()
	return int32(index)
}

type BloomFilter struct {
    // TODO: optimize storage to 1 bit for each key
	filter []bool
	size   int32
}

func newBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		filter: make([]bool, size),
		size:   size,
	}
}

func (bf *BloomFilter) Add(key string) {
	index := mHash(key, bf.size)
	bf.filter[index] = true
}

func (bf *BloomFilter) Exists(key string) bool {
	index := mHash(key, bf.size)
	return bf.filter[index]
}

func main() {
	filter := newBloomFilter(16)

	keys := []string{"foo", "bar", "ez", "gg"}
	for _, key := range keys {
		filter.Add(key)
	}
    
    fmt.Println(filter.Exists("abc"))
    fmt.Println(filter.Exists("bar"))
    fmt.Println(filter.Exists("jj"))
}

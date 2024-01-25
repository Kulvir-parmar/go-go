package main

import (
	"fmt"
	"hash"
	"math/rand"

	"github.com/google/uuid"
	"github.com/twmb/murmur3"
)

var hasher []hash.Hash32

func init() {
	hasher = []hash.Hash32{
		murmur3.SeedNew32(rand.Uint32()),
		murmur3.SeedNew32(rand.Uint32()),
		murmur3.SeedNew32(rand.Uint32()),
	}
}

/*
Murmur Hash function is used to hash the values
https://en.wikipedia.org/wiki/MurmurHash

hash1(key) -> returns a number
hash2(key) -> returns a number
hash3(key) -> returns a number

insertion index = number % size
filter[index] = true
*/
func mHash(key string, size int32, hashIndex int) int32 {
	hasher[hashIndex].Write([]byte(key))
	index := hasher[hashIndex].Sum32() % uint32(size)
	defer hasher[hashIndex].Reset()

	return int32(index)
}

type BloomFilter struct {
	filter []uint8
	size   int32
}

func newBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		filter: make([]uint8, size),
		size:   size,
	}
}

func (bf *BloomFilter) PrintBloomFilter() {
	fmt.Println(bf.filter)
}

func (bf *BloomFilter) Add(key string, numHashes int) {
	for idx := 0; idx < numHashes; idx++ {
		hashVal := mHash(key, bf.size, idx)
		index := hashVal / 8
		bit := hashVal % 8
		bf.filter[index] |= (1 << bit)
	}
}

func (bf *BloomFilter) Exists(key string, numHashes int) bool {
	for idx := 0; idx < numHashes; idx++ {
		hashVal := mHash(key, bf.size, idx)
		index := hashVal / 8
		bit := hashVal % 8
		if bf.filter[index]&(1<<bit) == 0 {
			return false
		}
	}
	return true
}

func main() {
	filterSize := 1000

	dataset := make([]string, 0)
	exists := make(map[string]bool)

	for idx := 0; idx < filterSize; idx++ {
		id := uuid.New()
		dataset = append(dataset, id.String())
		exists[id.String()] = true
	}

	for idx := 0; idx < filterSize/2; idx++ {
		id := uuid.New()
		dataset = append(dataset, id.String())
	}

	filter := newBloomFilter(int32(filterSize))

	for i := 0; i < len(hasher); i++ {

		for id := range exists {
			filter.Add(id, i+1)
		}

		falsePositive := 0
		for _, id := range dataset {
			if filter.Exists(id, i) && !exists[id] {
				falsePositive++
			}
		}

		// NOTE: There are no false negatives in bloom filters
		// dataset = falsePositives + trueNegatives
		// Hence fpr = falsePositive / len(dataset)

		fpr := float64(falsePositive) / float64(len(dataset))
		fmt.Printf("Hasher %d: False Positive Rate: %f\n", i, fpr)
	}
}

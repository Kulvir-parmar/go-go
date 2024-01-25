package main

import (
	"fmt"

	"github.com/twmb/murmur3"
    "github.com/google/uuid"
)

var hasher = murmur3.SeedNew32(10)

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

func (bf *BloomFilter) Add(key string) {
	hashVal := mHash(key, bf.size)
    index := hashVal / 8
    bit := hashVal % 8
    bf.filter[index] |= (1 << bit)
}

func (bf *BloomFilter) Exists(key string) bool {
	hashVal := mHash(key, bf.size)
    index := hashVal / 8
    bit := hashVal % 8

    return bf.filter[index] & (1 << bit) != 0
}

func main() {
    filterSize := 1000

    dataset := make([]string, 0)
    existMap := make(map[string]bool)
    notExistMap := make(map[string]bool)

    for idx := 0; idx < filterSize/2; idx++{
        id := uuid.New()
        dataset = append(dataset, id.String())
        existMap[id.String()] = true
    }

    for idx := 0; idx < filterSize/2; idx++{
        id := uuid.New()
        dataset = append(dataset, id.String())
        notExistMap[id.String()] = true
    }

    for j := 500; j < 2000; j+=100 {
        filter := newBloomFilter(int32(j))

        for id := range existMap {
            filter.Add(id)
        }

        falsePositive := 0
        for _, id := range dataset {
            if filter.Exists(id) && notExistMap[id] {
                falsePositive++
            }
        }

        fmt.Printf("False Positive: %f\n", float64(falsePositive)/float64(len(dataset)))
    }
}

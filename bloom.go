package bloom

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bitset []bool
	size   uint
	hashes int
}

func NewBloomFilter(size uint, hashes int) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		size:   size,
		hashes: hashes,
	}
}

func (bf *BloomFilter) Insert(item interface{}) {
	for i := 0; i < bf.hashes; i++ {
		index := bf.hash(item, i) % bf.size
		bf.bitset[index] = true
	}
}

func (bf *BloomFilter) Find(item interface{}) bool {
	for i := 0; i < bf.hashes; i++ {
		index := bf.hash(item, i) % bf.size
		if !bf.bitset[index] {
			return false
		}
	}
	return true
}
func (bf *BloomFilter) hash(item interface{}, seed int) uint {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%d%v", seed, item)))
	return uint(h.Sum32())
}

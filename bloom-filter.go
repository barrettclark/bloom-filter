package main

import (
	"hash/fnv"
)

type HashFunction func(str string) uint32

type BloomFilter struct {
	HashFunctions []HashFunction
	ByteArray     []byte
}

func HashSum(str string) uint32 {
	var hashSum uint32
	for _, character := range str {
		hashSum += uint32(character)
	}
	return hashSum
}

func HashProduct(str string) uint32 {
	var hashProduct uint32 = 1
	for _, character := range str {
		hashProduct *= uint32(character)
	}
	return hashProduct
}

func HashHash(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

func (bf BloomFilter) Add(str string) {
	for _, hf := range bf.HashFunctions {
		hashValue := hf(str) % uint32(len(bf.ByteArray)*8)
		bucket := hashValue / 8
		bitPosition := hashValue % 8
		// fmt.Println(bucket, bitPosition)
		mask := 1 << bitPosition
		bf.ByteArray[bucket] = (bf.ByteArray[bucket] | byte(mask))
	}
}

func (bf BloomFilter) Contains(str string) bool {
	for _, hf := range bf.HashFunctions {
		hashValue := hf(str) % uint32(len(bf.ByteArray)*8)
		bucket := hashValue / 8
		bitPosition := hashValue % 8
		mask := 1 << bitPosition
		if bf.ByteArray[bucket]&byte(mask) == 0 {
			return false
		}
	}
	return true
}

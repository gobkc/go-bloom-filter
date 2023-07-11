package gobloom

import (
	"fmt"
	"hash/fnv"
)

type Bloom struct {
	hashes []int64
	maxBit int64
	bits   []bool
	count  int
}

func NewBloom() *Bloom {
	return &Bloom{}
}

func (b *Bloom) Add(dest any) {
	hash64 := b.Hash(dest)
	doubleHash := b.Hash(fmt.Sprintf(`%v%v`, dest, dest))
	b.hashes = append(b.hashes, hash64, doubleHash)
	b.maxBit += 4
}

func (b *Bloom) Delete(dest any) {
	hash64 := b.Hash(dest)
	doubleHash := b.Hash(fmt.Sprintf(`%v%v`, dest, dest))
	newHashesLen := len(b.hashes)
	if newHashesLen >= 2 {
		newHashesLen -= 2
	}
	newHashes := make([]int64, newHashesLen)
	for _, hash := range b.hashes {
		if hash != hash64 && hash != doubleHash {
			newHashes = append(newHashes, hash)
		}
	}
	b.hashes = newHashes
	b.maxBit -= 4
}

func (b *Bloom) Has(dest any) bool {
	if hashLen := len(b.hashes); hashLen > b.count {
		b.maxBit *= 10
		b.bits = make([]bool, b.maxBit)
		for _, hash := range b.hashes {
			index := hash % b.maxBit
			b.bits[index] = true
		}
		b.count = hashLen
	}
	// hash function 1
	hash := b.Hash(dest) % b.maxBit
	if !b.bits[hash] {
		return false
	}
	// hash function 2
	hash = b.Hash(fmt.Sprintf(`%v%v`, dest, dest)) % b.maxBit
	if !b.bits[hash] {
		return false
	}
	return true
}

// Hash the Hash function converts "dest" to a string and returns a hash int value
func (b *Bloom) Hash(dest any) int64 {
	s := fmt.Sprintf(`%v`, dest)
	h := fnv.New32a()
	h.Write([]byte(s))
	return int64(h.Sum32())
}

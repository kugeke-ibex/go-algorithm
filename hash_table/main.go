package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/binary"
)

func main() {
	h := HashTable{}
	h.init(10)
	h.add("car", "Telsa")
	h.add("car", "Toyota")
	h.add("pc", "Mac")
	h.add("sns", "Youtube")
	h.add("cooking", "soup")
	h.print()
	value, ok := h.get("car")
	fmt.Printf("key: car, value: %s, ok: %t\n", value, ok)
	value, ok = h.get("game")
	fmt.Printf("key: game, value: %s, ok: %t\n", value, ok)
}

type KeyValuePair struct {
	key string
	value string
}

type HashTable struct {
	size int
	table [][]*KeyValuePair
}

func (h *HashTable) append(index int, pair *KeyValuePair) {
	h.table[index] = append(h.table[index], pair)
}

func (h *HashTable) init(size int) {
	h.size = size
	h.table = make([][]*KeyValuePair, size)
	for i := range h.table {
		h.table[i] = []*KeyValuePair{}
	}
}

func (h *HashTable) hash(key string) int {
	hash := sha256.Sum256([]byte(key))
	hashInt := binary.BigEndian.Uint64(hash[:])
	return int(hashInt % uint64(h.size))
}

func (h *HashTable) add(key, value string) {
	index := h.hash(key)

	for _, pair := range h.table[index] {
		if pair.key == key {
			pair.value = value
			return
		}
	}

	h.append(index, &KeyValuePair{key: key, value: value})
}

func (h HashTable) print() {
	for index, pairs := range h.table {
		fmt.Print(index)
		for _, pair := range pairs {
			fmt.Printf(" --> %+v", pair)
		}
		fmt.Println()
	}
}

func (h HashTable) get(key string) (string, bool) {
	index := h.hash(key)
	for _, pair := range h.table[index] {
		if pair.key == key {
			return pair.value, true
		}
	}

	return "", false
}

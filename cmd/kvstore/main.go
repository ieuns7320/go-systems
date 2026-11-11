package main

import (
	"fmt"

	"github.com/ieuns7320/go-systems/internal/kvstore"
)

func main() {
	scenarioA()
	scenarioB()
	scenarioC()
}

func scenarioA() {
	store := kvstore.New()

	buf := []byte("hello")
	store.Set("k1", buf)
	buf[0] = 'j'

	value, _ := store.Get("k1")
	fmt.Println(string(value))
}

func scenarioB() {
	store := kvstore.New()
	store.Set("k2", []byte("world"))
	value, _ := store.Get("k2")
	value[0] = 'W'

	value, _ = store.Get("k2")

	fmt.Println(string(value))

}

func scenarioC() {
	store := kvstore.New()

	store.Set("carol", []byte("Nice to meet you"))
	store.Set("alice", []byte("My name is alice"))
	store.Set("bob", []byte("Hello, I'm bob"))

	sorted := store.Keys()

	fmt.Println(sorted)

}

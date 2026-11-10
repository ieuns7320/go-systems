package main

import (
	"errors"
	"fmt"

	"github.com/ieuns7320/go-systems/internal/kvstore"
)

func main() {
	store := kvstore.New()
	store.Set("alice", []byte("hello"))

	value, _ := store.Get("alice")
	_, err := store.Get("missing")

	fmt.Println(string(value))
	fmt.Println(err)
	fmt.Println(errors.Is(err, kvstore.ErrNotFound))
}

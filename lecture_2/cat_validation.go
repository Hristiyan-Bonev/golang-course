package main

import (
	"fmt"
	"log"
)

type Cat struct {
	Name string
}

func validName(name string) error {
	if name != "Jessie" {
		return fmt.Errorf("name is not jessie")
	}
	return nil
}

func NewCat(name string) (*Cat, error) {
	if err := validName(name); err != nil {
		return nil, err
	}
	return &Cat{name}, nil
}

func main() {
	cat, _ := NewCat("Jessie")
	log.Println(cat)
}

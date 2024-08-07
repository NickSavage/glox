package interpreter

import (
	"fmt"
	"log"
)

func (i *Interpreter) Put(key string, value interface{}) error {
	log.Printf("put %v %v", key, value)
	i.Memory.Memory[key] = value
	return nil
}

func (i *Interpreter) Get(key string) (interface{}, error) {
	log.Printf("memory %v", i.Memory.Memory)
	value, ok := i.Memory.Memory[key]
	if !ok {
		return nil, fmt.Errorf("undefined variable: %s", key)
	}
	return value, nil
}

package interpreter

import (
	"fmt"
)

func (i *Interpreter) Assign(key string, value interface{}) error {
	_, ok := i.Memory.Memory[key]
	if !ok {
		return fmt.Errorf("undefined variable: %s", key)
	}
	return i.Put(key, value)
}

func (i *Interpreter) Define(key string, value interface{}) error {
	return i.Put(key, value)

}

func (i *Interpreter) Put(key string, value interface{}) error {
	i.Memory.Memory[key] = value
	return nil
}

func (i *Interpreter) Get(key string) (interface{}, error) {
	value, ok := i.Memory.Memory[key]
	if !ok {
		return nil, fmt.Errorf("undefined variable: %s", key)
	}
	return value, nil
}

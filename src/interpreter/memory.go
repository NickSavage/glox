package interpreter

func (i *Interpreter) Put(key string, value interface{}) error {
	i.Memory[key] = value
	return nil
}

func (i *Interpreter) Get(key string) (interface{}, error) {
	return i.Memory[key], nil
}

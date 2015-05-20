package vm

import (
	"reflect"
)

type Env struct {
	Env map[string]reflect.Value
}

func NewEnv() *Env {
	return &Env{Env: make(map[string]reflect.Value)}
}

func (e *Env) Set(k string, v reflect.Value) error {
	e.Env[k] = v
	return nil
}

func (e *Env) Get(k string) (reflect.Value, error) {
	value := e.Env[k]
	return value, nil
}

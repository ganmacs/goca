package vm

import (
	"reflect"
)

type Env struct {
	env map[string]reflect.Value
}

func NewEnv() *Env {
	return &Env{}
}

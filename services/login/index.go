
package main

import "github.com/google/uuid"

type Index struct {
	Root string
	IndexName string
}

func (i Index) New() (Object) {
	return Object{ i, uuid.NewString() }
}

package storage

import "diplom/pkg/data"

type Storage map[string]*data.Codes

func New() Storage {
	return make(map[string]*data.Codes
}

package storage

import "diplom/pkg/data"

type Storage map[string]*data.Iso3166

func New() Storage {
	return make(map[string]*data.Iso3166)
}

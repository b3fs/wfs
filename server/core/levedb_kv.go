package core

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type leveldbKV struct {
	db *leveldb.DB
}

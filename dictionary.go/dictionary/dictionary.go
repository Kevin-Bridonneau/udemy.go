package dictionary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

// Dictionary type is a link for access to badger db
type Dictionary struct {
	db *badger.DB
}

// Entry type is like a schema for entry badger db
type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

// New is for create the db
func New(dir string) (*Dictionary, error) {
	opts := badger.DefaultOptions("./")
	opts.Dir = dir
	opts.ValueDir = dir

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	dict := &Dictionary{
		db: db,
	}
	return dict, nil
}

// Close the db
func (d *Dictionary) Close() {
	d.db.Close()
}

package way

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mattn/go-sqlite3"
)

type Explorer struct {
	path string
	db   *sql.DB
}

func (e Explorer) OpenBlockChain(path string) (*Explorer, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS blockchain(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		prevhash BLOB NOT NULL,
		hash BLOB NOT NULL,
		date TEXT NOT NULL UNIQUE,
		data BLOB NOT NULL
);
	CREATE INDEX IF NOT EXISTS idx_date ON blockchain(date);
	`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}

	return &Explorer{path: path, db: db}, nil
}

func (e Explorer) SaveBlock(block Block, time_utc time.Time) (id int64, err error) {
	stmt, err := e.db.Prepare("INSERT INTO blockchain(prevhash, hash, date, data) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(block.PrevHash, block.Hash, time_utc, block.Data)
	if err != nil {
		if sqlLiteErr, ok := err.(sqlite3.Error); ok && sqlLiteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, errors.New("this block is exist")
		}
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

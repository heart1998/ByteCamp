package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

type DB struct {
	*postDB
	*topicDB
}

func initDB(filePath string, f func(db io.ReadWriteCloser) error) error {
	db, err := openFile(filePath)
	if err != nil {
		return err
	}
	if err := f(db); err != nil {
		return err
	}
	return nil
}

func InitByPath(basePath string) (*DB, error) {
	var err error
	var db = new(DB)
	if err := initDB(path.Join(basePath, "topicDB"), func(open io.ReadWriteCloser) error {
		db.topicDB, err = initTopicDB(open)
		return err
	}); err != nil {
		return nil, err
	}
	if err := initDB(path.Join(basePath, "postDB"), func(open io.ReadWriteCloser) error {
		db.postDB, err = initPostDB(open)
		return err
	}); err != nil {
		return nil, err
	}
	return db, nil
}

func (db *DB) Close() {
	db.topicDB.Close()
	db.postDB.Close()
}

func openFile(path string) (io.ReadWriteCloser, error) {
	open, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return open, nil
}

func dataToFile(data interface{}, file io.ReadWriteCloser) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintln(file, string(bs)); err != nil {
		return err
	}
	return nil
}

func fileToData[T any](file io.ReadWriteCloser, f func(data T)) error {
	var data T
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if err := json.Unmarshal([]byte(text), &data); err != nil {
			return err
		}
		f(data)
	}
	return nil
}

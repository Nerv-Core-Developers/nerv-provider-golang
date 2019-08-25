package database

import (
	"fmt"
	"log"
	"strings"

	badger "github.com/dgraph-io/badger"
	"github.com/naokichau/nerv-provider-golang/shared"
)

type DBpersistent struct {
	Name    string
	DataDir string
	DB      *badger.DB
	Logger  shared.LoggerInterface
}

func (db *DBpersistent) Init(dbname string, dbpath string) error {
	opts := badger.DefaultOptions(dbpath + "/" + dbname)
	db.Logger.Init("persistentDB")
	opts.Logger = &logAdapter{db.Logger}
	database, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	db.DB = database
	db.Name = dbname
	db.DataDir = dbpath + "/" + dbname
	return nil
}

func (db *DBpersistent) Close() error {
	return db.DB.Close()
}

func (db *DBpersistent) isOpen() bool {
	if db.DB != nil {
		return true
	}
	return false
}

func (db *DBpersistent) Export(path string) error {
	return nil
}
func (db *DBpersistent) Put(key string, value []byte, overwrite bool) error {
	return nil
}
func (db *DBpersistent) Batch([]shared.DBPatch) error {
	return nil
}
func (db *DBpersistent) Get(key string) ([]byte, error) {
	return nil, nil
}

func (db *DBpersistent) Delete(key string) error {
	return nil
}

type logAdapter struct {
	Logger shared.LoggerInterface
}

func (l *logAdapter) Errorf(f string, v ...interface{}) {
	l.Logger.Debug(strings.TrimSuffix(fmt.Sprintf("ERROR: "+f, v...), "\n"))
}
func (l *logAdapter) Warningf(f string, v ...interface{}) {
	l.Logger.Debug(strings.TrimSuffix(fmt.Sprintf("WARNING: "+f, v...), "\n"))
}
func (l *logAdapter) Infof(f string, v ...interface{}) {
	l.Logger.Log(strings.TrimSuffix(fmt.Sprintf("INFO: "+f, v...), "\n"))
}
func (l *logAdapter) Debugf(f string, v ...interface{}) {
	l.Logger.Debug(strings.TrimSuffix(fmt.Sprintf("DEBUG: "+f, v...), "\n"))
}

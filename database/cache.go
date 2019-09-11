package database

import (
	"errors"
	"time"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
	cache "github.com/patrickmn/go-cache"
)

type DBcache struct {
	Name          string
	DataDir       string
	ExpTime       time.Duration
	CleanInterval time.Duration
	DB            *cache.Cache
	Logger        shared.LoggerInterface
}

func (db *DBcache) Delete(key string) error {
	if db.isOpen() {
		return errors.New("init db first")
	}
	db.DB.Delete(key)
	return nil
}

func (db *DBcache) Get(key string) (interface{}, int, error) {
	if db.isOpen() {
		return nil, 0, errors.New("init db first")
	}
	if value, ttl, exist := db.DB.GetWithExpiration(key); exist {
		return value, ttl.Minute(), nil
	}
	return nil, 0, errors.New("key not exist")
}

func (db *DBcache) Put(key string, value interface{}, timeout int, overwrite bool) error {
	if db.isOpen() {
		return errors.New("init db first")
	}

	if overwrite {
		if _, ok := db.DB.Get(key); ok {
			return errors.New("key already exist")
		}
		db.DB.Set(key, value, time.Duration(timeout)*time.Second)
	} else {
		db.DB.Set(key, value, time.Duration(timeout)*time.Second)
	}
	return nil
}

func (db *DBcache) Export(des string) error {
	if db.isOpen() {
		return errors.New("init db first")
	}
	return nil
}

func (db *DBcache) Load(src string) error {
	if db.DB == nil {
		return errors.New("init db first")
	}
	return nil
}

func (db *DBcache) Init(dbname string, dbpath string) error {
	db.Name = dbname
	db.DataDir = dbpath + "/" + dbname
	db.DB = cache.New(db.ExpTime, db.CleanInterval)
	return nil
}

func (db *DBcache) Close() error {
	db.DB = nil
	return nil
}

func (db *DBcache) isOpen() bool {
	if db.DB != nil {
		return true
	}
	return false
}

func (db *DBcache) Batch([]shared.DBPatch) error {
	return nil
}

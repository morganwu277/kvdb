package db

import (
	"fmt"
	"log"
	"sync"

	"errors"

	"github.com/DataDog/leveldb"
)

type KVDB struct {
	sync.Mutex
	_db       *levigo.DB
	_rOptions *levigo.ReadOptions
	_wOptions *levigo.WriteOptions
}

func Init(path string) (*KVDB, error) {
	ver := fmt.Sprintf("%v.%v", levigo.GetLevelDBMajorVersion(), levigo.GetLevelDBMinorVersion())
	log.Println(ver)
	env := levigo.NewDefaultEnv()
	cache := levigo.NewLRUCache(16 * (1 << 20)) // 16MB read cache
	options := levigo.NewOptions()
	options.SetCache(cache)
	options.SetWriteBufferSize(16 * (1 << 20)) // write buffer 16MB
	options.SetBlockSize(4 * (1 << 10))        // 1KB
	options.SetMaxOpenFiles(256)
	options.SetBlockRestartInterval(8)

	// if DB already exist, doesn't throw error
	options.SetErrorIfExists(true)
	options.SetEnv(env)
	options.SetInfoLog(nil)
	options.SetParanoidChecks(true)

	roptions := levigo.NewReadOptions()
	roptions.SetVerifyChecksums(true)
	roptions.SetFillCache(false)

	woptions := levigo.NewWriteOptions()
	woptions.SetSync(true)

	// _ = levigo.DestroyDatabase(path, options)
	log.Println("Open DB PATH:", path)

	options.SetCreateIfMissing(true)
	options.SetErrorIfExists(false)

	_db, err := levigo.Open(path, options)
	if err != nil {
		log.Printf("Open failed: %v\n", err)
		return nil, err
	}
	kvdb := &KVDB{
		_db:       _db,
		_rOptions: roptions,
		_wOptions: woptions,
	}
	return kvdb, nil
}

func (kvdb *KVDB) Read(k string) (string, error) {
	if k == "" {
		err := errors.New("Key is empty! ")
		log.Printf("Get failed, k: %v,  err: %v\n", k, err)
		return "", err
	}
	kvdb.Lock()
	defer kvdb.Unlock()
	getValue, err := kvdb._db.Get(kvdb._rOptions, []byte(k))
	if err != nil {
		log.Printf("Get failed, k: %v, err: %v\n", k, err)
		return "", err
	}
	if getValue == nil {
		err := errors.New("Key doesn't exist in DB! ")
		log.Printf("Get failed, k: %v doesn't exist in DB. \n", k)
		return "", err
	}
	return string(getValue), nil
}

func (kvdb *KVDB) Write(k string, v string) error {
	if k == "" {
		err := errors.New("Key is empty! ")
		log.Printf("Get failed, k: %v,  err: %v\n", k, err)
		return err
	}
	putKey := []byte(k)
	putValue := []byte(v)
	kvdb.Lock()
	defer kvdb.Unlock()
	err := kvdb._db.Put(kvdb._wOptions, putKey, putValue)
	if err != nil {
		log.Printf("Put failed: %v\n", err)
		return err
	}
	return nil
}

func (kvdb *KVDB) Close() {
	kvdb._db.Close()
}

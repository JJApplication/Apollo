package kv

import (
	"strconv"
	"time"

	"github.com/JJApplication/Apollo/config"
	"go.etcd.io/bbolt"
)

var (
	KV *bbolt.DB
)

func NewKV() error {
	kvPath := config.ApolloConf.DB.KV
	if kvPath == "" {
		kvPath = "kv.db"
	}
	db, err := bbolt.Open(kvPath, 0600, nil)
	if err != nil {
		return err
	}
	KV = db
	return initBucket(KV)
}

// 初始化bucket
func initBucket(kv *bbolt.DB) error {
	var err error
	err = kv.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("internal"))
		return err
	})
	if err != nil {
		return err
	}

	err = kv.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("CPU"))
		return err
	})
	if err != nil {
		return err
	}

	err = kv.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("MEMORY"))
		return err
	})
	if err != nil {
		return err
	}

	err = kv.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("IO"))
		return err
	})
	if err != nil {
		return err
	}

	err = kv.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("NETWORK"))
		return err
	})
	if err != nil {
		return err
	}

	err = kv.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("LOAD"))
		return err
	})
	if err != nil {
		return err
	}

	err = kv.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("TOP10"))
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

// AddQuick 所有时序数据使用时间戳作为键
func AddQuick(kv *bbolt.DB, bucket string, value []byte) error {
	PreHookAdd(kv, bucket)
	key := strconv.Itoa(int(time.Now().Unix()))
	return kv.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put([]byte(key), value)
	})
}

// Add 所有时序数据使用时间戳作为键
func Add(kv *bbolt.DB, bucket, key string, value []byte) error {
	PreHookAdd(kv, bucket)
	return kv.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put([]byte(key), value)
	})
}

// PreHookAdd 插入前hook 每个bucket的容量为100，超过后从头部开始进行清理
func PreHookAdd(kv *bbolt.DB, bucket string) error {
	return kv.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		c := b.Cursor()
		// 是否超过了100个
		keys := b.Stats().KeyN
		if keys >= 100 {
			key, _ := c.First()
			return b.Delete(key)
		}
		return nil
	})
}

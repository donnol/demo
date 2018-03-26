package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("demo.db", 0700, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Printf("%+v\n", db)

	// 插入或更新
	bname := []byte("jd")
	if err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket(bname)
		if err != nil {
			if err == bolt.ErrBucketExists {
				b = tx.Bucket(bname)
			} else {
				return err
			}
		}
		if err := b.Put([]byte("foo"), []byte("bar")); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// 获取
	if err = db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(bname).Get([]byte("foo"))
		fmt.Printf("The value of 'foo' is: %s\n", value)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// 删除
	if err = db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(bname)
	}); err != nil {
		log.Fatal(err)
	}

	a := ^uint8(0) // 异或自身，相当于取反码；^uint8(0)+1 = 补码 = 反码 + 1
	log.Println(a)
}

package boltdb

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb/helpers"
	"github.com/rs/zerolog/log"
	bolt "go.etcd.io/bbolt"
)

func Create(dbPath string, data map[string]interface{}) (err error) {
	log.Debug().
		Str("dbPath", dbPath).
		Msg("create database")

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) (err error) {
		for bucketName, bucketData := range data {
			b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return
			}

			log.Debug().Msgf("bucket name: %s", bucketName)

			if bucketData == nil {
				log.Debug().Msg("bucket is null")
				continue
			}

			for key, value := range bucketData.(map[string]interface{}) {
				v, err := helpers.MarshalObject(value)
				if err != nil {
					return
				}

				err = b.Put([]byte(key), v)
				if err != nil {
					return
				}
			}
		}
		return
	})

	return
}

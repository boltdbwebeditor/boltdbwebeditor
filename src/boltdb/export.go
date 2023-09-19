package boltdb

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	bolt "go.etcd.io/bbolt"
)

func ExportJSON(dbPath string, data map[string]interface{}) error {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		for bucketName, bucketData := range data {
			b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return err
			}

			log.Debug().Msgf("bucket name: %s", bucketName)

			if bucketData == nil {
				log.Debug().Msg("bucket is null")
				continue
			}

			for key, value := range bucketData.(map[string]interface{}) {
				v, err := MarshalObject(value)
				if err != nil {
					return err
				}

				err = b.Put([]byte(key), v)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	return err
}

func MarshalObject(data interface{}) ([]byte, error) {
	if data == nil {
		return []byte(""), nil
	}

	return json.Marshal(data)
}

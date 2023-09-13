package boltdb

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/rs/zerolog/log"
	bolt "go.etcd.io/bbolt"
)

func ExportJSON(dbPath string, data map[string]interface{}, metadata bool) error {

	err := os.Rename(dbPath, dbPath+".bak")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return errors.New("failed to rename db file")
	}

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

	if err != nil {
		// restore backup
		log.Info().Err(err).Msg("failed to update db, restoring backup")
		os.Rename(dbPath+".bak", dbPath)
	}

	err = os.Remove(dbPath + ".bak")
	if err != nil {
		log.Info().Err(err).Msg("failed to remove backup")
	}

	return nil
}

func MarshalObject(data interface{}) ([]byte, error) {
	if data == nil {
		return []byte(""), nil
	}

	ret, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return ret, nil

}

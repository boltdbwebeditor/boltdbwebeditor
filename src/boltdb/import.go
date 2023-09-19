package boltdb

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"

	"github.com/rs/zerolog/log"
	bolt "go.etcd.io/bbolt"
)

func backupMetadata(connection *bolt.DB) (map[string]interface{}, error) {
	buckets := map[string]interface{}{}

	err := connection.View(func(tx *bolt.Tx) error {
		err := tx.ForEach(func(name []byte, bucket *bolt.Bucket) error {
			bucketName := string(name)
			seqId := bucket.Sequence()
			buckets[bucketName] = int(seqId)
			return nil
		})

		return err
	})

	return buckets, err
}

func ImportJSON(dbPath string, metadata bool) (backup map[string]interface{}, err error) {
	log.Debug().
		Str("dbPath", dbPath).
		Msg("importJson")

	backup = make(map[string]interface{})

	option := &bolt.Options{
		Timeout:  1 * time.Second,
		ReadOnly: true,
	}

	conn, err := bolt.Open(dbPath, 0600, option)
	if err != nil {
		return backup, err
	}
	defer conn.Close()

	if metadata {
		meta, err := backupMetadata(conn)
		if err != nil {
			log.Error().Err(err).Msg("failed exporting metadata")
		}

		backup["__metadata"] = meta
	}

	err = conn.View(func(tx *bolt.Tx) error {
		err = tx.ForEach(func(name []byte, bucket *bolt.Bucket) error {
			bucketName := string(name)
			cursor := bucket.Cursor()

			length := bucket.Stats().KeyN
			list := make(map[string]interface{}, length)

			for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
				if v == nil {
					continue
				}

				var obj interface{}
				err := UnmarshalObject(v, &obj)
				if err != nil {
					log.Error().
						Str("bucket", bucketName).
						Str("object", string(v)).
						Err(err).
						Msg("failed to unmarshal")

					obj = v
				}

				list[string(k)] = obj
			}

			backup[bucketName] = list
			return nil
		})

		return err
	})

	return backup, err
}

func UnmarshalObject(data []byte, object interface{}) error {
	var err error

	e := json.Unmarshal(data, object)
	if e != nil {
		// Special case for the VERSION bucket. Here we're not using json
		// So we need to return it as a string
		s, ok := object.(*string)
		if !ok {
			return errors.Wrap(err, e.Error())
		}

		*s = string(data)
	}
	return err
}

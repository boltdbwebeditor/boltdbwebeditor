package boltdb

import (
	helpers "github.com/boltdbwebeditor/boltdbwebeditor/src/libs/boltdb/helpers"
	"time"

	"github.com/rs/zerolog/log"
	bolt "go.etcd.io/bbolt"
)

func Read(dbPath string, metadata bool) (data map[string]interface{}, err error) {
	log.Debug().
		Str("dbPath", dbPath).
		Msg("read database")

	data = make(map[string]interface{})

	option := &bolt.Options{
		Timeout:  1 * time.Second,
		ReadOnly: true,
	}

	conn, err := bolt.Open(dbPath, 0600, option)
	if err != nil {
		return
	}
	defer conn.Close()

	if metadata {
		data[helpers.MetadataKey], err = helpers.ReadMetadata(conn)
		if err != nil {
			return
		}
	}

	err = conn.View(func(tx *bolt.Tx) (err error) {
		err = tx.ForEach(func(name []byte, bucket *bolt.Bucket) (err error) {
			bucketName := string(name)
			cursor := bucket.Cursor()

			length := bucket.Stats().KeyN
			list := make(map[string]interface{}, length)

			for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
				if v == nil {
					continue
				}

				var obj interface{}
				err := helpers.UnmarshalObject(v, &obj)
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

			data[bucketName] = list
			return nil
		})

		return
	})

	return
}

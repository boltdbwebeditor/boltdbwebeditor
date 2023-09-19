package helpers

import "go.etcd.io/bbolt"

func ExportMetadata(connection *bbolt.DB) (map[string]interface{}, error) {
	buckets := map[string]interface{}{}

	err := connection.View(func(tx *bbolt.Tx) error {
		err := tx.ForEach(func(name []byte, bucket *bbolt.Bucket) error {
			bucketName := string(name)
			seqId := bucket.Sequence()
			buckets[bucketName] = int(seqId)
			return nil
		})

		return err
	})

	return buckets, err
}

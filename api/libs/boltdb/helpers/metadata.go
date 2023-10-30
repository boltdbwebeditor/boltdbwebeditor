package helpers

import "go.etcd.io/bbolt"

func ReadMetadata(connection *bbolt.DB) (metadata map[string]interface{}, err error) {
	metadata = make(map[string]interface{})
	err = connection.View(func(tx *bbolt.Tx) (err error) {
		err = tx.ForEach(func(name []byte, bucket *bbolt.Bucket) (err error) {
			bucketName := string(name)
			seqId := bucket.Sequence()
			metadata[bucketName] = uint64(seqId)
			return
		})
		return
	})
	return
}

func WriteMetadata(connection *bbolt.DB, metadata map[string]interface{}) (err error) {
	err = connection.Update(func(tx *bbolt.Tx) (err error) {
		err = tx.ForEach(func(name []byte, bucket *bbolt.Bucket) (err error) {
			bucketName := string(name)
			value, ok := metadata[bucketName]
			if ok {
				seqId, ok := value.(float64)
				if ok {
					bucket.SetSequence(uint64(seqId))
				}
			}
			return
		})
		return
	})
	return
}

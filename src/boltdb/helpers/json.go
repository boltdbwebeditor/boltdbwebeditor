package helpers

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func MarshalObject(data interface{}) ([]byte, error) {
	if data == nil {
		return []byte(""), nil
	}

	return json.Marshal(data)
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

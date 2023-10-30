package boltdb

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/api/helpers"
)

func ForceRead(dbPath string, metadata bool) (data map[string]interface{}, err error) {
	tempDbPath, err := helpers.CopyDbToTemp(dbPath)
	if err != nil {
		return
	}

	data, err = Read(tempDbPath, metadata)

	return
}

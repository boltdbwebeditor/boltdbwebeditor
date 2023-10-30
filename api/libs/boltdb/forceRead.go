package boltdb

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/api/libs/tempFile"
)

func ForceRead(dbPath string, metadata bool) (data map[string]interface{}, err error) {
	tempDbPath, err := tempFile.CopyDbToTemp(dbPath)
	if err != nil {
		return
	}

	data, err = Read(tempDbPath, metadata)

	return
}

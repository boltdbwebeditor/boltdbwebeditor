package helpers

import (
	"fmt"
	"time"
)

func GenerateDbTmpFilePath() string {
	return fmt.Sprintf("/tmp/bold.db.%d", time.Now().UnixMilli())
}

func CopyDbToTemp(dbPath string) (tempDbPath string, err error) {
	tempDbPath = GenerateDbTmpFilePath()

	return tempDbPath, CopyFile(dbPath, tempDbPath)
}

package utils

import "os"

func CreateFolderIfNotExists(path string) error {
	err := os.MkdirAll(path, 0750)
	if os.IsExist(err) {
		return nil
	}
    return err
}
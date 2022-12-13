package file

import "os"

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func NewFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewDir(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

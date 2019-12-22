package main

import (
	"common/random"
	"os"
)

func SaveFileToLocal(data []byte) (string, error) {
	path := "https://www.yycaptain.com/resources/images/" + random.GetRandomString(8) + ".jpg"
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		return "", err
	}

	_, err = f.WriteAt(data, 0)
	if err != nil {
		return "", err
	}
	return path, nil
}

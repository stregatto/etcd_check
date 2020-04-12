// Package utils contains all functions can be useful but not ETCD related
package utils

import (
	"io/ioutil"
	"log"
)

//GetFile returns a file content in []byte format form a given path
func GetFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

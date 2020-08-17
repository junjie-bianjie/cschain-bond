package utils

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// GetFromUrl return the responseBody from input url
func GetFromUrl(url string) ([]byte, error) {
	// get resource of url
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

// RandStringOfLength return a random string
func RandStringOfLength(l int) string {
	var result []byte
	bytes := []byte(charset)
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetScheme() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "/scripts/schema.json")
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bz)
}

func GetPrivKeyArmor() []byte {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "/scripts/priv.key")
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bz
}

func GetOriginTest() []byte {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "/scripts/origin.text")
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bz
}

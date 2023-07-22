package textfilekv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type KeyValueStore struct {
	filePath string
	store    map[string]string
}

func NewKeyValueStore(filePath string) *KeyValueStore {
	kvs := &KeyValueStore{
		filePath: filePath,
		store:    make(map[string]string),
	}
	kvs.loadFromFile()
	return kvs
}

func (kvs *KeyValueStore) Set(key, value string) {
	kvs.store[key] = value
	kvs.saveToFile()
}

func (kvs *KeyValueStore) Get(key string) (string, bool) {
	value, ok := kvs.store[key]
	return value, ok
}

func (kvs *KeyValueStore) saveToFile() {
	file, err := os.OpenFile(kvs.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for key, value := range kvs.store {
		fmt.Fprintf(writer, "%s=%s\n", key, value)
	}

	if err := writer.Flush(); err != nil {
		log.Fatal(err)
	}
}

func (kvs *KeyValueStore) loadFromFile() {
	file, err := os.Open(kvs.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			kvs.store[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

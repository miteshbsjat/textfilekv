package textfilekv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type KeyValueStore struct {
	filePath string
	store    map[string]string
	mu       sync.RWMutex
}

func NewKeyValueStore(filePath string) (*KeyValueStore, error) {
	kvs := &KeyValueStore{
		filePath: filePath,
		store:    make(map[string]string),
	}
	// Load existing data from file
	if err := kvs.loadFromFile(); err != nil {
		log.Printf("Failed load file %s %v\n", filePath, err)
		return nil, err
	}

	return kvs, nil

}

func (kvs *KeyValueStore) Set(key, value string) error {
	kvs.mu.RLock()
	defer kvs.mu.RUnlock()

	kvs.store[key] = value
	return kvs.saveToFile()
}

func (kvs *KeyValueStore) Get(key string) (string, bool) {
	kvs.mu.RLock()
	defer kvs.mu.RUnlock()

	value, ok := kvs.store[key]
	return value, ok
}

func (kvs *KeyValueStore) Keys() []string {
	kvs.mu.RLock()
	defer kvs.mu.RUnlock()

	keys := make([]string, 0, len(kvs.store))
	for key := range kvs.store {
		keys = append(keys, key)
	}
	return keys
}

func (kvs *KeyValueStore) Delete(key string) error {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	delete(kvs.store, key)
	return kvs.saveToFile()
}

func (kvs *KeyValueStore) saveToFile() error {
	file, err := os.OpenFile(kvs.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("Error in saving to file %s %v\n", kvs.filePath, err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for key, value := range kvs.store {
		fmt.Fprintf(writer, "%s=%s\n", key, value)
	}

	if err := writer.Flush(); err != nil {
		log.Printf("Failed to write flush %v\n", err)
		return err
	}
	return nil
}

func (kvs *KeyValueStore) loadFromFile() error {
	file, err := os.Open(kvs.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		log.Printf("Failed to open textfilekv file %v\n", err)
		return err
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
		log.Printf("Failed to scanner Err %v\n", err)
		return err
	}
	return nil
}

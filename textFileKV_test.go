package textfilekv

import "testing"

func TestTextFileKV(t *testing.T) {
	filePath := "/tmp/data.txt"
	kvs, _ := NewKeyValueStore(filePath)

	kvs.Set("name", "Mitesh")
	kvs.Set("age", "41")
	kvs.Set("data", "{\"height\": 167}")

	name, _ := kvs.Get("name")
	if name != "Mitesh" {
		t.Errorf("Name: got %s instead of Mitesh", name)
	}

	age, _ := kvs.Get("age")
	if age != "41" {
		t.Errorf("Age: got %s instead of 41", age)
	}
}

func TestTextFileKVRead(t *testing.T) {
	filePath := "/tmp/data.txt"
	kvs, _ := NewKeyValueStore(filePath)

	name, _ := kvs.Get("name")
	if name != "Mitesh" {
		t.Errorf("Name: got %s instead of Mitesh", name)
	}

	age, _ := kvs.Get("age")
	if age != "41" {
		t.Errorf("Age: got %s instead of 41", age)
	}

	// delete check
	kvs.Delete("age")
	age, exist := kvs.Get("age")
	if exist != false {
		t.Errorf("Age: %s should not be present after Delete", age)
	}
}

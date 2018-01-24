package archive

import "testing"

func TestSimplePutAndGet(t *testing.T) {
	expectedResult := "value"

	archive := createArchive(10, 10, 1, 1, t)
	archive.Put(Entry{Key: "key", Value: "value"})
	actualResult, _ := archive.Get("key")
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestSimpleEvict(t *testing.T) {
	archive := createArchive(10, 10, 1, 1, t)
	archive.Put(Entry{Key: "key1", Value: "value1"})
	archive.Put(Entry{Key: "key2", Value: "value2"})
	_, in1 := archive.Get("key1")
	v2, _ := archive.Get("key2")

	if in1 || v2 != "value2" {
		t.Fatalf("Eviction poliy error")
	}
}

func TestUpdate(t *testing.T) {
	expectedResult := "updated"
	archive := createArchive(10, 10, 1, 1, t)
	archive.Put(Entry{Key: "key", Value: "original"})
	archive.Put(Entry{Key: "key", Value: expectedResult})
	actualResult, _ := archive.Get("key")
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestDelete(t *testing.T) {
	archive := createArchive(10, 10, 1, 1, t)
	archive.Put(Entry{Key: "key", Value: "value"})
	archive.Delete("key")
	_, in := archive.Get("key")
	if in {
		t.Fatalf("KVP not deleted from archive")
	}
}

func TestMaxKeyValueLength(t *testing.T) {
	archive := createArchive(1, 1, 1, 1, t)
	err := archive.Put(Entry{Key: "key", Value: "Value"})
	if err == nil {
		t.Fatalf("Key/value exceed maximum length, but no error was outputted.")
	}
}

func createArchive(mkl int, mvl int, mkvps int, policy int, t *testing.T) Archive {
	archive, err := NewArchive(mkl, mvl, mkvps, policy)
	if err != nil {
		t.Fatalf("Problem with Archive creation.")
	}
	return archive
}

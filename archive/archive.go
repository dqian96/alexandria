package archive

import (
	"errors"
	"sync"
)

// An Archive is interacted similarly to a map.
type Archive interface {
	Get(string) (string, bool)
	Put(Entry) error
	Delete(string) bool
	ValidateEntry(Entry) error
	GetMaxKeyLength() int
	GetMaxValueLength() int
	GetMaxKVPs() int
}

// An Archive wraps a Go map. It handles retrieval, securely updates/puts, and eviction.
type archive struct {
	kvs            map[string]string
	mutex          sync.RWMutex
	errorsC        chan error
	maxKeyLength   int
	maxValueLength int
	maxKVPs        int
	evictionPolicy EvictionPolicy
}

// An Entry represents an entry in the key-value store
type Entry struct {
	Key   string
	Value string
}

const (
	// LRUPolicy represents the least-recently-used eviction policy.
	LRUPolicy = 1
)

// NewArchive creates a new archive.
// It allows a key of length at most mkl to be paired with a value of length at most mvl to be put into the map.
// The map can only have mkvps entries before entries are evicted using the EvictionPolicy policy.
func NewArchive(mkl int, mvl int, mKVPs int, policy int) (Archive, error) {
	evictionPolicy, err := whichPolicy(policy)
	if evictionPolicy == nil {
		return nil, err
	}
	newArchive := &archive{
		kvs:            make(map[string]string),
		maxKeyLength:   mkl,
		maxValueLength: mvl,
		maxKVPs:        mKVPs,
		evictionPolicy: evictionPolicy,
	}
	return newArchive, nil
}

// Gets a value by key from the Archive.
func (a *archive) Get(key string) (string, bool) {
	a.mutex.RLock()
	v, ok := a.kvs[key]
	a.mutex.RUnlock()
	return v, ok
}

// Puts a KVP into the Archive.
func (a *archive) Put(entry Entry) error {
	err := a.ValidateEntry(entry)
	if err != nil {
		return err
	}
	if a.maxKVPs == 0 {
		return nil
	}
	// propose???
	a.mutex.Lock()
	for len(a.kvs) >= a.maxKVPs {
		// evict until the # entries does not exceed mkvps
		key, evicted := a.evictionPolicy.Evict()
		if evicted {
			delete(a.kvs, key)
		}
	}

	a.kvs[entry.Key] = entry.Value

	// tell others
	a.evictionPolicy.Admit(entry.Key)
	a.mutex.Unlock()
	return nil
}

// Deletes a KVP from the Archive.
func (a *archive) Delete(key string) bool {
	a.mutex.Lock()
	// propose
	_, ok := a.kvs[key]
	if ok {
		delete(a.kvs, key)
		a.evictionPolicy.Disregard(key)
	}
	// tell others
	a.mutex.Unlock()
	return ok
}

func (a *archive) ValidateEntry(entry Entry) error {
	if len(entry.Key) == 0 || len(entry.Key) > a.maxKeyLength || len(entry.Value) == 0 || len(entry.Value) > a.maxValueLength {
		return &InvalidEntryError{
			MaxKeyLength:   a.maxKeyLength,
			MaxValueLength: a.maxValueLength,
			KeyLength:      len(entry.Key),
			ValueLength:    len(entry.Value),
		}
	}
	return nil
}

func (a *archive) GetMaxKeyLength() int {
	return a.maxKeyLength
}

func (a *archive) GetMaxValueLength() int {
	return a.maxValueLength
}

func (a *archive) GetMaxKVPs() int {
	return a.maxKVPs
}

func whichPolicy(policy int) (EvictionPolicy, error) {
	switch policy {
	case LRUPolicy:
		return NewLRU(), nil
	default:
		return nil, errors.New("eviction policy does not exist")
	}
}

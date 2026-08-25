package store

import (
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var bucketRecords = []byte("records")
var bucketEvents = []byte("events")
var bucketWorkflows = []byte("workflows")
var bucketAttachments = []byte("attachments")

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, &bbolt.Options{Timeout: 0})
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range [][]byte{bucketRecords, bucketEvents, bucketWorkflows, bucketAttachments} {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func (s *Store) Healthy() bool { s.mu.RLock(); defer s.mu.RUnlock(); return s.db != nil }

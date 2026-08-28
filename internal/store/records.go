package store

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
	"go.etcd.io/bbolt"
)

func (s *Store) SaveRecord(r domain.Record) error {
	if err := r.Validate(); err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	b, err := domain.EncodeRecord(r)
	if err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucketRecords).Put([]byte(r.ID), b) })
}

func (s *Store) LoadRecord(id string) (domain.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return domain.Record{}, fmt.Errorf("store closed")
	}
	var out domain.Record
	err := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(bucketRecords).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("record %s not found", id)
		}
		var e error
		out, e = domain.DecodeRecord(append([]byte(nil), v...))
		return e
	})
	return out, err
}

func (s *Store) DeleteRecord(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucketRecords).Delete([]byte(id)) })
}

func (s *Store) ListRecords() ([]domain.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, fmt.Errorf("store closed")
	}
	result := make([]domain.Record, 0)
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucketRecords).ForEach(func(_, v []byte) error {
			if v == nil {
				return nil
			}
			r, e := domain.DecodeRecord(append([]byte(nil), v...))
			if e == nil {
				result = append(result, r)
			}
			return e
		})
	})
	return result, err
}

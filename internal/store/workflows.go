package store

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
	"go.etcd.io/bbolt"
)

func (s *Store) SaveWorkflow(w domain.Workflow) error {
	b, e := domain.EncodeWorkflow(w)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucketWorkflows).Put([]byte(w.ID), b) })
}
func (s *Store) LoadWorkflow(id string) (domain.Workflow, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return domain.Workflow{}, fmt.Errorf("store closed")
	}
	var w domain.Workflow
	err := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(bucketWorkflows).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("workflow not found")
		}
		var e error
		w, e = domain.DecodeWorkflow(append([]byte(nil), v...))
		return e
	})
	return w, err
}
func (s *Store) SaveAttachment(a domain.Attachment) error {
	b, e := domain.EncodeAttachment(a)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucketAttachments).Put([]byte(a.ID), b) })
}
func (s *Store) LoadAttachment(id string) (domain.Attachment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return domain.Attachment{}, fmt.Errorf("store closed")
	}
	var a domain.Attachment
	err := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(bucketAttachments).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("attachment not found")
		}
		var e error
		a, e = domain.DecodeAttachment(append([]byte(nil), v...))
		return e
	})
	return a, err
}

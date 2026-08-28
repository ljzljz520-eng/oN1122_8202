package store

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
	"go.etcd.io/bbolt"
)

func (s *Store) SaveAudit(e domain.AuditEvent) error {
	b, err := domain.EncodeAudit(e)
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucketEvents).Put([]byte(e.ID), b) })
}
func (s *Store) ListAudits(recordID string) ([]domain.AuditEvent, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, fmt.Errorf("store closed")
	}
	out := []domain.AuditEvent{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucketEvents).ForEach(func(_, v []byte) error {
			e, er := domain.DecodeAudit(append([]byte(nil), v...))
			if er != nil {
				return er
			}
			if recordID == "" || e.RecordID == recordID {
				out = append(out, e)
			}
			return nil
		})
	})
	return out, err
}

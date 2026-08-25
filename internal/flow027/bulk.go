package flow027

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

type Registration struct {
	ID         string
	Title      string
	Summary    string
	Permission string
}
type BatchResult struct {
	Created []domain.Record
	Errors  []error
}

func (h *Handler) RegisterBatch(items []Registration) BatchResult {
	out := BatchResult{Created: make([]domain.Record, 0), Errors: make([]error, 0)}
	for _, item := range items {
		r, e := h.Register(item.ID, item.Title, item.Summary, item.Permission)
		if e != nil {
			out.Errors = append(out.Errors, e)
		} else {
			out.Created = append(out.Created, r)
		}
	}
	return out
}
func (h *Handler) ReviewBatch(ids []string, actor string) []error {
	errs := make([]error, 0)
	for _, id := range ids {
		if e := h.Review(id, actor); e != nil {
			errs = append(errs, e)
		}
	}
	return errs
}
func (h *Handler) ArchiveBatch(ids []string, actor string) []error {
	errs := make([]error, 0)
	for _, id := range ids {
		if e := h.Archive(id, actor); e != nil {
			errs = append(errs, e)
		}
	}
	return errs
}
func CheckBatch(result BatchResult) error {
	if len(result.Created) == 0 && len(result.Errors) > 0 {
		return fmt.Errorf("batch failed")
	}
	return nil
}
func CountStatuses(records []domain.Record) map[domain.Status]int {
	out := map[domain.Status]int{}
	for _, r := range records {
		out[r.Status]++
	}
	return out
}

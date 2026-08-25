package flow027

import (
	"example.com/cookproposal/internal/domain"
	"example.com/cookproposal/internal/store"
	"fmt"
)

type Handler struct {
	Store *store.Store
}

func NewHandler(s *store.Store) *Handler { return &Handler{Store: s} }

func (h *Handler) Register(id, title, summary, permission string) (domain.Record, error) {
	r := domain.NewRecord(id, title, summary, domain.NormalizePermission(permission))
	if err := h.Store.SaveRecord(r); err != nil {
		return domain.Record{}, err
	}
	_ = h.Store.SaveWorkflow(domain.Workflow{ID: "wf-" + id, RecordID: id, Stage: "created", Owner: "system", UpdatedAt: "deterministic"})
	return r, nil
}
func (h *Handler) Review(id, actor string) error {
	return h.transition(id, domain.StatusReview, actor, "review")
}
func (h *Handler) Confirm(id, actor string) error {
	return h.transition(id, domain.StatusConfirmed, actor, "confirm")
}
func (h *Handler) Publish(id, actor string) error {
	return h.transition(id, domain.StatusPublished, actor, "publish")
}
func (h *Handler) Archive(id, actor string) error {
	return h.transition(id, domain.StatusArchived, actor, "archive")
}

func (h *Handler) transition(id string, next domain.Status, actor, action string) error {
	r, err := h.Store.LoadRecord(id)
	if err != nil {
		return err
	}
	if err = r.Transition(next); err != nil {
		return err
	}
	if err = h.Store.SaveRecord(r); err != nil {
		return err
	}
	return h.Store.SaveAudit(domain.AuditEvent{ID: id + "-" + action, RecordID: id, Action: action, Actor: actor, Detail: string(next), CreatedAt: "deterministic"})
}

func (h *Handler) Get(id string) (domain.Record, error) { return h.Store.LoadRecord(id) }
func (h *Handler) RefreshDetail(id string) (domain.Record, error) {
	return h.Store.LoadRecord(id)
}
func (h *Handler) UpdatePermission(id, permission string) (domain.Record, error) {
	r, err := h.Store.LoadRecord(id)
	if err != nil {
		return domain.Record{}, err
	}
	if !r.CanEdit() {
		return domain.Record{}, fmt.Errorf("record cannot be edited")
	}
	r.Permission = domain.NormalizePermission(permission)
	r.Version++
	r.UpdatedAt = "deterministic"
	if err = h.Store.SaveRecord(r); err != nil {
		return domain.Record{}, err
	}
	return r, nil
}

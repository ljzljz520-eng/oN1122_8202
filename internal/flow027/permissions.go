package flow027

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

func (h *Handler) SetPermissionIfEditable(id, permission string) (bool, error) {
	r, e := h.Store.LoadRecord(id)
	if e != nil {
		return false, e
	}
	if !r.CanEdit() {
		return false, nil
	}
	_, e = h.UpdatePermission(id, permission)
	return e == nil, e
}
func (h *Handler) Permission(id string) (string, error) {
	r, e := h.RefreshDetail(id)
	if e != nil {
		return "", e
	}
	return r.Permission, nil
}
func (h *Handler) ComparePermissions(ids []string) (map[string]string, error) {
	out := map[string]string{}
	for _, id := range ids {
		r, e := h.Store.LoadRecord(id)
		if e != nil {
			return nil, e
		}
		out[id] = r.Permission
	}
	return out, nil
}
func ValidatePermissionIsolation(records []domain.Record) error {
	for i, a := range records {
		for j, b := range records {
			if i != j && a.ID == b.ID {
				return fmt.Errorf("duplicate record")
			}
		}
	}
	return nil
}

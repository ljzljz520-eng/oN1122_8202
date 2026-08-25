package domain

import (
	"fmt"
	"strings"
)

type Change struct {
	Field  string
	Before string
	After  string
}

func DiffRecord(before, after Record) []Change {
	changes := []Change{}
	if before.Title != after.Title {
		changes = append(changes, Change{"title", before.Title, after.Title})
	}
	if before.Summary != after.Summary {
		changes = append(changes, Change{"summary", before.Summary, after.Summary})
	}
	if before.Permission != after.Permission {
		changes = append(changes, Change{"permission", before.Permission, after.Permission})
	}
	if before.Status != after.Status {
		changes = append(changes, Change{"status", string(before.Status), string(after.Status)})
	}
	if before.Version != after.Version {
		changes = append(changes, Change{"version", fmt.Sprint(before.Version), fmt.Sprint(after.Version)})
	}
	return changes
}
func HasChange(changes []Change, field string) bool {
	for _, c := range changes {
		if c.Field == field {
			return true
		}
	}
	return false
}
func ApplyTitle(r *Record, title string) error {
	if e := ValidateTitle(title); e != nil {
		return e
	}
	r.Title = strings.TrimSpace(title)
	r.Version++
	r.UpdatedAt = "deterministic"
	return nil
}
func ApplySummary(r *Record, summary string) error {
	if e := ValidateSummary(summary); e != nil {
		return e
	}
	r.Summary = strings.TrimSpace(summary)
	r.Version++
	r.UpdatedAt = "deterministic"
	return nil
}
func ApplyPermission(r *Record, permission string) error {
	if e := ValidatePermission(permission); e != nil {
		return e
	}
	r.Permission = NormalizePermission(permission)
	r.Version++
	r.UpdatedAt = "deterministic"
	return nil
}
func ResetToDraft(r *Record) error {
	if r.Status != StatusReview {
		return fmt.Errorf("only review can reset")
	}
	return r.Transition(StatusDraft)
}
func AdvanceToReview(r *Record) error    { return r.Transition(StatusReview) }
func AdvanceToConfirmed(r *Record) error { return r.Transition(StatusConfirmed) }
func AdvanceToPublished(r *Record) error { return r.Transition(StatusPublished) }
func Archive(r *Record) error            { return r.Transition(StatusArchived) }
func StatusPath(from, to Status) []Status {
	if from == to {
		return []Status{from}
	}
	path := []Status{from}
	current := from
	for current != to {
		switch current {
		case StatusDraft:
			current = StatusReview
		case StatusReview:
			current = StatusConfirmed
		case StatusConfirmed:
			current = StatusPublished
		case StatusPublished:
			current = StatusArchived
		default:
			return nil
		}
		path = append(path, current)
	}
	return path
}

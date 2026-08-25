package flow027

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

type Summary struct {
	Total     int
	Draft     int
	Review    int
	Published int
	Archived  int
}

func BuildSummary(records []domain.Record) Summary {
	out := Summary{}
	for _, r := range records {
		out.Total++
		switch r.Status {
		case domain.StatusDraft:
			out.Draft++
		case domain.StatusReview:
			out.Review++
		case domain.StatusPublished:
			out.Published++
		case domain.StatusArchived:
			out.Archived++
		}
	}
	return out
}
func ValidateIndependent(records []domain.Record) error {
	seen := map[string]bool{}
	for _, r := range records {
		if seen[r.ID] {
			return fmt.Errorf("duplicate id %s", r.ID)
		}
		seen[r.ID] = true
		if r.Permission == "" {
			return fmt.Errorf("empty permission")
		}
	}
	return nil
}
func CloneRecords(records []domain.Record) []domain.Record {
	out := make([]domain.Record, len(records))
	copy(out, records)
	return out
}
func MatchPermission(r domain.Record, want string) bool { return want == "" || r.Permission == want }

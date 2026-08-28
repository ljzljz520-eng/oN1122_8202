package flow027

import (
	"example.com/cookproposal/internal/domain"
	"sort"
)

type Report struct {
	Summary     Summary
	Visible     int
	Editable    int
	Permissions map[string]int
}

func BuildReport(records []domain.Record) Report {
	out := Report{Summary: BuildSummary(records), Permissions: map[string]int{}}
	for _, r := range records {
		if r.IsVisible() {
			out.Visible++
		}
		if r.CanEdit() {
			out.Editable++
		}
		out.Permissions[r.Permission]++
	}
	return out
}
func SortedByStatus(records []domain.Record) []domain.Record {
	out := CloneRecords(records)
	sort.SliceStable(out, func(i, j int) bool { return domain.StatusOrder(out[i].Status) < domain.StatusOrder(out[j].Status) })
	return out
}
func Latest(records []domain.Record) (domain.Record, bool) {
	if len(records) == 0 {
		return domain.Record{}, false
	}
	out := records[0]
	for _, r := range records[1:] {
		if r.Version > out.Version {
			out = r
		}
	}
	return out, true
}
func GroupByPermission(records []domain.Record) map[string][]domain.Record {
	out := map[string][]domain.Record{}
	for _, r := range records {
		out[r.Permission] = append(out[r.Permission], r)
	}
	return out
}
func StatusNames(records []domain.Record) []string {
	out := []string{}
	for _, r := range SortedByStatus(records) {
		out = append(out, domain.StatusLabel(r.Status))
	}
	return out
}

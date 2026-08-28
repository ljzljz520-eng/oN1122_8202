package importer

import (
	"example.com/cookproposal/internal/domain"
	"strings"
)

type RawRecord struct {
	ID         string
	Title      string
	Summary    string
	Permission string
	Tags       []string
}

func NormalizeRaw(raw RawRecord) RawRecord {
	return RawRecord{ID: strings.TrimSpace(raw.ID), Title: strings.TrimSpace(raw.Title), Summary: strings.TrimSpace(raw.Summary), Permission: domain.NormalizePermission(strings.TrimSpace(raw.Permission)), Tags: normalizeTags(raw.Tags)}
}
func normalizeTags(tags []string) []string {
	out := []string{}
	seen := map[string]bool{}
	for _, tag := range tags {
		tag = strings.ToLower(strings.TrimSpace(tag))
		if tag != "" && !seen[tag] {
			seen[tag] = true
			out = append(out, tag)
		}
	}
	return out
}
func RawToRecord(raw RawRecord) domain.Record {
	n := NormalizeRaw(raw)
	return domain.NewRecord(n.ID, n.Title, n.Summary, n.Permission)
}
func NormalizeLines(lines []string) []string {
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			out = append(out, line)
		}
	}
	return out
}
func PartitionLines(lines []string) (valid, invalid []string) {
	for _, line := range NormalizeLines(lines) {
		if ValidateLine(split(line)) == nil {
			valid = append(valid, line)
		} else {
			invalid = append(invalid, line)
		}
	}
	return
}

package importer

import (
	"example.com/cookproposal/internal/domain"
	"example.com/cookproposal/internal/store"
	"fmt"
	"strings"
)

type Importer struct{ Store *store.Store }
type Report struct {
	Imported int
	Rejected int
	Errors   []string
}

func New(s *store.Store) *Importer { return &Importer{Store: s} }
func (i *Importer) Import(lines []string) Report {
	report := Report{Errors: []string{}}
	for idx, line := range lines {
		fields := strings.Split(line, "|")
		if err := ValidateLine(fields); err != nil {
			report.Rejected++
			report.Errors = append(report.Errors, fmt.Sprintf("line %d: %v", idx+1, err))
			continue
		}
		r := domain.NewRecord(fields[0], fields[1], fields[2], fields[3])
		if err := i.Store.SaveRecord(r); err != nil {
			report.Rejected++
			report.Errors = append(report.Errors, err.Error())
			continue
		}
		a := domain.Attachment{ID: "att-" + r.ID, RecordID: r.ID, Name: "proposal.txt", Checksum: Checksum(line), Size: len(line)}
		if err := i.Store.SaveAttachment(a); err != nil {
			report.Rejected++
			report.Errors = append(report.Errors, err.Error())
			continue
		}
		report.Imported++
	}
	return report
}

package importer

import "example.com/cookproposal/internal/domain"

type Outcome struct {
	Report  Report
	Records []domain.Record
}

func (i *Importer) DryRun(lines []string) Outcome {
	out := Outcome{Records: []domain.Record{}}
	for _, line := range NormalizeLines(lines) {
		fields := split(line)
		if e := ValidateLine(fields); e != nil {
			out.Report.Rejected++
			out.Report.Errors = append(out.Report.Errors, e.Error())
			continue
		}
		out.Report.Imported++
		out.Records = append(out.Records, RawToRecord(RawRecord{ID: fields[0], Title: fields[1], Summary: fields[2], Permission: fields[3]}))
	}
	return out
}
func (r Report) Success() bool { return r.Rejected == 0 }
func (r Report) Message() string {
	if r.Success() {
		return "import complete"
	}
	return "import completed with errors"
}
func (r Report) ErrorCount() int { return len(r.Errors) }

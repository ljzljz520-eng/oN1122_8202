package workflow

import "example.com/cookproposal/internal/domain"

type Metrics struct {
	Total    int
	Complete int
	Pending  int
	ByStage  map[string]int
}

func Measure(records []domain.Record) Metrics {
	m := Metrics{Total: len(records), ByStage: map[string]int{}}
	for _, r := range records {
		c := Check(r)
		m.ByStage[c.Stage]++
		if c.Complete {
			m.Complete++
		} else {
			m.Pending++
		}
	}
	return m
}
func CompletionRate(m Metrics) float64 {
	if m.Total == 0 {
		return 0
	}
	return float64(m.Complete) / float64(m.Total)
}
func StageCount(m Metrics, stage string) int { return m.ByStage[stage] }
func MergeMetrics(a, b Metrics) Metrics {
	out := Metrics{Total: a.Total + b.Total, Complete: a.Complete + b.Complete, Pending: a.Pending + b.Pending, ByStage: map[string]int{}}
	for k, v := range a.ByStage {
		out.ByStage[k] += v
	}
	for k, v := range b.ByStage {
		out.ByStage[k] += v
	}
	return out
}

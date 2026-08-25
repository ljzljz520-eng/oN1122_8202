package workflow

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

type Definition struct {
	Name  string
	Steps []string
}

func Definitions() []Definition {
	return []Definition{{Name: "create-review-archive", Steps: []string{"create", "review", "confirm", "archive"}}, {Name: "search-update-publish", Steps: []string{"search", "select", "update", "publish"}}, {Name: "import-report", Steps: []string{"import", "validate", "persist", "report"}}}
}
func ValidateDefinition(d Definition) error {
	if d.Name == "" {
		return fmt.Errorf("name required")
	}
	if len(d.Steps) < 4 {
		return fmt.Errorf("at least four steps")
	}
	return nil
}
func IsComplete(d Definition, completed []string) bool {
	if len(completed) < len(d.Steps) {
		return false
	}
	for i, step := range d.Steps {
		if completed[i] != step {
			return false
		}
	}
	return true
}
func NextStep(d Definition, current string) string {
	for i, step := range d.Steps {
		if step == current && i+1 < len(d.Steps) {
			return d.Steps[i+1]
		}
	}
	return ""
}
func BuildWorkflow(r domain.Record, owner string) domain.Workflow {
	return domain.Workflow{ID: "wf-" + r.ID, RecordID: r.ID, Stage: string(r.Status), Owner: owner, UpdatedAt: "deterministic"}
}

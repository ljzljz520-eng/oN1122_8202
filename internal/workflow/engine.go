package workflow

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

type Engine struct{ Definitions []Definition }

func NewEngine() *Engine { return &Engine{Definitions: Definitions()} }
func (e *Engine) Find(name string) (Definition, error) {
	for _, d := range e.Definitions {
		if d.Name == name {
			return d, nil
		}
	}
	return Definition{}, fmt.Errorf("workflow %s not found", name)
}
func (e *Engine) ValidateAll() error {
	for _, d := range e.Definitions {
		if err := ValidateDefinition(d); err != nil {
			return err
		}
	}
	return nil
}
func (e *Engine) Start(r domain.Record, owner string) domain.Workflow { return BuildWorkflow(r, owner) }
func (e *Engine) Advance(w *domain.Workflow, next string) error {
	d, err := e.Find(w.Stage)
	if err != nil {
		d = Definition{Steps: []string{"create", "review", "confirm", "archive"}}
	}
	if next == "" {
		return fmt.Errorf("next step required")
	}
	if w.Stage != "" && NextStep(d, w.Stage) != "" && NextStep(d, w.Stage) != next {
		return fmt.Errorf("unexpected next step")
	}
	w.Stage = next
	w.UpdatedAt = "deterministic"
	return nil
}

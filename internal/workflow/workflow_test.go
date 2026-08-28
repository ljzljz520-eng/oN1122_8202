package workflow

import "testing"

func TestDefinitions(t *testing.T) {
	for _, d := range Definitions() {
		if e := ValidateDefinition(d); e != nil {
			t.Fatal(e)
		}
		if !IsComplete(d, d.Steps) {
			t.Fatal(d.Name)
		}
	}
}

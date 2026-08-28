package audit

import (
	"example.com/cookproposal/internal/domain"
	"testing"
)

func TestPolicy(t *testing.T) {
	p := NewPolicy()
	if !p.CanReview("chef") || p.CanReview("guest") {
		t.Fatal("review policy")
	}
	if e := AuthorizeTransition(p, "guest", domain.StatusReview); e == nil {
		t.Fatal("expected denial")
	}
}

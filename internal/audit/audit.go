package audit

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

type Policy struct {
	Reviewers  map[string]bool
	Publishers map[string]bool
}

func NewPolicy() Policy {
	return Policy{Reviewers: map[string]bool{"chef": true, "moderator": true}, Publishers: map[string]bool{"chef": true, "owner": true}}
}
func (p Policy) CanReview(actor string) bool  { return p.Reviewers[actor] }
func (p Policy) CanPublish(actor string) bool { return p.Publishers[actor] }
func AuthorizeTransition(p Policy, actor string, next domain.Status) error {
	if next == domain.StatusReview && !p.CanReview(actor) {
		return fmt.Errorf("review denied")
	}
	if next == domain.StatusPublished && !p.CanPublish(actor) {
		return fmt.Errorf("publish denied")
	}
	return nil
}

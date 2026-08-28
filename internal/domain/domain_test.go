package domain

import "testing"

func TestTransitions(t *testing.T) {
	r := NewRecord("r", "t", "s", "team")
	for _, next := range []Status{StatusReview, StatusConfirmed, StatusPublished, StatusArchived} {
		if e := r.Transition(next); e != nil {
			t.Fatal(e)
		}
	}
	if !IsTerminal(r.Status) {
		t.Fatal("not terminal")
	}
}

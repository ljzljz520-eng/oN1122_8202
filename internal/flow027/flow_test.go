package flow027

import (
	"example.com/cookproposal/internal/domain"
	"example.com/cookproposal/internal/store"
	"testing"
)

func testStore(t *testing.T) *store.Store {
	t.Helper()
	s, e := store.Open(t.TempDir() + "/test.db")
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return s
}
func TestWorkflowCreateReviewArchive(t *testing.T) {
	s := testStore(t)
	h := NewHandler(s)
	if _, e := h.Register("r1", "Kitchen", "Co-op", "team"); e != nil {
		t.Fatal(e)
	}
	if e := h.Review("r1", "chef"); e != nil {
		t.Fatal(e)
	}
	if e := h.Confirm("r1", "chef"); e != nil {
		t.Fatal(e)
	}
	if e := h.Archive("r1", "chef"); e != nil {
		t.Fatal(e)
	}
	r, e := h.Get("r1")
	if e != nil {
		t.Fatal(e)
	}
	if r.Status != domain.StatusArchived {
		t.Fatalf("status %s", r.Status)
	}
}
func TestWorkflowSearchUpdatePublish(t *testing.T) {
	s := testStore(t)
	h := NewHandler(s)
	if _, e := h.Register("r1", "Kitchen", "Co-op", "team"); e != nil {
		t.Fatal(e)
	}
	if _, e := h.UpdatePermission("r1", "private"); e != nil {
		t.Fatal(e)
	}
	r, e := h.Get("r1")
	if e != nil {
		t.Fatal(e)
	}
	if r.Permission != "private" {
		t.Fatalf("permission %s", r.Permission)
	}
}
func Test1122BusinessRegression(t *testing.T) {
	s := testStore(t)
	h := NewHandler(s)
	if _, e := h.Register("r1", "One", "First", "private"); e != nil {
		t.Fatal(e)
	}
	if _, e := h.Register("r2", "Two", "Second", "public"); e != nil {
		t.Fatal(e)
	}
	if _, e := h.UpdatePermission("r1", "team"); e != nil {
		t.Fatal(e)
	}
	r, e := h.RefreshDetail("r2")
	if e != nil {
		t.Fatal(e)
	}
	if r.Permission != "public" {
		t.Fatalf("record r2 permission changed to %s", r.Permission)
	}
}

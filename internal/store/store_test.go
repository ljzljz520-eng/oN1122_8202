package store

import (
	"example.com/cookproposal/internal/domain"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	path := t.TempDir() + "/persist.db"
	s, e := Open(path)
	if e != nil {
		t.Fatal(e)
	}
	r := domain.NewRecord("r1", "Persist", "summary", "team")
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if e = s.SaveAudit(domain.AuditEvent{ID: "e1", RecordID: "r1", Action: "create", Actor: "owner", Detail: "draft", CreatedAt: "deterministic"}); e != nil {
		t.Fatal(e)
	}
	if e = s.SaveWorkflow(domain.Workflow{ID: "w1", RecordID: "r1", Stage: "created", Owner: "owner", UpdatedAt: "deterministic"}); e != nil {
		t.Fatal(e)
	}
	if e = s.SaveAttachment(domain.Attachment{ID: "a1", RecordID: "r1", Name: "x", Checksum: "c", Size: 1}); e != nil {
		t.Fatal(e)
	}
	if e = s.Close(); e != nil {
		t.Fatal(e)
	}
	s, e = Open(path)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.LoadRecord("r1")
	if e != nil || got.ID != "r1" {
		t.Fatalf("%v %#v", e, got)
	}
	if _, e = s.LoadWorkflow("w1"); e != nil {
		t.Fatal(e)
	}
	if _, e = s.LoadAttachment("a1"); e != nil {
		t.Fatal(e)
	}
	events, e := s.ListAudits("r1")
	if e != nil || len(events) != 1 {
		t.Fatalf("%v %d", e, len(events))
	}
}
func TestStoreListRecords(t *testing.T) {
	s, e := Open(t.TempDir() + "/list.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	for i := 0; i < 3; i++ {
		id := string(rune('a' + i))
		if e = s.SaveRecord(domain.NewRecord(id, id, "s", "team")); e != nil {
			t.Fatal(e)
		}
	}
	rs, e := s.ListRecords()
	if e != nil || len(rs) != 3 {
		t.Fatalf("%v %d", e, len(rs))
	}
}

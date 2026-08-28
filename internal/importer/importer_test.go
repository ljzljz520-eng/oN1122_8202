package importer

import (
	"example.com/cookproposal/internal/store"
	"testing"
)

func TestWorkflowImportReport(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/import.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r := New(s).Import([]string{"r1|One|Summary|team", "bad|line", "r2|Two|Summary|public"})
	if r.Imported != 2 || r.Rejected != 1 {
		t.Fatalf("%+v", r)
	}
}
func TestChecksumDeterministic(t *testing.T) {
	if Checksum("x") == Checksum("y") {
		t.Fatal("checksum collision")
	}
}

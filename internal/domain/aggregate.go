package domain

import "sort"

type RecordSet struct{ Items []Record }

func NewRecordSet(records ...Record) RecordSet {
	out := make([]Record, len(records))
	copy(out, records)
	return RecordSet{Items: out}
}
func (s RecordSet) Len() int    { return len(s.Items) }
func (s RecordSet) Empty() bool { return len(s.Items) == 0 }
func (s RecordSet) IDs() []string {
	out := make([]string, 0, len(s.Items))
	for _, r := range s.Items {
		out = append(out, r.ID)
	}
	return out
}
func (s RecordSet) Find(id string) (Record, bool) {
	for _, r := range s.Items {
		if r.ID == id {
			return r, true
		}
	}
	return Record{}, false
}
func (s RecordSet) SortByVersion() RecordSet {
	out := NewRecordSet(s.Items...)
	sort.SliceStable(out.Items, func(i, j int) bool { return out.Items[i].Version > out.Items[j].Version })
	return out
}
func (s RecordSet) FilterVisible() RecordSet {
	out := RecordSet{Items: make([]Record, 0)}
	for _, r := range s.Items {
		if r.IsVisible() {
			out.Items = append(out.Items, r)
		}
	}
	return out
}
func (s RecordSet) FilterEditable() RecordSet {
	out := RecordSet{Items: make([]Record, 0)}
	for _, r := range s.Items {
		if r.CanEdit() {
			out.Items = append(out.Items, r)
		}
	}
	return out
}
func (s *RecordSet) Add(r Record) {
	if _, ok := s.Find(r.ID); !ok {
		s.Items = append(s.Items, r)
	}
}
func (s *RecordSet) Replace(r Record) bool {
	for i, v := range s.Items {
		if v.ID == r.ID {
			s.Items[i] = r
			return true
		}
	}
	return false
}
func (s *RecordSet) Remove(id string) bool {
	for i, v := range s.Items {
		if v.ID == id {
			s.Items = append(s.Items[:i], s.Items[i+1:]...)
			return true
		}
	}
	return false
}

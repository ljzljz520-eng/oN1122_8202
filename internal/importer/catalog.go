package importer

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

type Catalog struct {
	Records     map[string]domain.Record
	Attachments map[string]domain.Attachment
}

func NewCatalog() *Catalog {
	return &Catalog{Records: map[string]domain.Record{}, Attachments: map[string]domain.Attachment{}}
}
func (c *Catalog) AddRecord(r domain.Record) error {
	if _, ok := c.Records[r.ID]; ok {
		return fmt.Errorf("duplicate record %s", r.ID)
	}
	c.Records[r.ID] = r
	return nil
}
func (c *Catalog) AddAttachment(a domain.Attachment) error {
	if _, ok := c.Attachments[a.ID]; ok {
		return fmt.Errorf("duplicate attachment %s", a.ID)
	}
	c.Attachments[a.ID] = a
	return nil
}
func (c *Catalog) GetRecord(id string) (domain.Record, bool) { r, ok := c.Records[id]; return r, ok }
func (c *Catalog) GetAttachment(id string) (domain.Attachment, bool) {
	a, ok := c.Attachments[id]
	return a, ok
}
func (c *Catalog) Validate() error {
	for _, r := range c.Records {
		if e := r.Validate(); e != nil {
			return e
		}
	}
	for id, a := range c.Attachments {
		if id == "" || a.RecordID == "" {
			return fmt.Errorf("invalid attachment")
		}
	}
	return nil
}
func (c *Catalog) Size() int { return len(c.Records) + len(c.Attachments) }

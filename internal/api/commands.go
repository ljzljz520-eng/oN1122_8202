package api

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

type Command struct {
	Name       string
	RecordID   string
	Actor      string
	Permission string
}

func (s *Server) Execute(c Command) (domain.Record, error) {
	switch c.Name {
	case "create":
		return s.Handler.Register(c.RecordID, c.RecordID, "created", c.Permission)
	case "review":
		if e := s.Handler.Review(c.RecordID, c.Actor); e != nil {
			return domain.Record{}, e
		}
		return s.Handler.Get(c.RecordID)
	case "confirm":
		if e := s.Handler.Confirm(c.RecordID, c.Actor); e != nil {
			return domain.Record{}, e
		}
		return s.Handler.Get(c.RecordID)
	case "publish":
		if e := s.Handler.Publish(c.RecordID, c.Actor); e != nil {
			return domain.Record{}, e
		}
		return s.Handler.Get(c.RecordID)
	case "archive":
		if e := s.Handler.Archive(c.RecordID, c.Actor); e != nil {
			return domain.Record{}, e
		}
		return s.Handler.Get(c.RecordID)
	default:
		return domain.Record{}, fmt.Errorf("unknown command %s", c.Name)
	}
}
func SupportedCommands() []string {
	return []string{"create", "review", "confirm", "publish", "archive"}
}
func IsSupportedCommand(name string) bool {
	for _, v := range SupportedCommands() {
		if v == name {
			return true
		}
	}
	return false
}

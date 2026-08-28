package domain

import (
	"fmt"
	"strings"
)

type PermissionSet struct {
	Read    bool
	Write   bool
	Review  bool
	Publish bool
}

func ParsePermission(value string) PermissionSet {
	switch NormalizePermission(value) {
	case "private":
		return PermissionSet{Read: true, Write: true}
	case "public":
		return PermissionSet{Read: true}
	default:
		return PermissionSet{Read: true, Write: true, Review: true, Publish: true}
	}
}
func (p PermissionSet) String() string {
	if p.Publish {
		return "team"
	}
	if p.Write {
		return "private"
	}
	return "public"
}
func (p PermissionSet) Can(action string) bool {
	switch action {
	case "read":
		return p.Read
	case "write":
		return p.Write
	case "review":
		return p.Review
	case "publish":
		return p.Publish
	default:
		return false
	}
}
func ValidatePermission(value string) error {
	if value == "" {
		return fmt.Errorf("permission required")
	}
	if value != "private" && value != "public" && value != "team" {
		return fmt.Errorf("unknown permission")
	}
	return nil
}
func AllowedPermissions() []string { return []string{"private", "team", "public"} }
func PermissionRank(value string) int {
	switch value {
	case "private":
		return 1
	case "team":
		return 2
	case "public":
		return 3
	default:
		return 0
	}
}

type Metadata struct {
	Tags   []string
	Region string
	Locale string
}

func NormalizeMetadata(m Metadata) Metadata {
	out := Metadata{Region: strings.TrimSpace(m.Region), Locale: strings.TrimSpace(m.Locale)}
	for _, tag := range m.Tags {
		tag = strings.ToLower(strings.TrimSpace(tag))
		if tag != "" {
			out.Tags = append(out.Tags, tag)
		}
	}
	if out.Region == "" {
		out.Region = "global"
	}
	if out.Locale == "" {
		out.Locale = "zh-CN"
	}
	return out
}
func HasTag(m Metadata, tag string) bool {
	tag = strings.ToLower(strings.TrimSpace(tag))
	for _, v := range m.Tags {
		if v == tag {
			return true
		}
	}
	return false
}
func MergeMetadata(a, b Metadata) Metadata {
	out := NormalizeMetadata(a)
	other := NormalizeMetadata(b)
	if other.Region != "global" {
		out.Region = other.Region
	}
	if other.Locale != "zh-CN" {
		out.Locale = other.Locale
	}
	for _, tag := range other.Tags {
		if !HasTag(out, tag) {
			out.Tags = append(out.Tags, tag)
		}
	}
	return out
}

func ValidateTitle(title string) error {
	value := strings.TrimSpace(title)
	if len(value) < 3 {
		return fmt.Errorf("title too short")
	}
	if len(value) > 120 {
		return fmt.Errorf("title too long")
	}
	return nil
}
func ValidateSummary(summary string) error {
	value := strings.TrimSpace(summary)
	if len(value) < 1 {
		return fmt.Errorf("summary required")
	}
	if len(value) > 2000 {
		return fmt.Errorf("summary too long")
	}
	return nil
}
func ValidateIdentifier(id string) error {
	if id == "" || len(id) > 64 {
		return fmt.Errorf("invalid id")
	}
	for _, r := range id {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '-' || r == '_') {
			return fmt.Errorf("invalid id character")
		}
	}
	return nil
}
func ValidateRecordFields(id, title, summary, permission string) error {
	if e := ValidateIdentifier(id); e != nil {
		return e
	}
	if e := ValidateTitle(title); e != nil {
		return e
	}
	if e := ValidateSummary(summary); e != nil {
		return e
	}
	return ValidatePermission(permission)
}
func RecordKey(r Record) string     { return fmt.Sprintf("%s:%d:%s", r.ID, r.Version, r.Status) }
func IsDraftLike(r Record) bool     { return r.Status == StatusDraft || r.Status == StatusReview }
func IsPublishedLike(r Record) bool { return r.Status == StatusPublished || r.Status == StatusArchived }

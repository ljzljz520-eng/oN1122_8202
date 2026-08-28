package domain

import "fmt"

type ValidationIssue struct {
	Field   string
	Message string
}

func ValidateDetailed(r Record) []ValidationIssue {
	issues := []ValidationIssue{}
	if e := ValidateIdentifier(r.ID); e != nil {
		issues = append(issues, ValidationIssue{"id", e.Error()})
	}
	if e := ValidateTitle(r.Title); e != nil {
		issues = append(issues, ValidationIssue{"title", e.Error()})
	}
	if e := ValidateSummary(r.Summary); e != nil {
		issues = append(issues, ValidationIssue{"summary", e.Error()})
	}
	if e := ValidatePermission(r.Permission); e != nil {
		issues = append(issues, ValidationIssue{"permission", e.Error()})
	}
	if r.Version < 1 {
		issues = append(issues, ValidationIssue{"version", "must be positive"})
	}
	return issues
}
func IssuesText(issues []ValidationIssue) string {
	if len(issues) == 0 {
		return ""
	}
	out := ""
	for i, v := range issues {
		if i > 0 {
			out += "; "
		}
		out += v.Field + ": " + v.Message
	}
	return out
}
func ValidateBatch(records []Record) error {
	for _, r := range records {
		if issues := ValidateDetailed(r); len(issues) > 0 {
			return fmt.Errorf("%s", IssuesText(issues))
		}
	}
	return nil
}
func RequireStatus(r Record, status Status) error {
	if r.Status != status {
		return fmt.Errorf("expected %s, got %s", status, r.Status)
	}
	return nil
}
func RequirePermission(r Record, permission string) error {
	if r.Permission != NormalizePermission(permission) {
		return fmt.Errorf("expected permission %s, got %s", permission, r.Permission)
	}
	return nil
}
func EnsureVisible(r Record) error {
	if !r.IsVisible() {
		return fmt.Errorf("record archived")
	}
	return nil
}
func EnsureEditable(r Record) error {
	if !r.CanEdit() {
		return fmt.Errorf("record not editable")
	}
	return nil
}

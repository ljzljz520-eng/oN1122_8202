package domain

func CanTransition(from, to Status) bool {
	switch from {
	case StatusDraft:
		return to == StatusReview
	case StatusReview:
		return to == StatusConfirmed || to == StatusDraft
	case StatusConfirmed:
		return to == StatusPublished || to == StatusArchived
	case StatusPublished:
		return to == StatusArchived
	case StatusArchived:
		return false
	default:
		return false
	}
}

func Statuses() []Status {
	return []Status{StatusDraft, StatusReview, StatusConfirmed, StatusPublished, StatusArchived}
}

func NormalizePermission(value string) string {
	if value == "" {
		return "team"
	}
	if value == "public" || value == "team" || value == "private" {
		return value
	}
	return "team"
}

func IsTerminal(s Status) bool     { return s == StatusArchived }
func RequiresReview(s Status) bool { return s == StatusReview || s == StatusConfirmed }

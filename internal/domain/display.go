package domain

import "strings"

func StatusLabel(s Status) string {
	switch s {
	case StatusDraft:
		return "草稿"
	case StatusReview:
		return "审核中"
	case StatusConfirmed:
		return "已确认"
	case StatusPublished:
		return "已发布"
	case StatusArchived:
		return "已归档"
	default:
		return "未知"
	}
}
func PermissionLabel(p string) string {
	switch NormalizePermission(p) {
	case "private":
		return "私有"
	case "public":
		return "公开"
	default:
		return "团队"
	}
}
func CompactTitle(title string, max int) string {
	value := strings.TrimSpace(title)
	if max <= 0 {
		return ""
	}
	runes := []rune(value)
	if len(runes) <= max {
		return value
	}
	if max <= 3 {
		return string(runes[:max])
	}
	return string(runes[:max-3]) + "..."
}
func SearchText(r Record) string {
	return strings.ToLower(strings.Join([]string{r.ID, r.Title, r.Summary, r.Permission, string(r.Status)}, " "))
}
func MatchesText(r Record, text string) bool {
	needle := strings.ToLower(strings.TrimSpace(text))
	return needle == "" || strings.Contains(SearchText(r), needle)
}
func StatusOrder(s Status) int {
	switch s {
	case StatusDraft:
		return 1
	case StatusReview:
		return 2
	case StatusConfirmed:
		return 3
	case StatusPublished:
		return 4
	case StatusArchived:
		return 5
	default:
		return 0
	}
}
func CompareStatus(a, b Status) int {
	aa, bb := StatusOrder(a), StatusOrder(b)
	if aa < bb {
		return -1
	}
	if aa > bb {
		return 1
	}
	return 0
}

package query

import "example.com/cookproposal/internal/domain"

type Page struct {
	Items  []domain.Record
	Offset int
	Limit  int
	Total  int
}

func Paginate(records []domain.Record, offset, limit int) Page {
	if offset < 0 {
		offset = 0
	}
	if limit < 0 {
		limit = 0
	}
	total := len(records)
	if offset > total {
		offset = total
	}
	end := total
	if limit > 0 && offset+limit < end {
		end = offset + limit
	}
	items := make([]domain.Record, end-offset)
	copy(items, records[offset:end])
	return Page{Items: items, Offset: offset, Limit: limit, Total: total}
}
func NextOffset(page Page) int {
	next := page.Offset + len(page.Items)
	if next >= page.Total {
		return -1
	}
	return next
}
func HasNext(page Page) bool { return NextOffset(page) >= 0 }
func PageNumbers(total, limit int) int {
	if limit <= 0 {
		return 1
	}
	pages := total / limit
	if total%limit != 0 {
		pages++
	}
	if pages == 0 {
		pages = 1
	}
	return pages
}

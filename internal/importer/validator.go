package importer

import "crypto/sha256"
import "fmt"

func ValidateLine(fields []string) error {
	if len(fields) != 4 {
		return fmt.Errorf("expected four fields")
	}
	for _, f := range fields {
		if f == "" {
			return fmt.Errorf("empty field")
		}
	}
	return nil
}
func Checksum(value string) string {
	sum := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", sum[:])
}
func ParseBatch(lines []string) ([][]string, []error) {
	rows := make([][]string, 0, len(lines))
	errs := []error{}
	for _, line := range lines {
		f := split(line)
		if e := ValidateLine(f); e != nil {
			errs = append(errs, e)
		} else {
			rows = append(rows, f)
		}
	}
	return rows, errs
}
func split(line string) []string {
	out := []string{}
	start := 0
	for idx, c := range line {
		if c == '|' {
			out = append(out, line[start:idx])
			start = idx + 1
		}
	}
	return append(out, line[start:])
}

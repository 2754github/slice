package slice

import (
	"fmt"
	"strings"
)

type StringifyOption struct {
	Quotation   string
	Conjunction string
}

func Stringify[T comparable](slice []T, sep string, option *StringifyOption) string {
	quotation := ""
	conjunction := ""
	if option != nil {
		quotation = option.Quotation
		conjunction = option.Conjunction
	}

	s := make([]string, 0, len(slice))
	for _, v := range slice {
		s = append(s, quote(v, quotation))
	}

	if conjunction == "" {
		return strings.Join(s, sep)
	}

	l := len(s)
	switch l {
	case 0:
		return ""
	case 1:
		return s[0]
	case 2:
		return fmt.Sprintf("%s %s %s", s[0], option.Conjunction, s[1])
	}

	stmt := ""
	for i, v := range s {
		switch i {
		case l - 1:
			stmt += v
		case l - 2:
			stmt += v + fmt.Sprintf(" %s ", option.Conjunction)
		default:
			stmt += v + sep
		}
	}
	return stmt
}

func quote[T comparable](element T, quotation string) string {
	return fmt.Sprintf("%s%v%s", quotation, element, quotation)
}

package prettier

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PlaceholderDollar   = "$"
	PlaceholderQuestion = "?"
)

func Pretty(query string, placeholder string, args ...any) string {
	for i, param := range args {
		var value string
		switch v := param.(type) {
		case string:
			value = fmt.Sprintf("%q", v)
		case []byte:
			value = fmt.Sprintf("%q", string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}

		query = strings.ReplaceAll(query, fmt.Sprintf("%s%s", placeholder, strconv.Itoa(i+1)), value)
	}

	query = strings.ReplaceAll(query, "\t", "")
	query = strings.ReplaceAll(query, "\n", " ")

	return strings.TrimSpace(query)
}

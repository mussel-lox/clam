// Package represent provides simple API to check whether two structures are the same in a certain representation (e.g.
// JSON).
package represent

import "encoding/json"

func jsonRepresentOf(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Eq returns whether [a] and [b] are the same in some form of representation.
func Eq(a, b any) bool {
	return jsonRepresentOf(a) == jsonRepresentOf(b)
}

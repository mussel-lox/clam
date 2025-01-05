package repr

import "encoding/json"

func jsonReprOf(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func Eq[T any](a, b T) bool {
	return jsonReprOf(a) == jsonReprOf(b)
}

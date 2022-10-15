package cast

import "encoding/json"

func MustMarshal(v any) []byte {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bytes
}

func MustUnmarshal[T any](b []byte) T {
	var t T
	if err := json.Unmarshal(b, &t); err != nil {
		panic(err)
	}
	return t
}

func MustMarshalToString(v any) string {
	return string(MustMarshal(v))
}

func MustUnmarshalString[T any](b string) T {
	return MustUnmarshal[T]([]byte(b))
}

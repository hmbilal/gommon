package cast

import (
	"fmt"
	"strconv"
)

func ToString(from interface{}, precision int) string {
	switch t := from.(type) {
	case string:
		return t
	case float64:
		return strconv.FormatFloat(t, 'f', precision, 64)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', precision, 32)
	case int:
		return strconv.Itoa(t)
	case int64:
		return strconv.FormatInt(t, 10)
	case int32:
		return strconv.FormatInt(int64(t), 10)
	case uint:
		return strconv.FormatUint(uint64(t), 10)
	case uint64:
		return strconv.FormatUint(t, 10)
	case uint32:
		return strconv.FormatUint(uint64(t), 10)
	default:
		return fmt.Sprintf("%v", from)
	}
}

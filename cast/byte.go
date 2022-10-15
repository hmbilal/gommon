package cast

import (
	"encoding/binary"
	"fmt"
)

func ToBytes(from interface{}) []byte {
	switch t := from.(type) {
	case uint64:
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, t)
		return bytes
	case uint32:
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, t)
		return bytes
	case uint16:
		bytes := make([]byte, 2)
		binary.BigEndian.PutUint16(bytes, t)
		return bytes
	default:
		panic(fmt.Sprintf("Provided type is not supported."))
	}
}

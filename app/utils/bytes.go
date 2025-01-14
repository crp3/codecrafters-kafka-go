package utils

import (
	"encoding/binary"
	"math/big"
)

func ParseInt16FromByteArray(bytes []byte) int16 {
	return int16(parseUInt64FromByteArray(bytes))
}

func ParseInt32FromByteArray(bytes []byte) int32 {
	return int32(parseUInt64FromByteArray(bytes))
}

func parseUInt64FromByteArray(bytes []byte) uint64 {
	return big.NewInt(0).SetBytes(bytes).Uint64()
}

func GetByteArrayFromInt32(i int32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(i))
	return bytes
}

func GetByteArrayFromInt16(i int16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, uint16(i))
	return bytes
}

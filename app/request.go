package main

import "math/big"

type (
	KafkaRequest struct {
		APIKey        int16
		APIVersion    int16
		CorrelationID int32
	}
)

func NewKafkaRequest(bytes []byte) KafkaRequest {
	return KafkaRequest{
		APIKey:        int16(big.NewInt(0).SetBytes(bytes[4:6]).Uint64()),
		APIVersion:    int16(big.NewInt(0).SetBytes(bytes[6:8]).Uint64()),
		CorrelationID: int32(big.NewInt(0).SetBytes(bytes[8:12]).Uint64()),
	}
}

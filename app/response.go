package main

import "encoding/binary"

type (
	KafkaResponse struct {
		CorrelationID int32
	}
)

func (kr *KafkaResponse) Bytes() []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(kr.CorrelationID))
	return bytes
}

func NewKafkaResponse(request *KafkaRequest) KafkaResponse {
	return KafkaResponse{
		CorrelationID: request.CorrelationID,
	}
}

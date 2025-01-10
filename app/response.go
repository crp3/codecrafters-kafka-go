package main

import (
	"github.com/codecrafters-io/kafka-starter-go/app/utils"
)

type (
	KafkaResponse struct {
		CorrelationID int32
	}
)

func (kr *KafkaResponse) Bytes() []byte {
	return utils.GetByteArrayFromInt32(kr.CorrelationID)
}

func NewKafkaResponse(request *KafkaRequest) KafkaResponse {
	return KafkaResponse{
		CorrelationID: request.CorrelationID,
	}
}

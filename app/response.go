package main

import (
	"github.com/codecrafters-io/kafka-starter-go/app/utils"
)

type (
	KafkaResponse struct {
		CorrelationID int32
		ErrorCode     int16
	}
)

func (kr *KafkaResponse) Bytes() []byte {
	bytes := make([]byte, 6)
	copy(bytes, utils.GetByteArrayFromInt32(kr.CorrelationID))
	copy(bytes[4:], utils.GetByteArrayFromInt16(kr.ErrorCode))
	return bytes
}

func NewKafkaResponse(request *KafkaRequest, errorCode int16) KafkaResponse {
	response := KafkaResponse{
		CorrelationID: request.CorrelationID,
	}
	if errorCode != 0 {
		response.ErrorCode = errorCode
	}
	return response
}

package main

import (
	"github.com/codecrafters-io/kafka-starter-go/app/utils"
)

type (
	KafkaRequest struct {
		APIKey        int16
		APIVersion    int16
		CorrelationID int32
	}
)

func NewKafkaRequest(bytes []byte) KafkaRequest {
	return KafkaRequest{
		APIKey:        utils.ParseInt16FromByteArray(bytes[4:6]),
		APIVersion:    utils.ParseInt16FromByteArray(bytes[6:8]),
		CorrelationID: utils.ParseInt32FromByteArray(bytes[8:12]),
	}
}

func (kr *KafkaRequest) Supported() bool {
	return kr.APIVersion >= 0 && kr.APIVersion <= 4
}

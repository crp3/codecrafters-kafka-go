package main

import (
	"github.com/codecrafters-io/kafka-starter-go/app/utils"
)

type (
	KafkaResponse struct {
		CorrelationID int32
		ErrorCode     int16
		Versions      []Version
		ThrottleTime  int32
	}

	Version struct {
		APIKey        int16
		APIMinVersion int16
		APIMaxVersion int16
	}
)

const (
	APIVersionsAPIKey = 18

	APIVersionsMaxVersion = 4
)

func (v *Version) bytes() ([]byte, int) {
	bytes := make([]byte, 0)
	APIKeyBytes := utils.GetByteArrayFromInt16LittleEndian(v.APIKey)
	APIMinVersionBytes := utils.GetByteArrayFromInt16LittleEndian(v.APIMinVersion)
	APIMaxVersionBytes := utils.GetByteArrayFromInt16LittleEndian(v.APIMaxVersion)
	bytes = append(bytes, APIKeyBytes...)
	bytes = append(bytes, APIMinVersionBytes...)
	bytes = append(bytes, APIMaxVersionBytes...)
	bytes = append(bytes, 0) //tag buffer
	messageLength := len(APIKeyBytes) + len(APIMinVersionBytes) + len(APIMaxVersionBytes) + 1

	return bytes, messageLength

}

func (kr *KafkaResponse) Bytes() []byte {
	correlationIDBytes := utils.GetByteArrayFromInt32(kr.CorrelationID)
	errorCodeBytes := utils.GetByteArrayFromInt16BigEndian(kr.ErrorCode)
	versionsLengthBytes := byte(len(kr.Versions) + 1)
	bytes := make([]byte, 0)
	messageLength := len(correlationIDBytes)
	messageLength += len(errorCodeBytes) + 1 // error + versionsLength
	bytes = append(bytes, correlationIDBytes...)
	bytes = append(bytes, errorCodeBytes...)
	bytes = append(bytes, versionsLengthBytes)
	bytes = append(bytes, byte(0)) // tag buffer
	for _, v := range kr.Versions {
		vBytes, vSize := v.bytes()
		bytes = append(bytes, vBytes...)
		messageLength += vSize
	}
	throttleTimeBytes := utils.GetByteArrayFromInt32(kr.ThrottleTime)
	bytes = append(bytes, throttleTimeBytes...)
	bytes = append(bytes, 0) // tag buffer
	messageLength += len(throttleTimeBytes) + 1
	messageLengthBytes := utils.GetByteArrayFromInt32(int32(messageLength))

	bytes = append(messageLengthBytes, bytes...)

	return bytes
}

func NewKafkaResponse(request *KafkaRequest, errorCode int16) KafkaResponse {
	response := KafkaResponse{
		CorrelationID: request.CorrelationID,
		Versions: []Version{
			{
				APIKey:        18,
				APIMinVersion: 0,
				APIMaxVersion: 4,
			},
			{
				APIKey:        1,
				APIMinVersion: 1,
				APIMaxVersion: 4,
			},
		},
		ErrorCode: errorCode,
	}
	return response
}

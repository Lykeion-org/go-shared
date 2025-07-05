package mq

import (
	"encoding/json"
	"errors"
)

func ConvertMessageToStruct[T any](msg []byte) (T, error) {
	var result T
	err := json.Unmarshal(msg, &result)
	if err != nil {
		return result, errors.New("failed to convert message to struct")
	}
	return result, nil
}

func ConverStructToMessage[T any](structData T) ([]byte, error) {
	var jsonData []byte
	
	jsonData, err := json.Marshal(structData)
	if err != nil {
		return jsonData, errors.New("failed to convert struct to JSON")
	}
	return jsonData, nil
}
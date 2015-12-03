package main

import "encoding/json"

type StructuredLog struct {
	Logger     string
	Msg        string
	Properties map[string]interface{}
}

func parseLog(input []byte) (*StructuredLog, error) {
	// pipeline processing
	data := make(map[string]interface{})
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, err
	}

	// TODO dehardcode those fields
	namespace := data["logger"].(string)
	delete(data, "logger")
	msg := data["msg"].(string)
	delete(data, "msg")

	return &StructuredLog{
		Logger:     namespace,
		Msg:        msg,
		Properties: data,
	}, nil
}

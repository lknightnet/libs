package tg_logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func convert(chtoto any) bytes.Buffer {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(chtoto)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func (sl *ServiceLogger) AddMessage(msg *MessageRequest, timeout time.Duration) (*Response, error) {
	client := http.Client{
		Timeout: timeout,
	}
	buf := convert(msg)
	request, err := http.NewRequest(http.MethodPost, sl.API+"api/message/set", &buf)
	if err != nil {
		return nil, NewError("libs/tg-logger/AddMessage: failed to create request", err)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, NewError("libs/tg-logger/AddMessage: failed to create response", err)
	}
	var loggerResponse Response
	err = json.NewDecoder(response.Body).Decode(&loggerResponse)
	if err != nil {
		log.Println(NewError("libs/tg-logger/AddMessage: failed to decode response", err))
		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(NewError("libs/tg-logger/AddMessage: failed to decode response to string", err))
		}
		fmt.Println(string(data))
	}
	return &loggerResponse, nil
}

package model

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"testing"
	"time"
)

func TestCreateMessage(t *testing.T) {

	//traceId := "65545f194a454634a5e5f6ffc7ecd8a9"
	//requestId := "33c450e9-9f12-6454-1a76-b64590a3c5df"
	//conversionIdSuffix := "85B0E75B24D92FABACCF255483B1921F4451E30AFFEDAADBE8B081D51CB79903"

	// Generate a traceId using uuid4
	traceId := uuid.New().String()

	// Generate a requestId using uuid4
	requestId := uuid.New().String()

	// Generate a conversionIdSuffix using the current timestamp and uuid4
	currentTime := time.Now()
	conversionIdSuffix := fmt.Sprintf("%s%s", currentTime.Format("20060102150405"), uuid.New().String())

	text := "gpt的参数量是多少"
	conversationId := "51D|BingProd|" + conversionIdSuffix
	invocationId := "0"
	participantId := "985155416400228"

	askMessage := CreateMessageOfHonolulu(traceId, requestId, text, conversationId, participantId, invocationId)

	// 打印出结构体的 JSON 表示
	jsonData, err := json.Marshal(askMessage)

	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	log.Println("data length:", len(jsonData))
	fmt.Println(string(jsonData))
}

package main

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go-bing-chat-client/model"
	"log"
)

var emptyMessage = `{}` + "\x1E"
var typeSixMessage = `{"type":6}` + "\x1E"

func main() {
	// 连接的 WebSocket 服务器 URL
	endPoint := "wss://sydney.bing.com/sydney/ChatHub"
	secAccessToken := "KIg2SJOPjuOxHABg53gwmc%2BJXPr3zgTBNpZXTwtTw4bEln27O%2B8binqZSKz%2BoMYtErWi%2FznRdOXhIue23VrxzMCkUkFXsrsMzQHfAIBjP6e8ZNI2%2FXWkWDU8PJTN%2B0udfG%2BRo7hei5c0yRjQ2BVJz8W7qxofm7NRA8fxSt5QdA7loLFcN8IDFIPNjhEw2YQZnp%2BfMxLup1QvMk2tJrl6Da62uCxm1ttDerX1voeNFul5%2FFqG7TIRfuPCLODWeVbulVcq1avim5evG4uspzHog6fJwYyriDIAZHQyGrks02K2OfnuhEHJ%2Fi9ylO9T1TwS1fxHrG6cqyRQuSAeqN4TARw81BYvuHNnfhl1X1wH%2BpVJnZksLhJ3UPVKcLiD%2B8%2Bzj4yKbOFITz7n1H3pnA%2BgpNmP7htizCrDHTZKK%2FBbQhYUm77ft4%2BjBjlO72zuUvn15ew2kvLAGbfFGkZWrWRCmmluzVBJpO%2B%2B6r688%2FzwJwjuem69tff7ZAWcFxoTDBWLxNOTf8fpyqUC5Ntfzif72HMoFXb6CROynuk%2FTVQBr9bAEAf03YPU5%2FgSUT00Vs74mTTxjFyMTCf3iEtEn%2Bl%2BDdNKgfn%2BsXfz%2FmIa56M7cUvkEnvMp6z2fP%2B63FDdSBSILI9qJ3L3cDRlW9Awc9XAb2gjljhPCJyLXSHc2BgqR2%2FY4mtYdk%2F5ogXRLMVei1B0nDF0rNtHeqlDRCabGMvBOunwftXpPHmiU3Qb4nYtC75pRCZqpPpRPcJtKPnxI%2FaSv5qOWVq7aq7kju1buHK9Bal7z1bGRDDxNStIo%2FUlvVo%2BtY1jr9nM6efe4b5AO3CXc9B88riSipDbO%2FY2Iq8A%2F1o8Ls23ZfTHPbpvEeKLPgtOUXJmDwUXhvbnpCdu9CVxCpaQ9VleubkEHAp%2Burr5IdbWlgimDvnVvmljbW2Cm1fnkDCJ7O7m1W48ox7N4APzRXQHni7WuKuo%2BmHy1ozop%2FzMLAEYOyNUsQ7GA8TIQXkBMkeZFXVkjW%2BQpDsGHxi5l%2Fb3oTDKGOtEo2qC4RBHUak0ByzaZjKOGQuXpFmK9RnfYtAjTRPMFrVFxWkKbNvvmQ%2FmdFHHRninbnYwmXRtZg%3D%3D"

	url := endPoint + "?sec_access_token=" + secAccessToken
	//log.Println("url:", url)

	// 创建 WebSocket 连接
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// 发送消息
	protocolMessage := `{"protocol":"json","version":1}` + "\x1E"

	err = c.WriteMessage(websocket.TextMessage, []byte(protocolMessage))
	log.Printf("sned: %s,length: %d", protocolMessage, len(protocolMessage))
	if err != nil {
		log.Fatal("write:", err)
	}

	// 接收消息
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
		}
		log.Printf("recv: %s,length: %d", message, len(message))
		var clientMessage = dispatcherMessage(message)
		if clientMessage == emptyMessage {
			continue
		}
		sendMessage(c, clientMessage)
		if clientMessage == typeSixMessage {
			//发送一个问题
			traceId := uuid.New().String()

			// Generate a requestId using uuid4
			requestId := uuid.New().String()

			// Generate a conversionIdSuffix using the current timestamp and uuid4
			// currentTime := time.Now()
			//conversionIdSuffix := fmt.Sprintf("%s%s", currentTime.Format("20060102150405"), uuid.New().String())
			conversionIdSuffix := "4196A5FE69AD263F14CE489998D0B7929FF6E0A3A1386F13ED290EDA9E0A2574"

			text := "gpt的参数量是多少"
			conversationId := "51D|BingProd|" + conversionIdSuffix
			invocationId := "0"
			participantId := "985155416400228"

			askMessage := model.CreateMessageOfHonolulu(traceId, requestId, text, conversationId, participantId, invocationId)

			// 打印出结构体的 JSON 表示
			jsonData, err := json.Marshal(askMessage)
			if err != nil {
				log.Println(err)
			}
			// 将 \x1E 字符添加到 jsonData 的末尾
			jsonDataWithSuffix := append(jsonData, '\x1E')
			sendByteMessage(c, jsonDataWithSuffix)
		}
	}

}
func sendByteMessage(c *websocket.Conn, message []byte) {
	log.Printf("sned: %s,length: %d", string(message), len(message))
	var err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal("write:", err)
	}
}

func sendMessage(c *websocket.Conn, message string) {
	log.Printf("sned: %s,length: %d", message, len(message))
	var err = c.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Fatal("write:", err)
	}
}

func dispatcherMessage(message []byte) string {

	defaultMessageBytes := []byte(emptyMessage)

	if bytes.Equal(message, defaultMessageBytes) {
		return typeSixMessage
	} else {
		return emptyMessage
	}
}

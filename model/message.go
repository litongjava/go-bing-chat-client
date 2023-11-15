package model

import "time"

// Argument 包含与消息相关的所有参数。
type Argument struct {
	Source                         string      `json:"source"`
	OptionsSets                    []string    `json:"optionsSets"`
	AllowedMessageTypes            []string    `json:"allowedMessageTypes"`
	SliceIds                       []string    `json:"sliceIds"`
	Verbosity                      string      `json:"verbosity"`
	Scenario                       string      `json:"scenario"`
	Plugins                        []string    `json:"plugins"`
	TraceId                        string      `json:"traceId"`
	ConversationHistoryOptionsSets []string    `json:"conversationHistoryOptionsSets"`
	IsStartOfSession               bool        `json:"isStartOfSession"`
	RequestId                      string      `json:"requestId"`
	Message                        Message     `json:"message"`
	Tone                           string      `json:"tone"`
	SpokenTextMode                 string      `json:"spokenTextMode"`
	ConversationId                 string      `json:"conversationId"`
	Participant                    Participant `json:"participant"`
}

// Message 描述了消息的具体内容。
type Message struct {
	Locale        string         `json:"locale"`
	Market        string         `json:"market"`
	Region        string         `json:"region"`
	Location      string         `json:"location"`
	LocationHints []LocationHint `json:"locationHints"`
	UserIpAddress string         `json:"userIpAddress"`
	Timestamp     string         `json:"timestamp"`
	Author        string         `json:"author"`
	InputMethod   string         `json:"inputMethod"`
	Text          string         `json:"text"`
	MessageType   string         `json:"messageType"`
	RequestId     string         `json:"requestId"`
	MessageId     string         `json:"messageId"`
}

// LocationHint 描述了消息中提供的地理位置提示。
type LocationHint struct {
	SourceType               int     `json:"SourceType"`
	RegionType               int     `json:"RegionType"`
	Center                   Center  `json:"Center"`
	Radius                   int     `json:"Radius"`
	Name                     string  `json:"Name"`
	Accuracy                 int     `json:"Accuracy"`
	FDConfidence             float64 `json:"FDConfidence"`
	CountryName              string  `json:"CountryName"`
	CountryConfidence        int     `json:"CountryConfidence"`
	Admin1Name               string  `json:"Admin1Name"`
	PopulatedPlaceName       string  `json:"PopulatedPlaceName"`
	PopulatedPlaceConfidence int     `json:"PopulatedPlaceConfidence"`
	PostCodeName             string  `json:"PostCodeName"`
	UtcOffset                int     `json:"UtcOffset"`
	Dma                      int     `json:"Dma"`
}

// Center 描述了地理位置的经纬度。
type Center struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

// Participant 描述了参与者的信息。
type Participant struct {
	ID string `json:"id"`
}

// AskMessage 是最外层的JSON结构体。
type AskMessage struct {
	Arguments    []Argument `json:"arguments"`
	InvocationId string     `json:"invocationId"`
	Target       string     `json:"target"`
	Type         int        `json:"type"`
}

var optionsSets = []string{
	"nlu_direct_response_filter",
	"deepleo",
	"disable_emoji_spoken_text",
	"responsible_ai_policy_235",
	"enablemm",
	"dv3sugg",
	"iyxapbing",
	"iycapbing",
	"h3precise",
	"clgalileo",
	"gencontentv3",
}

var allowedMessageTypes = []string{
	"ActionRequest",
	"Chat",
	"ConfirmationCard",
	"Context",
	"InternalSearchQuery",
	"InternalSearchResult",
	"Disengaged",
	"InternalLoaderMessage",
	"InvokeAction",
	"Progress",
	"RenderCardRequest",
	"RenderContentRequest",
	"AdsQuery",
	"SemanticSerp",
	"GenerateContentQuery",
	"SearchQuery",
}
var sliceIds = []string{
	"gbaa",
	"gba",
	"gbapa",
	"bggpt",
	"caccnctat1",
	"errorhandle",
	"styleoffall",
	"rwcf",
	"713logprobss0",
	"1024noarchive",
	"119wcphi",
	"1113gldcl1s0",
	"10312tlocrets0",
	"plgbd2",
	"727nrprdrt7",
	"727nrprdrt3",
}
var conversationHistoryOptionsSets = []string{
	"autosave",
	"savemem",
	"uprofupd",
	"uprofgen",
}

func CreateMessageOfHonolulu(traceId string, requestId string, text string, conversationId string, participantId string, invocationId string) AskMessage {
	userIpAddress := "66.75.89.81"
	locale := "en-US"
	region := "US"

	location := "lat:47.639557;long:-122.128159;re=1000m;"

	center := Center{
		Latitude:  21.283700942993164,
		Longitude: -157.7978057861328,
	}

	locationHint := []LocationHint{
		{
			SourceType:               1,
			RegionType:               2,
			Center:                   center,
			Radius:                   24902,
			Name:                     "Honolulu, Hawaii",
			Accuracy:                 24902,
			FDConfidence:             0.8999999761581421,
			CountryName:              "United States",
			CountryConfidence:        9,
			Admin1Name:               "Hawaii",
			PopulatedPlaceName:       "Honolulu",
			PopulatedPlaceConfidence: 9,
			PostCodeName:             "96816",
			UtcOffset:                -10,
			Dma:                      744,
		},
	}

	// 创建时间对象，代表 JSON 字符串中的时间戳
	// timestamp, _ := time.Parse(time.RFC3339, "2023-11-14T20:03:00-10:00")
	timestamp := time.Now()
	askMessage := createMessage1(traceId, requestId, locale, region, location, locationHint, userIpAddress, timestamp, text, conversationId, participantId, invocationId)
	return askMessage
}

func createMessage1(traceId string, requestId string, locale string, region string, location string, locationHint []LocationHint, userIpAddress string, timestamp time.Time, text string, conversationId string, participantId string, invocationId string) AskMessage {
	// 手动创建 AskMessage 对象，并为字段赋值
	askMessage := createMessage0(optionsSets, allowedMessageTypes, sliceIds, conversationHistoryOptionsSets, traceId, requestId, locale, region, location, locationHint, userIpAddress, timestamp, text, conversationId, participantId, invocationId)
	return askMessage
}

func createMessage0(optionsSets []string, allowedMessageTypes []string, sliceIds []string, conversationHistoryOptionsSets []string, traceId string, requestId string, locale string, region string, location string, locationHint []LocationHint, userIpAddress string, timestamp time.Time, text string, conversationId string, participantId string, invocationId string) AskMessage {
	formattedTimestamp := timestamp.Format("2006-01-02T15:04:05-07:00")
	askMessage := AskMessage{
		Arguments: []Argument{
			{
				Source:                         "cib",
				OptionsSets:                    optionsSets,
				AllowedMessageTypes:            allowedMessageTypes,
				SliceIds:                       sliceIds,
				Verbosity:                      "verbose",
				Scenario:                       "SERP",
				Plugins:                        []string{}, // 空的插件列表
				TraceId:                        traceId,
				ConversationHistoryOptionsSets: conversationHistoryOptionsSets,
				IsStartOfSession:               false,
				RequestId:                      requestId,
				Message: Message{
					Locale:        locale,
					Market:        locale,
					Region:        region,
					Location:      location,
					LocationHints: locationHint,
					UserIpAddress: userIpAddress,
					Timestamp:     formattedTimestamp,
					Author:        "user",
					InputMethod:   "Keyboard",
					Text:          text,
					MessageType:   "Chat",
					RequestId:     requestId,
					MessageId:     requestId,
				},
				Tone:           "Precise",
				SpokenTextMode: "None",
				ConversationId: conversationId,
				Participant: Participant{
					ID: participantId,
				},
			},
		},
		InvocationId: invocationId,
		Target:       "chat",
		Type:         4,
	}
	return askMessage
}

package model

import (
	"encoding/json"
	"testing"
)

// 假设这是您从上面的JSON中获取的字符串。
const jsonString = `{"arguments":[{"source":"cib","optionsSets":["nlu_direct_response_filter","deepleo","disable_emoji_spoken_text","responsible_ai_policy_235","enablemm","dv3sugg","iyxapbing","iycapbing","h3precise","clgalileo","gencontentv3"],"allowedMessageTypes":["ActionRequest","Chat","ConfirmationCard","Context","InternalSearchQuery","InternalSearchResult","Disengaged","InternalLoaderMessage","InvokeAction","Progress","RenderCardRequest","RenderContentRequest","AdsQuery","SemanticSerp","GenerateContentQuery","SearchQuery"],"sliceIds":["gbaa","gba","gbapa","bggpt","caccnctat1","errorhandle","styleoffall","rwcf","713logprobss0","1024noarchive","119wcphi","1113gldcl1s0","10312tlocrets0","plgbd2","727nrprdrt7","727nrprdrt3"],"verbosity":"verbose","scenario":"SERP","plugins":[],"traceId":"65545f194a454634a5e5f6ffc7ecd8a9","conversationHistoryOptionsSets":["autosave","savemem","uprofupd","uprofgen"],"isStartOfSession":false,"requestId":"33c450e9-9f12-6454-1a76-b64590a3c5df","message":{"locale":"en-US","market":"en-US","region":"US","location":"lat:47.639557;long:-122.128159;re=1000m;","locationHints":[{"SourceType":1,"RegionType":2,"Center":{"Latitude":21.283700942993164,"Longitude":-157.7978057861328},"Radius":24902,"Name":"Honolulu, Hawaii","Accuracy":24902,"FDConfidence":0.8999999761581421,"CountryName":"United States","CountryConfidence":9,"Admin1Name":"Hawaii","PopulatedPlaceName":"Honolulu","PopulatedPlaceConfidence":9,"PostCodeName":"96816","UtcOffset":-10,"Dma":744}],"userIpAddress":"66.75.89.81","timestamp":"2023-11-14T20:03:00-10:00","author":"user","inputMethod":"Keyboard","text":"gpt的参数量是多少","messageType":"Chat","requestId":"33c450e9-9f12-6454-1a76-b64590a3c5df","messageId":"33c450e9-9f12-6454-1a76-b64590a3c5df"},"tone":"Precise","spokenTextMode":"None","conversationId":"51D|BingProd|85B0E75B24D92FABACCF255483B1921F4451E30AFFEDAADBE8B081D51CB79903","participant":{"id":"985155416400228"}}],"invocationId":"4","target":"chat","type":4}`

func TestJSONParsingAndConversion(t *testing.T) {
	// 解析 JSON 字符串到 model
	var model AskMessage
	err := json.Unmarshal([]byte(jsonString), &model)
	if err != nil {
		t.Fatalf("Unmarshal json failed: %v", err)
	}

	// 验证model的值
	if model.Arguments[0].Source != "cib" {
		t.Errorf("Expected source to be 'cib', got '%s'", model.Arguments[0].Source)
	}

	// 将 model 转换回 JSON 字符串
	jsonBytes, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("Could not marshal model back to json: %v", err)
	}

	// 将 JSON bytes 转换成字符串
	jsonStringFromModel := string(jsonBytes)

	// 输出转换回的 JSON 字符串，以便查看
	t.Logf("Model converted back to JSON string: %s", jsonStringFromModel)
}

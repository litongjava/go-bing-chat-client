package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateConversation(t *testing.T) {

	cookie := "MUID=1145CD7AA8A4645E149ADEB0A9E465F0; MUIDB=1145CD7AA8A4645E149ADEB0A9E465F0; _EDGE_V=1; SRCHD=AF; SRCHUID=V; USRLOC=HS; ANON=A; NAP=V; PPLState=1; KievRPSSecAuth=FAByBBRaTOJILtFsMkpLVWSG6AN6C/svRwNmAAAEgAAACI80btDHZElKMATBen/IwIcgwUB/rfvPzvamnzDBLNRoYtKUHXdVPLPtMTVyDqlR149pTbH1DeiwxJu7UCtJmZL376leBu2y8CeGgJzpa3ql2LXYauAZggwEZN66WYz1gJLMxpeRV8/dv0ex6lq3rNpFJYxlobbNkbyQakWdQguZ5F3PCY3ule4Oi8MYH00yQMG1941k/0KCVV16G/mdWRBUWzuazO0K93HeeFKjns8fiXx97aitzuswNT7m123ld1QCCsB9tVSCjvvI7Lnhbd8K/ydOvygX/to0qfGg1kMwgctwEQ/c2mxb7G+25IqeXJvz11nC27uCNYVQP1urcc0AFp7YPm2nScTHvXtvT7Z3slt/44R7rTGqJ5ocmEPYfC4XyfOIU/R9DUHESR5PdPEOPAMXgLo9TnqEwKLrZMfKGdDf7ANcfk8JIuPW7esPZFJ8hLx7B5wkhWCdU621MHtMNGpI0tpJODu9LpaWO/BRmFw6vc1s+yx3l+36x+lCtZjEmv1hLCchIgEakK/LWELGUHG+rVSNn1idxyjmbrh/GUvBIla+wrvJhxOgeX1FtmmhUEQ6ajLFGCedLLv50ZeTNu7QqZt77JAEndM6rjhIr8u8qwqWWv18jkmhQKigZ3gfyevrkZh6PIW0MiJDlPEn7f/cw1IQU0C08WFYXkbMswmVsubPAOvXwTxvF4IMu6n+pAjBaLRqUTG2exTomX47uOo3DrHXfOb8ll85THutJhXGOtitYv7U1mz9b/haJ/r6NQjeqNmnzIo+5cR8/Zwrse10M0BSFkuYe+0RZCh13/cVU8TKwU0DVbhp1QgD7r7kOTL0nl8A9JOAQ85Nu+FK1bmk3otAsHkkftWdn345NCFkhp0qZXJFu5c45BxZ/rnl5n+EdFAnf6Jk3Qkb8OLjZqlC38sDDW6xsBNFf5Jo+02TDvDrlNMdn2WHYPfXRNPQGA58242pxC74jzvFQT5j5tFdL1NohqJrSp2PAtqVkTPHeC8LoMQGKshMjidmIHT+k7iRxtY9bOQqvub7bwsTAiQn7Tm8Q23jyC7caPOzZB0/w2lvWu83X1ug39TjH8ikHsSvsdV6j3eJl+WhaOm/2qegpNaYja70s3R93/izHw8gAEKXZk5uOe7PrXeg8ZDMgctLe0lNWhwuae6q3Z9327id4buQmqBAH2odqpAu8FgKl+I6DJZZ8KFgUw7IkUM6FVlks47cm3dTk++LBZMXdZmsVpmeQgwnPyvqFig24cvsti1tlGmXL1fXC4B519SwaETbh9Az3CFWz6/9BFuquPtT4dYqyqsCFXPUHxTWGqYKKDnkYDHP/mhIZ8FL/3B2xzMTW34wtwO5bUZoQSvQhfp6CnckRcWyALwpHQIPGLbI4Mp1AACC2PFEpiG+qSs0zECWRkSVlH+ETQZnsFirYhMfb4E/1aliFADVovI1KIJcLB4EBC4JPICyQcpIUw; _U=1aDN7NFyNRDx1r4M3EORGGMQLmg2Dr8Zgt78Mi8I6n34Z3fzlwITJEytR-xqbiChzdOa_PY5J50bX1Cj_MhVjj7dxO-8jKIjP9TgirxFlzATeyM0fTpnba2YwC1qNFliPZmJwba9mD4S5dlBPyz9dxYsbMt4cZNH8a6Ksqx7h6VnhMXavo2YEgcrVYLcp3rSqzBdY5JhEn37QwOcnLirkPg; WLID=OL2S9PIXlxnn5oVst5/VMOOM1/s3+GQgZPC4sYg0qBDpnboyVgGjLRIJPQQ+oOYixJdsKxv4fl/fOq3368OYRRbcxmoTBQ9dSgSe0vau8GE; SnrOvr=X; _UR=; _EDGE_S=SID; WLS=C; _Rwho=u; _SS=SID; ABDEF=V; SRCHUSR=DOB; ipv6=hit; cct=cGX2o-UKDtdOhfYvWSi2ihWdxhiZq3xBBtlTnRCcesoK9G8qYzJdyTsGoFX-HndT_TvxhAABDgCIxb1U5U0mJw; _RwBf=r; SRCHHPGUSR=SRCHLANG"

	url := "https://www.bing.com/turing/conversation/create?bundleVersion=1.1359.3"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "www.bing.com")
	req.Header.Add("sec-ch-ua-arch", "\"x86\"")
	req.Header.Add("sec-ch-ua-bitness", "\"64\"")
	req.Header.Add("sec-ch-ua-full-version", "\"119.0.6045.124\"")
	req.Header.Add("sec-ch-ua-full-version-list", "\"Google Chrome\";v=\"119.0.6045.124\", \"Chromium\";v=\"119.0.6045.124\", \"Not?A_Brand\";v=\"24.0.0.0\"")
	req.Header.Add("sec-ch-ua-model", "\"\"")
	req.Header.Add("sec-ch-ua-platform-version", "\"15.0.0\"")
	req.Header.Add("x-ms-client-request-id", "820c2b85-5d5f-473c-a00e-88409e23e244")
	req.Header.Add("x-ms-useragent", "azsdk-js-api-client-factory/1.0.0-beta.1 core-rest-pipeline/1.12.0 OS/Windows")
	req.Header.Add("Cookie", cookie)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

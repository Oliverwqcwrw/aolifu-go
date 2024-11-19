package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	xhs()
}

func xhs() {
	url := "https://edith.xiaohongshu.com/api/sns/web/v1/homefeed"
	method := "POST"

	payload := `{"cursor_score":"","num":18,"refresh_type":1,"note_index":35,"unread_begin_note_id":"","unread_end_note_id":"","unread_note_count":0,"category":"homefeed_recommend","search_key":"","need_num":8,"image_formats":["jpg","webp","avif"],"need_filter_image":false}`

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("cookie", "abRequestId=1f62cf94-b816-50e9-93e6-583389cfb770; a1=192e0f41eb28rl2s06t7jprodoz8mem7zinadmvn730000385524; webId=05a58a300f8dfe475e293030645e5ec6; gid=yjJd8i4JK0uSyjJd8i4yd36hDJYF1JM8K9WVIIWUFlflu0q8Y6d6Wf888qY22J4882i8W2fK; webBuild=4.43.0; xsecappid=xhs-pc-web; web_session=030037a0105b1130c69b5d0e44204a1302bc69; websectiga=3633fe24d49c7dd0eb923edc8205740f10fdb18b25d424d2a2322c6196d2a4ad; sec_poison_id=d43ef9b0-ba26-41da-b2cd-f5a6db8d1a7c; acw_tc=0a0b142517320065135623150e3b8e821d62ec740810dbd641897fd3bd610a; unread={%22ub%22:%22671dc7a200000000260362c4%22%2C%22ue%22:%226716931b000000001402dd15%22%2C%22uc%22:31}")
	req.Header.Add("origin", "https://www.xiaohongshu.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://www.xiaohongshu.com/")
	req.Header.Add("sec-ch-ua", `"Chromium";v="130", "Google Chrome";v="130", "Not?A_Brand";v="99"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"macOS"`)
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
	req.Header.Add("x-b3-traceid", "30e5cc5c6d787949")
	req.Header.Add("x-s", `XYW_eyJzaWduU3ZuIjoiNTUiLCJzaWduVHlwZSI6IngyIiwiYXBwSWQiOiJ4aHMtcGMtd2ViIiwic2lnblZlcnNpb24iOiIxIiwicGF5bG9hZCI6ImJhMTA4MzI4MzgwMzA4MWVhNDBlMDdlMzllMTJhMDU5NzIxNzQ3ZTI5MDljODg3ZmY0OTVjN2Q3YTJkYTAwNWZkMDQwNDJiZWZjYzk3NDZjMDkxN2E5YTZjY2Q2ZmE4YzdhMGM5NmEzM2NhMmU1NTcxZWQ1YjE1ZTJlYmRhMDFlOTExMTAxMDUyNGQ3ZGFiNmYyZTEwYzcxNGJkYjg2NjUzYTFkNTBkNTNlZjY4MzY3YTc1OGI0YTM2ZjA3MmI0ZmFmMGE2ODcyMWJkNGUxOTk5NTU3MmUzMDViYjJhYWE4YTUxMzIyZGNiNTQ1MjFkMGRhM2UzNjMyOWUzYWVhNWM5NzRjYzlkMjRlNTc1ZDg5NzdhMDRkOTliMTAwMTZjYmU2ZmViMGNmMWVjMGZmOWI2YzcxMjA4MDFlNzkwMDk1OTM0MmUxMmE4YTBmMWY2MmVkN2IxOTIyNjQ4NzhjNzA1YmJiMDcyNWQxOWE0ZjU4NDliODhmZjNjNDQwNDkxZWU4ZTA4ZDI1MGEzNjAyNjI5MDBkZTUwZjJjZTMwMWNjIn0=`)
	req.Header.Add("x-s-common", `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`)
	req.Header.Add("x-t", "1732006625939")
	req.Header.Add("x-xray-traceid", "c9a1d1e143afd7173016ad07747eae42")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	// 打印返回数据中json的data的值
	fmt.Println("Response Status:", res.Body)
	// 读取body，将它转为字符串可读
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response Body:", string(body))
}

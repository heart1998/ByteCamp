package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DictRequest struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

type DictResponse struct {
	Words []struct {
		Source  int    `json:"source"`
		Text    string `json:"text"`
		PosList []struct {
			Type      int `json:"type"`
			Phonetics []struct {
				Type int    `json:"type"`
				Text string `json:"text"`
			} `json:"phonetics"`
			Explanations []struct {
				Text     string `json:"text"`
				Examples []struct {
					Type      int `json:"type"`
					Sentences []struct {
						Text      string `json:"text"`
						TransText string `json:"trans_text"`
					} `json:"sentences"`
				} `json:"examples"`
				Synonyms []interface{} `json:"synonyms"`
			} `json:"explanations"`
			Relevancys []interface{} `json:"relevancys"`
		} `json:"pos_list"`
	} `json:"words"`
	Phrases  []interface{} `json:"phrases"`
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

func query(word string) {
	client := &http.Client{}
	request := DictRequest{Text: word, Language: "en"}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "http://translate.volcengine.com/web/dict/match/v1/?msToken=&X-Bogus=DFSzswVLQDcCwW3KSWM/bVXAIQ2k&_signature=_02B4Z6wo00001nuc4VwAAIDDGJYhNGnFuAZ7nOXAAPyfWzrohK8qd0PAYqSOXT7-tj0M5JcALrjIIqFcEy5BN24k6fvK.2QZF1RqxZfI.leKAHotTOSC3er71akm60fvgg-xsKjkPFLcsk1sdd", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "i18next=zh-CN; ttcid=a98e5300d2cb42a7aa0c170939ecb3d334; tt_scid=JntuGncA7QiybZp3XW94iR-jpGACp0wR4ealtIAQO8ZQDxfxRruq4xoo4ovP7x5Rcdfd; s_v_web_id=verify_acd72d47959c637ea72b09d1035fcaf7; _tea_utm_cache_2018=undefined")
	req.Header.Set("Origin", "http://translate.volcengine.com")
	req.Header.Set("Referer", "http://translate.volcengine.com/translate?category=&home_language=zh&source_language=detect&target_language=zh&text=hello")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	for _, word := range dictResponse.Words {
		for _, poslist := range word.PosList {
			for _, explanation := range poslist.Explanations {
				fmt.Println(explanation.Text)
			}
		}
	}
}

func main() {
	var word string
	fmt.Println("请输入你的单词:")
	fmt.Scanf("%s", &word)
	query(word)
}

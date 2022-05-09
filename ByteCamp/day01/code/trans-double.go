package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DictRequest1 struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

type DictResponse1 struct {
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

func query1(word string, Engine chan<- int) {
	client := &http.Client{}
	request := DictRequest1{Text: word, Language: "en"}
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
	var dictResponse DictResponse1
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	Engine <- 1
	for _, word := range dictResponse.Words {
		for _, poslist := range word.PosList {
			for _, explanation := range poslist.Explanations {
				fmt.Println(explanation.Text)
			}
		}
	}
}

type DictRequest2 struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponse2 struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query2(word string, Engine chan<- int) {
	client := &http.Client{}
	request := DictRequest2{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
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
	var dictResponse DictResponse2
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	Engine <- 2
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	var word string
	fmt.Scanf("%s", &word)
	Engine := make(chan int)
	go query1(word, Engine)
	go query2(word, Engine)
	x := <-Engine
	//if you don't wanna see which trasnslater is faster, you can just send a struct{} to Engine
	switch x {
	case 2:
		fmt.Printf("Caiyun Translation is faster")
	case 1:
		fmt.Printf("Volcano Translation is faster")
	}
}

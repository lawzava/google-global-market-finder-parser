package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Resp struct {
	Result struct {
		Data []struct {
			Text struct {
				Keyword string `json:"1"`
				Extra   int    `json:"2"`
			} `json:"1"`
			Value struct {
				Searches    string  `json:"1"`
				Currency    string  `json:"2"`
				Bid         float64 `json:"3"`
				Opportunity float64 `json:"4"`
				Competition float64 `json:"5"`
			}
		} `json:"1"`
	} `json:"result"`
}

func main() {
	//var Response Resp

	CaptchaToken := "AAJi84FEGD4JGcvhLnqvy9-HbRLATkVip5WDirI3rcTVXQBo15KeiL-sUqwoQfYz-HpdpqsxGd0Ti-pBE8c1iW1uGxCYnvzRc0f6vi41or541oxyKeGg2KTDtAsgaX6QTsRCTUUktqMCaPfb2RiJ8V4OIURrGoGe5_ourGtCK6T8HhSecaB9704T_E7iXHFkHhy9tT-MgwrlLOEm3t30haM4lxd-Bx1PDUNBYVgom9vHe9oecI3ruqKgZI0WxQDFEF6e1pD6eYyWyijWuIsvja5HwYcltR9EAQ6TfmresntFm4RZ_1OPdqPty3o26abXVQ2SY6S2vjclw8_xvwBG3PqsTiNz5Dgx4A"

	rawURL := "http://translate.google.com"
	resource := "/globalmarketfinder/g/search"

	Cookie := fmt.Sprintf("CONSENT=YES+LT.en+20150726-13-0; SID=DQXn8Fo0yaCPYjS5EJaajR1i5zMeoyhGpCj9bM-jPGl83ZKmazL4mhUO3VuQHGBM-rZVgw.; HSID=AF2u56Jy8FqfA_ibT; APISID=CgGgR9ZzCzFD94Ho/Aqmdd8skbMQeS2mbC; NID=110=h1rGQyTUpBxW2k-wbE1RzYyiVuca88P4U8LOwz0hUn9jPjT3WPHZileJ_invd5Ygxhm02N9kUuW1wDWlM0ZfkzLnE6uSiNO3MSNXPmgXA2QDmHk2e13X1o8gDxWCmfDW1316u2_FAIOOolTFNadUDN4DJuQLbpaTAxqnYlt2qSWalUz8XO86_ppAgMDTjkNWYiSxHf7ECmmcvtiKf_BxxeJ2Uuf1F_-ZdGYFzWEvy6JtQdqZ; CaptchaToken=%s; SIDCC=AA248bdD7v2RGshFMGP5Yc-Cr-7atiaoy_QgFpgmQoo-dGborOMjdhwnFsAtFOHMVIS5P5mBOEpRl3u0EIbR'", CaptchaToken)

	data := url.Values{}
	data.Add("method", "expand")
	data.Add("params", "{\"1\":[{\"1\":\"weather\",\"2\":1}],\"2\":\"en\",\"3\":\"US\",\"4\":\"USD\"}")

	u, _ := url.ParseRequestURI(rawURL)
	u.Path = resource
	u.RawQuery = data.Encode()
	//	urlStr := fmt.Sprintf("%v", u)

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", rawURL+resource, nil)
	if err != nil {
		log.Print(err)
	}

	req.PostForm = data
	req.Header.Add("Cookie", Cookie)
	req.Header.Add("Origin", "http://translate.google.com")
	//req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8,lt;q=0.6,nl;q=0.4,fr;q=0.2,ru;q=0.2,de;q=0.2")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36")
	req.Header.Add("Content-Type", "application/javascript; charset=UTF-8")
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-GWT-Module-Base", "http://translate.google.com/globalmarketfinder/g/")
	req.Header.Add("X-GWT-Permutation", "4ED82CD23774DB3736F80066F98DF5CE")
	req.Header.Add("Referer", "http://translate.google.com/globalmarketfinder/g/index.html?locale=en")
	req.Header.Add("Connection", "kepp-alive")
	req.Header.Add("DNT", "1")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}

	rsp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
	}

	res.Body.Close()

	fmt.Println(res.Status)
	fmt.Println(string(rsp))
	fmt.Println(req)
	fmt.Println(data)
	/*
		e := json.Unmarshal([]byte(resp), &Response)
		if e != nil {
			log.Print(e)
		}*/
	//fmt.Println(string(resp[:]))
	//fmt.Println(req)

	// if flickResponse.Code == 200
}

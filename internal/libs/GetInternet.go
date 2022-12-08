package libs

import (
	"ParsingAcatOnline/internal/ChangeData"
	"ParsingAcatOnline/internal/constData"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type URLInputGet struct {
	URLIn string
}

func (s URLInputGet) Geter() (doc *goquery.Document) {

	body := getReq(s.URLIn)
	/*	res, err := http.Get(s.URLIn)
		if err != nil {
			fmt.Println(s.URLIn)
			fmt.Println(err)
			fmt.Println(`Не считали данные`)
			return doc
		}
		defer res.Body.Close()*/

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println(`Не верные данные`)
		log.Fatal(err)
	}
	return doc
}

func getReq(sURL string) (body []byte) {
	time.Sleep(constData.TimeOutRequest * time.Second)
	sURL = ChangeData.Replacer(sURL)

	for rep := 0; rep < constData.ReplyGetRequest; rep++ {

		if rep > 0 {
			fmt.Printf("Ошибка сети. Запуск цил повтора. Попытка номер - %v \nURL - %v \n", rep, sURL)
		}

		req, err := http.NewRequest("GET", sURL, nil)
		if err != nil {
			fmt.Println(err)
			fmt.Println(`Не Создали запрос`)
			time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
			continue
		}

		req.Header.Set("User-Agent", constData.UserAgert)
		req.Header.Set("Cache-Control", "no-cache")

		fmt.Println(`*`)

		client := &http.Client{Timeout: (time.Second * constData.TimeOutRequestGeter)}
		res, err := client.Do(req)
		if err != nil {
			//	fmt.Println(sURL)
			fmt.Println(err)
			fmt.Println(`Не считали данные`)
			time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
			continue
		}

		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(`Error reading body. `)
			fmt.Println(err)
			time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
			res.Body.Close()
			continue
		}
		res.Body.Close()
		break
	}
	return body
}

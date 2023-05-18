package libs

import (
	"bytes"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/MakeConfiger"

	"context"
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/ChangeData"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/constData"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type URLInputGet struct {
	URLIn  string
	Config *MakeConfiger.Config
}

func (s URLInputGet) Geter() (doc *goquery.Document) {

	s.Config.URLReq <- s.URLIn

	//вот тут мы должны вернуть массив байт
	//	body := getReq(s.URLIn, s.Config)

	body := <-s.Config.DataOUTReq

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println(`Не верные данные`)
		log.Fatal(err)
	}
	return doc
}

func getReq(sURL string, cfg *MakeConfiger.Config) (body []byte) {
	time.Sleep(constData.TimeOutRequest * time.Second)
	sURL = ChangeData.Replacer(sURL)

	var res string

	for rep := 0; rep < constData.ReplyGetRequest; rep++ {

		if rep > 0 {
			fmt.Printf("Ошибка сети. Запуск цил повтора. Попытка номер - %v \nURL - %v \n", rep, sURL)
		}

		var err1 error
		func() {

			opts := append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.UserDataDir("C:\\temp\\asat_temp"),
				chromedp.Flag("headless", cfg.Debug),
				chromedp.Flag("disable-gpu", false),
				chromedp.Flag("enable-automation", false),
				chromedp.Flag("disable-extensions", true),
			)

			allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
			defer cancel()

			// create context
			ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
			defer cancel()

			// run task list

			err := chromedp.Run(ctx) /*	chromedp.Navigate(sURL),
				chromedp.WaitReady(`#oemwidget_oem_catalog`),
				chromedp.Sleep(5*time.Second),
				chromedp.Outp
			*/

			//	scrapIt(sURL, &res),

			if err != nil {
				err1 = err
			}

		}()

		fmt.Println(`++`)
		time.Sleep(10 * time.Second)

		fmt.Println(`*`)
		if err1 != nil {
			continue
		}

		body = []byte(res)
		break
		//Правим тут

		/*
			req, err := http.NewRequest("GET", sURL, nil)
			if err != nil {
				fmt.Println(err)
				fmt.Println(`Не Создали запрос`)
				time.Sleep(constData.ReplyGetRequestTimeOut * time.Second)
				continue
			}

			req.Header.Set("User-Agent", constData.UserAgert)
			req.Header.Set("Cache-Control", "no-cache")
		*/

		/*
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

		*/
	}

	return body
}

func scrapRun() (body []byte) {

	return
}

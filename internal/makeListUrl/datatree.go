package makeListUrl

import (
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/ChangeData"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/MakeConfiger"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/libs"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/scraping"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
	"time"
)

func dataTree(sURL string, textI string, cfg *MakeConfiger.Config) {

	getQ := libs.URLInputGet{URLIn: sURL, Config: cfg}
	doc := getQ.Geter()

	//var doc *goquery.Document
	//copy(docI, doc)
	var bool1 bool
	// рекурсивно вызываться до получения страницы списка с товарами.
	doc.Find("div.fiat_units").Each(func(i int, s *goquery.Selection) {
		bool1 = true
		s.Find("span.fiat_unit").Each(func(i2 int, l1 *goquery.Selection) {
			urlOut, _ := l1.Parent().Attr("href")
			text := textI + " | " + l1.Text()

			var URLRequestItem1 string

			str2 := "ea.acat.online/catalogs/"
			if strings.Contains(sURL, str2) {
				URLRequestItem1 = sURL + "/." + ChangeData.Replacer(urlOut)
				//		fmt.Println(URLRequestItem1 + " | " + text)
			} else {
				// Разбираем URL
				u, err := url.Parse(sURL)
				if err != nil {
					panic(err)
				}

				// Получаем адрес домена и протокол подключения
				domain := fmt.Sprintf("%s://%s", u.Scheme, u.Hostname())

				URLRequestItem1 = domain + ChangeData.Replacer(urlOut)
			}

			dataTree(URLRequestItem1, text, cfg)
			URLRequestItem1 = ""
			//		time.Sleep(5 * time.Second)

		})
	})

	if !bool1 {

		/**/ // Тут конечная точка, в которой нужно отправить данные на парсинг
		temp1 := scraping.DataSc{
			Doc: doc,
			//	TextBefore: textI,
			//	URLScriping: URLRequestItem1,
			Text2File: textI,
			TimeStart: time.Now(),
		}
		temp1.Scriping()

	}

}

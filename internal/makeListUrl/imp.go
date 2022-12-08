package makeListUrl

import (
	"ParsingAcatOnline/internal/constData"
	"ParsingAcatOnline/internal/endApp"
	"ParsingAcatOnline/internal/libs"
	"ParsingAcatOnline/internal/scraping"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"time"
)

//var counter int

func MakeList(s []string) {

	_, err := url.ParseRequestURI(s[1])
	if err != nil {
		fmt.Println(`Ошибка в файле импорта.`)
		fmt.Println(s[1])
		fmt.Println(`Не является URL`)
		endApp.Fin()
	}

	textBefore := s[0]
	if len(textBefore) < 1 {
		fmt.Println(`Нет дополнительного значения для URL`)
		fmt.Println(s[1])
		endApp.Fin()
	}

	getQ := libs.URLInputGet{URLIn: s[1]}
	doc := getQ.Geter()

	doc.Find("ul.levamCatalogTree").Each(func(i int, s *goquery.Selection) {
		s.Find("a.ml-3").Each(func(i3 int, l1 *goquery.Selection) {
			urlOut, _ := l1.Attr("href")
			text := ""
			for iT1 := 0; iT1 < 8; iT1++ {
				l1 = l1.Parent()
				title_1, iT1B := l1.Attr("data-name")
				if iT1B {
					if iT1 == 0 {
						text = title_1
					} else {
						text = title_1 + constData.SepatatorName + text
					}
				}
			}

			//		replacer := strings.NewReplacer(" ", "%20")
			//			urlOut = replacer.Replace(urlOut)

			temp1 := scraping.DataSc{
				TextBefore:  textBefore,
				URLScriping: urlOut,
				Text2File:   text,
				TimeStart:   time.Now(),
			}

			temp1.Scriping()

		})

	})

}

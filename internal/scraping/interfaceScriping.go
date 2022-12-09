package scraping

import (
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/constData"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/writeOutFiles"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

type DataSc struct {
	Doc *goquery.Document
	//	TextBefore string
	//	URLScriping string
	Text2File string
	TimeStart time.Time
}

func (data DataSc) Scriping() {
	parsingDataFromURL(data)
	return
}

func parsingDataFromURL(strData DataSc) {

	t1Time := time.Now()
	var allStr []string
	replacer := strings.NewReplacer("<strong>", "", "</strong>", "", "&#39;", "", "&gt;", ">", "  ", " ", "Code:", "", "Additional info:", "", ";", "|")
	delBlack := strings.NewReplacer(" ", "")
	//Страница
	//	fmt.Println(strData.URLScriping)

	/*	getQ := libs.URLInputGet{URLIn: strData.URLScriping}
		doc := getQ.Geter()*/

	doc := strData.Doc

	doc.Find("tbody.table-body").Each(func(i int, i1 *goquery.Selection) {
		i1.Find("tr.table-row").Each(func(i int, i2 *goquery.Selection) {

			numberParts, err := i2.Before("td.number-info-cell").Children().Next().Html()
			if err != nil {
				numberParts = ""
			}

			nameParts, err := i2.Before("td.number-info-cell").Children().Next().Next().Children().Html()
			if err != nil {
				nameParts = ""
			}
			/*			fmt.Println(numberParts)
						fmt.Println(nameParts)
						fmt.Println(`__________`)*/

			strOutW :=
				strData.Text2File +
					constData.Separator +
					//strData.Text2File + constData.SepatatorName + replacer.Replace(textAdd) +
					//	constData.Separator +
					replacer.Replace(numberParts) +
					constData.Separator +
					replacer.Replace(nameParts)
			//constData.Separator +
			//	replacer.Replace(NSchemeAdd) +
			//	constData.Separator +
			//	replacer.Replace(QAdd) +
			//	constData.Separator +
			//	replacer.Replace(InfoAdd)

			_ = delBlack.Replace("")
			_ = replacer.Replace("")

			//		strOutW := fmt.Sprintf("%v%v%v", delBlack.Replace(numberParts), constData.Separator, delBlack.Replace(nameParts))

			allStr = append(allStr, strOutW)

			// ТУТ ВЫЗВАТЬ ИНТЕРФЕЙС ЗАПИСИ ДАННЫХ

		})
	})
	/*
		doc.Find("#partContainer").Each(func(i int, s *goquery.Selection) {

			s.Find("li.list-group-item").Each(func(i3 int, l1 *goquery.Selection) {

				// старая версия
				l1.Find("div.tooltip-part").Each(func(i4 int, itm1 *goquery.Selection) {

					fmt.Print(`OLD Version`)

					textI, _ := l1.Find("div.col-8>a.row>p").Html()
					textAdd, _ := itm1.Children().Html()
					CodeAdd, _ := itm1.Children().Next().Html()
					NSchemeAdd, _ := itm1.Children().Next().Next().Html()
					QAdd, _ := itm1.Children().Next().Next().Next().Children().Html()
					InfoAdd, _ := itm1.Children().Next().Next().Next().Next().Children().Html()

					strOutW :=
						strData.TextBefore +
							constData.Separator +
							strData.Text2File + constData.SepatatorName + replacer.Replace(textAdd) +
							constData.Separator +
							replacer.Replace(textI) +
							constData.Separator +
							delBlack.Replace(replacer.Replace(CodeAdd)) +
							constData.Separator +
							//	replacer.Replace(NSchemeAdd) +
							//	constData.Separator +
							//	replacer.Replace(QAdd) +
							//	constData.Separator +
							replacer.Replace(InfoAdd)

					_ = NSchemeAdd
					_ = QAdd

					allStr = append(allStr, strOutW)

					//				time.Sleep(1 * time.Second)
				})

				//новая версия
				//	l1.Find("div.tooltip-part").Each(func(i4 int, itm1 *goquery.Selection) {
				l1.Find("div.row").Each(func(i4 int, itm1 *goquery.Selection) {
					fmt.Print(` NEW Version `)

					/*
						//		textI, _ := l1.Find("div.col-8>a.row>p").Html()
						//	fmt.Println(itm1.Html())

									//Пробуем вытянуть
										CodeAdd, _ := itm1.Children().Next().AddClass("mp-0").Html()
										//		fmt.Println(`===3`)
										//	fmt.Println(CodeAdd)
										_ = CodeAdd
										CodeAdd, _ = itm1.Children().Next().Html()
										//  <p style="color: #000; font-size: 16px;font-weight: 400;" class="w-100 d-block">Grip disc</p> = Grip disc
										CodeAdd, _ = itm1.Children().Next().Children().Children().Html()






					textI, _ := itm1.Children().Next().Children().Children().Html()
					CodeAddNEW, _ := itm1.Children().Next().Children().Children().Next().Html()
					CodeAddNEWClean := strings.Fields(CodeAddNEW)
					textAdd, _ := itm1.Children().Next().Children().Children().Html() //itm1.Children().Html()
					NSchemeAdd, _ := itm1.Children().Next().Next().Html()
					QAdd, _ := itm1.Children().Next().Next().Next().Children().Html()
					//	InfoAdd, _ := itm1.Children().Next().Next().Next().Next().Children().Html()

					strOutW :=
						strData.TextBefore +
							constData.Separator +
							strData.Text2File + constData.SepatatorName + replacer.Replace(textAdd) +
							constData.Separator +
							replacer.Replace(textI) +
							constData.Separator +
							delBlack.Replace(replacer.Replace(CodeAddNEWClean[0]))
				 	constData.Separator +
							replacer.Replace(NSchemeAdd) +
							constData.Separator +
							replacer.Replace(QAdd) +
							constData.Separator +
					delBlack.Replace(replacer.Replace(CodeAddNEWClean[0]))
								replacer.Replace(InfoAdd)

					//constData.Separator +
					//	delBlack.Replace(replacer.Replace(CodeAddNEWClean[1]))
					if len(replacer.Replace(CodeAddNEWClean[1])) > 2 {
						strOutW = strOutW + constData.Separator + replacer.Replace(CodeAddNEWClean[1])
					}

					_ = NSchemeAdd
					_ = QAdd

					allStr = append(allStr, strOutW)

					//				time.Sleep(1 * time.Second)
				})

			})
		})

	*/

	formattedT := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t1Time.Year(), t1Time.Month(), t1Time.Day(), t1Time.Hour(), t1Time.Minute(), t1Time.Second())

	fin := writeOutFiles.ImpDataArr{ArrayStrings: allStr, TimeScraping: formattedT}
	fin.WriteToFile()

	return
}

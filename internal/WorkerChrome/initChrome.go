package WorkerChrome

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/ChangeData"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/MakeConfiger"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/constData"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/libs"
	"github.com/chromedp/chromedp"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

func InitChrome(cfg *MakeConfiger.Config) {

	//var err1 error

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		//		chromedp.UserDataDir("C:\\temp\\asat_temp"),
		//chromedp.Flag("headless", cfg.Debug),
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", true),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.Sleep(2*time.Second),
	)
	if err != nil {
		log.Fatal("Работа программы завершена не корректно. Нужно закрыть все лишние окна браузера Хром")
	}

	var keyURL int
	var valueURL string
	var res string

	reader := bufio.NewReader(os.Stdin)

	for {
		time.Sleep(constData.TimeOutRequest * time.Second)

		res = ""
		sURLc := <-cfg.URLReq

		sURLc = ChangeData.Replacer(sURLc)

	changeURL: //подготовка URL - c заменой адреса

		sURL := sURLc

		valueURL = ""
		keyURL = 0

		for key, value := range cfg.ListUrl {
			keyURL = key
			valueURL = value
			break
		}
		if keyURL == 0 {

			log.Fatal(`ВСЕ САЙТЫ ЗАБЛОКИРОВАНЫ. Продолжение работы не возможно. Последний адрес - `, sURLc)
		}
		valueURL = strings.TrimSpace(valueURL)

		// Разбор URL
		parsedURL, err := url.Parse(sURL)
		if err != nil {
			log.Fatal(`ОШИБКА В URL`, sURL, parsedURL, err)
			return
		}

		parsedURL.Host = valueURL
		sURL = parsedURL.String()

	scrap: //сбор данных
		err = chromedp.Run(ctx, libs.ScrapIt(sURL, &res, 10*time.Second))
		if err != nil {
			if errors.Is(err, context.Canceled) {
				fmt.Println(sURL)
				log.Fatal("Произошла не предвиденная ошибка. Последняя страница с ошибкой")
			}
		}

		if strings.Contains(res, "Доступ к каталогу заблокирован") || strings.Contains(res, "Страница не найдена") {
			log.Println(valueURL + " Был удалён из очереди - ЗАБЛОКИРОВАН")
			delete(cfg.ListUrl, keyURL)

			goto changeURL
		}

		// Проверяем наличие подстроки "Робота" в строке
		if strings.Contains(res, constData.SearchRobot) {

			if strings.Contains(res, "Подтверждаю") {
				fmt.Println("Пробуем пройти капчу")
				time.Sleep(2 * time.Second)
				err = chromedp.Run(ctx,
					chromedp.Navigate(sURL),
					chromedp.Sleep(5*time.Second),
					chromedp.Click(`div.d-flex`),
				)
				time.Sleep(2 * time.Second)
				fmt.Println("попробовали")

				if _ = chromedp.Run(ctx, libs.ScrapIt(sURL, &res, 15*time.Second)); strings.Contains(res, constData.SearchRobot) {
					goto needMan
				} else {
					goto scrap
				}

			}

		needMan:
			fmt.Println("Когда прошли капчу - Нажмите Enter")
			//fmt.Scanln()
			_, _ = reader.ReadString('\n')
			goto scrap
		}
		/*
			// Проверяем наличие подстроки "Робота" в строке
			if strings.Contains(res, constData.SearchRobot) {

				fmt.Println("Когда прошли капчу - Нажмите Enter")
				fmt.Scanln()
				goto scrap
			}
		*/
		if !strings.Contains(res, "acat.online") {
			goto scrap
		}

		cfg.DataOUTReq <- []byte(res)

		log.Println(`Открыли страницу`)
	}

}

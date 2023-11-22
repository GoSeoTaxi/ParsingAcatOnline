package WorkerChrome

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/ChangeData"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/MakeConfiger"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/constData"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

func InitChrome(cfg *MakeConfiger.Config) {

	//var err1 error

	opts := append(chromedp.DefaultExecAllocatorOptions[:],

		chromedp.UserDataDir("C:\\Users\\info\\AppData\\Local\\Google\\Chrome\\User Data\\"),
		chromedp.Flag("profile-directory", "Profile 6"),
		//chromedp.Flag("headless", cfg.Debug),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.WindowSize(1000, 768),
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

		//Счётчик повторов
		tryCounter := constData.CounterTryCounter

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
			for k, v := range cfg.ListUrlFromFile {
				cfg.ListUrl[k] = v
			}

			log.Println(`САЙТ ЗАПОДОЗРИЛ АВТОМАТИЧЕСКИЕ ДЕЙСТВИЯ.`)
			log.Println(`нужно попробовать сменить IP`)

			cfg.Pause = true
			fmt.Println("НАЖМИ ENTER")
			//	_, _ = reader.ReadString('\n')
			cfg.Pause = false
			res = ""
			tryCounter = constData.CounterTryCounter
			goto changeURL
			//log.Fatal(`ВСЕ САЙТЫ ЗАБЛОКИРОВАНЫ. Продолжение работы не возможно. Последний адрес - `, sURLc)
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
		tryCounter -= 1
		res = ""

		//Если счётчик уже близок к нулю, пробуем сменить URL

		if tryCounter == 0 {

			log.Println(`САЙТ ЗАПОДОЗРИЛ АВТОМАТИЧЕСКИЕ ДЕЙСТВИЯ.`)
			log.Println(`нужно зайти в гугл, открыть 2-3 страницы и т.п.`)

			cfg.Pause = true
			fmt.Println("НАЖМИ ENTER")
			_, _ = reader.ReadString('\n')
			cfg.Pause = false
			res = ""
			tryCounter = constData.CounterTryCounter
			goto changeURL
		}

		//	err = chromedp.Run(ctx, libs.ScrapIt(sURL, &res, 10*time.Second))

		err = chromedp.Run(ctx,
			chromedp.Navigate(sURL),
			chromedp.Sleep(time.Duration(constData.TimeOutStartCaptcha+rand.Intn(constData.TimeOutStopCaptcha-constData.TimeOutStartCaptcha+1))*time.Second),
			chromedp.ActionFunc(
				func(ctx context.Context) error {
					node, err := dom.GetDocument().Do(ctx)
					if err != nil {
						return err
					}
					res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)

					fmt.Println(node.DocumentURL)

					return err
				},
			),
		)

		if err != nil {
			if errors.Is(err, context.Canceled) {
				fmt.Println(sURL)
				log.Fatal("Произошла не предвиденная ошибка. Последняя страница с ошибкой")
			}
		}

		if strings.Contains(res, "Internal Server Error") {
			log.Println("На странице " + sURL + " БД не содержит данных")
			goto EndPoint
		}

		if strings.Contains(res, "Доступ к каталогу заблокирован") /*|| strings.Contains(res, "Страница не найдена") */ {
			log.Println(valueURL + " Был удалён из очереди - ЗАБЛОКИРОВАН")
			delete(cfg.ListUrl, keyURL)
			res = ""
			tryCounter = constData.CounterTryCounter
			goto changeURL
		}

		// Проверяем наличие подстроки "Робота" в строке
		if strings.Contains(res, constData.SearchRobot) {

			if strings.Contains(res, "Подтверждаю") || strings.Contains(res, "Страница не найдена") {

				fmt.Println("++++++ СПИМ 20 С")
				time.Sleep(30 * time.Second)
				fmt.Println("Проснулись")

				log.Printf("Пробуем перезагрузить страницу - %v \n", sURL)

				time.Sleep(2 * time.Second)
				err = chromedp.Run(ctx,
					chromedp.Navigate(sURL),
					chromedp.Sleep(time.Duration(constData.TimeOutStartCaptcha+rand.Intn(constData.TimeOutStopCaptcha-constData.TimeOutStartCaptcha+1))*time.Second),
					//	chromedp.Click(`div.d-flex`),
					chromedp.Sleep(time.Duration(constData.TimeOutStartCaptcha+rand.Intn(constData.TimeOutStopCaptcha-constData.TimeOutStartCaptcha+1))*time.Second),

					chromedp.ActionFunc(

						func(ctx context.Context) error {
							node, err := dom.GetDocument().Do(ctx)
							if err != nil {
								return err
							}
							res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
							return err
						},
					),
				)

				if err != nil {
					log.Printf("++++++++++\n ВАЖНАЯ ОШИБКА \n Сделай фото для Ильи %v \n ++++++++++\n", err)
				}
				err = nil

				time.Sleep(3 * time.Second)
				fmt.Println("попробовали")

				//_ = chromedp.Run(ctx, libs.ScrapIt(sURL, &res, 30*time.Second))

				if strings.Contains(res, "Проверка не пройдена, попробуйте еще раз.") /*|| strings.Contains(res, "Страница не найдена") */ {

					log.Println(valueURL + " Был удалён из очереди - блокировка")
					delete(cfg.ListUrl, keyURL)
					res = ""
					tryCounter = constData.CounterTryCounter
					goto changeURL
				}

				if strings.Contains(res, constData.SearchRobot) && tryCounter == 0 {
					tryCounter = constData.CounterTryCounter
					goto needMan
				} else {
					res = ""
					goto scrap
				}

			}

		needMan:
			cfg.Pause = true
			fmt.Println("НАЖМИ ENTER")
			//fmt.Scanln()
			_, _ = reader.ReadString('\n')
			cfg.Pause = false
			res = ""
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
			fmt.Println(`Не нашли акат Онлайн`)
			res = ""
			goto scrap
		}

	EndPoint:
		cfg.DataOUTReq <- []byte(res)
		log.Println(`Открыли страницу`)
	}

}

package WorkerChrome

import (
	"context"
	"errors"
	"fmt"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/ChangeData"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/MakeConfiger"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/constData"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/libs"
	"github.com/chromedp/chromedp"
	"log"
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
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://ya.ru"),
		chromedp.Sleep(2*time.Second),
	)
	if err != nil {
		log.Fatal("Работа программы завершена не корректно. Нужно закрыть все лишние окна браузера Хром")
	}

	for {

		var res string

		time.Sleep(constData.TimeOutRequest * time.Second)
		sURL := <-cfg.URLReq
		sURL = ChangeData.Replacer(sURL)

	scrap:
		err = chromedp.Run(ctx, libs.ScrapIt(sURL, &res))
		if err != nil {
			if errors.Is(err, context.Canceled) {
				fmt.Println(sURL)
				log.Fatal("Произошла не предвиденная ошибка. Последняя страница с ошибкой")

				//22

				//var err1 error

				//11
				time.Sleep(10 * time.Second)
			}
		}

		// Проверяем наличие подстроки "Робота" в строке
		if strings.Contains(res, constData.SearchRobot) {
			fmt.Println("Когда прошли капчу - Нажмите Enter")
			fmt.Scanln()
			goto scrap
		}
		if !strings.Contains(res, "acat.online") {
			goto scrap
		}

		cfg.DataOUTReq <- []byte(res)

		fmt.Println(`Открыли страницу`)
	}

}

package libs

import (
	"context"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"time"
)

func ScrapIt(url string, str *string, timeout time.Duration) chromedp.Tasks {
	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			*str, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)

			// Проверяем время и прерываем выполнение, если время превышено
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return err
			}
		}),
	}
}

/*
func ScrapIt(url string, str *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			*str, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	}
}
*/

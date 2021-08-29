package weather

import (
	"context"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Args struct {
	Area string
}

type Reply struct {
	Weather string
}

type Weather int

func (t *Weather) Get(ctx context.Context, args *Args, reply *Reply) error {
	reply.Weather = getWeather(args.Area)
	return nil
}

// 通过爬取查找地区拼音，直接转会出现 成都 -> chengdou 的情况
func getPinyin(area string) (pinyin string) {
	url := "http://lishi.tianqi.com/"
	c := colly.NewCollector()
	c.OnHTML("a[title='"+area+"历史天气']", func(h *colly.HTMLElement) {
		href := h.Attr("href")
		index := strings.Index(href, "/")
		pinyin = href[:index]
	})
	c.Visit(url)
	c.Wait()
	if len(pinyin) == 0 {
		pinyin = area
	}
	return
}

func getWeather(area string) (weather string) {
	pinyin := area
	if len([]rune(area)) != len(area) {
		pinyin = getPinyin(area)
	}
	fmt.Println(pinyin)
	if strings.EqualFold(pinyin, "") {
		return
	}
	url := "https://www.tianqi.com/" + pinyin + "/7/"

	c := colly.NewCollector()
	c.OnHTML("ul[class='weaul']", func(h *colly.HTMLElement) {
		h.ForEach("li", func(i int, h *colly.HTMLElement) {
			if i <= 2 {
				weather += h.ChildText("span[class='fr']")
				weather += " " + h.ChildText("span[class='fl']")
				h.ForEach("div[class='weaul_z']", func(i int, h *colly.HTMLElement) {
					weather += " " + h.Text
				})
			}
			if i < 2 {
				weather += "\n"
			}
		})
	})
	c.OnError(func(r *colly.Response, e error) {
		weather = ""
	})
	c.Visit(url)

	c.Wait()
	return
}

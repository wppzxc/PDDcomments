package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

const (
	GoodsInfoUrl = "https://mobile.yangkeduo.com/goods.html?goods_id=%s"
)

func GetGoodsPrice(itemId string, discount string) string {
	url := fmt.Sprintf(GoodsInfoUrl, itemId)
	dom, err := goquery.NewDocument(url)
	if err != nil {
		return "0"
	}
	tmpPrice := dom.Find("div[data-active=red]").Find("span").First().Text()
	tmpPrice = strings.Replace(tmpPrice, "￥", "", -1)
	oldPrice, _ := strconv.ParseFloat(tmpPrice, 32)
	dis, _ := strconv.ParseFloat(discount, 32)
	newPrice := fmt.Sprintf("%.2f", oldPrice * dis)
	return newPrice
}

//func Get

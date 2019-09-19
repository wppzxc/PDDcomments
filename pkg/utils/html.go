package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

const (
	GoodsInfoUrl = "https://mobile.yangkeduo.com/goods.html?goods_id=%s"
)

func GetGoodsPrice(key, itemId string, discount string) string {
	url := fmt.Sprintf(GoodsInfoUrl, itemId)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "0"
	}
	cookie := &http.Cookie{Name: "PDDAccessToken", Value: key, HttpOnly: false}
	req.AddCookie(cookie)
	resp, err := client.Do(req)
	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return "0"
	}
	tmpPrice := dom.Find("div[data-active=red]").Find("span").First().Text()
	tmpPrice = strings.Replace(tmpPrice, "￥", "", -1)
	oldPrice, _ := strconv.ParseFloat(tmpPrice, 32)
	dis, _ := strconv.ParseFloat(discount, 32)
	newPrice := fmt.Sprintf("%.1f", oldPrice * dis)
	return newPrice
}

//func Get

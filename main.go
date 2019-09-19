package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/wpp/PDDComments/javascripts"
	"github.com/wpp/PDDComments/pages"
	"github.com/wpp/PDDComments/pkg/utils"
	"github.com/wpp/PDDComments/types"
	"github.com/zserge/webview"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var logined = false
var AK = ""
var newCommentCh = make(chan string)
var newPriceCh = make(chan string)

func main() {
	rand.Seed(time.Now().Unix())
	// 初始化配置文件
	utils.InitConfig("./PDDComments.json")
	loginPage := webview.New(webview.Settings{
		Title:                  "Login",
		URL:                    "https://mobile.yangkeduo.com/login.html",
		ExternalInvokeCallback: eventHandler,
	})
	login := types.Application{
		WebApp: loginPage,
	}
	go func() {
		fmt.Println("reset ak ...")
		time.Sleep(1 * time.Second)
		login.RestAK()
		fmt.Println("start login ...")
		for {
			login.Login()
			if logined {
				login.CloseLoginPage()
				return
			}
		}
	}()
	fmt.Println("start login ...")
	login.WebApp.Run()
	login.WebApp.Exit()
	if len(AK) <= 0 {
		fmt.Println("get AccessKey error : there is no AccessKey!")
		os.Exit(0)
	}
	fmt.Println("finish login ...")
	// 加载配置文件
	pd := &types.PageData{}
	data, _ := ioutil.ReadFile("./PDDComments.json")
	_ = json.Unmarshal(data, pd)
	html := fmt.Sprintf(pages.IndexHtml, pd.PicPriceX, pd.PicPriceY, pd.PicPriceSize, pd.PicAccountName, pd.PicAccountX, pd.PicAccountY,pd.PicAccountSize)
	fmt.Println(html)
	mainPage := webview.New(webview.Settings{
		Title:                  "PDDComments",
		URL:                    "data:text/html," + url.PathEscape(html),
		Width:                  1200,
		Height:                 620,
		ExternalInvokeCallback: eventHandler,
	})
	main := types.Application{
		WebApp: mainPage,
	}
	
	go func() {
		fmt.Println("start output comment")
		for {
			select {
			case newPrice := <- newPriceCh:
				fmt.Println("get new price : ", newPrice)
				main.WebApp.Dispatch(func() {
					if err := main.WebApp.Eval(fmt.Sprintf(javascripts.OutputPriceJS, newPrice)); err != nil {
						fmt.Println(err)
					}
				})
			case newComment := <- newCommentCh:
				fmt.Println("get new comment : ", newComment)
				main.WebApp.Dispatch(func() {
					if err := main.WebApp.Eval(fmt.Sprintf(javascripts.OutputCommentJS, newComment)); err != nil {
						fmt.Println(err)
					}
				})
			}
		}
	}()
	defer main.WebApp.Exit()
	main.WebApp.Run()
}

func eventHandler(w webview.WebView, data string) {
	strs := strings.Split(data, "|||")
	fmt.Println("event is : ", data)
	switch strs[0] {
	case "cookie":
		cookies := parseData(data)
		ak, ok := cookies["PDDAccessToken"]
		if ok {
			fmt.Println("-------------------------AK is :", ak)
			AK = ak
			logined = true
			return
		}
	case "generate":
		utils.SaveConfig("./PDDComments.json", strs[1])
		comment := generateComment(strs[1])
		price := generatePrice(strs[1])
		newCommentCh <- comment
		newPriceCh <- price
	case "autoDownloadPic":
		if err := autoDownloadPic(strs[1]); err != nil {
			fmt.Println(err)
		}
	}
}

func parseData(data string) map[string]string {
	kvs := strings.Split(data, ";")
	cookies := make(map[string]string)
	for _, kv := range kvs {
		kv = strings.Trim(kv, " ")
		strs := strings.Split(kv, "=")
		cookies[strs[0]] = strs[1]
	}
	return cookies
}

func generateComment(data string) string {
	pd := &types.PageData{}
	if err := json.Unmarshal([]byte(data), pd); err != nil {
		return ""
	}
	u, err := url.Parse(pd.ItemLink)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	param, _ := url.ParseQuery(u.RawQuery)
	itemId := param["goods_id"][0]
	if len(itemId) == 0 {
		fmt.Println("商品链接错误，无法解析商品id！")
		return ""
	}
	fmt.Println("pd is :", pd)
	result := getCommentResult(itemId, pd.CommentNumber, pd.CommentHead, pd.CommentFoot, pd.CommentFilter)
	fmt.Println("result comment is : ", result)
	return result
}

func generatePrice(data string) string {
	pd := &types.PageData{}
	if err := json.Unmarshal([]byte(data), pd); err != nil {
		return ""
	}
	u, err := url.Parse(pd.ItemLink)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	param, _ := url.ParseQuery(u.RawQuery)
	itemId := param["goods_id"][0]
	if len(itemId) == 0 {
		fmt.Println("商品链接错误，无法解析商品id！")
		return ""
	}
	fmt.Println("pd is :", pd)
	result := utils.GetGoodsPrice(itemId, pd.CommentDiscount)
	fmt.Println("result price is : ", result)
	return result
}

func getCommentResult(itemId, minLength, commentPrefix, commentSuffix, commentFilter string) string {
	comment := exeComment(AK, itemId, minLength, commentFilter, "", 0)
	comment = strings.Replace(comment, "'", "‘", -1)
	comment = strings.Replace(comment, "\"", "“", -1)
	return commentPrefix + comment + commentSuffix
}

func exeComment(key string, itemId string, minLength string, filter string, comment string, failTimes int) string {
	if failTimes > 30 {
		return "登录失效，请重新登录！"
	}
	page := rand.Intn(29) + 1
	length, _ := strconv.Atoi(minLength)
	cms := utils.GetOnePageComments(key, itemId, page)
	if len(cms) == 0 {
		failTimes ++
		return exeComment(key, itemId, minLength, filter, comment, failTimes)
	}
	cms = utils.RemoveEmptyComments(cms)
	cms = utils.RemoveLowScoreComments(cms)
	cms = utils.FilterKeys(filter, cms)
	index := rand.Intn(len(cms) - 1)
	comment = comment + cms[index].Comment + "，"
	fmt.Println("tmp comment is ：", comment)
	if utf8.RuneCountInString(comment) >= length {
		return comment
	} else {
		return exeComment(key, itemId, minLength, filter, comment, failTimes)
	}
}

func autoDownloadPic(base64Str string) error {
	d, _ := base64.StdEncoding.DecodeString(base64Str)
	timestamp := time.Now().Unix()
	filename := strconv.FormatInt(timestamp, 10) + ".jpeg"
	os.Remove(filename)
	if _, err := os.Create(filename); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, d, 0755); err != nil {
		fmt.Println(err)
	}
	return nil
}

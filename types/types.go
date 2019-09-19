package types

import (
	"github.com/wpp/PDDComments/javascripts"
	"github.com/wpp/PDDComments/pkg/log"
	"github.com/zserge/webview"
	"time"
)

var logger = log.Logger

type Application struct {
	WebApp    webview.WebView
	Logined   bool
	AccessKey string
}

func (app *Application) Login() {
	time.Sleep(1 * time.Second)
	app.WebApp.Dispatch(func() {
		if err := app.WebApp.Eval(javascripts.LoginJS); err != nil {
			logger.Printf("Error : %s",err)
		}
	})
}

func (app *Application) RestAK() {
	app.WebApp.Dispatch(func() {
		if err := app.WebApp.Eval(javascripts.ResetAKJS); err != nil {
			logger.Printf("Error : %s",err)
		}
		logger.Println("初始化AK成功!")
	})
}

func (app *Application) CloseLoginPage() {
	app.WebApp.Dispatch(func() {
		if err := app.WebApp.Eval(javascripts.CloseLoginPage); err != nil {
			logger.Printf("Error : %s",err)
		}
		logger.Println("登录成功!")
	})
}

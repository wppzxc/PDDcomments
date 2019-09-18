package types

import (
	"fmt"
	"github.com/wpp/PDDComments/javascripts"
	"github.com/zserge/webview"
	"time"
)

type Application struct {
	WebApp    webview.WebView
	Logined   bool
	AccessKey string
}

func (app *Application) Login() {
	time.Sleep(1 * time.Second)
	app.WebApp.Dispatch(func() {
		if err := app.WebApp.Eval(javascripts.LoginJS); err != nil {
			fmt.Println(err)
		}
	})
}

func (app *Application) RestAK() {
	app.WebApp.Dispatch(func() {
		if err := app.WebApp.Eval(javascripts.ResetAKJS); err != nil {
			fmt.Println(err)
		}
		fmt.Println("初始化AK成功!")
	})
}

func (app *Application) CloseLoginPage() {
	app.WebApp.Dispatch(func() {
		if err := app.WebApp.Eval(javascripts.CloseLoginPage); err != nil {
			fmt.Println(err)
		}
		fmt.Println("登录成功!")
	})
}

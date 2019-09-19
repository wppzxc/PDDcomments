package main

import "github.com/zserge/webview"

func main() {
	webview.Open("Minimal webview example",
		"https://mobile.yangkeduo.com/login.html", 800, 600, true)
}

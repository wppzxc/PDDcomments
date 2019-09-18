package javascripts

var ResetAKJS = `!function () {document.cookie = "PDDAccessToken=;expires=Thu, 01 Jan 1970 00:00:00 GMT;"}();`

var LoginJS = `!function () {window.external.invoke('cookie|||'+document.cookie)}();`

var CloseLoginPage = `!function () {document.body.innerText="登录成功，请关闭此页面！"}()`

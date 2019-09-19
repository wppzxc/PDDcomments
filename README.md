##build
```
# rsrc.exe -arch amd64 -manifest PDDComments.exe.manifest -o rsrc.syso
# // 指定程序icon
# // rsrc.exe -arch amd64 -manifest PDDComments.exe.manifest -ico ./assets/img/icon.ico -o rsrc.syso
# go build -ldflags="-H windowsgui"
``` 

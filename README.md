trf
===

**windows timer resolution fix for go**

_this issue is fixed by https://go-review.googlesource.com/#/c/17164/ in go 1.6 , it is no longer needed after go 1.6._

in windows, [go runtime](http://golang.org/src/pkg/runtime/os_windows.c?s=#L99) calls [ timeBeginPeriod ](http://msdn.microsoft.com/en-us/library/windows/desktop/dd757624(v=vs.85).aspx ) when initialization, and set the system timer resolution to 1ms.

actually, all .exe files build by go, will hold a 1ms system timer resolution until quit.

[ such behavior](https://code.google.com/p/chromium/issues/detail?id=153139), which chrome used to have, is proved to be not battery friendly.


i have reported a [issue](https://github.com/golang/go/issues/8687) here.

before this behavior is changed in future, you can simply import this package in your project to solve it.

	import _ "github.com/defia/trf"

that's all. the package calls [timeEndPeriod](http://msdn.microsoft.com/en-us/library/windows/desktop/dd757626%28v=vs.85%29.aspx) on start, so the system timer resolution change made by runtime will be discarded.


it may have impact on performance, but if your client stays background for a long time, such as [cow](https://github.com/cyfdecyf/cow) , [shadowsocks-go](https://github.com/shadowsocks/shadowsocks-go) , [gocode](https://github.com/nsf/gocode) on your windows laptop, i don't see any slow down but the battery life will be definitely better.


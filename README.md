trf
===

**windows timer resolution fix for go**  

  

in windows, [go runtime](http://golang.org/src/pkg/runtime/os_windows.c?s=#L99) calls [ timeBeginPeriod ](http://msdn.microsoft.com/en-us/library/windows/desktop/dd757624(v=vs.85).aspx ) when initialization, and set the system timer resolution to 1ms.  

actually, all .exe files build by go, will hold a 1ms system timer resolution until quit.

[ such behavior](https://code.google.com/p/chromium/issues/detail?id=153139), which chrome used to have, is proved to be not battery friendly.  


i have reported a [issue](https://code.google.com/p/go/issues/detail?id=8687) here.  

before this behavior is changed in future, you can simply import this package in your project to solve it.

	import _ "github.com/defia/trf"

that's all. the package calls [timeEndPeriod](http://msdn.microsoft.com/en-us/library/windows/desktop/dd757626%28v=vs.85%29.aspx) on start, so the system timer resolution change made by runtime will be discarded.


i don't know whether it will have impact on performance, but if you runs a client background, such as cow,shadowsocks,gocode on your windows laptop, i don't see any slow down but the battery life will definitely better.


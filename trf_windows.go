package trf

import "syscall"

const (
	DLL      = "winmm.dll"
	FUNCTION = "timeEndPeriod"
)

func init() {
	syscall.NewLazyDLL("winmm.dll").NewProc("timeEndPeriod").Call(1)

}

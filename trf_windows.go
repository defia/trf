package trf

import (
	"log"
	"syscall"
)

const (
	DLL      = "winmm.dll"
	FUNCTION = "timeEndPeriod"
)

func init() {
	winmm, err := syscall.LoadLibrary(DLL)
	if err != nil {
		log.Println(err)
		return
	}
	c, err := syscall.GetProcAddress(syscall.Handle(winmm), FUNCTION)
	if err != nil {
		log.Println(err)
		return
	}
	syscall.Syscall(c, 1, uintptr(1), 0, 0)

}

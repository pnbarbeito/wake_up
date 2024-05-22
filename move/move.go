package move

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
	"github.com/go-vgo/robotgo"
)

func Move_win() {
	userDll := syscall.NewLazyDLL("user32.dll")
	getWindowRectProc := userDll.NewProc("GetCursorPos")
	type POINT struct {
		X, Y int32
	}
	var pt POINT
	setWindowRectProc := userDll.NewProc("SetCursorPos")

	for range time.Tick(time.Second * 60) {
		_, _, eno := syscall.SyscallN(getWindowRectProc.Addr(), uintptr(unsafe.Pointer(&pt)))
		if eno != 0 {
			fmt.Println(eno)
		}
		x, y := pt.X, pt.Y
		fmt.Printf("[cursor.Pos] X:%d Y:%d\n", pt.X, pt.Y)
		syscall.SyscallN(getWindowRectProc.Addr(), uintptr(unsafe.Pointer(&pt)))

		rx := int(x)
		ry := int(y)
		num := robotgo.DisplaysNum()
		sx, _ := robotgo.GetScreenSize()
		if num > 1 {
			rx += (sx * (num - 1))
			fmt.Printf("[cursor.Pos] X:%d Y:%d\n", rx+1, ry+1)
		}
		robotgo.Move(rx, ry)

		time.Sleep(1 * time.Millisecond)
		setWindowRectProc.Call(uintptr(x), uintptr(y))
	}
}

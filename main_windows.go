// go:build windows

package main

import (
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/atotto/clipboard"
)

// Win32 API MessageBox() in Golang https://gist.github.com/NaniteFactory/0bd94e84bbe939cda7201374a0c261fd
// MessageBox of Win32 API.
func MessageBox(hwnd uintptr, caption, title string, flags uint) uintptr {
	captionPtr, _ := syscall.UTF16PtrFromString(caption)
	titlePtr, _ := syscall.UTF16PtrFromString(title)
	r, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(captionPtr)),
		uintptr(unsafe.Pointer(titlePtr)),
		uintptr(flags))
	return r
}

// MessageBoxPlain of Win32 API.
func MessageBoxPlain(title, caption string) uintptr {
	const (
		NULL  = 0
		MB_OK = 0
	)
	return MessageBox(NULL, caption, title, MB_OK)
}

func main() {
	execPath := strings.Split(os.Args[0], "\\")
	args := "nil"
	if len(os.Args) > 1 {
		args = strings.Join(os.Args[1:], "\" \"")
		if len(os.Args) > 2 {
			args = "\"" + args + "\""
		}
		clipboard.WriteAll(args)
	}
	MessageBoxPlain(execPath[len(execPath)-1], args)
}

//go:build windows

package main

import (
	"github.com/atotto/clipboard"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// Win32 API MessageBox() in Golang https://gist.github.com/NaniteFactory/0bd94e84bbe939cda7201374a0c261fd
// MessageBox of Win32 API.
func MessageBox(hwnd uintptr, caption, title string, flags uint) int {
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(flags))

	return int(ret)
}

// MessageBoxPlain of Win32 API.
func MessageBoxPlain(title, caption string) int {
	const (
		NULL  = 0
		MB_OK = 0
	)
	return MessageBox(NULL, caption, title, MB_OK)
}

func main() {
	args := strings.Join(os.Args[1:], " ")
	clipboard.WriteAll(args)
	MessageBoxPlain("ARGP", args)
}

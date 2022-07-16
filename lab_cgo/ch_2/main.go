package main

/*

#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lsdk
#include <stdio.h>
#include "libsdk.h"
*/
import "C"
import (
	"fmt"
)

func main() {
	Query()
}
func Query() {
	result := C.GetDeviceId()
	ans := C.GoString(result)
	fmt.Println(ans)
}

package iupcontrols

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/alanyuen/iup"
)

/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -L${SRCDIR}/../lib/cd -L${SRCDIR}/../lib/im
#cgo LDFLAGS: -liupcontrols -liupcd -liup
#cgo LDFLAGS: -lcd -lfreetype6 -lz
#cgo LDFLAGS: -lgdi32 -luser32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iupcontrols.h>
*/
import "C"

func ListView() iup.Ihandle {
	ih := Matrix()
	ih.SetAttributes(map[string]interface{}{
		"CURSOR":       "ARROW",
		"READONLY":     "YES",
		"MARKMODE":     "LIN",
		"MARKATTITLE":  "NO",
		"MARKAREA":     "NOT_CONTINUOUS",
		"MARKMULTIPLE": "YES",
		"HLCOLOR":      "124 180 255",
	})
	ih.SetCallback("CLICK_CB", func(handle iup.Ihandle, lin int, col int, status uintptr) int {
		if lin < 1 {
			return iup.DEFAULT
		}

		marked := handle.GetAttribute("MARKED")
		if len(marked) > lin {
			if !iup.IsControl(status) && !iup.IsShift(status) &&
				!(iup.IsButton3(status) && string(marked[lin]) == "1") {
				handle.SetAttribute("MARKED", nil)
			}
		}

		numcol := handle.GetInt("NUMCOL")
		if iup.IsShift(status) && iup.IsButton1(status) {
			lastLin := handle.GetInt("LAST_SELECT_LIN")
			if lastLin > 0 {
				minLin := int(math.Min(float64(lastLin), float64(lin)))
				maxLin := int(math.Max(float64(lastLin), float64(lin)))

				for j := minLin; j <= maxLin; j++ {
					for i := 0; i < numcol; i++ {
						handle.SetAttribute(fmt.Sprintf("MARK%d:%d", j, i+1), "1")
					}
				}
			}
		}

		for i := 0; i < numcol; i++ {
			handle.SetAttribute(fmt.Sprintf("MARK%d:%d", lin, i+1), "1")
		}
		handle.SetAttribute("LAST_SELECT_LIN", lin)

		if iup.IsButton1(status) {
			callBack := handle.GetCallback("MOUSE_LEFT_CLICK")
			fn := (*func(iup.Ihandle, int, int, uintptr) int)(unsafe.Pointer(callBack))
			if fn != nil {
				(*fn)(handle, lin, col, status)
			}
		}

		if iup.IsButton3(status) {
			callBack := handle.GetCallback("MOUSE_RIGHT_CLICK")
			fn := (*func(iup.Ihandle, int, int, uintptr) int)(unsafe.Pointer(callBack))
			if fn != nil {
				(*fn)(handle, lin, col, status)
			}
		}
		return iup.DEFAULT
	})
	return ih
}

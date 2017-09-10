package main

import (
	"fmt"
	"strconv"
	//"unsafe"

	"github.com/alanyuen/iup"
	"github.com/alanyuen/iup/iupcontrols"
)

var listView *iup.Ihandle

func main() {
	iup.Open()
	defer iup.Close()

	iupcontrols.Open()
	iup.SetGlobal("UTF8MODE", "YES")

	layout := iup.Vbox(
		iupcontrols.ListView().AssignTo(&listView),
	)

	count := 10
	listView.SetAttributes(map[string]interface{}{
		"NUMCOL": "2",
		//"NUMCOL_VISIBLE": "2",
		"NUMLIN": strconv.Itoa(count),
		"WIDTH1": "60",
		"WIDTH2": "20",
		"EXPAND": "VERTICAL",
		"0:1":    "Item",
		"0:2":    "Qty",
	})
	for i := 1; i <= count; i++ {
		listView.SetAttributes(fmt.Sprintf("%d:0=%d", i, i))
	}

	//	listView.SetCallback("CLICK_CB", func(handle iup.Ihandle, lin int, col int, status uintptr) int {
	//		if lin < 1 {
	//			return iup.DEFAULT
	//		}

	//		marked := handle.GetAttribute("MARKED")
	//		if len(marked) > lin {
	//			if iup.IsButton3(status) {
	//				if string(marked[lin]) != "1" {
	//					handle.SetAttribute("MARKED", nil)
	//				}
	//			}
	//		}
	//		handle.SetAttribute("MARK"+strconv.Itoa(lin)+":1", "1")
	//		handle.SetAttribute("MARK"+strconv.Itoa(lin)+":2", "1")

	//		return iup.DEFAULT
	//	})
	callback := func(handle iup.Ihandle, lin int, col int, status uintptr) int {
		fmt.Println("MOUSE_LEFT_CLICK")
		return iup.DEFAULT
	}
	listView.SetCallbackPtr("MOUSE_LEFT_CLICK", &callback)

	callback2 := func(handle iup.Ihandle, lin int, col int, status uintptr) int {
		fmt.Println("MOUSE_RIGHT_CLICK")
		return iup.DEFAULT
	}
	listView.SetCallbackPtr("MOUSE_RIGHT_CLICK", &callback2)

	/* Creating main dialog */
	dialog := iup.Dialog(layout)
	dialog.SetAttribute("SIZE", "320x240")

	iup.Show(dialog)
	iup.MainLoop()
}

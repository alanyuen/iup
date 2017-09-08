package main

import (
	"github.com/alanyuen/iup"
)

var img = []byte{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 6, 6, 6, 6, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 6, 6, 6, 6, 6, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 6, 6, 6, 6, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 6, 5, 5, 6, 6, 6, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 6, 5, 5, 6, 6, 6, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 6, 6, 6, 6, 6, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 6, 6, 5, 5, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 6, 6, 6, 6, 6, 6, 5, 5, 6, 6, 6, 6, 6, 6, 5, 5, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 6, 6, 6, 6, 6, 6, 5, 5, 5, 6, 6, 6, 6, 5, 5, 5, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 2,
	1, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 2,
	1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
}

func main() {
	iup.Open()
	defer iup.Close()

	/* Creating dialog's icon */
	icon := iup.Image(32, 32, img).SetAttributes(map[string]string{
		"1": "255 255 255",
		"2": "000 000 000",
		"3": "226 226 226",
		"4": "128 128 128",
		"5": "192 192 192",
		"6": "000 000 255",
	})
	iup.SetHandle("icon", icon)

	/* Creating dialog's content */
	quit_bt := iup.Button("Quit").SetCallback("ACTION", quit_cb)
	iup.SetHandle("quit", quit_bt)

	/* Creating dialog's menu */
	options := iup.Menu(iup.Item("Quit").SetCallback("ACTION", quit_cb))
	submenu := iup.Submenu("File", options)
	menu := iup.Menu(submenu)
	iup.SetHandle("menu", menu)

	/* Creating main dialog */
	dialog := iup.Dialog(iup.Vbox(quit_bt))
	iup.SetAttribute(dialog, "TITLE", "IupDialog")
	iup.SetAttribute(dialog, "MENU", "menu")
	iup.SetAttribute(dialog, "CURSOR", "CROSS")
	iup.SetAttribute(dialog, "ICON", "icon")
	iup.SetAttribute(dialog, "DEFAULTESC", "quit")

	iup.Show(dialog)
	iup.MainLoop()
}

func quit_cb(ih iup.Ihandle) int {
	return iup.CLOSE
}

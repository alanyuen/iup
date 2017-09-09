package main

import (
	"fmt"

	"github.com/alanyuen/iup"
	"github.com/alanyuen/iup/iupdraw"
	"github.com/alanyuen/iup/iupgl"
)

func main() {
	iup.Open()
	defer iup.Close()

	iupgl.CanvasOpen()

	canvas := iupgl.Canvas("ACTION")
	iup.SetCallback(canvas, "ACTION", redraw)
	//iup.SetAttribute(canvas, "BUFFER", "DOUBLE")
	iup.SetAttribute(canvas, "RASTERSIZE", "123x200")

	/* Creating main dialog */
	dialog := iup.Dialog(iup.Hbox(canvas))
	iup.SetAttribute(dialog, "TITLE", "IupGLDialog")

	iup.Show(dialog)
	iup.MainLoop()
}

func redraw(self iup.Ihandle, x float32, y float32) int {

	iupdraw.Begin(self)
	var w int
	var h int
	iupdraw.GetSize(self, &w, &h)
	fmt.Printf("%dX%d\n", w, h)
	iup.SetAttribute(self, "DRAWCOLOR", "255 0 0")
	iupdraw.Line(self, 0, 0, 100, 100)
	iup.SetAttribute(self, "DRAWCOLOR", "0 255 0")
	iupdraw.Rectangle(self, 0, 0, 100, 100)

	//	iup.SetAttribute(self, "DRAWCOLOR", "0 0 255")
	//	iupdraw.Arc(self, 0, 0, 100, 100, -90, 90)

	iup.SetAttribute(self, "DRAWCOLOR", "0 0 255")
	text := "01234567890"
	iupdraw.Text(self, text, len(text), 0, 110)

	//	iup.SetAttribute(self, "DRAWCOLOR", "255 0 0")
	//	points := []int{300, 310, 10, 10, 100, 10}
	//	iupdraw.Polygon(self, points, len(points)/3)

	iupdraw.End(self)
	return iup.DEFAULT
}

func quit_cb(ih iup.Ihandle) int {
	return iup.CLOSE
}

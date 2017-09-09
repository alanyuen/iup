package iupdraw

import (
	"unsafe"

	"github.com/alanyuen/iup"
)

/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -L${SRCDIR}/../lib
#cgo LDFLAGS: -liupgl -liup
#cgo LDFLAGS: -lopengl32 -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iupdraw.h>
*/
import "C"

//void IupDrawBegin(Ihandle* ih);
func Begin(ih iup.Ihandle) {
	C.IupDrawBegin(pih(ih))
}

//void IupDrawEnd(Ihandle* ih);
func End(ih iup.Ihandle) {
	C.IupDrawEnd(pih(ih))
}

/* all primitives can be called only between calls to Begin and End */

//void IupDrawSetClipRect(Ihandle* ih, int x1, int y1, int x2, int y2);
func SetClipRect(ih iup.Ihandle, x1 int, y1 int, x2 int, y2 int) {
	C.IupDrawSetClipRect(pih(ih), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

//void IupDrawResetClip(Ihandle* ih);
func ResetClip(ih iup.Ihandle) {
	C.IupDrawResetClip(pih(ih))
}

/* color controlled by the attribute DRAWCOLOR */
/* line style or fill controlled by the attribute DRAWSTYLE */

//void IupDrawParentBackground(Ihandle* ih);
func ParentBackground(ih iup.Ihandle) {
	C.IupDrawParentBackground(pih(ih))
}

//void IupDrawLine(Ihandle* ih, int x1, int y1, int x2, int y2);
func Line(ih iup.Ihandle, x1 int, y1 int, x2 int, y2 int) {
	C.IupDrawLine(pih(ih), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

//void IupDrawRectangle(Ihandle* ih, int x1, int y1, int x2, int y2);
func Rectangle(ih iup.Ihandle, x1 int, y1 int, x2 int, y2 int) {
	C.IupDrawRectangle(pih(ih), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

//void IupDrawArc(Ihandle* ih, int x1, int y1, int x2, int y2, double a1, double a2);
func Arc(ih iup.Ihandle, x1 int, y1 int, x2 int, y2 int, a1 float64, a2 float64) {
	C.IupDrawArc(pih(ih), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.double(a1), C.double(a2))
}

//void IupDrawPolygon(Ihandle* ih, int* points, int count);
func Polygon(ih iup.Ihandle, points []int, count int) {
	C.IupDrawPolygon(pih(ih), (*C.int)(unsafe.Pointer(&points[0])), C.int(count))
}

//void IupDrawText(Ihandle* ih, const char* text, int len, int x, int y);
func Text(ih iup.Ihandle, text string, length int, x int, y int) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.IupDrawText(pih(ih), c_text, C.int(length), C.int(x), C.int(y))
}

//void IupDrawImage(Ihandle* ih, const char* name, int make_inactive, int x, int y);
func Image(ih iup.Ihandle, name string, make_inactive int, x int, y int) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.IupDrawImage(pih(ih), c_name, C.int(make_inactive), C.int(x), C.int(y))
}

//void IupDrawSelectRect(Ihandle* ih, int x1, int y1, int x2, int y2);
func SelectRect(ih iup.Ihandle, x1 int, y1 int, x2 int, y2 int) {
	C.IupDrawSelectRect(pih(ih), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

//void IupDrawFocusRect(Ihandle* ih, int x1, int y1, int x2, int y2);
func FocusRect(ih iup.Ihandle, x1 int, y1 int, x2 int, y2 int) {
	C.IupDrawFocusRect(pih(ih), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

//void IupDrawGetSize(Ihandle* ih, int *w, int *h);
func GetSize(ih iup.Ihandle, w *int, h *int) {
	C.IupDrawGetSize(pih(ih), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//void IupDrawGetTextSize(Ihandle* ih, const char* str, int *w, int *h);
func GetTextSize(ih iup.Ihandle, str string, w *int, h *int) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.IupDrawGetTextSize(pih(ih), c_str, (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//void IupDrawGetImageInfo(const char* name, int *w, int *h, int *bpp);
func GetImageInfo(name string, w *int, h *int, bpp *int) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.IupDrawGetImageInfo(c_name, (*C.int)(unsafe.Pointer(w)),
		(*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(bpp)))
}

/* -------------------------------------------------------------------------- */

func pih(ih iup.Ihandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(ih))
}

func mkih(p *C.Ihandle) iup.Ihandle {
	return iup.Ihandle(unsafe.Pointer(p))
}

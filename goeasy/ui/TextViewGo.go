package ui

//
//	"image"
//	//	"image/color"
//	"image/png"
//	"io/ioutil"
//	"os"
import (
	"fmt"
	"image"

	"GoEasy/goeasy/uimanager/layoutparam"

	"GoEasy/goeasy/util"
	//	"../uimanager/layoutparam"
	"io/ioutil"

	//	"image/color"
	"image/draw"
	"log"

	"github.com/golang/freetype/truetype"

	//	"math"

	"github.com/golang/freetype"
	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/geom"
)

const (
//	dx = 100 // 图片的大小 宽度
//	dy = 40  // 图片的大小 高度
//	fontFile = "Lantingcuhei.TTF" // 需要使用的字体文件
//	fontSize = 18 // 字体尺寸
//	fontDPI  = 72 // 屏幕每英寸的分辨率
)

type TextViewGo struct {
	SizeEntry   *layoutparam.SizeEntry
	fontFile    string
	fontSize    float64 // 字体尺寸
	fontDPI     float64 // 屏幕每英寸的分辨率
	font        *truetype.Font
	sz          size.Event
	freeContext *freetype.Context
	text        string
	image       *glutil.Image
	X           int
	Y           int
	//	backgroundColor *Uniform
}

func (context *TextViewGo) SetSize(width, height int) {
	if context.SizeEntry == nil {
		context.SizeEntry = layoutparam.NewLayoutParam(width, height)
	} else {
		context.SizeEntry.SetSize(width, height)
		fmt.Println(context.SizeEntry.Width)
	}
}

func (context *TextViewGo) SetPosition(x, y int) {
	context.X = x
	context.Y = y
}
func (context *TextViewGo) SetX(x int) {
	context.X = x
}
func (context *TextViewGo) SetY(y int) {
	context.Y = y
}

func (context *TextViewGo) SetTypeface(fontFilePath string) {
	context.fontFile = fontFilePath
	//	fontBytes, err := ioutil.ReadFile(context.fontFile)
	fontFile, err := asset.Open(context.fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	fontBytes, err := ioutil.ReadAll(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	context.font = font
	context.freeContext.SetFont(context.font)
}

func (context *TextViewGo) SetTextSize(textSize float64) {
	context.fontSize = textSize
	context.freeContext.SetFontSize(float64(context.fontSize))
}

func (context *TextViewGo) SetFontDpi(dpi float64) {
	context.fontDPI = dpi
	context.freeContext.SetDPI(float64(context.fontDPI))
}

func (context *TextViewGo) SetText(text string) {
	context.text = text
}

func (context *TextViewGo) Draw(canvas *glutil.Images, sz size.Event) {
	if sz.WidthPx == 0 && sz.HeightPx == 0 || context.font == nil {
		log.Println("sz is width ==0 height ==0 or font is nil")
		return
	}
	ptScale := float32(sz.HeightPx) / float32(sz.HeightPt)
	scWidth, scHeight, fontSize := int(float32(context.SizeEntry.Width)*ptScale), int(float32(context.SizeEntry.Height)*ptScale), (float64(context.fontSize) * float64(ptScale))
	if context.sz != sz {
		context.sz = sz
		if context.image != nil {
			context.image.Release()
		}
		context.image = canvas.NewImage(scWidth, scHeight)
	}
	draw.Draw(context.image.RGBA, context.image.RGBA.Bounds(), image.Black, image.Point{}, draw.Src)

	freeCon := context.freeContext
	freeCon.SetClip(context.image.RGBA.Bounds())
	freeCon.SetDst(context.image.RGBA)
	freeCon.SetSrc(image.White)
	freeCon.SetFontSize(fontSize)

	//	pt := freetype.Pt(10, (40-util.GetRealSize(context.fontSize, freeCon))/2+util.GetRealSize(context.fontSize, freeCon))
	//	log.Println((40-util.GetRealSize(context.fontSize, freeCon))/2 + util.GetRealSize(context.fontSize, freeCon))
	pt := freetype.Pt(0, 0)
	//	log.Println((10 + util.GetRealSize(context.fontSize, freeCon)))
	left, top := 10, (scHeight-util.GetRealSize(fontSize, freeCon))/2
	_, maxOffsetY, _ := util.GetTextLength(freeCon, context.text, pt)
	//	log.Println("length111111:X:%d;Y:%d", ptest.X, ptest.Y)
	_, err := freeCon.DrawString(context.text, pt, maxOffsetY, left, top)
	//	log.Println("length:X:%d;Y:%d", length.X, length.Y)
	if err != nil {
		log.Println(err)
		return
	}
	//	log.Println("sz:width", sz.HeightPt, sz.HeightPx)
	context.image.Upload()
	//	ptScale := float32(sz.HeightPt) / float32(sz.HeightPx)
	//	context.image.Draw(
	//		sz,
	//		geom.Point{sz.WidthPt - geom.Pt(float32(context.SizeEntry.Width)*ptScale), sz.HeightPt - geom.Pt(float32(context.SizeEntry.Height)*float32(ptScale))},
	//		geom.Point{sz.WidthPt, sz.HeightPt - geom.Pt(float32(context.SizeEntry.Height)*ptScale)},
	//		geom.Point{sz.WidthPt - geom.Pt(float32(context.SizeEntry.Width)*ptScale), sz.HeightPt},
	//		context.image.RGBA.Bounds(),
	//	)
	log.Println("Y:", context.Y)
	context.image.Draw(
		sz,
		geom.Point{geom.Pt(context.X), geom.Pt(context.Y)},
		geom.Point{geom.Pt(context.SizeEntry.Width + context.X), geom.Pt(context.Y)},
		geom.Point{geom.Pt(context.X), geom.Pt(context.Y + context.SizeEntry.Height)},
		context.image.RGBA.Bounds(),
	)
}

func NewTextView() *TextViewGo {
	context := &TextViewGo{
		fontSize:    12,
		fontDPI:     72,
		X:           0,
		Y:           0,
		freeContext: freetype.NewContext(),
	}
	context.SetTypeface(util.NormalFontFile)
	context.freeContext.SetDPI(float64(context.fontDPI))
	context.freeContext.SetFont(context.font)
	context.freeContext.SetFontSize(float64(context.fontSize))
	return context
}

func NewTextViewWithSize(sizeEntry *layoutparam.SizeEntry) *TextViewGo {
	context := &TextViewGo{
		SizeEntry:   sizeEntry,
		fontSize:    12,
		fontDPI:     72,
		X:           0,
		Y:           0,
		freeContext: freetype.NewContext(),
	}
	context.SetTypeface(util.NormalFontFile)
	return context
}

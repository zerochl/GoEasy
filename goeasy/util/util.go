package util

import (
	"errors"

	//	"log"

	"GoEasy/goeasy/baseutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	NormalFontFile = "Lantingcuhei.TTF" // 需要使用的字体文件
)

func GetRealSize(oldSize float64, c *freetype.Context) int {
	return int(c.PointToFixed(oldSize) >> 6)
}

func GetTextLength(context *freetype.Context, s string, p fixed.Point26_6) (fixed.Point26_6, int, error) {
	if context.GetFont() == nil {
		return fixed.Point26_6{}, 0, errors.New("freetype: DrawText called with a nil font")
	}
	prev, hasPrev := truetype.Index(0), false
	maxOffsetY := 0
	for _, rune := range s {
		index := context.GetFont().Index(rune)
		//		log.Println("index:", index)
		if hasPrev {
			kern := context.GetFont().Kern(context.GetScale(), prev, index)
			if context.GetHinting() != font.HintingNone {
				kern = (kern + 32) &^ 63
			}
			p.X += kern
		}
		advanceWidth, _, offset, err := context.Glyph(index, p)
		//		log.Println("advanceWidth:%d", advanceWidth)
		if err != nil {
			return fixed.Point26_6{}, 0, err
		}
		//		offset.Y = 0
		if baseutil.Abs(offset.Y) > maxOffsetY {
			maxOffsetY = baseutil.Abs(offset.Y)
		}
		p.X += advanceWidth
		//glyphRect := mask.Bounds().Add(offset)
		//dr := c.clip.Intersect(glyphRect)
		//		if !dr.Empty() {
		//			mp := image.Point{0, dr.Min.Y - glyphRect.Min.Y}
		//			draw.DrawMask(c.dst, dr, c.src, image.ZP, mask, mp, draw.Over)
		//		}
		prev, hasPrev = index, true
	}
	//	log.Println("maxOffsetY:", maxOffsetY)
	return p, maxOffsetY, nil
}

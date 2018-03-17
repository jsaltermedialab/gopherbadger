package main

import (
	"errors"
	"fmt"

	"github.com/fogleman/gg"
)

func drawBadge(coveragePct float64, filename string) error {
	//Grey
	colorGrey := "#777"
	colorDarkGrey := "#333"
	//Green: >= 80% overall coverage
	colorGreen := "#00cc1e"
	colorDarkGreen := "#049100"
	//Yellow: 65% <= overall coverage < 80%
	colorYellow := "#e2bd00"
	colorDarkYellow := "#c6a601"
	//Red: < 65% overall coverage
	colorRed := "#db1a08"
	colorDarkRed := "#a31204"
	var accentColor, accentBorderColor string
	if coveragePct >= 80 {
		accentColor = colorGreen
		accentBorderColor = colorDarkGreen
	} else if coveragePct >= 55 {
		accentColor = colorYellow
		accentBorderColor = colorDarkYellow
	} else if coveragePct >= 0 {
		accentColor = colorRed
		accentBorderColor = colorDarkRed
	} else {
		return errors.New("Coverage value must be >= 0%")
	}
	//Create graphics context
	dc := gg.NewContext(600, 120)

	//Draw background rectangle
	dc.DrawRoundedRectangle(6, 6, 600-6*2, 120-6*2, 10)
	dc.SetHexColor(accentColor)
	dc.FillPreserve()
	dc.SetHexColor(accentBorderColor)
	dc.SetLineWidth(6.0)
	dc.Stroke()

	//Draw coverage background rectangle
	dc.DrawRoundedRectangle(10, 10, 410-10*2, 120-10*2, 5)
	dc.SetHexColor(colorDarkGrey)
	dc.FillPreserve()
	dc.SetHexColor(colorGrey)
	dc.SetLineWidth(2.0)
	dc.Stroke()

	//Drawing text
	err := dc.LoadFontFace("fonts/luxisr.ttf", 82)
	errCheck("Loading font", err)
	dc.SetHexColor("#ffffffff")
	dc.DrawString("Coverage:", 5+10, 120-5*2-27.5)
	covPctString := fmt.Sprintf("%2.f", coveragePct) + "%"
	dc.DrawString(covPctString, 410+5, 120-5*2-22)
	//Save to file
	err = dc.SavePNG(filename)
	errCheck("Saving image file", err)
	return err
}

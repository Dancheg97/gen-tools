package img

import "testing"

func TestPng(t *testing.T) {
	SvgToPng(`test.svg`, `out.png`, 512)
	PngToJpeg(`out.png`, `wht.png`, 180)
}

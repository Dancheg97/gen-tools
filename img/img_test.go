package svg

import "testing"

func TestPng(t *testing.T) {
	SvgToPng(`test.svg`, `out.png`, 512, 512)
}

package svg

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"

	"dancheg97.ru/dancheg97/gen-tools/utils"
	"github.com/sirupsen/logrus"
)

func SvgToPng(inp, outp string, w, h int) {
	in, err := os.Open(inp)
	utils.CheckErr(err)

	icon, err := oksvg.ReadIconStream(in)
	utils.CheckErr(err)
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	out, err := os.Create(outp)
	utils.CheckErr(err)

	err = png.Encode(out, rgba)
	utils.CheckErr(err)
}

func ResizeSvg(in, out string, w, h int) {
	inp, err := os.ReadFile(in)
	utils.CheckErr(err)

	first := strings.Split(string(inp), `<svg `)
	if len(first) != 2 {
		logrus.Error(`wrong svg format`)
		os.Exit(1)
	}
	second := strings.Split(first[1], `>`)

	second[0] = UpdateArrt(second[0], `width`, fmt.Sprint(w))
	second[0] = UpdateArrt(second[0], `height`, fmt.Sprint(h))

	first[1] = strings.Join(second, `>`)

	output := strings.Join(first, `<svg `)

	err = os.WriteFile(out, []byte(output), 0o600)
	utils.CheckErr(err)
}

func UpdateArrt(in, attr, val string) string {
	if strings.Contains(in, attr) {
		first := strings.Split(in, attr+`="`)
		second := strings.Split(first[1], `"`)
		second[0] = val
		first[1] = strings.Join(second, `"`)
		return strings.Join(first, attr+`="`)
	}
	return attr + `="` + val + `" ` + in
}

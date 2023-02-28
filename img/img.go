package img

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"

	"dancheg97.ru/dancheg97/gen-tools/utils"
	"github.com/sirupsen/logrus"
)

func PngToJpeg(in, out string, size int) {
	pngImgFile, err := os.Open(in)
	utils.CheckErr(err)

	imgSrc, err := png.Decode(pngImgFile)
	utils.CheckErr(err)

	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	jpgImgFile, err := os.Create(out)
	utils.CheckErr(err)

	var opt jpeg.Options
	opt.Quality = size
	err = jpeg.Encode(jpgImgFile, newImg, &opt)
	utils.CheckErr(err)
}

func SvgToPng(inp, out string, size int) {
	in, err := os.Open(inp)
	utils.CheckErr(err)

	icon, err := oksvg.ReadIconStream(in)
	utils.CheckErr(err)
	icon.SetTarget(0, 0, float64(size), float64(size))
	rgba := image.NewRGBA(image.Rect(0, 0, size, size))
	gvscanner := rasterx.NewScannerGV(size, size, rgba, rgba.Bounds())
	icon.Draw(rasterx.NewDasher(size, size, gvscanner), 1)

	output, err := os.Create(out)
	utils.CheckErr(err)

	err = png.Encode(output, rgba)
	utils.CheckErr(err)
}

func ResizeSvg(in, out string, size int) {
	inp, err := os.ReadFile(in)
	utils.CheckErr(err)

	first := strings.Split(string(inp), `<svg `)
	if len(first) != 2 {
		logrus.Error(`wrong svg format`)
		os.Exit(1)
	}
	second := strings.Split(first[1], `>`)

	second[0] = UpdateArrt(second[0], `width`, fmt.Sprint(size))
	second[0] = UpdateArrt(second[0], `height`, fmt.Sprint(size))

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

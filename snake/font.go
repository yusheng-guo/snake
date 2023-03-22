package snake

import (
	_ "embed"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed `..\assets\Comic Sans MS.ttf`
var comic []byte

func loadLocalFont(size float64) (font.Face, error) {
	font, err := freetype.ParseFont(comic)
	return truetype.NewFace(font, &truetype.Options{
		Size: size,
		DPI:  72,
	}), err
}

// func loadGoregularFont(size float64) (font.Face, error) {
// 	ttfFont, err := truetype.Parse(goregular.TTF)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return truetype.NewFace(ttfFont, &truetype.Options{
// 		Size:    size,
// 		DPI:     72,
// 		Hinting: font.HintingFull,
// 	}), err
// }

// func loadLocalFont(fontFilename string, size float64) (font.Face, error) {
// 	fontBytes, err := os.ReadFile(fontFilename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	font, err := freetype.ParseFont(fontBytes)

// 	return truetype.NewFace(font, &truetype.Options{
// 		Size: size,
// 		DPI:  72,
// 	}), err
// }

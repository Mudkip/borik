package bot

import (
	"fmt"

	"gopkg.in/gographics/imagick.v2/imagick"
)

type MagikArgs struct {
	ImageURL string  `default:"" description:"URL to the image to process. Leave blank to automatically attempt to find an image."`
	Scale    float64 `default:"1" description:"Scale of the magikification. Larger numbers produce more destroyed images."`
}

func (args MagikArgs) GetImageURL() string {
	return args.ImageURL
}

func magikHelper(wand *imagick.MagickWand, args MagikArgs, wMultiplier float64, hMultiplier float64) ([]*imagick.MagickWand, error) {
	width := wand.GetImageWidth()
	height := wand.GetImageHeight()

	err := wand.LiquidRescaleImage(uint(float64(width)*wMultiplier), uint(float64(height)*hMultiplier), args.Scale, 0)
	if err != nil {
		return nil, fmt.Errorf("error while attempting to liquid rescale: %w", err)
	}

	err = wand.ResizeImage(width, height, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		return nil, fmt.Errorf("error while attempting to resize image: %w", err)
	}

	return []*imagick.MagickWand{wand}, nil
}

// Magik runs content-aware scaling on an image.
func Magik(wand *imagick.MagickWand, args MagikArgs) ([]*imagick.MagickWand, error) {
	return magikHelper(wand, args, 0.5, 0.5)
}

// Lagik runs content-aware scaling on an image.
func Lagik(wand *imagick.MagickWand, args MagikArgs) ([]*imagick.MagickWand, error) {
	return magikHelper(wand, args, 1.5, 1.5)
}

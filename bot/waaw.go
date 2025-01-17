package bot

import (
	"fmt"

	"gopkg.in/gographics/imagick.v2/imagick"
)

type mirrorDirection string

const (
	mirrorDirectionVertical   mirrorDirection = "vertical"
	mirrorDirectionHorizontal mirrorDirection = "horizontal"
)

func mirrorImage(wand *imagick.MagickWand, direction mirrorDirection, flipped bool) ([]*imagick.MagickWand, error) {
	var desiredGravity imagick.GravityType
	if direction == mirrorDirectionHorizontal {
		if flipped {
			desiredGravity = imagick.GRAVITY_EAST
		} else {
			desiredGravity = imagick.GRAVITY_WEST
		}
	} else {
		if flipped {
			desiredGravity = imagick.GRAVITY_SOUTH
		} else {
			desiredGravity = imagick.GRAVITY_NORTH
		}
	}

	err := wand.SetImageGravity(desiredGravity)
	if err != nil {
		return nil, fmt.Errorf("error setting gravity: %w", err)
	}

	var half *imagick.MagickWand
	var xOffset, yOffset int

	if direction == mirrorDirectionHorizontal {
		half = wand.TransformImage("50%x100%", "")
		err = half.FlopImage()
		if err != nil {
			return nil, fmt.Errorf("error flipping image: %w", err)
		}

		if flipped {
			xOffset = 0
			yOffset = 0
		} else {
			xOffset = int(half.GetImageWidth())
			yOffset = 0
		}
	} else {
		half = wand.TransformImage("100%x50%", "")
		err = half.FlipImage()
		if err != nil {
			return nil, fmt.Errorf("error flipping image: %w", err)
		}

		if flipped {
			xOffset = 0
			yOffset = 0
		} else {
			xOffset = 0
			yOffset = int(half.GetImageHeight())
		}
	}

	err = wand.CompositeImage(half, imagick.COMPOSITE_OP_ATOP, xOffset, yOffset)
	if err != nil {
		return nil, fmt.Errorf("error compositing image: %w", err)
	}

	return []*imagick.MagickWand{wand}, nil

}

type WaawArgs struct {
	ImageURL string `default:"" description:"URL to the image to process. Leave blank to automatically attempt to find an image."`
}

func (args WaawArgs) GetImageURL() string {
	return args.ImageURL
}

func Waaw(wand *imagick.MagickWand, args WaawArgs) ([]*imagick.MagickWand, error) {
	return mirrorImage(wand, mirrorDirectionHorizontal, true)
}

type HaahArgs struct {
	ImageURL string `default:"" description:"URL to the image to process. Leave blank to automatically attempt to find an image."`
}

func (args HaahArgs) GetImageURL() string {
	return args.ImageURL
}

func Haah(wand *imagick.MagickWand, args HaahArgs) ([]*imagick.MagickWand, error) {
	return mirrorImage(wand, mirrorDirectionHorizontal, false)
}

type WoowArgs struct {
	ImageURL string `default:"" description:"URL to the image to process. Leave blank to automatically attempt to find an image."`
}

func (args WoowArgs) GetImageURL() string {
	return args.ImageURL
}

func Woow(wand *imagick.MagickWand, args WoowArgs) ([]*imagick.MagickWand, error) {
	return mirrorImage(wand, mirrorDirectionVertical, false)
}

type HoohArgs struct {
	ImageURL string `default:"" description:"URL to the image to process. Leave blank to automatically attempt to find an image."`
}

func (args HoohArgs) GetImageURL() string {
	return args.ImageURL
}

func Hooh(wand *imagick.MagickWand, args HoohArgs) ([]*imagick.MagickWand, error) {
	return mirrorImage(wand, mirrorDirectionVertical, true)
}

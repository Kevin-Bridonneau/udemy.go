package filter

import (
	"github.com/disintegration/imaging"
)

// Filter interface
type Filter interface {
	Process(srcPath, dstPath string) error
}

// Grayscale struct
type Grayscale struct{}

// Process grayscale
func (g Grayscale) Process(srcPath, dstPath string) error {

	// Open the src image.
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	dst := imaging.Grayscale(src)
	err = imaging.Save(dst, dstPath+"_grayscale.jpg")
	if err != nil {
		return err
	}
	return nil
}

// Blur struct
type Blur struct{}

// Process Blur
func (b Blur) Process(srcPath, dstPath string) error {

	// Open the src image.
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	dst := imaging.Blur(src, 1)
	err = imaging.Save(dst, dstPath+"_blur.jpg")
	if err != nil {
		return err
	}
	return nil
}

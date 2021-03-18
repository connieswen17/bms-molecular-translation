package main

import (
    "fmt"
    "errors"
    "gocv.io/x/gocv"
)

// By Preprocess, we mean:
//   + ensuring that the image is RGB
//   + ensuring that the image is (inverted) White on Black
// This function returns an image loaded from disk into memory
func loadAndPreprocessImage(listing string, isBlackOnWhite bool) (gocv.Mat, error) {
    img := gocv.IMRead(listing, gocv.IMReadUnchanged)
    defer img.Close()
    // when IMRead fails it returns an empty Mat, but we want
    // to give the caller the option to case on errors instead
    // of only checking for empty Mats
    if img.Empty() {
        err := errors.New("IMRead returned an empty image")
        return img, err
    }

    // most image processing functions require RGB images, not BGR
    gocv.CvtColor(img, &img, gocv.ColorBGRToRGB)

    // when we open (morph) the image we must
    // provide a white on black image
    if isBlackOnWhite {
        gocv.BitwiseNot(img, &img)
    }

    return img, nil
}

// Typically, the competition data has noise in its images
// This function passes an Opening morphology transformation
// over the given image in an attempt to remove noisy grain
func removeNoise(img gocv.Mat) (gocv.Mat, error) {
    // TODO
    return img, nil
}

func main() {
    fmt.Println("Hello World!")
}

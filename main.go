package main

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println("Error: Cannot open webcam")
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Object Detection")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	grayImg := gocv.NewMat()
	defer grayImg.Close()

	hsvImg := gocv.NewMat()
	defer hsvImg.Close()

	mask := gocv.NewMat()
	defer mask.Close()

	red := color.RGBA{255, 0, 0, 0}

	faceCascade := gocv.NewCascadeClassifier()
	defer faceCascade.Close()

	xmlFile := "haarcascade_frontalface_default.xml"

	if !faceCascade.Load(xmlFile) {
		fmt.Printf("Error reading cascade file %v\n", xmlFile)
		return
	}

lowerBound := gocv.NewScalar(35, 100, 100, 0) // Green
upperBound := gocv.NewScalar(85, 255, 255, 0)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device")
			return
		}

		if img.Empty() {
			continue
		}

		gocv.CvtColor(img, &grayImg, gocv.ColorBGRToGray)

		rects := faceCascade.DetectMultiScale(grayImg)

		for _, r := range rects {
			gocv.Rectangle(&img, r, red, 5)
		}

		gocv.CvtColor(img, &hsvImg, gocv.ColorBGRToHSV)
		gocv.InRangeWithScalar(hsvImg, lowerBound, upperBound, &mask)

		window.IMShow(img)

		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
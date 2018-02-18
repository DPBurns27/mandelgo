package main

import "fmt"
import "math"
import "github.com/fogleman/gg"
//import "sync"

func main() {
    fmt.Println("Running...")
    frames := 1
    zoomFactor := 0.5
    focalPointRe := 0.0
    focalPointIm := 1.0
    size := 1000
    escapeLoops := 100
    threshhold := 4.0
    canvas := gg.NewContext(size, size)
    //canvas.DrawCircle(500,500,400)
    //canvas.SetRGB(0,0,0)
    //canvas.Fill()
    
    for k := 0; k < frames; k++ {
	// how the size of the window changes each time it zooms in
	windowSize := 4.0*math.Pow(zoomFactor,float64(k))
    
	complexDomainRe, complexDomainIm := complexDomain(windowSize, focalPointRe, focalPointIm, size)

        // i dont know if i is the horizontal or vertical axis
        // Loop the canvas
        for i := 0; i < size; i++ {
	    for j := 0; j < size; j++ {
		// find the complex representation of the canvas position
	    	// also need to be careful about how i and j grow across the canvas
	    	colour := escapeTest(complexDomainRe[i], complexDomainIm[j], escapeLoops, threshhold)
	    	canvas.SetRGB(colour, colour, colour)
	    	canvas.SetPixel(i, j)
	    }
        }

	filename := fmt.Sprintf("/Users/David/go/src/github.com/DPBurns27/mandelgo/out/out%d.png", k)
	fmt.Println("Writing file:", filename)
        canvas.SavePNG(filename)
    }
}

// Work out the conversion factors for moving from the pixel positions to the imaginary and real domain values
// TODO: generate arrays using the size of the canvas 
func complexDomain(windowSize, focalPointRe, focalPointIm float64, size int) ([]float64, []float64) {
	cReMin := focalPointRe - windowSize/2
	cReMax := focalPointRe + windowSize/2
	cImMin := focalPointIm - windowSize/2
	cImMax := focalPointIm + windowSize/2

        canvasConversionRe := (cReMax - cReMin)/float64(size)
        canvasConversionIm := (cImMax - cImMin)/float64(size)

	complexDomainRe := make([]float64, size)
	complexDomainIm := make([]float64, size)

	for i := 0; i < size; i++ {
	    complexDomainRe[i] = cReMin + float64(i)*canvasConversionRe
	    complexDomainIm[i] = cImMax - float64(i)*canvasConversionIm
	}

	return complexDomainRe, complexDomainIm
}

// zIm,a is the imaginary component, zRe,cRe is the real
func zCubed(zRe, zIm, cRe, cIm float64) (float64, float64) {
    zReNext := cRe - 3*math.Pow(zIm, 2)*zRe + math.Pow(zRe, 3)
    zImNext := cIm - math.Pow(zIm, 3) + 3*zIm*math.Pow(zRe, 2)
    return zReNext, zImNext 
}

// zIm,cIm is the imaginary component, zRe,cRe is the real
func zSquared(zRe, zIm, cRe, cIm float64) (float64, float64) {
    zReNext := cRe - math.Pow(zIm, 2) + math.Pow(zRe, 2)
    zImNext := cIm + 2*zIm*zRe 
    return zReNext, zImNext 
}

func escapeTest(cRe, cIm float64, maxLoop int, threshhold float64) float64 {
    zRe := 0.0
    zIm := 0.0
    for i := 0; i < maxLoop; i++ {
	zRe, zIm = zSquared(zRe, zIm, cRe, cIm)
	if (math.Pow(zRe,2)+math.Pow(zIm,2)) > threshhold {
	    return float64(i)/float64(maxLoop)
	}
    }
    //fmt.Println(math.Hypot(zRe, zIm))
    return 0.0
}

//func makeMovie(file string) {
    //cmd := exec.Command("ffmpeg -r 25 -i %04d.bmp -f mp4 -c:v libx264 -preset ultrafast -qp 0 -pix_fmt yuv420p output3.mp4")
//}

package main

import "fmt"
import "math"
import "github.com/fogleman/gg"

func main() {
    fmt.Println("Running...")
    frames := 10
    zoomFactor := 0.5
    zoomRe := 0.25
    zoomIm := 0.0
    size := 500
    canvas := gg.NewContext(size, size)
    //canvas.DrawCircle(500,500,400)
    //canvas.SetRGB(0,0,0)
    //canvas.Fill()

    for k := 0; k < frames; k++ {
        // work out conversion factors for going from the canvas position to the actual imaginary and 
        // real axis values you want
        // cIm is imaginary, cRe is real
        cReMin := -3.0*math.Pow(zoomFactor,float64(k)) 
        cReMax := 1.0*math.Pow(zoomFactor,float64(k))
        cImMin := -2.0*math.Pow(zoomFactor,float64(k)) 
        cImMax := 2.0*math.Pow(zoomFactor,float64(k)) 
        reCanvasConversion := (cReMax - cReMin)/float64(size)
        imCanvasConversion := (cImMax - cImMin)/float64(size)
    
        // i dont know if i is the horizontal or vertical axis
        // Loop the canvas
        for i := 0; i < size; i++ {
    	cRe := cReMin + float64(i)*reCanvasConversion
    	for j := 0; j < size; j++ {
    	    // find the complex representation of the canvas position
    	    // also need to be careful about how i and j grow across the canvas
    	    cIm := cImMax - float64(j)*imCanvasConversion
    	    colour := escapeTest(cRe, cIm, 100, 4)
    	    canvas.SetRGB(colour, colour, colour)
    	    canvas.SetPixel(i, j)
    	}
        }
	filename := fmt.Sprintf("out/out%d.png", k) 
        canvas.SavePNG(filename)
    }
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

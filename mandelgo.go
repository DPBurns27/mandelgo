package main

//import "fmt"
import "math"
import "github.com/fogleman/gg"

func main() {
    //fmt.Println(complexFunction(2,3))
    size := 1000
    canvas := gg.NewContext(size, size)
    //canvas.DrawCircle(500,500,400)
    //canvas.SetRGB(0,0,0)
    //canvas.Fill()

    // work out conversion factors for going from the canvas position to the actual imaginary and 
    // real axis values you want
    // cIm is imaginary, cRe is real
    cImMin := -2.0
    cImMax := 2.0
    cReMin := -3.0
    cReMax := 1.0
    reCanvasConversion := (cReMax - cReMin)/float64(size)
    imCanvasConversion := (cImMax - cImMin)/float64(size)

    // i dont know if i is the horizontal or vertical axis
    // Loop the canvas
    for i := 0; i < size; i++ {
	for j := 0; j < size; j++ {
	    // find the complex representation of the canvas position
	    // also need to be careful about how i and j grow across the canvas
	    cRe := cReMin + float64(i)*reCanvasConversion
	    cIm := cImMax - float64(j)*imCanvasConversion
	    colour := escapeTest(cIm, cRe, 20, 0.1)
	    canvas.SetRGB(colour, colour, colour)
	    canvas.SetPixel(i, j)
	}
    }

    canvas.SavePNG("out/out.png")
}

// zIm,a is the imaginary component, zRe,cRe is the real
func zCubed(zIm, zRe, cIm, cRe float64) (float64, float64) {
    zReNext := cRe - 3*math.Pow(zIm, 2)*zRe + math.Pow(zRe, 3)
    zImNext := cIm - math.Pow(zIm, 3) + 3*zIm*math.Pow(zRe, 2)
    return zImNext, zReNext
}

// zIm,cIm is the imaginary component, zRe,cRe is the real
func zSquared(zIm, zRe, cIm, cRe float64) (float64, float64) {
    zReNext := cRe - math.Pow(zIm, 2) + math.Pow(zRe, 2)
    zImNext := cIm + 2*zIm*zRe 
    return zImNext, zReNext
}

func escapeTest(cIm, cRe float64, maxLoop int, threshhold float64) float64 {
    zIm := 0.0
    zRe := 0.0
    for i := 0; i < maxLoop; i++ {
	zIm, zRe = zSquared(zIm, zRe, cIm, cRe)
    }
    if (math.Hypot(zIm, zRe)) < threshhold {
	return 0.0
    }
    return 1.0
}

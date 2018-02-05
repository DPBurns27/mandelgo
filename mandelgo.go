package main

//import "fmt"
//import "math"
import "github.com/fogleman/gg"

func main() {
    //fmt.Println(complexFunction(2,3))
    var size int = 1000
    canvas := gg.NewContext(size,size)
    //canvas.DrawCircle(500,500,400)
    //canvas.SetRGB(0,0,0)
    //canvas.Fill()

    // work out conversion factors for going from the canvas position to the actual imaginary and 
    // real axis values you want
    var a_min float64 = -3
    var a_max float64 = 1
    var b_min float64 = -2
    var b_max float64 = 2
    iCanvasConversion := (a_max - a_min)/float64(size)
    jCanvasConversion := (b_max - b_min)/float64(size)

    // Loop the canvas
    for i := 0; i < size; i++ {
	for j := 0; j < size; j++ {
	    // find the complex representation of the canvas position
	    a := a_min + float64(i)*iCanvasConversion
	    b := b_max - float64(j)*jCanvasConversion // this is because y decreases downwards but x increases across
	    var colour float64 = escapeTest(a,b,20,0.0001)
	    canvas.SetRGB(colour,colour,colour)
	    canvas.SetPixel(i,j)
	}
    }

    canvas.SavePNG("out/out.png")
}

// x,a is the imaginary component, y,b is the real
func zCubed(x, y, a, b float64) (float64, float64) {
    var xn float64 = b - 3*x*x*y + y*y*y
    var yn float64 = a - x*x*x + 3*x*y*y
    return xn, yn
}

func escapeTest(a, b float64, maxLoop int, threshHold float64) float64 {
    var x float64 = 0
    var y float64 = 0
    for i := 0; i < maxLoop; i++ {
	x, y = zCubed(x, y, a, b)
    }
    if x*x + y*y < threshHold {
	return 0
    } else {
	return 1
    }
}

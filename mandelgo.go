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
    a_min := -3.0
    a_max := 1.0
    b_min := -2.0
    b_max := 2.0
    iCanvasConversion := (a_max - a_min)/float64(size)
    jCanvasConversion := (b_max - b_min)/float64(size)

    // Loop the canvas
    for i := 0; i < size; i++ {
	for j := 0; j < size; j++ {
	    // find the complex representation of the canvas position
	    a := a_min + float64(i)*iCanvasConversion
	    b := b_max - float64(j)*jCanvasConversion // this is because y decreases downwards but x increases across
	    colour := escapeTest(a,b,10,0.001)
	    canvas.SetRGB(colour,colour,colour)
	    canvas.SetPixel(i,j)
	}
    }

    canvas.SavePNG("out/out.png")
}

// x,a is the imaginary component, y,b is the real
func zCubed(x, y, a, b float64) (float64, float64) {
    var yn float64 = b - 3*x*x*y + y*y*y
    var xn float64 = a - x*x*x + 3*x*y*y
    return xn, yn
}

// x,a is the imaginary component, y,b is the real
func zSquared(x, y, a, b float64) (float64, float64) {
    var yn float64 = b - x*x + y*y
    var xn float64 = a + 2*x*y 
    return xn, yn
}

func escapeTest(a, b float64, maxLoop int, threshhold float64) float64 {
    var x float64 = 0
    var y float64 = 0
    for i := 0; i < maxLoop; i++ {
	x, y = zSquared(x, y, a, b)
    }
    if (x*x + y*y) < threshhold {
	return 0
    }
    return 1
}

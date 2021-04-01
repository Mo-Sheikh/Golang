package main
import (
	"net/http"
	"fmt"
	"log"
	"sync"

	"io"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"

	// "strings"
	"strconv"
	

	
)

var mu sync.Mutex
var count int
var srv http.Server
var query string


func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/funny", 	alrightyThen)
	http.HandleFunc("/showme", showme)

	log.Fatal(http.ListenAndServe("localhost:9010", nil))
}

func handler(w http.ResponseWriter, r *http.Request){

	// fmt.Fprintf(w, "ALRIGHTY THEN URL.path = %q\n", r.URL.Path)

	cycles := r.URL.Query().Get("cycles")







	cyclesCount, err := strconv.Atoi(cycles)
	fmt.Println("cycles count is ", cyclesCount)

	if err != nil{
		fmt.Fprintf(w, "%s","error")
}

	lissajous(w,cyclesCount)

	

	}
	


func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func alrightyThen(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Jim carrey is %d", 59)
}

func showme(w http.ResponseWriter, r *http.Request){

lissajous(w, 15)
	
}
var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0xd7, 0x00, 0xff},
	color.RGBA{0xff, 0xc0, 0xcb, 0xff},
}
// var palette = []color.Color{
// 	color.RGBA{0x00, 0x00, 0x00, 0xFF},
// 	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
// }

const (
whiteIndex = 0 // first color in palette
blackIndex = 1 // next color in palette
redIndex = 2
greyIndex = 3
blueIndex = 4
)

   
   
func lissajous(out io.Writer, cyclesCount int) {
	fmt.Println("received", cyclesCount)
	const (
	// cycles = 5 // number of complete x oscillator revolutions
	res = 0.001 // angular resolution
	size = 100 // image canvas covers [-size..+size]
	nframes = 64 // number of animation frames
	delay = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	fmt.Println("haa")


	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	img := image.NewPaletted(rect, palette)
	for t := 0.0; t < float64(cyclesCount)*2*math.Pi; t += res {
	x := math.Sin(t)
	y := math.Sin(t*freq + phase)

	img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 
	uint8(	rand.Intn(4 - 0) + 0))
	}
	phase += 0.1
	anim.Delay = append(anim.Delay, delay)
	anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = iota
	blackIndex
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for key, strings := range request.Header {
			fmt.Fprintf(writer, "Header[%q]=%q\n", key, strings)
		}
		fmt.Fprintf(writer, "addr: %q\n", request.RemoteAddr)
		fmt.Fprintf(writer, "host: %q\n", request.Host)
		fmt.Fprintf(writer, "url :%q\n", request.URL.Path)
		fmt.Fprintf(writer, "proto :%q\n", request.Proto)
		fmt.Fprintf(writer, "method :%q\n", request.Method)
		request.ParseForm()
		request.ParseMultipartForm(2048 * 2048)
		for key, strings := range request.Form {
			fmt.Fprintf(writer, "Form[%q]=%q\n", key, strings)
		}
	})

	http.HandleFunc("/lissajous", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		fmt.Println(request.FormValue("cycle"))
		float, err := strconv.ParseFloat(request.FormValue("cycle"), 64)
		if err != nil {
			lissajous(writer, cycles)
			return
		}
		lissajous(writer, float)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const (
	cycles  = 5
	res     = 0.001
	size    = 100
	nframes = 64
	delay   = 8
)

func lissajous(out io.Writer, cycle float64) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Println(err.Error())
	}
}

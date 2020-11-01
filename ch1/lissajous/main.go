// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
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
	"os"
	"strconv"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type lisConfig struct {
	cycles  float64
	res     float64
	sizef   float64
	size    int
	nframes int
	delay   int
}

var defaultConfig = &lisConfig{
	cycles:  5,     // number of complete x oscillator revolutions
	res:     0.001, // angular resolution
	sizef:   100,   // image canvas covers [-size..+size]
	size:    100,   // image canvas covers [-size..+size]
	nframes: 64,    // number of animation frames
	delay:   8,
}

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 {
		if os.Args[1] == "web" {
			//!+http
			handler := func(w http.ResponseWriter, r *http.Request) {
				lissajous(w, readConfig(r))
			}
			http.HandleFunc("/", handler)
			//!-http
			log.Fatal(http.ListenAndServe("localhost:8000", nil))
			return
		}
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "liss: %v\n", err)
		}
		lissajous(f, defaultConfig)
	} else {
		lissajous(os.Stdout, defaultConfig)
	}
}

func readConfig(r *http.Request) *lisConfig {
	var conf = &lisConfig{
		cycles:  5,     // number of complete x oscillator revolutions
		res:     0.001, // angular resolution
		sizef:   100,   // image canvas covers [-size..+size]
		size:    100,   // image canvas covers [-size..+size]
		nframes: 64,    // number of animation frames
		delay:   8,
	}
	keys, ok := r.URL.Query()["cycles"]
	if ok && len(keys[0]) > 1 {
		i, err := strconv.Atoi(keys[0])
		if err != nil {
			fmt.Printf("Can't parse res query param: %s", err)
			os.Exit(1)
		}
		conf.cycles = float64(i)
	}
	keys, ok = r.URL.Query()["res"]
	if ok && len(keys[0]) > 1 {
		i, err := strconv.ParseFloat(keys[0], 64)
		if err != nil {
			fmt.Printf("Can't parse res query param: %s", err)
			os.Exit(1)
		}
		conf.res = i
	}
	return conf
}

func lissajous(out io.Writer, config *lisConfig) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: config.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < config.nframes; i++ {
		rect := image.Rect(0, 0, 2*config.size+1, 2*config.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < config.cycles*2*math.Pi; t += config.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(config.size+int(x*config.sizef+0.5), config.size+int(y*config.sizef+0.5),
				uint8(nextColor(t)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, config.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func nextColor(cycle float64) int {
	return round(cycle)%30/10 + 1
}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

//!-main

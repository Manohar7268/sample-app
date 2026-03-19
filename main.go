package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
)

func main() {
	http.HandleFunc("/blue", blueHandler) // existing
	http.HandleFunc("/red", redHandler)   // new route
	http.ListenAndServe(":8080", nil)
}

// Existing blueHandler
func blueHandler(w http.ResponseWriter, r *http.Request) {
	// your blueHandler code here
}

// New redHandler
func redHandler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

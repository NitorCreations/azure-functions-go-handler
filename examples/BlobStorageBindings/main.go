package HttpTriggerWithReturn

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"strings"
	"time"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

func makeImage(data *function.Binary) {
	width, height := 200, 200
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	fill := color.RGBA{
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		255,
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, fill)
		}
	}

	if err := png.Encode(data.Buffer, img); err != nil {
		panic(err)
	}
}

func imageInfo(data *function.Binary) string {
	if data.Buffer.Len() == 0 {
		return "no image"
	}

	buf := data.Buffer.Bytes()[:]
	img, format, err := image.Decode(bytes.NewBuffer(buf))
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("size=%d, format=%s, color=%v",
		data.Buffer.Len(), format, img.At(0, 0))
}

func Handle(
	ctx *function.Context,
	req *function.HttpRequest,
	txtin string, // string type input bindings can be accessed as strings
	txtout *string, // string type output bindings can be accessed by *string
	imgIn *function.Binary, // binary type input bindings can be accessed as []byte or *function.Binary
	imgOut *function.Binary, // binary type output bindings can be accessed by *[]byte or *function.Binary
) *function.HttpResponse {

	msg := fmt.Sprintf("Go Gophers %s", time.Now())

	// set the output value
	*txtout = msg
	// create a new image with random color
	makeImage(imgOut)

	return &function.HttpResponse{
		Body: function.H{
			"txtin":  strings.Trim(txtin, "\""), // string type bindings are wrapped in quotes
			"txtout": msg,
			"imgin":  imageInfo(imgIn),
			"imgOut": imageInfo(imgOut),
		},
		Headers: function.ContentTypeJson(),
	}
}

package transcoder

import (
	"image/gif"
	"net/http"

	"github.com/barnacs/compy/proxy"
	"github.com/chai2010/webp"
)

type Gif struct{}

func (t *Gif) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
	img, err := gif.Decode(r)
	if err != nil {
		return err
	}
	if SupportsWebP(headers) {
		w.Header().Set("Content-Type", "image/webp")
		options := webp.Options{
			Lossless: false,
		}
		if err = webp.Encode(w, img, &options); err != nil {
			return err
		}
	} else {
		if err = gif.Encode(w, img, nil); err != nil {
			return err
		}
	}
	return nil
}

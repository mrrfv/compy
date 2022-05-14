package transcoder

import (
	"net/http"

	"github.com/barnacs/compy/proxy"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
)

type Minifier struct {
	m *minify.M
}

func NewMinifier() *Minifier {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	return &Minifier{
		m: m,
	}
}

func (t *Minifier) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
	return t.m.Minify(r.ContentType(), w, r)
}

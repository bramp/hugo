package helpers

import (
	"bytes"
	"html"

	"github.com/russross/blackfriday"
	"github.com/spf13/viper"
)

// Wraps a blackfriday.Renderer, typically a blackfriday.Html
type HugoHtmlRenderer struct {
	blackfriday.Renderer
}

func (renderer *HugoHtmlRenderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	if viper.GetBool("PygmentsCodeFences") {
		opts := viper.GetString("PygmentsOptions")
		str := html.UnescapeString(string(text))
		out.WriteString(Highlight(str, lang, opts))
	} else {
		renderer.Renderer.BlockCode(out, text, lang)
	}
}

package handler

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"social-media-feed/internal/feed_data"
)

func (h *Handler) getFilledMainTemplate(postsWithLike []feed_data.PostWithLike) (string, error) {
	tmpl, err := os.ReadFile("./web/templates/main_tmpl.html")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Not read file: %s", err))
	}

	t, err := template.New("webpage").Parse(string(tmpl))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Bad create template: %s", err))
	}

	data := struct {
		PostsWithLike []feed_data.PostWithLike
	}{
		PostsWithLike: postsWithLike,
	};

	buf := new(bytes.Buffer)

	err = t.Execute(buf, data)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Bad parsing: %s", err))
	}

	return buf.String(), nil
}
package template

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"os"
)

func GetFilledMainTemplate(logger *slog.Logger) (string, error) {
	type Posts struct {
		Title string
		PostImg string
	}

	tmpl, err := os.ReadFile("./web/templates/main_tmpl.html")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Not read file: %s", err))
	}

	t, err := template.New("webpage").Parse(string(tmpl))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Bad create template: %s", err))
	}

	data := struct {
		Posts []Posts
	}{
		Posts: []Posts{
			Posts{
				Title: "Мафбосс 1",
				PostImg: "resources/posts_image/mathboss_1.jpg",
			},
			Posts{
				Title: "Мафбосс 2",
				PostImg: "resources/posts_image/mathboss_2.jpg",
			},
		},
	};

	buf := new(bytes.Buffer)

	err = t.Execute(buf, data)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Bad parsing: %s", err))
	}

	return buf.String(), nil
}
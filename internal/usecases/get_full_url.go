package usecases

import (
	"github.com/andranikuz/shortener/internal/app"
)

func GetFullURL(id string) string {
	url, err := app.App.DB.Get(id)
	if err != nil {
		return ""
	}

	return url.FullURL
}

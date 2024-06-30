package authorize

import (
	"net/http"

	"github.com/hashicorp/go-uuid"
)

// GetOrGenerateUserID функция получения или создания идентификатора пользователя. После генерации
// идентификатор помещается в куку Authorize и используется для авторизации
func GetOrGenerateUserID(res http.ResponseWriter, req *http.Request) string {
	cookie, err := req.Cookie("Authorize")
	if err != nil || cookie.Value == "" {
		userID, _ := uuid.GenerateUUID()
		cookie = &http.Cookie{Name: "Authorize", Value: userID}
		http.SetCookie(res, cookie)
	}

	return cookie.Value
}

// GetUserID функция получения пользователя из куки Authorize
func GetUserID(req *http.Request) (string, error) {
	cookie, err := req.Cookie("Authorize")
	if err != nil || cookie.Value == "" {
		return "", err
	}

	return cookie.Value, nil
}

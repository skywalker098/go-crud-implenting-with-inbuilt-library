package utils

import (
	"net/http"
	"strconv"
	"strings"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

func CheckMethod(httpMethod string, chkMethod HttpMethod) bool {
	return httpMethod == string(chkMethod)
}

// func GetUrlId(r *http.Request) (int, error) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	return id, err
// }

func GetUrlParmId(r *http.Request) (int, error) {
	path := r.URL.Path
	isStr := strings.Split(path, "/")
	id, err := strconv.Atoi(isStr[len(isStr)-1])
	return id, err
}

// func GetQueryParams(r *http.Request) (map[string]string, error) {
// 	queryParams := make(map[string]string)
// 	for key, value := range r.URL.Query() {
// 		queryParams[key] = strings.Join(value, ",")
// 	}
// 	return queryParams, nil
// }

func CustomRepsonseWriter(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

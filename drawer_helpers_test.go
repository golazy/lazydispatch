package lazydispatch

import (
	"net/http"
	"net/http/httptest"
)

func newFakeHandler(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(name))
	})
}

func isHandler(h http.Handler, name string) bool {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Body.String() == name {
		return true
	}
	return false

}

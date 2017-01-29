package app

import "net/http"

func serveProfile(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "user/profile.html", http.StatusOK, fakeData)
}

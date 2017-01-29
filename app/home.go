package app

import "net/http"

var fakeData = map[string]interface{}{
	"Role": map[string]bool{
		"Admin": true,
	},
	"User": map[string]string{
		"Name":  "acoshift",
		"Photo": "https://firebasestorage.googleapis.com/v0/b/acourse-d9d0a.appspot.com/o/user%2FHSq22OmOAeb4NuEP3ad09uQh83u2%2F1476559766339?alt=media&token=adb2fe10-ca11-4574-95b9-ae6edea21144",
	},
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	executeTemplate(w, "home.html", http.StatusOK, fakeData)
}

package app

import (
	"net/http"
	"sync"
	"time"

	"github.com/acoshift/ds"
)

var fakeData = map[string]interface{}{
	"Role": map[string]bool{
		"Admin":      true,
		"Instructor": true,
	},
	"User": map[string]interface{}{
		"ID":    func() string { return "HSq22OmOAeb4NuEP3ad09uQh83u2" },
		"Name":  "acoshift",
		"Photo": "https://firebasestorage.googleapis.com/v0/b/acourse-d9d0a.appspot.com/o/user%2FHSq22OmOAeb4NuEP3ad09uQh83u2%2F1476559766339?alt=media&token=adb2fe10-ca11-4574-95b9-ae6edea21144",
	},
	"OwnCourses": []map[string]interface{}{
		{
			"ID":               func() string { return "5701751084679168" },
			"Owner":            "HSq22OmOAeb4NuEP3ad09uQh83u2",
			"Photo":            "https://firebasestorage.googleapis.com/v0/b/acourse-d9d0a.appspot.com/o/user%2FHSq22OmOAeb4NuEP3ad09uQh83u2%2F1474647406971?alt=media&token=a6706d2f-199e-4153-b83c-f0510b4f41fe",
			"Price":            float64(60),
			"ShortDescription": "Create web application with Vue.js 2 and Firebase",
			"Start":            time.Date(2016, 11, 26, 0, 0, 0, 0, time.UTC),
			"Title":            "Vue.js 2 for Real World",
			"Type":             "video",
			"URL":              "vuejs2-for-real-world",
			"EnrollCount":      128,
		},
	},
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	data := map[string]interface{}{}
	{
		var courses []*Course
		err := client.Query(ctx, kindCourse, &courses,
			ds.Filter("Options.Public =", true),
		)
		must(err)
		m := &sync.WaitGroup{}
		m.Add(len(courses))
		for _, course := range courses {
			go func(course *Course) {
				var err error
				course.EnrollCount, err = client.QueryCount(ctx, kindEnroll, ds.Filter("CourseID =", course.ID()))
				must(err)
				m.Done()
			}(course)
		}
		m.Wait()
		data["PublicCourses"] = courses
	}

	executeTemplate(w, "home.html", http.StatusOK, data)
}

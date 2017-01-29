package app

import (
	"net/http"
	"path"
	"strings"

	"github.com/acoshift/ds"
)

func serveCourse(w http.ResponseWriter, r *http.Request) {
	p := path.Clean(r.URL.Path)
	p = p[1:]
	rIndex := strings.Index(p, "/")
	if rIndex >= 0 {
		// serve other
		return
	}

	data := map[string]interface{}{}
	courseID := p
	ctx := r.Context()
	var course Course
	var owner User

	// try find course by url
	err := client.QueryFirst(ctx, kindCourse, &course, ds.Filter("URL =", courseID))
	if ds.NotFound(err) {
		err = client.GetByStringID(ctx, kindCourse, courseID, &course)
	}
	err = ds.IgnoreFieldMismatch(err)
	must(err)
	err = client.GetByStringID(ctx, kindUser, course.Owner, &owner)
	err = ds.IgnoreFieldMismatch(err)
	err = ds.IgnoreNotFound(err)
	must(err)
	owner.Default()
	course.OwnerUser = &owner

	data["Course"] = &course

	executeTemplate(w, "course/index.html", http.StatusOK, data)
}

func serveCourseEdit(w http.ResponseWriter, r *http.Request) {

}

package app

import (
	"context"
	"net/http"
	"path"
	"strings"

	"github.com/acoshift/ds"
)

func serveCourse(w http.ResponseWriter, r *http.Request) {
	p := path.Clean(r.URL.Path)
	if p != r.URL.Path {
		http.Redirect(w, r, p, http.StatusMovedPermanently)
		return
	}
	p = strings.TrimPrefix(p, "/")
	rIndex := strings.Index(p, "/")
	courseID := p
	if rIndex >= 0 {
		courseID = p[:rIndex]
	}
	ctx := r.Context()

	var course Course
	// try find course by url
	err := client.QueryFirst(ctx, kindCourse, &course, ds.Filter("URL =", courseID))
	if ds.NotFound(err) {
		err = client.GetByStringID(ctx, kindCourse, courseID, &course)
	}
	err = ds.IgnoreFieldMismatch(err)
	must(err)

	if rIndex >= 0 {
		p = p[rIndex:]
		if strings.HasPrefix(p+"/", "/edit/") {
			ctx = context.WithValue(ctx, keyCourseID, courseID)
			ctx = context.WithValue(ctx, keyCourse, &course)
			http.StripPrefix("/"+courseID+"/edit", http.HandlerFunc(serveCourseEdit)).
				ServeHTTP(w, r.WithContext(ctx))
			return
		}
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{}
	var owner User

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
	ctx := r.Context()
	course, _ := ctx.Value(keyCourse).(*Course)
	data := map[string]interface{}{}
	data["Course"] = course
	executeTemplate(w, "course/edit.html", http.StatusOK, data)
}

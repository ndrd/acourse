package app

import (
	"log"
	"net/http"
)

// Run runs the app
func Run() {
	var err error

	err = parseTemplates([][]string{
		{"home.html", "common.html", "layout.html"},
		{"user/profile.html", "common.html", "layout.html"},
		{"privacy.html", "common.html", "layout.html"},
	})
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	// assets
	mux.Handle("/-/jquery-3.1.1.min.js", serveFile("bower_components/jquery/dist/jquery.min.js"))
	mux.Handle("/-/semantic-2.2.7-accordion.min.js", serveSemanticComponent("accordion"))
	mux.Handle("/-/semantic-2.2.7-checkbox.min.js", serveSemanticComponent("checkbox"))
	mux.Handle("/-/semantic-2.2.7-dimmer.min.js", serveSemanticComponent("dimmer"))
	mux.Handle("/-/semantic-2.2.7-dropdown.min.js", serveSemanticComponent("dropdown"))
	mux.Handle("/-/semantic-2.2.7-embed.min.js", serveSemanticComponent("embed"))
	mux.Handle("/-/semantic-2.2.7-modal.min.js", serveSemanticComponent("modal"))
	mux.Handle("/-/semantic-2.2.7-transition.min.js", serveSemanticComponent("transition"))
	mux.Handle("/-/semantic-2.2.7-progress.min.js", serveSemanticComponent("progress"))
	mux.Handle("/-/semantic-2.2.7-reset.min.css", serveSemanticComponentCSS("reset"))
	mux.Handle("/-/semantic-2.2.7-site.min.css", serveSemanticComponentCSS("site"))
	mux.Handle("/-/semantic-2.2.7-accordion.min.css", serveSemanticComponentCSS("accordion"))
	mux.Handle("/-/semantic-2.2.7-breadcrumb.min.css", serveSemanticComponentCSS("breadcrumb"))
	mux.Handle("/-/semantic-2.2.7-button.min.css", serveSemanticComponentCSS("button"))
	mux.Handle("/-/semantic-2.2.7-card.min.css", serveSemanticComponentCSS("card"))
	mux.Handle("/-/semantic-2.2.7-checkbox.min.css", serveSemanticComponentCSS("checkbox"))
	mux.Handle("/-/semantic-2.2.7-comment.min.css", serveSemanticComponentCSS("comment"))
	mux.Handle("/-/semantic-2.2.7-container.min.css", serveSemanticComponentCSS("container"))
	mux.Handle("/-/semantic-2.2.7-dimmer.min.css", serveSemanticComponentCSS("dimmer"))
	mux.Handle("/-/semantic-2.2.7-divider.min.css", serveSemanticComponentCSS("divider"))
	mux.Handle("/-/semantic-2.2.7-dropdown.min.css", serveSemanticComponentCSS("dropdown"))
	mux.Handle("/-/semantic-2.2.7-embed.min.css", serveSemanticComponentCSS("embed"))
	mux.Handle("/-/semantic-2.2.7-form.min.css", serveSemanticComponentCSS("form"))
	mux.Handle("/-/semantic-2.2.7-grid.min.css", serveSemanticComponentCSS("grid"))
	mux.Handle("/-/semantic-2.2.7-header.min.css", serveSemanticComponentCSS("header"))
	mux.Handle("/-/semantic-2.2.7-icon.min.css", serveSemanticComponentCSS("icon"))
	mux.Handle("/-/semantic-2.2.7-image.min.css", serveSemanticComponentCSS("image"))
	mux.Handle("/-/semantic-2.2.7-input.min.css", serveSemanticComponentCSS("input"))
	mux.Handle("/-/semantic-2.2.7-item.min.css", serveSemanticComponentCSS("item"))
	mux.Handle("/-/semantic-2.2.7-list.min.css", serveSemanticComponentCSS("list"))
	mux.Handle("/-/semantic-2.2.7-menu.min.css", serveSemanticComponentCSS("menu"))
	mux.Handle("/-/semantic-2.2.7-message.min.css", serveSemanticComponentCSS("message"))
	mux.Handle("/-/semantic-2.2.7-modal.min.css", serveSemanticComponentCSS("modal"))
	mux.Handle("/-/semantic-2.2.7-segment.min.css", serveSemanticComponentCSS("segment"))
	mux.Handle("/-/semantic-2.2.7-transition.min.css", serveSemanticComponentCSS("transition"))
	mux.Handle("/-/semantic-2.2.7-progress.min.css", serveSemanticComponentCSS("progress"))
	mux.Handle("/-/semantic-2.2.7-table.min.css", serveSemanticComponentCSS("table"))
	mux.Handle("/-/site.css", serveFile("site.css"))
	mux.Handle("/-/acourse.svg", serveFile("acourse.svg"))
	mux.Handle("/-/acourse-120.png", serveFile("acourse-120.png"))
	mux.Handle("/-/acourse-og.jpg", serveFile("acourse-og.jpg"))
	mux.Handle("/-/icons/ic_face_black_48px.svg", serveFile("icons/ic_face_black_48px.svg"))
	mux.Handle("/-/icons/ic_insert_drive_file_black_48px.svg", serveFile("icons/ic_insert_drive_file_black_48px.svg"))
	mux.Handle("/manifest.json", serveFile("manifest.json"))
	mux.Handle("/favicon.ico", serveFile("acourse-120.png"))

	mux.Handle("/_ah/health", http.HandlerFunc(serveHealthCheck))
	mux.Handle("/_ah/", http.NotFoundHandler())
	mux.Handle("/privacy", http.HandlerFunc(servePrivacy))
	mux.Handle("/profile", http.HandlerFunc(serveProfile))
	mux.Handle("/", http.HandlerFunc(serveHome))
	http.ListenAndServe(":8080", mux)
}

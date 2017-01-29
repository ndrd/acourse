package app

import (
	"context"
	"log"
	"net"
	"net/http"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/acoshift/ds"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// Config type
type Config struct {
	Debug     bool   `yaml:"debug"`
	Port      string `yaml:"port"`
	TLSPort   string `yaml:"tlsPort"`
	TLSCert   string `yaml:"tlsCert"`
	TLSKey    string `yaml:"tlsKey"`
	Domain    string `yaml:"domain"`
	ProjectID string `yaml:"projectId"`
	Email     struct {
		From     string `yaml:"from"`
		Server   string `yaml:"server"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"email"`
	Firebase struct {
		ProjectID      string `yaml:"projectId"`
		ServiceAccount string `yaml:"serviceAccount"`
	} `yaml:"firebase"`
	ServiceAccount string `yaml:"serviceAccount"`
}

var client *ds.Client

// Run runs the app
func Run(config Config) {
	ctx := context.Background()
	var err error

	var tokenSource oauth2.TokenSource
	{
		jwtConfig, err := google.JWTConfigFromJSON([]byte(config.ServiceAccount),
			datastore.ScopeDatastore,
			pubsub.ScopePubSub,
			storage.ScopeFullControl,
		)
		if err != nil {
			log.Fatal(err)
		}
		tokenSource = jwtConfig.TokenSource(ctx)
	}

	err = parseTemplates([][]string{
		{"home.html", "common.html", "layout.html"},
		{"user/profile.html", "common.html", "layout.html"},
		{"course/index.html", "course/layout.html", "common.html", "layout.html"},
		{"course/edit.html", "course/layout.html", "common.html", "layout.html"},
		{"privacy.html", "common.html", "layout.html"},
	})
	if err != nil {
		log.Fatal(err)
	}

	client, err = ds.NewClient(ctx, config.ProjectID, option.WithTokenSource(tokenSource))
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	// assets
	mux.Handle("/-/jquery-3.1.1.min.js", serveFile("bower_components/jquery/dist/jquery.min.js"))
	mux.Handle("/-/semantic-2.2.7/components/accordion.min.js", serveSemanticComponent("accordion"))
	mux.Handle("/-/semantic-2.2.7/components/checkbox.min.js", serveSemanticComponent("checkbox"))
	mux.Handle("/-/semantic-2.2.7/components/dimmer.min.js", serveSemanticComponent("dimmer"))
	mux.Handle("/-/semantic-2.2.7/components/dropdown.min.js", serveSemanticComponent("dropdown"))
	mux.Handle("/-/semantic-2.2.7/components/embed.min.js", serveSemanticComponent("embed"))
	mux.Handle("/-/semantic-2.2.7/components/modal.min.js", serveSemanticComponent("modal"))
	mux.Handle("/-/semantic-2.2.7/components/transition.min.js", serveSemanticComponent("transition"))
	mux.Handle("/-/semantic-2.2.7/components/progress.min.js", serveSemanticComponent("progress"))
	mux.Handle("/-/semantic-2.2.7/components/reset.min.css", serveSemanticComponentCSS("reset"))
	mux.Handle("/-/semantic-2.2.7/components/site.min.css", serveSemanticComponentCSS("site"))
	mux.Handle("/-/semantic-2.2.7/components/accordion.min.css", serveSemanticComponentCSS("accordion"))
	mux.Handle("/-/semantic-2.2.7/components/breadcrumb.min.css", serveSemanticComponentCSS("breadcrumb"))
	mux.Handle("/-/semantic-2.2.7/components/button.min.css", serveSemanticComponentCSS("button"))
	mux.Handle("/-/semantic-2.2.7/components/card.min.css", serveSemanticComponentCSS("card"))
	mux.Handle("/-/semantic-2.2.7/components/checkbox.min.css", serveSemanticComponentCSS("checkbox"))
	mux.Handle("/-/semantic-2.2.7/components/comment.min.css", serveSemanticComponentCSS("comment"))
	mux.Handle("/-/semantic-2.2.7/components/container.min.css", serveSemanticComponentCSS("container"))
	mux.Handle("/-/semantic-2.2.7/components/dimmer.min.css", serveSemanticComponentCSS("dimmer"))
	mux.Handle("/-/semantic-2.2.7/components/divider.min.css", serveSemanticComponentCSS("divider"))
	mux.Handle("/-/semantic-2.2.7/components/dropdown.min.css", serveSemanticComponentCSS("dropdown"))
	mux.Handle("/-/semantic-2.2.7/components/embed.min.css", serveSemanticComponentCSS("embed"))
	mux.Handle("/-/semantic-2.2.7/components/form.min.css", serveSemanticComponentCSS("form"))
	mux.Handle("/-/semantic-2.2.7/components/grid.min.css", serveSemanticComponentCSS("grid"))
	mux.Handle("/-/semantic-2.2.7/components/header.min.css", serveSemanticComponentCSS("header"))
	mux.Handle("/-/semantic-2.2.7/components/icon.min.css", serveSemanticComponentCSS("icon"))
	mux.Handle("/-/semantic-2.2.7/components/image.min.css", serveSemanticComponentCSS("image"))
	mux.Handle("/-/semantic-2.2.7/components/input.min.css", serveSemanticComponentCSS("input"))
	mux.Handle("/-/semantic-2.2.7/components/item.min.css", serveSemanticComponentCSS("item"))
	mux.Handle("/-/semantic-2.2.7/components/list.min.css", serveSemanticComponentCSS("list"))
	mux.Handle("/-/semantic-2.2.7/components/menu.min.css", serveSemanticComponentCSS("menu"))
	mux.Handle("/-/semantic-2.2.7/components/message.min.css", serveSemanticComponentCSS("message"))
	mux.Handle("/-/semantic-2.2.7/components/modal.min.css", serveSemanticComponentCSS("modal"))
	mux.Handle("/-/semantic-2.2.7/components/segment.min.css", serveSemanticComponentCSS("segment"))
	mux.Handle("/-/semantic-2.2.7/components/transition.min.css", serveSemanticComponentCSS("transition"))
	mux.Handle("/-/semantic-2.2.7/components/progress.min.css", serveSemanticComponentCSS("progress"))
	mux.Handle("/-/semantic-2.2.7/components/table.min.css", serveSemanticComponentCSS("table"))
	mux.Handle("/-/semantic-2.2.7/themes/default/assets/fonts/icons.woff2", serveFile("bower_components/semantic/dist/themes/default/assets/fonts/icons.woff2"))
	mux.Handle("/-/semantic-2.2.7/themes/default/assets/fonts/icons.woff", serveFile("bower_components/semantic/dist/themes/default/assets/fonts/icons.woff"))
	mux.Handle("/-/semantic-2.2.7/themes/default/assets/fonts/icons.ttf", serveFile("bower_components/semantic/dist/themes/default/assets/fonts/icons.ttf"))
	mux.Handle("/-/semantic-2.2.7/themes/default/assets/fonts/icons.svg", serveFile("bower_components/semantic/dist/themes/default/assets/fonts/icons.svg"))
	mux.Handle("/-/semantic-2.2.7/themes/default/assets/fonts/icons.otf", serveFile("bower_components/semantic/dist/themes/default/assets/fonts/icons.otf"))
	mux.Handle("/-/semantic-2.2.7/themes/default/assets/fonts/icons.eot", serveFile("bower_components/semantic/dist/themes/default/assets/fonts/icons.eot"))
	mux.Handle("/-/semantic-2.2.7/themes/default/assets/images/flags.png", serveFile("bower_components/semantic/dist/themes/default/assets/images/flags.png"))
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
	mux.Handle("/course", http.NotFoundHandler())
	mux.Handle("/course/", http.StripPrefix("/course", http.HandlerFunc(serveCourse)))
	mux.Handle("/", http.HandlerFunc(serveHome))

	handler := mux
	addr := net.JoinHostPort("0.0.0.0", config.Port)

	if len(config.TLSPort) > 0 {
		tlsAddr := net.JoinHostPort("0.0.0.0", config.TLSPort)
		go func() {
			log.Printf("Listening Redirect on %s\n", addr)
			log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "https://"+config.Domain, http.StatusMovedPermanently)
			})))
		}()
		log.Printf("Listening TLS on %s", tlsAddr)
		log.Fatal(http.ListenAndServeTLS(tlsAddr, config.TLSCert, config.TLSKey, handler))
	} else {
		log.Printf("Listening on %s", addr)
		log.Fatal(http.ListenAndServe(addr, handler))
	}
}

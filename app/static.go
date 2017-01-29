package app

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func serveHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Health check: ok\n"))
}

func servePrivacy(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "privacy.html", http.StatusOK, nil)
}

const (
	assetDir = "assets"
	maxAge   = 365 * 24 * time.Hour
)

func joinAssetDir(file string) string {
	return filepath.Join(assetDir, file)
}

func serveSemanticComponent(name string) http.Handler {
	return serveFile("bower_components/semantic/dist/components/" + name + ".min.js")
}

func serveSemanticComponentCSS(name string) http.Handler {
	return serveFile("bower_components/semantic/dist/components/" + name + ".min.css")
}

func serveFile(file string) http.Handler {
	cacheControl := fmt.Sprintf("public, max-age=%d", maxAge/time.Second)
	mimeType := mime.TypeByExtension(path.Ext(file))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	rawBuf, err := ioutil.ReadFile(joinAssetDir(file))
	if err != nil {
		log.Fatal(err)
	}
	rawLength := strconv.Itoa(len(rawBuf))

	var gzBuf []byte
	{
		buf := bytes.NewBuffer([]byte{})
		gzWriter := gzip.NewWriter(buf)
		_, err = io.Copy(gzWriter, ioutil.NopCloser(bytes.NewReader(rawBuf)))
		if err != nil {
			log.Fatal(err)
		}
		gzWriter.Flush()
		gzBuf = buf.Bytes()
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := path.Clean(r.URL.Path)
		if p != r.URL.Path {
			http.Redirect(w, r, p, http.StatusMovedPermanently)
			return
		}
		w.Header().Set("Cache-Control", cacheControl)
		w.Header().Set("Content-Type", mimeType)

		var reader io.Reader
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Add("Vary", "Accept-Encoding")
			w.Header().Set("Content-Encoding", "gzip")
			reader = ioutil.NopCloser(bytes.NewReader(gzBuf))
		} else {
			w.Header().Add("Content-Length", rawLength)
			reader = ioutil.NopCloser(bytes.NewReader(rawBuf))
		}

		w.WriteHeader(http.StatusOK)
		if r.Method != http.MethodHead {
			io.Copy(w, reader)
		}
	})
}

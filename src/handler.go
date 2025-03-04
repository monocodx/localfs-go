// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"localfs/util/fsutil"
	"localfs/util/netutil"
	"localfs/view"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

type fileInfo struct {
	name string
	size int64
	hash string
}

var prgCache = map[string]fileInfo{}

func uploadPageHandler(w http.ResponseWriter, r *http.Request) {
	// page navigation bar
	navBar := view.NavBar{
		ActiveItem: "Upload",
		NavItem: []view.NavItem{
			{Name: "Home", Link: "/"},
		},
	}

	path := appCache[ckey_storage]
	files, err := fsutil.FilesListing(path)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set headers
	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")

	fmap := template.FuncMap{
		"index":    view.ListingIndex,
		"zebraCss": view.ListingZebraCss,
	}

	t, err := template.New("uploadPage").Funcs(fmap).Parse(view.UploadPageTmpl)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, view.UploadPageViewModel{
		Build:  appBuild,
		Files:  files,
		NavBar: navBar,
	})
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		if strings.Contains(err.Error(), "no space left on device") {
			errorHandler(w, err.Error(), http.StatusInternalServerError)
			return
		}
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	path := appCache[ckey_storage]
	fname := fsutil.ResolveFileConflict(path, fileHeader.Filename)
	err = fsutil.WriteStreamToFile(path, fname, file)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uid := uuid.New().String()
	// back to the beginning of the file
	file.Seek(0, io.SeekStart)
	hash, _ := fsutil.Sha256sum(file)

	prgCache[uid] = fileInfo{
		name: fname,
		size: fileHeader.Size,
		hash: hash,
	}

	http.Redirect(w, r, fmt.Sprintf("/upload/status?uid=%s", uid), http.StatusSeeOther)
}

func uploadStatusPageHandler(w http.ResponseWriter, r *http.Request) {
	// page navigation bar
	navBar := view.NavBar{
		ActiveItem: "Status",
		NavItem: []view.NavItem{
			{Name: "Home", Link: "/"},
			{Name: "Upload", Link: "/upload"},
		},
	}

	// get uid from query param
	uid := r.URL.Query().Get("uid")
	// get file info from cache
	fi, ok := prgCache[uid]
	if !ok {
		errorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// clear cache
	delete(prgCache, uid)

	// do verification
	path := filepath.Join(appCache[ckey_storage], fi.name)
	file, err := os.Open(path)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// get hash
	hash, _ := fsutil.Sha256sum(file)

	mismatch := false
	message := ""
	if fi.hash != hash {
		mismatch = true
		message = "hash mismatch detected."
	}

	// set headers
	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.New("uploadStatusPage").Parse(view.UploadStatusPageTmpl)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, view.UploadStatusPageViewModel{
		Error:     mismatch,
		Message:   message,
		Filename:  fi.name,
		Size:      strconv.FormatInt(fi.size, 10),
		Sha256sum: hash,
		NavBar:    navBar,
	})
}

func indexPageHandler(w http.ResponseWriter, _ *http.Request) {
	// get ip address
	addrs, err := netutil.IPv4Address()
	if err != nil {
		// fallback to localhost
		addrs = []string{"localhost"}
	}

	// parse qrcode content
	host := net.JoinHostPort(addrs[0], appCache[ckey_port])
	content := url.URL{Scheme: "http", Host: host, Path: "upload"}
	// generate the QR code image as a byte slice (PNG format)
	byt, err := qrcode.Encode(content.String(), qrcode.Medium, 320)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// convert the byte slice to a base64 string
	base64png := base64.StdEncoding.EncodeToString(byt)

	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.New("IndexPage").Parse(view.IndexPageTmpl)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, view.IndexPageViewModel{
		Base64QRImage: base64png,
		Address:       addrs[0],
	})
}

func fileHandler(prefix string) http.Handler {
	fileDir := appCache[ckey_storage]
	return http.StripPrefix(prefix, http.FileServer(http.Dir(fileDir)))
}

func routesHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path != "/favicon.ico" {
		// route to index page
		if path == "/" {
			indexPageHandler(w, r)
			return
		}
		// handle trailing slashes
		if path == "/upload/" {
			uploadPageHandler(w, r)
			return
		}

		// handle all invalid routes
		errorHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

// extend http.error implementation
func errorHandler(w http.ResponseWriter, error string, code int) {
	h := w.Header()
	// Delete the Content-Length header, which might be for some other content.
	// Assuming the error string fits in the writer's buffer, we'll figure
	// out the correct Content-Length for it later.
	//
	// We don't delete Content-Encoding, because some middleware sets
	// Content-Encoding: gzip and wraps the ResponseWriter to compress on-the-fly.
	// See https://go.dev/issue/66343.
	h.Del("Content-Length")
	// There might be content type already set, but we reset it to
	// text/html for the error message.
	h.Set("Content-Type", "text/html; charset=utf-8")
	h.Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	t, err := template.New("uploadPage").Parse(view.ErrorPageTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, view.ErrorViewModel{
		Code:    code,
		Status:  http.StatusText(code),
		Message: error,
	})
}

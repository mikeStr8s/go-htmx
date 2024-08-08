package main

import (
	"net/http"

	"github.com/jritsema/gotoolbox/web"
)

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "index.html", data, nil)
}

func update(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "update.html", data, nil)
}

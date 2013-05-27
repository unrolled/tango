package main

import (
    "github.com/cojac/tango"
)

type IndexHandler struct{ tango.BaseHandler }

func (h *IndexHandler) New() tango.HandlerInterface {
    return &IndexHandler{}
}

func (h *IndexHandler) Get(request *tango.HttpRequest) *tango.HttpResponse {
    return tango.NewHttpResponse("Hello! No visit a <a href=\"/bad/\">bad page</a>.")
}

type NotFoundHandler struct{ tango.BaseHandler }

func (h *NotFoundHandler) New() tango.HandlerInterface {
    return &NotFoundHandler{}
}

func (h *NotFoundHandler) Get(request *tango.HttpRequest) *tango.HttpResponse {
    return tango.NewHttpResponse("Sorry... you hit a page. Don't forget to set this handler as <strong>404</strong>!!", 404)
}

func init() {
    tango.Settings.Set("debug", true)
    tango.Settings.Set("serve_address", ":8000")

    tango.Pattern("/", &IndexHandler{})
    tango.SetNotFoundHandler(&NotFoundHandler{})
}

func main() {
    tango.ListenAndServe()
}

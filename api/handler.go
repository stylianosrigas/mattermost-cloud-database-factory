package api

import (
	"net/http"

	"github.com/mattermost/mattermost-cloud-database-factory/model"
	log "github.com/sirupsen/logrus"
)

type contextHandlerFunc func(c *Context, w http.ResponseWriter, r *http.Request)

type contextHandler struct {
	context *Context
	handler contextHandlerFunc
}

// ServeHTTP gets the http Request
func (h contextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := h.context.Clone()
	context.RequestID = model.NewID()
	context.Logger = context.Logger.WithFields(log.Fields{
		"path":    r.URL.Path,
		"request": context.RequestID,
	})

	h.handler(context, w, r)
}

func newContextHandler(context *Context, handler contextHandlerFunc) *contextHandler {
	return &contextHandler{
		context: context,
		handler: handler,
	}
}

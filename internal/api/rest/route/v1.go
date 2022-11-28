package route

import (
	"net/http"
)

func (r *r) InitV1Route() {
	routes = append(
		routes,
		addRoute(http.MethodPost, "", r.h.CreateUser),
	)
}

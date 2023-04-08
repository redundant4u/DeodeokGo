package server_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/internal/server"
	"github.com/redundant4u/DeoDeokGo/mocks"
)

func TestRouter(t *testing.T) {
	r := server.Router(&mocks.MockMongoClient{})
	list := r.Routes()

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodGet,
		Path:   "/events",
	})

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodGet,
		Path:   "/events/id/:id",
	})

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodGet,
		Path:   "/events/name/:name",
	})

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodPost,
		Path:   "/events",
	})
}

func assertRoutePresent(t *testing.T, gotRoutes gin.RoutesInfo, wantRoute gin.RouteInfo) {
	for _, gotRoute := range gotRoutes {
		if gotRoute.Path == wantRoute.Path && gotRoute.Method == wantRoute.Method {
			return
		}
	}

	t.Errorf("Route not found: %v", wantRoute)
}

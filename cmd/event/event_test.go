package main_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/redundant4u/DeoDeokGo/mocks"
	"github.com/redundant4u/DeoDeokGo/routes"
)

func TestRouter(t *testing.T) {
	ctx := context.Background()

	r := routes.InitEventsRoutes(ctx, &mocks.MockMongoDatabase{}, &mocks.MockEventListener{}, &mocks.MockEventEmitter{})
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

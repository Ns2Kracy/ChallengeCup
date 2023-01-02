package middleware

import (
	"time"

	"github.com/kataras/iris/v12/middleware/monitor"
)

var Monitor = monitor.New(monitor.Options{
	RefreshInterval:     2 * time.Second,
	ViewRefreshInterval: 2 * time.Second,
	ViewTitle:           "Challenge Resource Use Monitor",
})

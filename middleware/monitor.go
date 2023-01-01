package middleware

import (
    "github.com/kataras/iris/v12/middleware/monitor"
    "time"
)

var Monitor = monitor.New(monitor.Options{
	RefreshInterval:     2 * time.Second,
	ViewRefreshInterval: 2 * time.Second,
	ViewTitle:           "MyServer Monitor",
})

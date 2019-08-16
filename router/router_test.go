package router

import (
	"net/http"
	"testing"
)

var defProc http.HandlerFunc

// todo http发起测试
func TestAddRoute(t *testing.T) {
	AddRoute("get", "/home", defProc)

	if len(routes) != 1 {
		t.Error("add route failed")
	}

	AddRoute("get", "/home", defProc)
	AddRoute("get", "/home", defProc)
	AddRoute("get", "/home", defProc)

	if len(routes) != 1 {
		t.Error("add route failed")
	}
}

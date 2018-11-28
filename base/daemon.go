package base

import (
	"image/color"
	"net"
	"net/http"
	"sync"

	"golang.org/x/crypto/acme/autocert"
)

type (
	// Feather is a light-weight golang framework
	Feather struct {
		StdLogger     *stdLog.Logger
		colorer       *color.Color
		premiddleware []MiddlewareFunc
		middleware    []MiddlewareFunc
		maxParam      *int
		router        *Router
		//notFoundHandler HandlerFunc
		pool           sync.Pool
		Server         *http.Server
		TLSServer      *http.Server
		Listener       net.Listener
		TLSListener    net.Listener
		AutoTLSManager autocert.Manager
		DisableHTTP2   bool
		Debug          bool
		HideBanner     bool
		HidePort       bool
		AuthMethod     string
		//HTTPErrorHandler HTTPErrorHandler
		Binder Binder
		//Validator        Validator
		Renderer Renderer
		Logger   Logger
	}
)

package base

import (
	"image/color"
	"net"
	"net/http"
	"sync"

	"github.com/labstack/gommon/log"
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

// New creates an instance of Echo.
func New() (e *Feather) {
	e = &Feather{
		Server:    new(http.Server),
		TLSServer: new(http.Server),
		AutoTLSManager: autocert.Manager{
			Prompt: autocert.AcceptTOS,
		},
		Logger:   log.New("feather"),
		colorer:  color.New(),
		maxParam: new(int),
	}
	e.Server.Handler = e
	e.TLSServer.Handler = e
	e.HTTPErrorHandler = e.DefaultHTTPErrorHandler
	e.Binder = &DefaultBinder{}
	e.Logger.SetLevel(log.ERROR)
	e.StdLogger = stdLog.New(e.Logger.Output(), e.Logger.Prefix()+": ", 0)
	e.pool.New = func() interface{} {
		return e.NewContext(nil, nil)
	}
	e.router = NewRouter(e)
	return
}

// NewContext returns a Context instance.
func (e *Feather) NewContext(r *http.Request, w http.ResponseWriter) Context {
	return &context{
		request:  r,
		response: NewResponse(w, e),
		store:    make(Map),
		echo:     e,
		pvalues:  make([]string, *e.maxParam),
		handler:  NotFoundHandler,
	}
}

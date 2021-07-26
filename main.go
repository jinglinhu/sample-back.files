package main

import (
	"fmt"

	"github.com/aws/aws-xray-sdk-go/xray"

	"github.com/fasthttp/router"
    "github.com/valyala/fasthttp"
)

func init() {
	xray.Configure(xray.Config{
		DaemonAddr:     "xray-service.default:2000",
		LogLevel:       "info",
	})
}


func main() {
    fh := xray.NewFastHTTPInstrumentor(nil)
    r := router.New()
    
    r.GET("/", middleware("x-ray-sample-back-k8s", backend, fh))

    fasthttp.ListenAndServe(":8080", r.Handler)
}

func middleware(name string, h fasthttp.RequestHandler, fh xray.FastHTTPHandler) fasthttp.RequestHandler {
    f := func(ctx *fasthttp.RequestCtx) {
        h(ctx)
    }
    return fh.Handler(xray.NewFixedSegmentNamer(name), f)
}

func backend(ctx *fasthttp.RequestCtx){
	fmt.Fprintf(ctx,"Backend Server")
}




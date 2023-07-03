package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/mayuka-c/opentelemetry-pract/tracing"
)

func setupWebServer() {
	r := gin.Default()
	r.Use(otelgin.Middleware("sample-service"))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"status": "OK"})
	})
	r.HandleMethodNotAllowed = true
	r.Run(":8081")
}

func main() {
	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	setupWebServer()
}

package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thinkgos/gin-contrib/gzap"
	"github.com/thinkgos/gin-contrib/oteltraceid"
	"github.com/thinkgos/gin-rest-kit/web"
	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm-webhook-example/webhook"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func main() {
	appName := "zlm-http-api-example"
	zlog := logger.NewLogger(logger.WithConfig(logger.Config{
		Level:   "debug",
		Adapter: logger.AdapterFile,
		File: logger.LumberjackFile{
			Path:       "/home/thinkgo/.log",
			Filename:   "zlm-http-api-example.log",
			MaxSize:    100,
			MaxAge:     7,
			MaxBackups: 10,
			LocalTime:  false,
			Compress:   true,
		},
	}))
	logger.ReplaceGlobals(zlog.ExtendDefaultHook(
		&logger.ImmutableString{Key: "app", Value: appName},
		&logger.MutableString{Key: "traceId", Fc: oteltraceid.FromTraceId},
		&logger.MutableString{Key: "spanId", Fc: oteltraceid.FromSpanId},
		&logger.ImmutableString{Key: "type", Value: "global"},
	))
	setTracerProvider("", appName)

	engine := gin.New()
	engine.Use(
		otelgin.Middleware(appName),
		oteltraceid.TraceId(),
		gzap.Recovery(
			logger.WithNewHook(
				&logger.ImmutableString{Key: "app", Value: appName},
				&logger.ImmutableString{Key: "type", Value: "http"},
				&logger.MutableString{Key: "traceId", Fc: oteltraceid.FromTraceId},
				&logger.MutableString{Key: "spanId", Fc: oteltraceid.FromSpanId},
			).
				SetNewCallerCore(logger.NewCallerCore()),
			true,
		),
		gzap.Logger(
			logger.WithNewHook(
				&logger.ImmutableString{Key: "app", Value: appName},
				&logger.ImmutableString{Key: "type", Value: "http"},
				&logger.MutableString{Key: "traceId", Fc: oteltraceid.FromTraceId},
				&logger.MutableString{Key: "spanId", Fc: oteltraceid.FromSpanId},
			).
				SetNewCallerCore(logger.NewCallerCore()),
			gzap.WithEnableBody(true),
			gzap.WithBodyLimit(8192),
			// gzap.WithEnableDebugCurl(deploy.IsTesting()),
		),
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: true,
			AllowWebSockets:  true,
			MaxAge:           12 * time.Hour,
		}),
	)
	g := engine.Group("")
	RegisterZlmWebhook(g, webhook.NewZlmHook())

	engine.Run(":9527")
}

func setTracerProvider(endpoint string, name string) *trace.TracerProvider {
	traceOptions := []trace.TracerProviderOption{
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
		)),
	}
	if endpoint != "" {
		exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
		if err != nil {
			log.Fatal(err.Error())
		}
		traceOptions = append(traceOptions, trace.WithBatcher(exp))
	}

	tp := trace.NewTracerProvider(traceOptions...)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
	return tp
}

func RegisterZlmWebhook(g *gin.RouterGroup, srv zlm_webhook.Webhook) {
	r := g.Group(zlm_webhook.UrlPathPrefix,
		web.TransportInterceptor(),             // gin.Context注入context.Context
		web.CarrierInterceptor(web.NewCarry()), // carrier注入到context.Context
	)
	{
		r.POST(zlm_webhook.UrlPath_OnFlowReport, web.Handler(srv.OnFlowReport))
		r.POST(zlm_webhook.UrlPath_OnHttpAccess, web.Handler(srv.OnHttpAccess))
		r.POST(zlm_webhook.UrlPath_OnPlay, web.Handler(srv.OnPlay))
		r.POST(zlm_webhook.UrlPath_OnPublish, web.Handler(srv.OnPublish))
		r.POST(zlm_webhook.UrlPath_OnRecordMp4, web.Handler(srv.OnRecordMp4))
		r.POST(zlm_webhook.UrlPath_OnRecordTs, web.Handler(srv.OnRecordTs))
		r.POST(zlm_webhook.UrlPath_OnRtpServerTimeout, web.Handler(srv.OnRtpServerTimeout))
		r.POST(zlm_webhook.UrlPath_OnSendRtpStopped, web.Handler(srv.OnSendRtpStopped))
		r.POST(zlm_webhook.UrlPath_OnRtspAuth, web.Handler(srv.OnRtspAuth))
		r.POST(zlm_webhook.UrlPath_OnRtspRealm, web.Handler(srv.OnRtspRealm))
		r.POST(zlm_webhook.UrlPath_OnServerKeepalive, web.Handler(srv.OnServerKeepalive))
		r.POST(zlm_webhook.UrlPath_OnServerStarted, web.Handler(srv.OnServerStarted))
		r.POST(zlm_webhook.UrlPath_OnServerExited, web.Handler(srv.OnServerExited))
		r.POST(zlm_webhook.UrlPath_OnShellLogin, web.Handler(srv.OnShellLogin))
		r.POST(zlm_webhook.UrlPath_OnStreamChanged, web.Handler(srv.OnStreamChanged))
		r.POST(zlm_webhook.UrlPath_OnStreamNoneReader, web.Handler(srv.OnStreamNoneReader))
		r.POST(zlm_webhook.UrlPath_OnStreamNotFound, web.Handler(srv.OnStreamNotFound))
	}
}

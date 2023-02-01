package api

import (
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	router *echo.Echo
	logger *log.Logger
	config *config
}

type config struct {
	port int
}

func NewServer(port int, logger *log.Logger) (*Server, error) {
	e := echo.New()
	e.HideBanner = true

	return &Server{
		router: e,
		logger: logger,
		config: &config{
			port: port,
		},
	}, nil
}

func (s *Server) ListenAndServe() error {
	s.registerMiddleware()
	s.registerHandlers()

	return s.router.Start(":" + strconv.Itoa(s.config.port))
}

func (s *Server) registerHandlers() {
	s.router.GET("/", Song)
}

func (s *Server) registerMiddleware() {
	s.router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(&LanguageContext{c})
		}
	})

	s.router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogHeaders:  []string{"Accept-Language"},
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				fmt.Printf("REQUEST: uri: %v, Headers: %v, status: %v\n", v.URI, v.Headers, v.Status)
			} else {
				fmt.Printf("REQUEST_ERROR: uri: %v, Headers: %v, status: %v, err: %v\n", v.URI, v.Headers, v.Status, v.Error)
			}

			return nil
		},
	}))
}

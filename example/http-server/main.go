package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"github.com/hadroncorp/geck/actuatorfx"
	"github.com/hadroncorp/geck/applicationfx"
	"github.com/hadroncorp/geck/observability/loggingfx"
	"github.com/hadroncorp/geck/securityfx"
	"github.com/hadroncorp/geck/transportfx"
	"github.com/hadroncorp/geck/validationfx"

	"http-server/pingfx"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg("error loading .env file, using system env variables")
	}

	app := fx.New(
		applicationfx.ApplicationModule,
		loggingfx.ZerologAppLoggerModule,
		actuatorfx.ActuatorModule,
		validationfx.GoPlaygroundValidationModule,
		securityfx.CognitoModule,
		transportfx.TransportModuleHTTP,
		transportfx.TransportJWTModuleHTTP,
		pingfx.PingModule,
	)
	app.Run()
}

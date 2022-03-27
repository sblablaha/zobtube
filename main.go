package main

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/sblablaha/zobtube/internal/api"
	"github.com/sblablaha/zobtube/internal/front"
	"github.com/sblablaha/zobtube/internal/webserver"
)



func main() {
	app := &cli.App{
		Name: "zobtube",
		Version: "v0.0.1",
		Usage: "passion of the zob, lube for the tube!",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:    "db-conn",
				Aliases: []string{"d"},
				Value:   "postgresql://user:pass@host:port/dbname",
				Usage:   "connection string for the database",
				EnvVars: []string{"ZBT_DB_CONN"},
				Required: true,
			},
			&cli.IntFlag{
				Name:    "listening-port",
				Aliases: []string{"p"},
				Value:   8000,
				Usage:   "port to listen to",
				EnvVars: []string{"ZBT_LISTEN_PORT"},
				DefaultText: "8000",
			},
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:   "server",
				Usage:  "start the server",
				Action: startServer,
			},
		},
	}

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal().Err(err).Msg("application failed")
  }
}

func startServer(c *cli.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("zobtube starting")

	// create web server
	log.Debug().Msg("create webserver")
	ws := webserver.New(c.Int("listening-port"))

	router := mux.NewRouter()

	// add backend part
	api.Setup(router)

	// add frontend part
	front.AddSPA(router)

	// assign router
	ws.AssignHandler(router)

	// and start
	log.Debug().Msg("start webserver")
	err := ws.ListenAndServe()
	log.Fatal().Err(err).Msg("http server closed")

	return nil
}

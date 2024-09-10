package main

import (
	"flag"
	"github.com/hayohtee/auth/internal/data"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

func main() {
	var cfg config
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file")
		os.Exit(1)
	}

	// Read the config values from the commandline flag
	flag.IntVar(&cfg.port, "port", 4000, "API server port")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("TEST_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConn, "db-max-open-conn", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConn, "db-max-idle-conn", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection established")

	app := application{
		config: cfg,
		logger: logger,
		users:  data.UserModel{DB: db},
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}

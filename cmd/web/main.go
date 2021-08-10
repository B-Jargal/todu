package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/B-Jargal/todu.git/pkg/application"
	"github.com/B-Jargal/todu.git/pkg/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var (
// 	DBConnection = "host=localhost port=5432 user=postgres password=U@>i@^<_f$xePBY3<29^ov>hG2GeNW!-qpSUL^HX+1v+U0qgzL%&#J+f@dbc dbname=console"
// 	ServicePort  = ":10000"
// 	// Bolor login
// 	ServiceName            = "QA"
// 	AuthAddress            = "https://qa.chimege.com/pub/authenticate"
// 	AuthTimeoutMinute      = 1440
// 	NotificationPrivateKey = "2NHzm7c3I-FRmZEFQ2zLpmKqP-EUKfex23vyc4Mf06w"
// )

func loadDebugSettings() {
	DBConnection = "host=localhost port=5432 user=postgres password=password dbname=console sslmode=disable"
	ServicePort = ":4000"
}

func main() {
	mode := flag.String("mode", "production", "Enable debug mode")
	flag.Parse()

	switch *mode {
	case "debug":
		loadDebugSettings()
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate/log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate/log.Ltime/log.Lshortfile)

	db, err := openDB(DBConnection)
	if err != nil {
		errorLog.Print(err)
	}

	loc, _ := time.LoadLocation("Asia/Ulaanbaatar")
	database := entity.New(db, loc)

	app := &application.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Location: loc,
		DataBase: database,
	}

	database.DB.AutoMigrate(entity.Owners{})
	srv := &http.Server{
		Handler:      routes(app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

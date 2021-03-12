package main

import (
	"autodom/services/advisor/middleware"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	kithttp "github.com/go-kit/kit/transport/http"
	_ "github.com/go-sql-driver/mysql"

	"autodom/services/advisor"
	impl "autodom/services/advisor/implementation"
	"autodom/services/advisor/sqldb"
	"autodom/services/advisor/transport"
	httptransport "autodom/services/advisor/transport/http"
)

var (
	userFlag     = flag.String("u", "", "DB user")
	passFlag     = flag.String("p", "", "DB password")
	dbNameFlag   = flag.String("d", "", "DB name")
	httpAddrFlag = flag.String("http.addr", ":8080", "HTTP listen address")
)

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *userFlag == "" || *passFlag == "" || *dbNameFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "solution",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error
		credits := fmt.Sprintf("%s:%s@/%s", *userFlag, *passFlag, *dbNameFlag)
		fmt.Println("CREDITS", credits)

		db, err = sql.Open("mysql", credits)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	var svc advisor.Service
	{
		repository, err := sqldb.New(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = impl.NewService(repository, logger)

		svc = middleware.LoggingMiddleware(logger)(svc)
	}

	var h http.Handler
	{
		endpoints := transport.MakeEndpoints(svc)
		serverOptions := []kithttp.ServerOption{}
		h = httptransport.NewService(endpoints, serverOptions, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddrFlag)
		server := &http.Server{
			Addr:    *httpAddrFlag,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
}

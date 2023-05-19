package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Zmey56/openai-api-proxy/authorization"
	"github.com/Zmey56/openai-api-proxy/log"
	"github.com/Zmey56/openai-api-proxy/repository"
	"github.com/Zmey56/openai-api-proxy/server/middlewares"
	"github.com/Zmey56/openai-api-proxy/server/proxy"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
)

var (
	logLevel = flag.String("log-level", "info", "the level of logging (debug, info, warning, error)")

	serverCmd     = flag.NewFlagSet("server", flag.ExitOnError)
	openaiToken   = serverCmd.String("openai-token", os.Getenv("OPENAI_TOKEN"), "the token used to communicate with OpenAI API")
	openaiAddress = serverCmd.String("openai-address", "https://api.openai.com", "the address of the OpenAI API")
	serverAddress = serverCmd.String("local-addr", "localhost:8080", "the binding for the server (host and port)")
	serverDBLoc   = serverCmd.String("db-location", "db.sqlite3", "the location of the database")

	initdbCmd    = flag.NewFlagSet("initdb", flag.ExitOnError)
	initdbDBLoc  = initdbCmd.String("db-location", "db.sqlite3", "the location of the database")
	addTestUsers = initdbCmd.Bool("add-test-users", false, "add test users to the database")
)

func printUsage() {
	fmt.Printf("Usage: %s <command> [options]\n\n", os.Args[0])
	fmt.Println("Commands:")
	fmt.Println("  server   start server with specified options")
	fmt.Println("  initdb   initialize the database")
	fmt.Println("  help     prints usage info")
	fmt.Println()

	flag.PrintDefaults()
	fmt.Println()

	fmt.Println("Server command flags:")
	serverCmd.PrintDefaults()
	fmt.Println()

	fmt.Println("InitDB command flags:")
	initdbCmd.PrintDefaults()
	fmt.Println()
}

func main() {
	flag.Parse()

	err := log.SetLevel(*logLevel)
	if err != nil {
		fmt.Printf("Failed to set log level: %s\n", err)
		printUsage()
		os.Exit(1)
	}

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "server":
		if err := serverCmd.Parse(args[1:]); err != nil {
			log.Error.Print(err)
			printUsage()
			os.Exit(1)
		}
		if err := runServer(); err != nil {
			log.Error.Fatal(err)
		}
	case "initdb":
		if err := initdbCmd.Parse(args[1:]); err != nil {
			log.Error.Print(err)
			printUsage()
			os.Exit(1)
		}
		if err := runInitDb(); err != nil {
			log.Error.Fatal(err)
		}
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func runServer() error {
	mux := http.NewServeMux()

	proxyInst, err := proxy.NewProxy(proxy.Configuration{
		OpenaiToken:   *openaiToken,
		OpenaiAddress: *openaiAddress,
	})
	if err != nil {
		return err
	}

	authService := authorization.StaticService{}

	fmt.Println("proxyInst", proxyInst)
	fmt.Println("authService", authService)

	mux.Handle("/openai/",
		middlewares.RemovePathPrefixMiddleware(
			middlewares.AuthorizationMiddleware(proxyInst, authService),
			"/openai/",
		),
	)

	return http.ListenAndServe(*serverAddress, mux)
}

func runInitDb() error {
	dbLocation := *initdbDBLoc

	log.Debug.Printf("creating database in file %s", dbLocation)

	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error.Printf("failed to close database: %s", err)
		}
	}()

	repository.CreatedTableUsers(db)

	if *addTestUsers {
		repository.AddTestUsers(db)
	}

	log.Debug.Printf("database created successfully at %s", dbLocation)
	return nil
}

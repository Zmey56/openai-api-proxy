package main

import (
	"flag"
	"fmt"
	"github.com/Zmey56/openai-api-proxy/authorization"
	"github.com/Zmey56/openai-api-proxy/log"
	"github.com/Zmey56/openai-api-proxy/server/middlewares"
	"github.com/Zmey56/openai-api-proxy/server/proxy"
	"net/http"
	"os"
)

var (
	logLevel = flag.String("log-level", "info", "the level of logging (debug, info, warning, error)")

	serverCmd     = flag.NewFlagSet("server", flag.ExitOnError)
	openaiToken   = serverCmd.String("openai-token", os.Getenv("OPENAI_TOKEN"), "the token used to communicate with OpenAI API")
	openaiAddress = serverCmd.String("openai-address", "https://api.openai.com", "the address of the OpenAI API")
	serverAddress = serverCmd.String("local-addr", "localhost:8080", "the binding for the server (host and port)")
	serverDBLoc   = serverCmd.String("db", "db.sqlite3", "the location of the database")

	initdbCmd   = flag.NewFlagSet("initdb", flag.ExitOnError)
	initdbDBLoc = initdbCmd.String("db", "db.sqlite3", "the location of the database")
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
		runInitDb()
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

	mux.Handle("/openai/",
		middlewares.RemovePathPrefixMiddleware(
			middlewares.AuthorizationMiddleware(proxyInst, authService),
			"/openai/",
		),
	)

	return http.ListenAndServe(*serverAddress, mux)
}

func runInitDb() {
	// TODO: when we will have a database
}

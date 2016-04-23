package main

import (
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/RichardKnop/logging"
	"github.com/codegangsta/cli"
)

var (
	cliApp     *cli.App
	httpClient *http.Client
	logger     *logging.Logger
)

func init() {
	// Initialise a CLI app
	cliApp = cli.NewApp()
	cliApp.Name = "Pinglist Proxy"
	cliApp.Usage = "pinglist-proxy"
	cliApp.Author = "Richard Knop"
	cliApp.Email = "risoknop@gmail.com"
	cliApp.Version = "0.0.0"

	// Initialise the logger
	logger = logging.New(nil, nil, new(logging.ColouredFormatter))
}

func main() {
	// Set the CLI app flags
	cliApp.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "timeout",
			Value: 10,
			Usage: "Timeout for requests in seconds",
		},
	}

	// Set the CLI app commands
	cliApp.Commands = []cli.Command{
		{
			Name:  "runserver",
			Usage: "run web server",
			Action: func(c *cli.Context) {
				// Prepare HTTP client
				httpClient = &http.Client{
					Timeout: time.Duration(c.Int("timeout")) * time.Second,
				}
				http.HandleFunc("/", handler)
				http.ListenAndServe(":8090", nil)
			},
		},
	}

	// Run the CLI app
	cliApp.Run(os.Args)
}

func handler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, false)
	if err != nil {
		logger.Infof("%v", reqDump)
	}

	// Get the requet URL query string parameter
	requestURL := r.URL.Query().Get("request_url")
	if requestURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare a request
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Make the request
	resp, _ := httpClient.Do(req)
	respDump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(respDump)
}

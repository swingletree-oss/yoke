package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/error418/yoke/internal/buildinfo"
	"github.com/error418/yoke/internal/prettylog"
	"github.com/error418/yoke/internal/swingletree"
	"github.com/urfave/cli"
)

var Version string

func main() {
	app := cli.NewApp()

	app.EnableBashCompletion = true

	app.Usage = "Swingletree client cli"
	app.Version = Version

	var endpoint string
	var configFile string
	var apiToken string

	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generates information for other build tools",
			Subcommands: []cli.Command{
				{
					Name:  "buildid",
					Usage: "calculates the build id for this build and prints it",
					Action: func(c *cli.Context) error {
						info, err := buildinfo.NewBuildInfo()
						if err != nil {
							return cli.NewExitError("Failed to generate build id", 100)
						}

						fmt.Println(info.BuildId())

						return nil
					},
				},
			},
		},

		{
			Name:    "publish",
			Aliases: []string{"p"},
			Usage:   "publishes a build report to Swingletree",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "endpoint, e",
					Usage:       "Swingletree Gate API endpoint base",
					EnvVar:      "YOKE_ENDPOINT",
					Destination: &endpoint,
				},
				cli.StringFlag{
					Name:        "config, c",
					Value:       ".swingletree.yml",
					Usage:       "Swingletree repository configuration file",
					Destination: &configFile,
				},
				cli.StringFlag{
					Name:        "token, t",
					Value:       "",
					Usage:       "Swingletree API token. Consider setting this using the env var",
					EnvVar:      "YOKE_TOKEN",
					Destination: &apiToken,
				},
				cli.BoolFlag{
					Name:  "insecure",
					Usage: "Enable sending HTTP Authentication data via unsecured http connections",
				},
			},
			Action: func(c *cli.Context) error {

				if endpoint == "" {
					return cli.NewExitError("Yoke api endpoint is not set. Please provide it.", 1)
				}

				if apiToken == "" {
					prettylog.Warn("No authentication token was provided.")
				}

				if !c.Bool("insecure") && apiToken != "" && strings.HasPrefix(endpoint, "http://") {
					return cli.NewExitError("Prevented sending credentials over insecure http connection. Use --insecure to send anyway, if you know what you are doing.", 100)
				}

				prettylog.Info("Running publish to %s using config %s", endpoint, configFile)

				conf, err := swingletree.LoadConf(configFile)
				if err != nil {
					return cli.NewExitError(fmt.Sprintf("Could not load config from %s: %s", configFile, err), 1)
				}

				info, err := buildinfo.NewBuildInfo()
				if err != nil {
					return cli.NewExitError(fmt.Sprintf("Failed to initialize build information collector %v", err), 1)
				}

				fmt.Printf("\nExtracted git information:\n%s\n\n", info.GitInfo.String())

				publishReport, err := info.Transmit(endpoint, apiToken, conf)
				if err != nil {
					return cli.NewExitError(fmt.Sprintf("Failed running command."), 1)
				}

				fmt.Println("\n\nPublishing summary ---")
				fmt.Println(publishReport)

				if publishReport.Missing > 0 || publishReport.Failures > 0 {
					return cli.NewExitError("Encountered problems while uploading reports", 50)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

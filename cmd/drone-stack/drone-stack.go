package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"megpoid.xyz/go/drone-stack"
)

func main() {
	// Load env-file if it exists first
	if env := os.Getenv("PLUGIN_ENV_FILE"); env != "" {
		if err := godotenv.Load(env); err != nil {
			fmt.Printf("Cannot load env file: %s", err.Error())
		}
	}

	app := cli.NewApp()
	app.Name = "drone-stack plugin"
	app.Usage = "drone-stack plugin"
	app.Action = run
	app.Version = Version
	cli.VersionPrinter = printVersion

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "host",
			Usage:   "docker host",
			EnvVars: []string{"PLUGIN_HOST"},
		},
		&cli.BoolFlag{
			Name:    "tls",
			Usage:   "docker tls",
			EnvVars: []string{"PLUGIN_TLS"},
		},
		&cli.BoolFlag{
			Name:    "tlsverify",
			Usage:   "docker tlsverify",
			EnvVars: []string{"PLUGIN_TLSVERIFY"},
		},
		&cli.StringSliceFlag{
			Name:    "compose",
			Usage:   "stack deploy compose",
			Value:   cli.NewStringSlice("docker-compose.yml"),
			EnvVars: []string{"PLUGIN_COMPOSE"},
		},
		&cli.StringFlag{
			Name:    "stack.name",
			Usage:   "stack deploy name",
			EnvVars: []string{"PLUGIN_STACK_NAME"},
		},
		&cli.BoolFlag{
			Name:    "prune",
			Usage:   "stack deploy prune",
			EnvVars: []string{"PLUGIN_PRUNE"},
		},
		&cli.StringFlag{
			Name:    "docker.registry",
			Usage:   "docker registry",
			Value:   "https://index.docker.io/v1/",
            EnvVars: []string{"PLUGIN_REGISTRY","PLUGIN_DOCKER_REGISTRY"},
		},
		&cli.StringFlag{
			Name:    "docker.username",
			Usage:   "docker username",
			EnvVars: []string{"PLUGIN_USERNAME","PLUGIN_DOCKER_USERNAME"},
		},
		&cli.StringFlag{
			Name:    "docker.password",
			Usage:   "docker password",
			EnvVars: []string{"PLUGIN_PASSWORD","PLUGIN_DOCKER_PASSWORD"},
		},
		&cli.StringFlag{
			Name:    "docker.email",
			Usage:   "docker email",
			EnvVars: []string{"PLUGIN_EMAIL","DOCKER_EMAIL"},
		},
		&cli.StringFlag{
			Name:    "docker.cacert",
			Usage:   "docker ca",
			EnvVars: []string{"PLUGIN_CACERT","DOCKER_CACERT"},
		},
		&cli.StringFlag{
			Name:    "docker.key",
			Usage:   "docker key",
			EnvVars: []string{"PLUGIN_KEY","DOCKER_KEY"},
		},
		&cli.StringFlag{
			Name:    "docker.cert",
			Usage:   "docker cert",
			EnvVars: []string{"PLUGIN_CERT","DOCKER_CERT"},
		},
		&cli.StringFlag{
			Name:    "sshkey",
			Usage:   "docker ssh key",
			EnvVars: []string{"PLUGIN_SSH_KEY"},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	cli.VersionPrinter(c)

	plugin := docker.Plugin{
		Login: docker.Login{
			Registry: c.String("docker.registry"),
			Username: c.String("docker.username"),
			Password: c.String("docker.password"),
			Email:    c.String("docker.email"),
		},
		Deploy: docker.Deploy{
			Name:    c.String("stack.name"),
			Compose: c.StringSlice("compose"),
			Prune:   c.Bool("prune"),
		},
		Certs: docker.Certs{
			TLSKey:    c.String("docker.key"),
			TLSCert:   c.String("docker.cert"),
			TLSCACert: c.String("docker.cacert"),
		},
		Host: docker.Host{
			Host:      c.String("host"),
			UseTLS:    c.Bool("tls"),
			TLSVerify: c.Bool("tlsverify"),
		},
		SSH: docker.SSH{
			Key: c.String("sshkey"),
		},
	}

	return plugin.Exec()
}

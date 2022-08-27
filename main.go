/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package main

import (
	"flag"
	"fmt"
	"github.com/rluisr/datadog-slo-insufflate/config"
	"github.com/rluisr/datadog-slo-insufflate/datadog"
	"github.com/rluisr/datadog-slo-insufflate/slack"
	"os"
)

var (
	version string
)

func main() {
	versionSFlag := flag.Bool("v", false, "version")
	versionLFlag := flag.Bool("version", false, "version")
	flag.Parse()
	if *versionSFlag || *versionLFlag {
		fmt.Println("version: ", version)
		os.Exit(0)
	}

	env, err := config.NewEnv()
	if err != nil {
		panic(err)
	}

	ddClient := datadog.NewClint()
	sloResults, err := ddClient.FetchSLO()
	if err != nil {
		panic(err)
	}

	slackClient, err := slack.NewClient(env)
	if err != nil {
		panic(err)
	}

	blocks := slackClient.BuildBlock(sloResults, env.DDSLOURL)
	err = slackClient.Send(blocks)
	if err != nil {
		panic(err)
	}
}

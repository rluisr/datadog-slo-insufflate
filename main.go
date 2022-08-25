/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package main

import (
	"github.com/rluisr/datadog-slo-insufflate/config"
	"github.com/rluisr/datadog-slo-insufflate/datadog"
	"github.com/rluisr/datadog-slo-insufflate/slack"
	"time"
)

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

func main() {
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

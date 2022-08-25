/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package config

import goEnv "github.com/Netflix/go-env"

type Env struct {
	SlackToken        string `env:"SLACK_TOKEN,required=true"`
	SlackChannelID    string `env:"SLACK_CHANNEL_ID,required=true"`
	SlackMessageTitle string `env:"SLACK_MESSAGE_TITLE,default=SLO daily reports :eyes:"`
	DDSLOURL          string `env:"DD_SLO_URL,default=https://app.datadoghq.com/slo"` // default https://app.datadoghq.com/slo
}

func NewEnv() (*Env, error) {
	var env Env

	_, err := goEnv.UnmarshalFromEnviron(&env)
	if err != nil {
		return nil, err
	}

	return &env, nil
}

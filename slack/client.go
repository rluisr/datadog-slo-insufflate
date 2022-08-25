/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package slack

import (
	"fmt"
	"github.com/rluisr/datadog-slo-insufflate/config"
	"github.com/slack-go/slack"
)

type Client struct {
	API       *slack.Client
	Token     string
	ChannelID string
	Title     string
}

func NewClient(env *config.Env) (*Client, error) {
	client := &Client{
		API:       slack.New(env.SlackToken),
		ChannelID: env.SlackChannelID,
		Title:     env.SlackMessageTitle,
	}

	r, err := client.API.AuthTest()
	if err != nil {
		return nil, fmt.Errorf("authentication failed slack %w, response: %s", err, r)
	}

	return client, nil
}

func (s *Client) Send(blocks []slack.Block) error {
	_, _, err := s.API.PostMessage(s.ChannelID, slack.MsgOptionBlocks(blocks...))
	if err != nil {
		return fmt.Errorf("failed PostMessage: %w", err)
	}

	return nil
}

/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package slack

import (
	"fmt"
	"github.com/rluisr/datadog-slo-insufflate/lib"
	"github.com/rluisr/datadog-slo-insufflate/models"
	"github.com/slack-go/slack"
	"math"
)

func (s *Client) BuildBlock(sloResults []models.SLOResult, ddSLOURL string) []slack.Block {
	headerBlock := &slack.SectionBlock{
		Type: slack.MBTHeader,
		Text: &slack.TextBlockObject{
			Type:  slack.PlainTextType,
			Text:  s.Title,
			Emoji: true,
		},
	}

	blocks := []slack.Block{
		headerBlock,
		slack.NewDividerBlock(),
	}

	var sloBlocks []slack.Block

	for i := range sloResults {
		slo := sloResults[i]

		// Name
		// https://app.datadoghq.com/slo?slo_id=id&tab=status_and_history&timeframe=30d
		sloURL := fmt.Sprintf("%s?slo_id=%s&tab=status_and_history&timeframe=%s", ddSLOURL, slo.ID, slo.TimeFrame)
		targetText, statusText, errorBudgetText := s.decorateText(&slo)

		sloBlock := &slack.SectionBlock{
			Type: slack.MBTSection,
			Text: &slack.TextBlockObject{
				Type:     slack.MarkdownType,
				Text:     fmt.Sprintf("*SLO: * <%s|%s> / %s\n%s", sloURL, slo.Name, slo.TimeFrame, slo.Description),
				Verbatim: true,
			},
			Fields: []*slack.TextBlockObject{
				// Target
				{
					Type: slack.MarkdownType,
					Text: targetText,
				},
				// Status
				{
					Type: slack.MarkdownType,
					Text: statusText,
				},
				// Error budge remaining
				{
					Type: slack.MarkdownType,
					Text: errorBudgetText,
				},
			},
		}
		sloBlocks = append(sloBlocks, sloBlock)
	}

	blocks = append(blocks, sloBlocks...)

	return blocks
}

func (s *Client) decorateText(sloResult *models.SLOResult) (target, status, errorBudget string) {
	// Target
	target = fmt.Sprintf(":eyes: *Target: * %s%%", lib.ConvertF64ToString(sloResult.Target))

	// Status
	if sloResult.Target > sloResult.Status {
		status = fmt.Sprintf(":red_circle: *Status: %s%%*", lib.ConvertF64ToString(sloResult.Status))
	} else {
		status = fmt.Sprintf(":green_heart: *Status: %s%%*", lib.ConvertF64ToString(sloResult.Status))
	}

	// Error budget
	if math.Signbit(sloResult.ErrorBudgetRemaining) {
		errorBudget = fmt.Sprintf(":red_circle: *Error budget: %s%%*", lib.ConvertF64ToString(sloResult.ErrorBudgetRemaining))
	} else {
		errorBudget = fmt.Sprintf(":green_heart: *Error budget: %s%%*", lib.ConvertF64ToString(sloResult.ErrorBudgetRemaining))
	}

	return target, status, errorBudget
}

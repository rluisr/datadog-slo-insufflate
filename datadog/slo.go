/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package datadog

import (
	"fmt"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/rluisr/datadog-slo-insufflate/models"
	"time"
)

func (c *Client) FetchSLO() ([]models.SLOResult, error) {
	slos, err := c.listSLOs()
	if err != nil {
		return nil, err
	}

	var sloResults []models.SLOResult

	for i := range slos {
		slo := slos[i]

		sloResult := models.SLOResult{
			Name:        slo.GetName(),
			ID:          slo.GetId(),
			Description: slo.GetDescription(),
		}

		thresholds := slo.GetThresholds()
		for _, threshold := range thresholds {
			sloResult.Target = threshold.GetTarget()
			sloResult.TimeFrame = string(threshold.GetTimeframe())
		}

		history, err := c.getSLOHistory(slo.GetId(), &sloResult.Target)
		if err != nil {
			return nil, err
		}
		overall := history.GetOverall()
		sloResult.Status = overall.GetSliValue()
		sloResult.ErrorBudgetRemaining = overall.GetErrorBudgetRemaining()["custom"]
		sloResults = append(sloResults, sloResult)
	}

	return sloResults, nil
}

func (c *Client) listSLOs() ([]datadogV1.ServiceLevelObjective, error) {
	resp, r, err := c.SLOAPIClient.ListSLOs(c.ctx, *datadogV1.NewListSLOsOptionalParameters())
	if err != nil {
		echoError(r)
		return nil, fmt.Errorf("ListSLOs err: %w", err)
	}

	return resp.Data, nil
}

func (c *Client) getSLOHistory(sloID string, targetSLO *float64) (*datadogV1.SLOHistoryResponseData, error) {
	resp, r, err := c.SLOAPIClient.GetSLOHistory(c.ctx, sloID, time.Now().AddDate(0, 0, -30).Unix(), time.Now().Unix(), datadogV1.GetSLOHistoryOptionalParameters{
		Target: targetSLO,
	})
	if err != nil {
		echoError(r)
		return nil, fmt.Errorf("GetSLOHistory err: %w", err)
	}

	return resp.Data, nil
}

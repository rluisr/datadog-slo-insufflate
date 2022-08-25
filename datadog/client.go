/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package datadog

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

type Client struct {
	ctx          context.Context
	APIClient    *datadog.APIClient
	SLOAPIClient *datadogV1.ServiceLevelObjectivesApi
}

func NewClint() *Client {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v1.SearchSLO", true)
	configuration.SetUnstableOperationEnabled("v1.GetSLOHistory", true)

	apiClient := datadog.NewAPIClient(configuration)
	sloAPI := datadogV1.NewServiceLevelObjectivesApi(apiClient)

	return &Client{
		ctx:          ctx,
		APIClient:    apiClient,
		SLOAPIClient: sloAPI,
	}
}

/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package models

type SLOResult struct {
	Name                 string
	ID                   string
	Description          string
	TimeFrame            string
	Target               float64 // in percent
	Status               float64 // in percent
	ErrorBudgetRemaining float64 // in percent
}

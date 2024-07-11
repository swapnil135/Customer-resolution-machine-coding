package main

import(
	"time"
)


type IssueType string

const (
	PaymentRelated   IssueType = "Payment Related"
	MutualFundRelated IssueType = "Mutual Fund Related"
	GoldRelated       IssueType = "Gold Related"
	InsuranceRelated  IssueType = "Insurance Related"
)

type IssueStatus string

const (
	Open       IssueStatus = "OPEN"
	InProgress IssueStatus = "IN_PROGRESS"
	Resolved   IssueStatus = "RESOLVED"
)

type Issue struct {
	ID           string
	TransactionID string
	IssueType    IssueType
	Subject      string
	Description  string
	Email        string
	Status       IssueStatus
	Resolution   string
	CreatedAt    time.Time
	AgentID      string
}
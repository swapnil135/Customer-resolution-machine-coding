package main

type Agent struct {
	ID          string
	Email       string
	Name        string
	Expertise   []IssueType
	AssignedIssues []string
}
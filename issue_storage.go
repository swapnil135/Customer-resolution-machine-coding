package main

import (
	"errors"
	"fmt"
)

type IIssueStorage interface {
	GetIssueById(issueId string) (*Issue, error)
	CreateIssue(issue *Issue) error
	UpdateIssue(issue *Issue) error
	GetIssuesBySpecification(spec *AndSpecification) ([]*Issue, error)
}

type MapIssueStorage struct {
	IdToIssueMapping map[string]*Issue
}

func NewMapIssueStorage() *MapIssueStorage {
	return &MapIssueStorage{
		IdToIssueMapping: make(map[string]*Issue),
	}
}

func (m *MapIssueStorage) GetIssueById(issueId string) (*Issue, error) {
	issue, ok := m.IdToIssueMapping[issueId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("issue not found for id: %v", issueId))
	}
	return issue, nil
}

func (m *MapIssueStorage) CreateIssue(issue *Issue) error {
	if issue == nil || issue.ID == "" {
		return errors.New("Invalid argument")
	}

	_, ok := m.IdToIssueMapping[issue.ID]
	if ok {
		return errors.New(fmt.Sprintf("issue already exists with id: %v", issue.ID))
	}

	m.IdToIssueMapping[issue.ID] = issue
	return nil
}

func (m *MapIssueStorage) UpdateIssue(issue *Issue) error {
	if issue == nil || issue.ID == "" {
		return errors.New("Invalid argument")
	}

	_, ok := m.IdToIssueMapping[issue.ID]
	if !ok {
		return errors.New(fmt.Sprintf("issue not found for id: %v", issue.ID))
	}

	m.IdToIssueMapping[issue.ID] = issue
	return nil
}

func (m *MapIssueStorage) GetIssuesBySpecification(spec *AndSpecification) ([]*Issue, error) {
	var result []*Issue
	for _, issue := range m.IdToIssueMapping {
		
		if spec.IsSatisfiedBy(issue) {
			result = append(result, issue)
		}
	}
	return result, nil
}
package main

import (
	"errors"
	"fmt"
	"time"
)

type ICustomerResolutionService interface {
	CreateIssue(id, transactionId, subject, description, email string, issueType IssueType) error
	AddAgent(agentId, agentEmail, agentName string, expertise []IssueType) error
	AssignIssue(issueId string, strategy IAssignStrategy) error
	GetIssues(specs []Specification) ([]*Issue, error)
	UpdateIssue(issueId string, status IssueStatus, resolution string) error
	ResolveIssue(issueId string, resolution string) error
}

type Service struct {
	agentStorage IAgentStorage
	issueStorage IIssueStorage
}

func NewService(agentStorage IAgentStorage, issueStorage IIssueStorage) *Service {
	return &Service{
		agentStorage: agentStorage,
		issueStorage: issueStorage,
	}
}

func (s *Service) CreateIssue(id, transactionId, subject, description, email string, issueType IssueType) error {
	issue := &Issue{
		ID: id,
		TransactionID: transactionId,
		Subject:       subject,
		Description:   description,
		Email:         email,
		IssueType:     issueType,
		Status:        Open,
		CreatedAt:     time.Now(),
	}

	err := s.issueStorage.CreateIssue(issue)
	if err != nil {
		fmt.Printf("Error creating issue: %v\n", err)
		return errors.New("error creating issue")
	}

	fmt.Printf("Issue created successfully: %s\n", issue.ID)
	return nil
}

func (s *Service) AddAgent(agentId, agentEmail, agentName string, expertise []IssueType) error {
	agent := &Agent{
		ID: 			agentId,
		Email:           agentEmail,
		Name:            agentName,
		Expertise:       expertise,
		AssignedIssues:  []string{},
	}

	err := s.agentStorage.CreateAgent(agent)
	if err != nil {
		fmt.Printf("Error adding agent: %v\n", err)
		return errors.New("error adding agent")
	}

	fmt.Printf("Agent added successfully: %s\n", agent.ID)
	return nil
}

func (s *Service) AssignIssue(issueId string, strategy IAssignStrategy) error {
	// making sure issue exists in system
	issue, err := s.issueStorage.GetIssueById(issueId)
	if err != nil {
		fmt.Printf("Error getting issue: %v\n", err)
		return errors.New("error getting issue")
	}

	agentId ,err := strategy.GetAgentIdForIssue(issue, s.agentStorage.GetAllAgents())
	if err != nil {
		fmt.Printf("Error getting agent for issue: %v\n", err)
		return errors.New("agent not found for assignment")
	}

	err = s.agentStorage.AssignIssue(agentId, issueId)
	if err != nil {
		fmt.Printf("Error assigning issue to agent: %v\n", err)
		return errors.New("error assigning issue")
	}

	fmt.Printf("Issue %s assigned to agent %s\n", issueId, agentId)
	return nil
}

func (s *Service) GetIssues(specs []Specification) ([]*Issue, error) {
	combinedSpec := &AndSpecification{Specifications: specs}

	issues, err := s.issueStorage.GetIssuesBySpecification(combinedSpec)
	if err != nil {
		fmt.Printf("Error getting issues: %v\n", err)
		return nil, errors.New("error getting issues")
	}

	fmt.Println("Issues retrieved successfully")
	return issues, nil
}

func (s *Service) UpdateIssue(issueId string, status IssueStatus, resolution string) error {
	issue, err := s.issueStorage.GetIssueById(issueId)
	if err != nil {
		fmt.Printf("Error getting issue: %v\n", err)
		return errors.New("error getting issue")
	}

	issue.Status = status
	issue.Resolution = resolution

	err = s.issueStorage.UpdateIssue(issue)
	if err != nil {
		fmt.Printf("Error updating issue: %v\n", err)
		return errors.New("error updating issue")
	}

	fmt.Printf("Issue %s updated: status - %s, resolution - %s\n", issueId, status, resolution)
	return nil
}

func (s *Service) ResolveIssue(issueId string, resolution string) error {
	err := s.UpdateIssue(issueId, Resolved, resolution)
	if err != nil {
		fmt.Printf("Error resolving issue: %v\n", err)
		return errors.New("error resolving issue")
	}

	fmt.Printf("Issue %s resolved: %s\n", issueId, resolution)
	return nil
}
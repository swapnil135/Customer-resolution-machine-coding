package main

import (
	"errors"
	"fmt"
)

type IAgentStorage interface {
	GetAgentById(agentId string) (*Agent, error)
	CreateAgent(agent *Agent) error
	// can be merged into one update issue method
	AssignIssue(agentId string, issueId string) error
	ResolveIssue(agentId string, issueId string) error
	// can be merged into one update issue method
	GetAllAgents() []*Agent
}

type MapAgentStorage struct {
	IdToAgentMapping map[string]*Agent
}

func NewMapAgentStorage() *MapAgentStorage {
	return &MapAgentStorage{
		IdToAgentMapping: make(map[string]*Agent),
	}
}

func (m *MapAgentStorage) GetAllAgents() []*Agent{
	res := []*Agent{}

	for _, existingAgent := range m.IdToAgentMapping{
		res = append(res, existingAgent)
	}

	return res
}

func (m *MapAgentStorage) AssignIssue(agentId string, issueId string) error {
	agent, ok := m.IdToAgentMapping[agentId]
	if !ok {
		return errors.New(fmt.Sprintf("agent nor found for id: %v", agentId))
	}

	// handle issue already assigned error

	agent.AssignedIssues = append(agent.AssignedIssues, issueId)

	return nil
}

func (m *MapAgentStorage) ResolveIssue(agentId string, issueId string) error {
	agent, ok := m.IdToAgentMapping[agentId]
	if !ok {
		return errors.New(fmt.Sprintf("agent nor found for id: %v", agentId))
	}

	for idx, existingIssueId := range agent.AssignedIssues {
		if existingIssueId == issueId {
			agent.AssignedIssues = append(agent.AssignedIssues[:idx], agent.AssignedIssues[idx+1:]...)
			return nil
		}
	}

	return errors.New("given issueId not assigned to given agent")
}

func (m *MapAgentStorage) GetAgentById(agentId string) (*Agent, error) {
	agent, ok := m.IdToAgentMapping[agentId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("agent nor found for id: %v", agentId))
	}

	return agent, nil
}

func (m *MapAgentStorage) CreateAgent(agent *Agent) error{
	if agent == nil || agent.ID == ""{
		return errors.New("Invalid argument")
	}

	_, ok := m.IdToAgentMapping[agent.ID]
	if ok {
		return errors.New(fmt.Sprintf("agent already exists with id: %v", agent.ID))
	}

	m.IdToAgentMapping[agent.ID] = agent

	return nil
}

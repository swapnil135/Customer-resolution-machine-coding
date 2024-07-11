package main

import "errors"

type IAssignStrategy interface {
	GetAgentIdForIssue(issue *Issue, agents []*Agent) (string, error)
}


type FirstFreeAgentAssignStrategy struct{}

func (s *FirstFreeAgentAssignStrategy) GetAgentIdForIssue(issue *Issue, agents []*Agent) (string, error) {
	for _, agent := range agents {
		if len(agent.AssignedIssues) == 0 {
			return agent.ID, nil
		}
	}
	return "", errors.New("no available agent found for the issue")
}
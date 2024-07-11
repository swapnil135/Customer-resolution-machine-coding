package main

import (
	"fmt"
)

func main() {
	// Initialize service with mock storage implementations
	agentStorage := NewMapAgentStorage()
	issueStorage := NewMapIssueStorage()
	service := NewService(agentStorage, issueStorage)

	// Test scenarios
	fmt.Println("--- Creating Issues ---")
	err := service.CreateIssue("I1", "T1", "Payment Related", "Payment Failed", "My payment failed but money is debited", "testUser1@test.com")
	if err != nil {
		fmt.Println("Error creating issue:", err)
	}
	err = service.CreateIssue("I2", "T2", "Mutual Fund Related", "Purchase Failed", "Unable to purchase Mutual Fund", "testUser2@test.com")
	if err != nil {
		fmt.Println("Error creating issue:", err)
	}
	err = service.CreateIssue("I3", "T3", "Payment Related", "Payment Failed", "My payment failed but money is debited", "testUser2@test.com")
	if err != nil {
		fmt.Println("Error creating issue:", err)
	}

	fmt.Println("--- Adding Agents ---")
	err = service.AddAgent("A1", "agent1@test.com", "Agent 1", []IssueType{"Payment Related", "Gold Related"})
	if err != nil {
		fmt.Println("Error adding agent:", err)
	}
	err = service.AddAgent("A2","agent2@test.com", "Agent 2", []IssueType{"Payment Related"})
	if err != nil {
		fmt.Println("Error adding agent:", err)
	}

	assignStrategy := FirstFreeAgentAssignStrategy{}

	fmt.Println("--- Assigning Issues ---")
	err = service.AssignIssue("I1", &assignStrategy)
	if err != nil {
		fmt.Println("Error assigning issue:", err)
	}
	err = service.AssignIssue("I2", &assignStrategy)
	if err != nil {
		fmt.Println("Error assigning issue:", err)
	}
	err = service.AssignIssue("I3", &assignStrategy)
	if err != nil {
		fmt.Println("Error assigning issue:", err)
	}

	fmt.Println("--- Getting Issues by Filter ---")
	emailSpec := EmailSpecification{
		Email: "testUser2@test.com",
	}
	issuesByEmail, err := service.GetIssues([]Specification{&emailSpec})
	if err != nil {
		fmt.Println("Error getting issues by email:", err)
	} else {
		fmt.Println("Issues by email:")
		for _, issue := range issuesByEmail {
			fmt.Printf("%+v\n", *issue)
		}
	}

	fmt.Println("--- Updating Issue ---")
	err = service.UpdateIssue("I3", InProgress, "Waiting for payment confirmation")
	if err != nil {
		fmt.Println("Error updating issue:", err)
	}

	fmt.Println("--- Resolving Issue ---")
	err = service.ResolveIssue("I3", "PaymentFailed debited amount will get reversed")
	if err != nil {
		fmt.Println("Error resolving issue:", err)
	}
}
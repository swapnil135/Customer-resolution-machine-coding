package main

// can be extended by various specifications based on filtering
type Specification interface {
	IsSatisfiedBy(issue *Issue) bool
}

type AndSpecification struct {
	Specifications []Specification
}

func (a *AndSpecification) IsSatisfiedBy(issue *Issue) bool {
	for _, spec := range a.Specifications {
		if !spec.IsSatisfiedBy(issue) {
			return false
		}
	}
	return true
}

// concrete email specification
type EmailSpecification struct {
	Email string
}

func (e *EmailSpecification) IsSatisfiedBy(issue *Issue) bool {
	return issue.Email == e.Email
}

type TypeSpecification struct {
	IssueType IssueType
}

func (t *TypeSpecification) IsSatisfiedBy(issue *Issue) bool {
	return issue.IssueType == t.IssueType
}
package superclog

// Component holds which component was affected
type Component map[string][]*Commit

type ConventionalCommit struct {
	Type      Category
	Component string
	Message   string
}

type Commit struct {
	Name         string
	Message      string
	Hash         string
	ShortHash    string
	ReleaseNote  string
	Conventional ConventionalCommit
}


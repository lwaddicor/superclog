package superclog

import (
	"bufio"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"regexp"
	"strings"
)

const (
	ReleaseNoteTag = "#r"
)

var (
	convCommitMatcher = regexp.MustCompile(`^(build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test)\((.+)\): (.+$)`)
)

func GetCommits(path string, fromHash, toHash string) ([]Commit, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	var to *object.Commit
	if toHash == "" {
		h, err := r.Head()
		if err != nil {
			return nil, err
		}
		if to, err = r.CommitObject(h.Hash()); err != nil {
			return nil, err
		}

	} else {
		if to, err = r.CommitObject(plumbing.NewHash(toHash)); err != nil {
			return nil, err
		}
	}

	from, err := r.CommitObject(plumbing.NewHash(fromHash))
	if err != nil {
		return nil, fmt.Errorf("from invalid: %w", err)
	}

	ancestor, err := from.IsAncestor(to)
	if err != nil {
		return nil, fmt.Errorf("is ancestor check: %w", err)
	}

	if !ancestor {
		return nil, fmt.Errorf("%s is not ancestor of %s", toHash, fromHash)
	}

	// Commits reachable from the "to" hash
	cIter, err := r.Log(&git.LogOptions{From: to.Hash})
	if err != nil {
		return nil, fmt.Errorf("log: %w", err)
	}

	commits := make([]Commit, 0, 16)
	err = cIter.ForEach(func(c *object.Commit) error {
		commits = append(commits, parseCommit(c))

		if c.Hash.String() == fromHash {
			return storer.ErrStop
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return commits, nil
}

func parseCommit(commit *object.Commit) Commit {
	// Use scanner so we handle both cr and crlf
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(commit.Message))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	hash := commit.Hash.String()
	c := Commit{
		Hash:      hash,
		ShortHash: hash[:7],
	}

	for i, line := range lines {
		if i == 0 {
			c.Message = line
			c.Conventional = parseMessage(line)
		}

		if strings.HasPrefix(line, ReleaseNoteTag) {
			c.ReleaseNote = strings.TrimSpace(strings.Replace(line, ReleaseNoteTag, "", 1))
		}
	}
	return c
}

func parseMessage(message string) ConventionalCommit {
	vals := convCommitMatcher.FindAllStringSubmatch(message, 3)

	if len(vals) == 0 {
		return ConventionalCommit{}
	}

	return ConventionalCommit{
		Type:      CategoryOf(vals[0][1]),
		Component: vals[0][2],
		Message:   vals[0][3],
	}
}

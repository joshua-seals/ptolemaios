// We need to think about this
// Helx-Apps really breaks an ideally managed rest-api.
// Instead of relying on database to persist list of applications
// which support crud functionality, here we would either
// embed or reach out to a github repository for the information.
// Benefits - Allows for dynamic loading into ptolemaios
// however the downside is
package core

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Clone the helx-apps branch from url and branchName
// provided. This will represent the applications
// available to the user in appstore.
func CloneBranch(url, branchName string) error {
	// Clone only the specified branch with a depth of 1
	_, err := git.PlainClone("./helx-apps", false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.NewBranchReferenceName(branchName),
		SingleBranch:  true,
		Depth:         1,
	})
	if err != nil {
		return err
	}

	return nil
}

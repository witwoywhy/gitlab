package gitlab

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

type CreateMergeRequest struct {
	Repository   Repository
	Title        string
	SourceBranch string
	TargetBranch string
	AssigneeID   int
	AutoTitle    bool
}

func CreateMerge(req *CreateMergeRequest, git *gitlab.Client) (int, error) {
	if req.AutoTitle {
		log.Printf("%s, get latest commit from branch \"%s\"\n", req.Repository.Name, req.SourceBranch)

		commits, _, err := git.Commits.ListCommits(
			req.Repository.Id,
			&gitlab.ListCommitsOptions{
				RefName: &req.SourceBranch,
			},
		)
		if err != nil {
			return 0, err
		}

		commit := commits[0]
		req.Title = commit.Title
		log.Printf("%s, branch \"%s\" latest commit: %s", req.Repository.Name, req.SourceBranch, req.Title)
	}

	log.Printf("%s, create merge request, %s => %s, title: %s\n", req.Repository.Name, req.SourceBranch, req.TargetBranch, req.Title)
	mergeRequest, _, err := git.MergeRequests.CreateMergeRequest(
		req.Repository.Id,
		&gitlab.CreateMergeRequestOptions{
			Title:           &req.Title,
			SourceBranch:    &req.SourceBranch,
			TargetBranch:    &req.TargetBranch,
			AssigneeID:      &req.AssigneeID,
			TargetProjectID: &req.Repository.Id,
		},
	)
	if err != nil {
		return 0, err
	}

	log.Printf("%s, create merge request, %s => %s, id: %d\n", req.Repository.Name, req.SourceBranch, req.TargetBranch, mergeRequest.IID)
	return mergeRequest.IID, nil
}

type MergeRequest struct {
	Repository     Repository
	MergeRequestId int
}

func Merge(req *MergeRequest, git *gitlab.Client) error {
	log.Printf("%s, merge id: %d\n", req.Repository.Name, req.MergeRequestId)

	_, _, err := git.MergeRequests.AcceptMergeRequest(req.Repository.Id, req.MergeRequestId, nil)
	return err
}

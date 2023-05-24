package gitlab

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/xanzy/go-gitlab"
)

type NewTagFunc func(string, string) string

type CreateTagRequest struct {
	Repository Repository
	Name       string
	Branch     string
	// default prefix is x.x. => 4.6.
	AutoNamePrefix string
	AutoName       bool
	AutoNameFunc   NewTagFunc
}

func CreateTag(req *CreateTagRequest, git *gitlab.Client) error {
	if IsEmptyString(req.Branch) {
		return errors.New("create tag, invalid branch")
	}

	if req.AutoName {
		log.Printf("%s, get latest tag\n", req.Repository.Name)

		tags, _, err := git.Tags.ListTags(req.Repository.Id, &gitlab.ListTagsOptions{
			Sort: ToPointer("desc"),
		})
		if err != nil {
			return err
		}

		if len(tags) == 0 {
			req.Name = req.AutoNameFunc(req.AutoNamePrefix, "")
			log.Printf("%s, initial tag: %s\n", req.Repository.Name, req.Name)
		} else {
			if req.AutoNameFunc == nil {
				req.AutoNameFunc = NewTag
			}

			tag := tags[0]
			log.Printf("%s, latest tag: %s\n", req.Repository.Name, tag.Name)
			
			req.Name = req.AutoNameFunc("", tag.Name)
		}
	}

	if IsEmptyString(req.Name) {
		return errors.New("create tag, invalid name of tag")
	}

	log.Printf("%s, create tag: %s\n", req.Repository.Name, req.Name)
	_, _, err := git.Tags.CreateTag(
		req.Repository.Id,
		&gitlab.CreateTagOptions{
			TagName: &req.Name,
			Ref:     &req.Branch,
		})
	return err
}

// for x.x.x
func NewTag(prefix, s string) string {
	if s != "" {
		sp := strings.Split(s, ".")
		latest := sp[2]
		nextTag, _ := strconv.Atoi(latest)
		return fmt.Sprintf("%s.%s.%d", sp[0], sp[1], nextTag+1)
	}

	if prefix == "" {
		prefix = "4.6."
	}

	return prefix + "1"

}

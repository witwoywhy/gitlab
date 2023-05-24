package gitlab

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func NewClient(u *User) (*gitlab.Client, error) {
	log.Printf("new client for: %s\n", u.Name)
	return gitlab.NewClient(u.Token, gitlab.WithBaseURL(u.Url))
}

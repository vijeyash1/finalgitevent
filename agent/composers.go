package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/webhooks/v6/bitbucket"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/google/uuid"
	"github.com/vijeyash1/gitevent/models"
)

var gitdatas models.Gitevent

func gitComposer(release interface{}, event string) *models.Gitevent {
	uuid := uuid.New()
	switch v := release.(type) {

	case github.PushPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.HTMLURL
		gitdatas.Event = event
		gitdatas.Eventid = v.Commits[0].ID
		gitdatas.Authorname = v.Pusher.Name
		gitdatas.Authormail = v.Pusher.Email
		gitdatas.DoneAt = v.HeadCommit.Timestamp
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Repository.DefaultBranch
		addedFilesSlice := v.Commits[0].Added
		addedFilesString := getStats(&addedFilesSlice)
		if addedFilesString == "" {
			gitdatas.Addedfiles = "---"
		} else {
			gitdatas.Addedfiles = addedFilesString
		}
		modifiedFilesSlice := v.Commits[0].Modified
		modifiedFilesString := getStats(&modifiedFilesSlice)
		if modifiedFilesString == "" {
			gitdatas.Modifiedfiles = "---"
		} else {
			gitdatas.Modifiedfiles = modifiedFilesString
		}
		removedFilesSlice := v.Commits[0].Removed
		removedFilesString := getStats(&removedFilesSlice)
		if removedFilesString == "" {
			gitdatas.Removedfiles = "---"
		} else {
			gitdatas.Removedfiles = removedFilesString
		}
		gitdatas.Message = v.Commits[0].Message
	case github.ForkPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.HTMLURL
		gitdatas.Event = event
		gitdatas.Eventid = strconv.Itoa(int(v.Forkee.ID))
		gitdatas.Authorname = v.Forkee.FullName
		gitdatas.Authormail = "---"
		gitdatas.DoneAt = fmt.Sprintf("%v", v.Forkee.CreatedAt)
		gitdatas.Branch = v.Repository.DefaultBranch
		gitdatas.Addedfiles = "---"
		gitdatas.Modifiedfiles = "---"
		gitdatas.Removedfiles = "---"
		gitdatas.Message = "---"
	case github.PullRequestPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.HTMLURL
		gitdatas.Event = event
		gitdatas.Eventid = strconv.Itoa(int(v.PullRequest.ID))
		gitdatas.Authorname = v.PullRequest.User.Login
		gitdatas.Authormail = "---"
		gitdatas.DoneAt = fmt.Sprintf("%v", v.PullRequest.CreatedAt)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Repository.DefaultBranch
		addedFilesSlice := strconv.Itoa(int(v.PullRequest.Additions))
		addedFilesString := addedFilesSlice
		if addedFilesString == "" {
			gitdatas.Addedfiles = "---"
		} else {
			gitdatas.Addedfiles = addedFilesString
		}
		modifiedFilesSlice := strconv.Itoa(int(v.PullRequest.ChangedFiles))
		modifiedFilesString := modifiedFilesSlice
		if modifiedFilesString == "" {
			gitdatas.Modifiedfiles = "---"
		} else {
			gitdatas.Modifiedfiles = modifiedFilesString
		}
		removedFilesSlice := strconv.Itoa(int(v.PullRequest.Deletions))
		removedFilesString := removedFilesSlice
		if removedFilesString == "" {
			gitdatas.Removedfiles = "---"
		} else {
			gitdatas.Removedfiles = removedFilesString
		}
		gitdatas.Message = v.PullRequest.Title

	case gitlab.PushEventPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Project.WebURL
		gitdatas.Event = event
		gitdatas.Eventid = v.Commits[0].ID
		gitdatas.Authorname = v.Commits[0].Author.Name
		gitdatas.Authormail = v.Commits[0].Author.Email
		gitdatas.DoneAt = fmt.Sprintf("%v", v.Commits[0].Timestamp)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Project.DefaultBranch
		addedFilesSlice := v.Commits[0].Added
		addedFilesString := getStats(&addedFilesSlice)
		if addedFilesString == "" {
			gitdatas.Addedfiles = "---"
		} else {
			gitdatas.Addedfiles = addedFilesString
		}
		modifiedFilesSlice := v.Commits[0].Modified
		modifiedFilesString := getStats(&modifiedFilesSlice)
		if modifiedFilesString == "" {
			gitdatas.Modifiedfiles = "---"
		} else {
			gitdatas.Modifiedfiles = modifiedFilesString
		}
		removedFilesSlice := v.Commits[0].Removed
		removedFilesString := getStats(&removedFilesSlice)
		if removedFilesString == "" {
			gitdatas.Removedfiles = "---"
		} else {
			gitdatas.Removedfiles = removedFilesString
		}
		gitdatas.Message = v.Commits[0].Message
	case gitlab.MergeRequestEventPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Project.URL
		gitdatas.Event = event
		gitdatas.Eventid = strconv.Itoa(int(v.ObjectAttributes.ID))
		gitdatas.Authorname = v.ObjectAttributes.LastCommit.Author.Name
		gitdatas.Authormail = v.ObjectAttributes.LastCommit.Author.Email
		gitdatas.DoneAt = fmt.Sprintf("%v", v.ObjectAttributes.CreatedAt)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Project.DefaultBranch
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		if addedFilesString == "" {
			gitdatas.Addedfiles = "---"
		} else {
			gitdatas.Addedfiles = addedFilesString
		}
		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		if modifiedFilesString == "" {
			gitdatas.Modifiedfiles = "---"
		} else {
			gitdatas.Modifiedfiles = modifiedFilesString
		}
		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		if removedFilesString == "" {
			gitdatas.Removedfiles = "---"
		} else {
			gitdatas.Removedfiles = removedFilesString
		}
		gitdatas.Message = v.ObjectAttributes.LastCommit.Message

	case bitbucket.RepoPushPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Push.Changes[0].New.Links.HTML.Href
		gitdatas.Event = event
		gitdatas.Eventid = v.Push.Changes[0].New.Target.Hash
		gitdatas.Authorname = v.Push.Changes[0].New.Target.Author.DisplayName
		gitdatas.Authormail = "---"
		gitdatas.DoneAt = fmt.Sprintf("%v", v.Push.Changes[0].New.Target.Date)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Push.Changes[0].New.Name
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		if addedFilesString == "" {
			gitdatas.Addedfiles = "---"
		} else {
			gitdatas.Addedfiles = addedFilesString
		}
		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		if modifiedFilesString == "" {
			gitdatas.Modifiedfiles = "---"
		} else {
			gitdatas.Modifiedfiles = modifiedFilesString
		}
		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		if removedFilesString == "" {
			gitdatas.Removedfiles = "---"
		} else {
			gitdatas.Removedfiles = removedFilesString
		}
		gitdatas.Message = v.Push.Changes[0].New.Target.Message

	}
	return &gitdatas
}

//getStats builds a string from the given slice
func getStats(stat *[]string) string {
	var sb strings.Builder
	for _, comm := range *stat {
		sb.WriteString(comm)
		sb.WriteString(",")
	}
	return sb.String()
}

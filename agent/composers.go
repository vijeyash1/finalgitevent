package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/webhooks/v6/bitbucket"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/google/uuid"
	"github.com/vijeyash1/gitevent/models"
)

func githubComposer(release github.PushPayload, event string) *models.Gitevent {
	uuid := uuid.New()
	gitdatas.Uuid = uuid
	gitdatas.Url = release.Repository.HTMLURL
	gitdatas.Event = event
	gitdatas.Eventid = release.Commits[0].ID
	gitdatas.Authorname = release.Pusher.Name
	gitdatas.Authormail = release.Pusher.Email
	gitdatas.DoneAt = (time.Unix(release.Repository.PushedAt, 0).Format("02.01.2022 15:04:05"))
	gitdatas.Repository = release.Repository.Name
	gitdatas.Branch = release.Repository.DefaultBranch
	addedFilesSlice := release.Commits[0].Added
	addedFilesString := getStats(&addedFilesSlice)
	gitdatas.Addedfiles = addedFilesString
	modifiedFilesSlice := release.Commits[0].Modified
	modifiedFilesString := getStats(&modifiedFilesSlice)
	gitdatas.Modifiedfiles = modifiedFilesString
	removedFilesSlice := release.Commits[0].Removed
	removedFilesString := getStats(&removedFilesSlice)
	gitdatas.Modifiedfiles = removedFilesString
	gitdatas.Message = release.Commits[0].Message

	return &gitdatas
}

func gitlabComposer(release gitlab.PushEventPayload, event string) *models.Gitevent {
	uuid := uuid.New()
	gitdatas.Uuid = uuid
	gitdatas.Url = release.Project.WebURL
	gitdatas.Event = event
	gitdatas.Eventid = release.Commits[0].ID
	gitdatas.Authorname = release.Commits[0].Author.Name
	gitdatas.Authormail = release.Commits[0].Author.Email
	gitdatas.DoneAt = fmt.Sprintf("%v", release.Commits[0].Timestamp)
	gitdatas.Repository = release.Repository.Name
	gitdatas.Branch = release.Project.DefaultBranch
	addedFilesSlice := release.Commits[0].Added
	addedFilesString := getStats(&addedFilesSlice)
	gitdatas.Addedfiles = addedFilesString
	modifiedFilesSlice := release.Commits[0].Modified
	modifiedFilesString := getStats(&modifiedFilesSlice)
	gitdatas.Modifiedfiles = modifiedFilesString
	removedFilesSlice := release.Commits[0].Removed
	removedFilesString := getStats(&removedFilesSlice)
	gitdatas.Modifiedfiles = removedFilesString
	gitdatas.Message = release.Commits[0].Message

	return &gitdatas
}

func bitbucketComposer(release bitbucket.RepoPushPayload, event string) *models.Gitevent {
	uuid := uuid.New()
	gitdatas.Uuid = uuid
	gitdatas.Event = event
	gitdatas.Eventid = release.Push.Changes[0].New.Target.Hash
	gitdatas.Authorname = release.Push.Changes[0].New.Target.Author.DisplayName
	gitdatas.DoneAt = fmt.Sprintf("%v", release.Push.Changes[0].New.Target.Date)
	gitdatas.Repository = release.Repository.Name
	gitdatas.Branch = release.Push.Changes[0].New.Name
	gitdatas.Message = release.Push.Changes[0].New.Target.Message

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

package main

import (
	"log"

	"net/http"

	"github.com/go-playground/webhooks/v6/bitbucket"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/go-playground/webhooks/v6/gitlab"
)

func (app *application) githubHandler(w http.ResponseWriter, r *http.Request) {

	hook, _ := github.New()
	payload, err := hook.Parse(r, github.PushEvent, github.ForkEvent, github.PullRequestEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}

	}

	switch value := payload.(type) {
	case github.PushPayload:
		release := value
		composed := gitComposer(release, "PushEvent")
		app.publish.JS.GitPublish(composed)
	case github.ForkPayload:
		release := value
		composed := gitComposer(release, "ForkEvent")
		app.publish.JS.GitPublish(composed)
	case github.PullRequestPayload:
		release := value
		composed := gitComposer(release, "PullRequest")
		app.publish.JS.GitPublish(composed)

	}

}

func (app *application) gitlabHandler(w http.ResponseWriter, r *http.Request) {

	hook, _ := gitlab.New()
	payload, err := hook.Parse(r, gitlab.PushEvents, gitlab.MergeRequestEvents)
	if err != nil {
		if err == gitlab.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {

	case gitlab.PushEventPayload:
		release := value
		composed := gitComposer(release, "PushEvent")
		app.publish.JS.GitPublish(composed)
	case gitlab.MergeRequest:
		release := value
		composed := gitComposer(release, "MergeRequest")
		app.publish.JS.GitPublish(composed)
	}
}

func (app *application) bitBucketHandler(w http.ResponseWriter, r *http.Request) {

	hook, _ := bitbucket.New()
	payload, err := hook.Parse(r, bitbucket.RepoPushEvent, bitbucket.RepoForkEvent, bitbucket.PullRequestCreatedEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {

	case bitbucket.RepoPushPayload:
		release := value
		composed := gitComposer(release, "PushEvent")
		app.publish.JS.GitPublish(composed)
	case bitbucket.RepoForkPayload:
		release := value
		composed := gitComposer(release, "ForkEvent")
		app.publish.JS.GitPublish(composed)
	case bitbucket.PullRequestCreatedPayload:
		release := value
		composed := gitComposer(release, "PullRequest")
		app.publish.JS.GitPublish(composed)

	}
}

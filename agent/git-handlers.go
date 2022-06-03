package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/go-playground/webhooks/v6/bitbucket"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/vijeyash1/gitevent/models"
)

var gitdatas models.Gitevent

func (app *application) githubHandler(w http.ResponseWriter, r *http.Request) {

	hook, _ := github.New()
	payload, err := hook.Parse(r, github.PushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error Event not found")
		}

	}

	switch value := payload.(type) {
	case github.PushPayload:
		release := value
		composed := githubComposer(release, "PushEvent")
		app.publish.JS.GitPublish(composed)
	}

}

func (app *application) gitlabHandler(w http.ResponseWriter, r *http.Request) {

	hook, _ := gitlab.New()
	payload, err := hook.Parse(r, gitlab.PushEvents)
	if err != nil {
		if err == gitlab.ErrEventNotFound {
			log.Print("Error Event not found")
		}
	}

	switch value := payload.(type) {

	case gitlab.PushEventPayload:
		release := value
		composed := gitlabComposer(release, "PushEvent")
		app.publish.JS.GitPublish(composed)
	}
}

func (app *application) bitBucketHandler(w http.ResponseWriter, r *http.Request) {

	hook, _ := bitbucket.New()
	payload, err := hook.Parse(r, bitbucket.RepoPushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error Event not found")
		}
	}

	switch value := payload.(type) {

	case bitbucket.RepoPushPayload:
		release := value
		fmt.Printf("url url %s\n", release.Repository.Website)
		composed := bitbucketComposer(release, "PushEvent")
		app.publish.JS.GitPublish(composed)

	}
}

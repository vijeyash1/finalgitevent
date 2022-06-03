package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/go-playground/webhooks/v6/bitbucket"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/go-playground/webhooks/v6/gitlab"
)

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
		// Do whatever you want from here...
		fmt.Printf("%s\n", release.Repository.Name)

		app.publish.JS.GitPublish()

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
		fmt.Printf("%s\n", release.Repository.Name)

		app.publish.JS.GitPublish()

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

		app.publish.JS.GitPublish()

	}
}

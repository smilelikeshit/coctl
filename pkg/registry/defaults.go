// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package registry

import (
	"embed"

	"github.com/AlecAivazis/survey/v2"
)

// CLIVersion version of Create Go App CLI.
const CLIVersion string = "0.0.1"

// Variables struct for Ansible variables (inventory, hosts).
type Variables struct {
	List map[string]interface{}
}

// CreateAnswers struct for a survey's answers for `create` command.
type CreateAnswers struct {
	Frontend      string
	AgreeCreation bool `survey:"agree"`
}

var (
	// EmbedMiscFiles misc files and configs.
	//go:embed misc/*
	EmbedMiscFiles embed.FS

	// EmbedRoles Ansible roles.
	//go:embed roles/*
	EmbedRoles embed.FS

	// EmbedTemplates template files.
	//go:embed templates/*
	EmbedTemplates embed.FS

	// CreateQuestions survey's questions for `create` command.
	CreateQuestions = []*survey.Question{
		{
			Name: "frontend",
			Prompt: &survey.Select{
				Message: "Choose a frontend framework/library:",
				Help:    "Option with a `*-ts` tail will create a TypeScript template.",
				Options: []string{
					"none",
					"react",
					"react-ts",
					"vue",
					"vue-ts",
				},
				Default:  "none",
				PageSize: 13,
			},
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}

	// CustomCreateQuestions survey's questions for `create -c` command.
	CustomCreateQuestions = []*survey.Question{
		{
			Name: "frontend",
			Prompt: &survey.Input{
				Message: "Enter URL to the custom frontend repository:",
				Help:    "No need to specify `http://` or `https://` protocol.",
				Default: "none",
			},
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}
)

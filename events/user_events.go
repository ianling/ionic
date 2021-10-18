package events

import (
	"encoding/json"
)

type (
	// UserEvent represents the user related segment of an Event within Ion Channel.
	// Action is the specific type of User Event that occurred,
	// Data is information relevant to the type of event, defined by one of the other structs in this file.
	UserEvent struct {
		Action string          `json:"action"`
		Data   json.RawMessage `json:"data"`
	}

	// ProjectFlippedData represents the Data portion of a ProjectFlipped event
	ProjectFlippedData struct {
		Projects []struct {
			ID  string
			URL string
		}
		Email string
	}

	// InviteDetails represents the Data portion of several events related to a user being invited
	InviteDetails struct {
		Email		string
		AcceptLink  string
		UserName    string
		AccountName string
	}

	// AccountCreatedData represents the Data portion of an AccountCreated event
	AccountCreatedData struct {
		InviteDetails
	}

	// UserSignupData represents the Data portion of an UserSignup event
	UserSignupData struct {
		InviteDetails
	}

	// UserSignupStartedData represents the Data portion of an UserSignupStarted event
	UserSignupStartedData struct {
		InviteDetails
	}
)

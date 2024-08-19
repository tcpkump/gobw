package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type listKeyMap struct {
	User key.Binding
	Pass key.Binding
}

var listKeys = listKeyMap{
	User: key.NewBinding(
		key.WithKeys("y"),
		key.WithHelp("y", "Copy Username"),
	),
	Pass: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "Copy Password"),
	),
}

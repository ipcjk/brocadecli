package main

import "io"

type RouterInt interface {
	CloseConnection()
	ConfigureTerminalMode() error
	ConnectPrivilegedMode() error
	GetPromptMode() error
	PasteConfiguration(io.Reader) error
	RunCommands(io.Reader) error
	SkipPageDisplayMode() (string, error)
	WriteConfiguration() error
}

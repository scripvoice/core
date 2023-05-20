package infrastructure

import (
	"context"
	"sync"
)

// Command defines the interface for commands.
type Command interface {
	GetName() string
}

// CommandHandler defines the interface for command handlers.
type CommandHandler interface {
	HandleCommand(command Command, context context.Context)
}

// CommandFactory represents the factory class that resolves the command handler based on the command type.
type CommandFactory struct {
	commandHandlers map[string]CommandHandler
}

// NewCommandFactory creates a new instance of the CommandFactory.
func NewCommandFactory() *CommandFactory {
	return &CommandFactory{
		commandHandlers: make(map[string]CommandHandler),
	}
}

var (
	commandfactoryinstance *CommandFactory
	oncecommandfactory     sync.Once
)

// GetInstance returns the singleton instance
func GetCommandFactoryInstance() *CommandFactory {
	oncecommandfactory.Do(func() {
		commandfactoryinstance = NewCommandFactory() // Create the singleton instance
	})
	return commandfactoryinstance
}

// RegisterCommandHandler registers an command handler for a specific command type.
func (factory *CommandFactory) RegisterCommandHandler(commandType string, handler CommandHandler) {
	factory.commandHandlers[commandType] = handler
}

// ResolveCommandHandler resolves the command handler based on the command type.
func (factory *CommandFactory) ResolveCommandHandler(commandType string) CommandHandler {
	handler, ok := factory.commandHandlers[commandType]
	if ok {
		return handler
	}
	return nil
}

type CommandRegistrator interface {
	Register(factory *CommandFactory)
}

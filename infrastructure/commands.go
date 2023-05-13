package infrastructure

// Command defines the interface for commands.
type Command interface {
	GetName() string
}

// CommandHandler defines the interface for command handlers.
type CommandHandler interface {
	HandleCommand(command Command)
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

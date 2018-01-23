package protocol

// Command enumerates protocol commands
type Command uint16

const (
	// CmdJoin for joining the network
	CmdJoin = Command(iota)

	// CmdAddresses for retrieving network active nodes
	CmdAddresses
)

// Message is the message structure used by sockets
type Message struct {
	Command Command
	Payload interface{}
}

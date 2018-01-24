package director

// Node represents a server hosting Director
type Node interface {
	IncrementMatchIndex() int64
	IncrementNextIndex() int64
	GetMatchIndex() int64
	GetNextIndex() int64
	GetClient() Client
}

type node struct {
	client     Client
	nextIndex  int64 // next index to send to node
	matchIndex int64 // index of highest index replicated on server
}

func (n *node) IncrementMatchIndex() int64 {
	n.matchIndex++
	return n.matchIndex
}

func (n *node) IncrementNextIndex() int64 {
	n.nextIndex++
	return n.nextIndex
}

func (n *node) GetMatchIndex() int64 {
	return n.matchIndex
}

func (n *node) GetNextIndex() int64 {
	return n.nextIndex
}

func (n *node) GetClient() Client {
	return n.client
}

// NewNode creates a new node from a given host
func NewNode(host string) Node {
	return &node{
		client:     NewGRPCClient(host),
		nextIndex:  0,
		matchIndex: 0,
	}
}

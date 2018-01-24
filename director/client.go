package director

import (
	"context"
	"log"

	a "github.com/dqian96/alexandria/archive"
	"github.com/dqian96/alexandria/utils"
	"google.golang.org/grpc"
)

// Client allows the Director to make calls to remote Directors
type Client interface {
	AppendEntriesRPC(string, uint64, uint64, []*a.Entry, *a.Entry, uint64, uint64) (bool, uint64, error)
	RequestVoteRPC(string, uint64, uint64, uint64) (bool, uint64, error)
	GetHost() string
	Close()
}

type grpcClient struct {
	conn   *grpc.ClientConn
	client DirectorClient
	host   string
}

const (
	retryTimeoutSeconds = 30
	retrySleepSeconds   = 5
)

// NewGRPCClient creates a new Client, with gRPC under the hood
func NewGRPCClient(host string) Client {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", host, err)
	}
	return &grpcClient{
		conn:   conn,
		client: NewDirectorClient(conn),
		host:   host,
	}
}

// AppendEntriesRPC is used by the leader to append (log) entries to a follower
// It also checks and maintains consistency between the nodes and can be used as the "heartbeat"
func (c *grpcClient) AppendEntriesRPC(leaderID string, commitIndex uint64, term uint64, entries []*a.Entry,
	lastEntry *a.Entry, lastIndex uint64, lastTerm uint64) (bool, uint64, error) {
	msgEntries := make([]*Entry, len(entries), len(entries))
	for i, entry := range entries {
		msgEntries[i] = convertToPbEntry(entry)
	}
	req := &AppendEntriesRequest{
		LeaderId:    leaderID,
		CommitIndex: commitIndex,
		Term:        term,
		Entries:     msgEntries,
		LastEntry:   convertToPbEntry(lastEntry),
		LastIndex:   lastIndex,
		LastTerm:    lastTerm,
	}
	var res *AppendEntriesReply
	var err error
	connErr := utils.Retry(retryTimeoutSeconds, retrySleepSeconds, func() error {
		res, err = c.client.AppendEntries(context.Background(), req)
		return err
	})
	if connErr != nil {
		return false, 0, &ConnectionError{
			Host:           c.host,
			TimeoutSeconds: retryTimeoutSeconds,
			SourceError:    connErr,
		}
	}
	return res.Success, res.Term, nil
}

// RequestVoteRPC is used by a candidate to request votes from followers in leader election
func (c *grpcClient) RequestVoteRPC(candidateID string, term uint64, lastTerm uint64, lastIndex uint64) (bool, uint64, error) {
	req := &RequestVoteRequest{
		CandidateId: candidateID,
		Term:        term,
		LastTerm:    lastTerm,
		LastIndex:   lastIndex,
	}
	var res *RequestVoteReply
	var err error
	connErr := utils.Retry(retryTimeoutSeconds, retrySleepSeconds, func() error {
		res, err = c.client.RequestVote(context.Background(), req)
		return err
	})
	if connErr != nil {
		return false, 0, &ConnectionError{
			Host:           c.host,
			TimeoutSeconds: retryTimeoutSeconds,
			SourceError:    connErr,
		}
	}
	return res.VoteGranted, res.Term, nil
}

// GetHost gets the host the client is connected to
func (c *grpcClient) GetHost() string {
	return c.host
}

// Close closes the connection
func (c *grpcClient) Close() {
	c.conn.Close()
}

func convertToPbEntry(entry *a.Entry) *Entry {
	return &Entry{
		Key:   entry.Key,
		Value: entry.Value,
	}
}

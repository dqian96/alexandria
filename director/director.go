package director

import (
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	a "github.com/dqian96/alexandria/archive"
)

const (
	// "leader" is the leader/master of the cluster; responsible for updating state and log replication
	leader = 1
	// "candidate" is a FOLLOWER that is trying to become a LEADER in an election following a timeout period
	candidate = 2
	// "follower" is a follower/slave of the cluster; responsible for servicing (majority of) reads
	// hosting redundancy, and monitoring the LEADER
	follower = 3

	// Amount of time to sleep when busy waiting
	busyWaitSleepTimeMillis = 500
)

// Director is responsible for handling consensus between nodes
// It is a gRPC server that listens to other directors, makes calls to other Directors,
// and receives local calls from the HTTP server to direct the Archive
type Director interface {
	Get(string, string) (string, bool, error)
	Put(string, a.Entry) error
	Delete(string, string) (bool, error)
	Serve(string) error
	GetLeader() string
}

type director struct {
	log                    commitLog
	currentTerm            uint64
	votedFor               int
	followers              []Node
	leader                 Node
	state                  int
	electionTimeoutSeconds int
	archive                a.Archive
	commitIndex            int64
	lastAppliedIndex       int64
	s                      *grpc.Server
	mutex                  sync.RWMutex
}

// Get returns a value from the archive given the key if the current node is a leader;
// otherwise it returns an error wrapping the leader's host
func (d *director) Get(reqID string, key string) (string, bool, error) {
	if d.state != leader {
		log.Printf("[%s] Only the leader can propose commands", reqID)
		// TODO dead/missing leader
		return "", false, &NonLeaderCmdError{leader: d.leader.GetClient().GetHost()}
	}
	value, in := d.archive.Get(key)
	if !in {
		return "", false, nil
	}
	return value, true, nil
}

func (d *director) Delete(reqID string, key string) (bool, error) {
	return true, nil
	// TODO some kind of error
}

func (d *director) Put(reqID string, entry a.Entry) error {
	if d.state != leader {
		log.Printf("[%s] Only the LEADER can propose commands", reqID)
		// TODO dead/missing leader
		return &NonLeaderCmdError{leader: d.leader.GetClient().GetHost()}
	}
	err := d.archive.ValidateEntry(entry)
	if err != nil {
		return err
	}
	var timesAppended uint64
	for _, n := range d.followers {
		go func() {
			// n.GetClient().AppendEntriesRPC(d.leader.GetClient().GetHost(),
			// d.commitIndex,
			// d.currentTerm,
			// entries,
			// lastEntry,
			// lastIndex,
			// lasterTerm
			// )
			n = n // TODO create RPC
			atomic.AddUint64(&timesAppended, 1)
		}()
	}
	for timesAppended <= uint64(len(d.followers)/2) {
		// wait
		time.Sleep(busyWaitSleepTimeMillis)
	}
	d.commitIndex++      // commit
	d.archive.Put(entry) // apply
	d.lastAppliedIndex++
	// let heartbeat commit to rest
	return nil
	// TODO some kind of error
}

func (d *director) GetLeader() string {
	return d.leader.GetClient().GetHost()
}

// TODO: special error type
func (d *director) AppendEntries(ctx context.Context, in *AppendEntriesRequest) (*AppendEntriesReply, error) {
	// append to log
	// TODO some kind of error
	return &AppendEntriesReply{}, nil
}

func (d *director) RequestVote(ctx context.Context, in *RequestVoteRequest) (*RequestVoteReply, error) {
	// request vote
	return &RequestVoteReply{}, nil
	// TODO some kind of error
}

func NewDirector(followers []string, timeoutSeconds uint64, leader string, archive a.Archive) Director {
	s := grpc.NewServer()
	d := &director{
		s: s,
	} // TODO set
	RegisterDirectorServer(s, d)
	reflection.Register(s)
	return d
}

func (d *director) Serve(port string) error {
	log.Printf("Server starting on port %s", port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}
	// go heartbeat listener ad election
	if err := d.s.Serve(lis); err != nil {
		log.Printf("director server failed: %v", err)
		return err
	}
	return nil
}

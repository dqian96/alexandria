# Design

Alexandria stores key-value pairs in a Go map.

Although Alexandria can be run as a single instance (centeralized), it's designed to be distributed using master-slave replication for high availiability and scaling reads.

## Components

* Client:
  * the only client is a CLI written in Go
  * future clients include: Go, Java, and Scala
  * HA/single-instance modes
* Alexandria Archive:
  * server responsible for storage of key-value pairs
  * CRUD API
* Alexandria Director
  * server responsible for "cluster management"
  * monitoring of master and slaves
  * election (automatic failover) in the case the master fails
  * replication
  * adding/removing slaves (i.e. pointing to master or notification)
  * keeps a commit log for consensus and consistency


Note: every node runs the archive and the director

## Architecture

![Alt text](architecture.png?raw=true "Architecture")

* Master
  * source of truth
  * solely responsible for writing to cache
  * replicates data to slaves
* Slave
  * responsible for servicing client reads
  * any slave can be selected for a read
  * high availability
  * monitors master and can trigger election if master fails
  * holds references to peers for gathering votes during election

## Considerations

* all communication is done using gRPC
* currently, the only values that can be stored are strings of a custom size
* multiple eviction policies are available, with LRU being the default
* implementation of [Raft](https://raft.github.io/) for consensus and election

## Areas of Concern

* master is the single point of failure
* master is the bottleneck for writes

## Future features

* sharding; expanding the cache size beyond that of the master
* authorization
* more eviction policies (i.e. LFU)
* faster consistency - slaves could update the state of other slaves? sacrifice availability? (CAP theorem)


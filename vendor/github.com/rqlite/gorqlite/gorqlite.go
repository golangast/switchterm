// Package gorqlite provieds a database/sql-like driver for rqlite,
// the distributed consistent sqlite.
//
// Copyright (c)2016 andrew fabbro (andrew@fabbro.org)
//
// See LICENSE.md for license. tl;dr: MIT. Conveniently, the same license as rqlite.
//
// Project home page: https://github.com/raindo308/gorqlite
//
// Learn more about rqlite at: https://github.com/rqlite/rqlite
package gorqlite

// this file contains package-level stuff:
//   consts
//   init()
//   Open, TraceOn(), TraceOff()

import (
	"crypto/rand"
	"fmt"
	"io"
	"strings"
)

type consistencyLevel int

const (
	// ConsistencyLevelNone provides no consistency to other nodes.
	ConsistencyLevelNone consistencyLevel = iota
	// ConsistencyLevelWeak provides a weak consistency that guarantees the
	// queries are sent to the leader.
	ConsistencyLevelWeak
	// ConsistencyLevelStrong provides a strong consistency and guarantees
	// that queries are sent and received by other nodes.
	ConsistencyLevelStrong
)

// used in several places, actually
var (
	consistencyLevelNames map[consistencyLevel]string
	consistencyLevels     map[string]consistencyLevel
)

type apiOperation int

const (
	api_QUERY apiOperation = iota
	api_STATUS
	api_WRITE
	api_NODES
	api_REQUEST
)

func init() {
	traceOut = io.Discard

	consistencyLevelNames = make(map[consistencyLevel]string)
	consistencyLevelNames[ConsistencyLevelNone] = "none"
	consistencyLevelNames[ConsistencyLevelWeak] = "weak"
	consistencyLevelNames[ConsistencyLevelStrong] = "strong"

	consistencyLevels = make(map[string]consistencyLevel)
	consistencyLevels["none"] = ConsistencyLevelNone
	consistencyLevels["weak"] = ConsistencyLevelWeak
	consistencyLevels["strong"] = ConsistencyLevelStrong
}

// Open creates and returns a "connection" to rqlite.
//
// Since rqlite is stateless, there is no actual connection.
// Open() creates and initializes a gorqlite Connection type,
// which represents various config information.
//
// The URL should be in a form like this:
//
//	http://localhost:4001
//
//	http://     default, no auth, localhost:4001
//	https://    default, no auth, localhost:4001, using https
//
//	http://localhost:1234
//	http://mary:secret2@localhost:1234
//
//	https://mary:secret2@somewhere.example.com:1234
//	https://mary:secret2@somewhere.example.com // will use 4001
func Open(connURL string) (*Connection, error) {
	var conn = &Connection{}

	// generate our uuid for trace
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return conn, err
	}
	conn.ID = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	trace("%s: Open() called for url: %s", conn.ID, connURL)

	// set defaults
	conn.hasBeenClosed = false

	// parse the URL given
	err = conn.initConnection(connURL)
	if err != nil {
		return conn, err
	}

	if !conn.disableClusterDiscovery {
		// call updateClusterInfo() to re-populate the cluster and discover peers
		// also tests the user's default
		if err := conn.updateClusterInfo(); err != nil {
			return conn, err
		}
	}

	return conn, nil
}

// trace adds a message to the trace output
//
// not a public function.  we (inside) can add - outside they can
// only see.
//
// Call trace as:     Sprintf pattern , args...
//
// This is done so that the more expensive Sprintf() stuff is
// done only if truly needed.  When tracing is off, calls to
// trace() just hit a bool check and return.  If tracing is on,
// then the Sprintf-ing is done at a leisurely pace because, well,
// we're tracing.
//
// Premature optimization is the root of all evil, so this is
// probably sinful behavior.
//
// Don't put a \n in your Sprintf pattern becuase trace() adds one
func trace(pattern string, args ...interface{}) {
	// don't do the probably expensive Sprintf() if not needed
	if !wantsTrace {
		return
	}

	// this could all be made into one long statement but we have
	// compilers to do such things for us. let's sip a mint julep
	// and spell this out in glorious exposition.

	// make sure there is one and only one newline
	nlPattern := strings.TrimSpace(pattern) + "\n"
	msg := fmt.Sprintf(nlPattern, args...)
	traceOut.Write([]byte(msg))
}

// TraceOn turns on tracing output to the io.Writer of your choice.
//
// Trace output is very detailed and verbose, as you might expect.
//
// Normally, you should run with tracing off, as it makes absolutely
// no concession to performance and is intended for debugging/dev use.
func TraceOn(w io.Writer) {
	traceOut = w
	wantsTrace = true
}

// TraceOff turns off tracing output. Once you call TraceOff(), no further
// info is sent to the io.Writer, unless it is TraceOn'd again.
func TraceOff() {
	wantsTrace = false
	traceOut = io.Discard
}

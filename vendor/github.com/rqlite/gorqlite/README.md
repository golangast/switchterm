# gorqlite - a golang client for rqlite
[![Circle CI](https://circleci.com/gh/rqlite/gorqlite/tree/master.svg?style=svg)](https://circleci.com/gh/rqlite/gorqlite/tree/master)

gorqlite is a golang client for rqlite that provides easy-to-use abstractions for working with the rqlite API.

It provides an idiomatic API specialized for rqlite and a database/sql driver (read below for more information on it). The main API provides similar semantics to database/sql, such as `Open()`, `Query()` and `QueryOne()`, `Next()`/`Scan()`/`Map()`, `Write()` and `WriteOne()`, etc.

rqlite is the distributed consistent sqlite database.  [Learn more about rqlite here](https://www.rqlite.io).

## Status

This client library is used in production by various groups, including Replicated. Check out [their blog post](https://www.replicated.com/blog/app-manager-with-rqlite) on their use of rqlite.

## Features

* Abstracts the rqlite http API interaction - the POSTs, JSON handling, etc.  You submit your SQL and get back an iterator with familiar database/sql semantics (`Next()`, `Scan()`, etc.) or a `map[column name as string]interface{}.
* Timings and other metadata (e.g., num rows affected, last insert ID, etc.) is conveniently available and parsed into appropriate types.
* A connection abstraction allows gorqlite to discover and remember the rqlite leader.  gorqlite will automatically try other peers if the leader is lost, enabling fault-tolerant API operations.
* Timeout can be set on a per-Connection basis to accommodate those with far-flung empires.
* Use familiar database URL connection strings to connection, optionally including rqlite authentication and/or specific rqlite consistency levels.
* Only a single node needs to be specified in the connection.  **By default gorqlite will talk to that node and figure out the rest of the cluster from its redirects and status API**. This is known as _Cluster Discovery_.
* Depending on your deployment, **you may wish to disable _Cluster Discovery_**. If you do disable it only the provided URL will be used to communicate with the API instead of discovering the leader and peers and retrying failed requests with different peers. To disable _Cluster Discovery_ add `disableClusterDiscovery=true` as a URL Query Parameter when connecting to rqlite e.g. `http://localhost:14001?disableClusterDiscovery=true`.
  * This is helpful, for example, when using a Kubernetes service to handle the load balancing of the requests across healthy nodes. 
* Support for several rqlite-specific operations:
  * `Leader()` and `Peers()` to examine the cluster.
  * `SetConsistencyLevel()` can be called at any time on a connection to change the consistency level for future operations.
  * `Timing` can be referenced on a per-result basis to retrieve the timings information for executed operations as float64, per the rqlite API. 
* `Trace(io.Writer)`/`Trace(nil)` can be used to turn on/off debugging information on everything gorqlite does to a io.Writer of your choice.
* No external dependencies. Uses only standard library functions.

## Install

`go get github.com/rqlite/gorqlite`

## Examples
```go
// these URLs are just generic database URLs, not rqlite API URLs,
// so you don't need to worry about the various rqlite paths ("/db/query"), etc.
// just supply the base url and not "db" or anything after it.

// yes, you need the http or https

// no, you cannot specify a database name in the URL (this is sqlite, after all).

conn, err := gorqlite.Open("http://") // connects to localhost on 4001 without auth
conn, err := gorqlite.Open("https://") // same but with https
conn, err := gorqlite.Open("https://localhost:4001/") // same only explicitly

// with auth:
conn, err := gorqlite.Open("https://mary:secret2@localhost:4001/")
// different server, setting the rqlite consistency level
conn, err := gorqlite.Open("https://mary:secret2@server1.example.com:4001/?level=none")
// same without auth, setting the rqlite consistency level
conn, err := gorliqte.Open("https://server2.example.com:4001/?level=weak")
// different port, setting the rqlite consistency level and timeout
conn, err := gorqlite.Open("https://localhost:2265/?level=strong&timeout=30")
// different port, disabling cluster discovery in the client
conn, err := gorqlite.Open("https://localhost:2265/?disableClusterDiscovery=true")

// change our minds
conn.SetConsistencyLevel("none")
conn.SetConsistencyLevel("weak")
conn.SetConsistencyLevel("strong")

// simulate database/sql Prepare()
statements := make ([]string,0)
pattern := "INSERT INTO secret_agents(id, hero_name, abbrev) VALUES (%d, '%s', '%3s')"
statements = append(statements,fmt.Sprintf(pattern,125718,"Speed Gibson","Speed"))
statements = append(statements,fmt.Sprintf(pattern,209166,"Clint Barlow","Clint"))
statements = append(statements,fmt.Sprintf(pattern,44107,"Barney Dunlap","Barney"))
results, err := conn.Write(statements)

// now we have an array of []WriteResult 

for n, v := range WriteResult {
	fmt.Printf("for result %d, %d rows were affected\n",n,v.RowsAffected)
	if ( v.Err != nil ) {
		fmt.Printf("   we have this error: %s\n",v.Err.Error())
	}
}

// or if we have an auto_increment column
res, err := conn.WriteOne("INSERT INTO foo (name) values ('bar')")
fmt.Printf("last insert id was %d\n",res.LastInsertID)

// just like database/sql, you're required to Next() before any Scan() or Map()

// note that rqlite is only going to send JSON types - see the encoding/json docs
// which means all numbers are float64s.  gorqlite will convert to int64s for you
// because it is convenient but other formats you will have to handle yourself

var id int64
var name string
rows, err := conn.QueryOne("select id, name from secret_agents where id > 500")
fmt.Printf("query returned %d rows\n",rows.NumRows)
for rows.Next() {
	err := rows.Scan(&id, &name)
	fmt.Printf("this is row number %d\n",rows.RowNumber)
	fmt.Printf("there are %d rows overall%d\n",rows.NumRows)
}

// just like WriteOne()/Write(), QueryOne() takes a single statement,
// while Query() takes a []string.  You'd only use Query() if you wanted
// to transactionally group a bunch of queries, and then you'd get back
// a []QueryResult

// alternatively, use Next()/Map()

for rows.Next() {
	m, err := rows.Map()
	// m is now a map[column name as string]interface{}
	id := m["name"].(float64) // the only json number type
	name := m["name"].(string)
}

// get rqlite cluster information
leader, err := conn.Leader()
// err could be set if the cluster wasn't answering, etc.
fmt.Println("current leader is"leader)
peers, err := conn.Peers()
for n, p := range peers {
	fmt.Printf("cluster peer %d: %s\n",n,p)
}

// turn on debug tracing to the io.Writer of your choice.
// gorqlite will verbosely write very granular debug information.
// this is similar to perl's DBI->Trace() facility.
// note that this is done at the package level, not the connection
// level, so you can debug Open() etc. if need be.

f, err := os.OpenFile("/tmp/deep_insights.log",OS_RDWR|os.O_CREATE|os.O_APPEND,0644)
gorqlite.TraceOn(f)

// change my mind and watch the trace
gorqlite.TraceOn(os.Stderr)

// turn off
gorqlite.TraceOff()

// using parameterized statements
wr, err := conn.WriteParameterized(
	[]gorqlite.ParameterizedStatement{
		{
			Query:     "INSERT INTO secret_agents(id, name, secret) VALUES(?, ?, ?)",
			Arguments: []interface{}{7, "James Bond", "not-a-secret"},
		},
	},
)
seq, err := conn.QueueParameterized(
	[]gorqlite.ParameterizedStatement{
		{
			Query:     "INSERT INTO secret_agents(id, name, secret) VALUES(?, ?, ?)",
			Arguments: []interface{}{7, "James Bond", "not-a-secret"},
		},
	},
)
qr, err := conn.QueryParameterized(
	[]gorqlite.ParameterizedStatement{
		{
			Query:     "SELECT id, name from secret_agents where id > ?",
			Arguments: []interface{}{3},
		},
	},
)

// alternatively
wr, err := conn.WriteOneParameterized(
	gorqlite.ParameterizedStatement{
		Query:     "INSERT INTO secret_agents(id, name, secret) VALUES(?, ?, ?)",
		Arguments: []interface{}{7, "James Bond", "not-a-secret"},
	},
)
seq, err := conn.QueueOneParameterized(
	gorqlite.ParameterizedStatement{
		Query:     "INSERT INTO secret_agents(id, name, secret) VALUES(?, ?, ?)",
		Arguments: []interface{}{7, "James Bond", "not-a-secret"},
	},
)
qr, err := conn.QueryOneParameterized(
	gorqlite.ParameterizedStatement{
		Query:     "SELECT id, name from secret_agents where id > ?",
		Arguments: []interface{}{3},
	},
)

// using nullable types
var name gorqlite.NullString
rows, err := conn.QueryOne("select name from secret_agents where id = 7")
for rows.Next() {
	err := rows.Scan(&name)
}
if name.Valid {
	// use name.String
} else {
	// NULL value
}

```

### Queued Writes
The client does support [Queued Writes](https://github.com/rqlite/rqlite/blob/master/DOC/QUEUED_WRITES.md). Instead of calling the `Write()` functions, call the queueing versions instead.
```go
var seq int64
var err error

seq, err = conn.QueueOne("CREATE TABLE " + testTableName() + " (id integer, name text)")
seq, err = conn.Queue(...)
```

## Important Notes

If you use access control, any user connecting will need the "status" permission in addition to any other needed permission.  This is so gorqlite can query the cluster and try other peers if the master is lost.

rqlite does not support iterative fetching from the DBMS, so your query will put all results into memory immediately.  If you are working with large datasets on small systems, your experience may be sub-optimal.

## Driver for database/sql

It is recommended that you use the main gorqlite-specific API when possible. However, if you need to use gorqlite with database/sql, you can import `github.com/rqlite/gorqlite/stdlib`. For example:

```go
package main

import (
	"database/sql"

	_ "github.com/rqlite/gorqlite/stdlib"
)

func main() {
	db, err := sql.Open("rqlite", "http://")
	if err != nil {
		panic(err)
	}
	_, err := db.Exec("CREATE TABLE users (id INTEGER, name TEXT)")
	if err != nil {
		panic(err)
	}
}
```

The following limitations apply when using the rqlite database/sql driver:

* rqlite supports transactions, but only in a single batch.  You can group many statements into a single transaction, but you must submit them as a single unit.  You cannot start a transaction, send some statements, come back later and submit some more, and then later commit.

* As a consequence, there is no rollback.

* The statement parsing/preparation API is not exposed at the SQL layer by sqlite, and hence it's not exposed by rqlite.  What this means is that there's no way to prepare a statement (`"INSERT INTO superheroes (?,?)"`) and then later bind executions to it.  (In case you're wondering, yes, it would be possible for gorqlite to include a copy of sqlite3 and use its engine, but the sqlite C call to `sqlite3_prepare_v2()` will fail because a local sqlite3 won't know your DB's schemas and the `sqlite3_prepare_v2()` call validates the statement against the schema.  We could open the local sqlite .db file maintained by rqlite and validate against that, but there is no way to make a consistency guarantee between time of preparation and execution, especially since the user can mix DDL and DML in a single transaction).

* Therefore, `Begin()`, `Rollback()`, `Commit()`, and `Prepare()` are all no-ops that return no errors but don't do anything.

## TODO

https has not been tested yet.  In theory, https should work just fine because it's just a URL to gorqlite, but it has not been.

Several features may be added in the future:

- support for the backup API

- support for expvars (debugvars)

- perhaps deleting a node (the remove API)

- since connections are just config info, it should be possible to clone them, which would save startup time for new connections.  This needs to be thread-safe, though, since a connection at any time could be updating its cluster info, etc.

- gorqlite always talks to the master (unless it's searching for a master).  In theory, you talk to a non-master in "none" consistency mode, but this adds a surprising amount of complexity.  gorqlite has to take note of the URL you call it with, then try to match that to the cluster's list to mark it as the "default" URL.  Then whenever it wants to do an operation, it has to carefully sort the peer list based on the consistency model, if the default URL has gone away, etc.  And when cluster info is rebuilt, it has to track the default URL through that.

## Other Design Notes

In standard `database/sql` drivers, `Open()` doesn't actually do anything.  You get a "connection" that doesn't connect until you `Ping()` or send actual work.  In gorqlite's case, it needs to connect to get cluster information, so this is done immediately and automatically open calling `Open()`.  By the time `Open()` is returned, gorqlite has full cluster info.

Unlike `database/sql` connections, a gorqlite connection in the main gorqlite-specific API is not thread-safe. However, a gorqlite database/sql connection through package `stdlib` *is* thread-safe.

`Close()` will set a flag so if you try to use the connection afterwards, it will fail.  But otherwise, you can merrily let your connections be garbage-collected with no harm, because they're just configuration tracking bundles and everything to the rqlite cluster is stateless.  Indeed, the true reason that `Close()` exists is the author's feeling that if you open something, you should be able to close it.  So why not `GetConnection()` then instead of `Open()`?  Or `GetClusterConfigurationTrackingObject()`?  I don't know.  Fork me.

`Leader()` and `Peers()` will both cause gorqlite to reverify its cluster information before return.  Note that if you call `Leader()` and then `Peers()` and something changes in between, it's possible to get inconsistent answers.

Since "weak" consistency is the default rqlite level, it is the default level for the client as well.  The user can change this at will (either in the connection string or via `SetConsistencyLevel()`, and then the new level will apply to all future calls).

## Tests

`go test` is used for testing.  A running cluster is required.

By default, gorqlite uses this config for testing:

	database URL : http://localhost:4001
	table name   : gorqlite_test

Also, the tests in package stdlib use the same config but a table name of `gorqlite_test_stdlib`.

These can overridden using the environment variables:

	GORQLITE_TEST_URL=https://somewhere.example.com:1234
	GORQLITE_TEST_URL=https//user:password@somewhere.example.com:1234
	etc.

	GORQLITE_TEST_TABLE=some_other_table
	GORQLITE_TEST_TABLE_STDLIB=some_other_table

## Pronunciation
rqlite is supposed to be pronounced "ree qwell lite" (though many people pronounce is "R Q lite" too).  So you could pronounce gorqlite as either "go ree kwell lite" or "gork lite".  The Klingon in me prefers the latter.  Really, isn't rqlite just the kind of battle-hardened, lean and mean system Klingons would use?  **Qapla'!**


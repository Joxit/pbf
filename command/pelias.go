package command

import (
	"fmt"
	"os"
	"sync"

	"github.com/missinglink/pbf/handler"
	"github.com/missinglink/pbf/leveldb"
	"github.com/missinglink/pbf/lib"
	"github.com/missinglink/pbf/parser"
	"github.com/missinglink/pbf/proxy"

	"github.com/codegangsta/cli"
)

// Pelias cli command
func Pelias(c *cli.Context) error {

	// validate args
	var argv = c.Args()
	if len(argv) != 1 {
		fmt.Println("invalid arguments, expected: {pbf}")
		os.Exit(1)
	}

	// create parser
	parser := parser.NewParser(argv[0])

	// -- bitmask --

	// bitmask is mandatory
	var bitmaskPath = c.String("bitmask")

	// bitmask file doesn't exist
	if _, err := os.Stat(bitmaskPath); err != nil {
		fmt.Println("bitmask file doesn't exist")
		os.Exit(1)
	}

	// debug
	fmt.Println("loaded bitmask:", bitmaskPath)

	// read bitmask from disk
	masks := lib.NewBitmaskMap()
	masks.ReadFromFile(bitmaskPath)

	// -- leveldb --

	// leveldb directory is mandatory
	var leveldbPath = c.String("leveldb")

	// stat leveldb destination
	lib.EnsureDirectoryExists(leveldbPath, "leveldb")

	// open database connection
	conn := &leveldb.Connection{}
	conn.Open(leveldbPath)
	defer conn.Close()

	// --

	// create parser handler
	var handle = &handler.DenormlizedJSON{
		Mutex: &sync.Mutex{},
		Conn:  conn,
	}

	// create filter proxy
	var filter = &proxy.WhiteList{
		Handler:      handle,
		NodeMask:     masks.Nodes,
		WayMask:      masks.Ways,
		RelationMask: masks.Relations,
	}

	// create store proxy
	var store = &proxy.StoreRefs{
		Handler: filter,
		Conn:    conn,
		Masks:   masks,
	}

	// Parse will block until it is done or an error occurs.
	parser.Parse(store)

	return nil
}
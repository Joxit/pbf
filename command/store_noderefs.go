package command

import "github.com/codegangsta/cli"

// StoreNodeRefs cli command
func StoreNodeRefs(c *cli.Context) error {

	// // validate args
	// var argv = c.Args()
	// if len(argv) != 1 {
	// 	log.Println("invalid arguments, expected: {pbf}")
	// 	os.Exit(1)
	// }
	//
	// // create parser
	// parser := parser.NewParser(argv[0])
	//
	// // bitmask is mandatory
	// var bitmaskPath = c.String("bitmask")
	// masks := lib.NewBitmaskMap()
	// masks.ReadFromFile(bitmaskPath)
	//
	// // leveldb directory is mandatory
	// var leveldbPath = c.String("leveldb")
	// lib.EnsureDirectoryExists(leveldbPath, "leveldb")
	//
	// // open database connection
	// conn := &leveldb.Connection{}
	// conn.Open(leveldbPath)
	// defer conn.Close()
	//
	// // create store proxy
	// var store = &proxy.StoreRefs{
	// 	Handler: &handler.Null{},
	// 	Conn:    conn,
	// 	Masks:   masks,
	// }
	//
	// // Parse will block until it is done or an error occurs.
	// parser.Parse(store)

	return nil
}

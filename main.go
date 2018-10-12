package main

// #cgo CFLAGS: -I/usr/lib/jvm/java-8-oracle/include/
// #cgo CFLAGS: -I/usr/lib/jvm/java-8-oracle/include/linux/
// #include <jni.h>
import "C"

import (
	"fmt"
	"github.com/pingcap/tidb/config"
	"github.com/pingcap/tidb/store/tikv"
)

//export echo
func echo(msg []byte) []byte {
	return msg
}

//export test
func test() {
	cli, err := tikv.NewRawKVClient([]string{"10.130.136.100:2379"}, config.Security{})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	fmt.Printf("cluster ID: %d\n", cli.ClusterID())

	key := []byte("Company")
	val := []byte("PingCAP")

	// put key into tikv
	err = cli.Put(key, val)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully put %s:%s to tikv\n", key, val)

	// get key from tikv
	val, err = cli.Get(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("found val: %s for key: %s\n", val, key)

	// delete key from tikv
	err = cli.Delete(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %s deleted\n", key)

	// get key again from tikv
	val, err = cli.Get(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("found val: %s for key: %s\n", val, key)
}

// main function is required, don't know why!
func main() {} // a dummy function

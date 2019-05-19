package op

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"fmt"
)

func Get(conn *zk.Conn, node string) {
	data, stat, err := conn.Get(node)
	if err != nil {
		log.Fatalf("get node %s failed, %s", node, err)
	}
	if stat.NumChildren != 0 {
		fmt.Printf("not a leaf node, list:\n")
		List(conn, node)
	} else {
		log.Printf("%s\n", data)
	}
}

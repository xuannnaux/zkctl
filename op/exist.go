package op

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
)

func Exists(conn *zk.Conn, node string) {
	exist, _, err := conn.Exists(node)
	if err != nil {
		log.Fatalf("check node %s failed, %s", node, err)
	}
	if exist {
		log.Printf("node %s exists\n", node)
	} else {
		log.Printf("node %s does not exists\n", node)
	}
}

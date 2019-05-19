package op

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
)

func Delete(conn *zk.Conn, node string) {
	exists, stat, err := conn.Exists(node)
	if err != nil {
		log.Fatalf("check node %s exists failed, %s", node, err)
	}
	if !exists {
		log.Fatalf("node %s does not exists", node)
	}
	err = conn.Delete(node, stat.Version)
	if err != nil {
		log.Fatalf("delete node %s failed, %s", node, err)
	}
	log.Printf("node %s has been deleted\n", node)

}

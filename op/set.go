package op

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
)

func Set(conn *zk.Conn, node string, data string){
	exists, stat, err := conn.Exists(node)
	if err != nil {
		log.Fatalf("check node %s exists failed, %s", node, err)
	}
	if !exists {
		log.Fatalf("node %s does not exists, use create command instead", node)
	}
	
	
	if stat.NumChildren != 0 {
		log.Printf("set failed, node is not a leaf node, please delete children first\n")
		List(conn, node)
		return
	}
	
	_, err = conn.Set(node, []byte(data), stat.Version)
	if err != nil {
		log.Fatalf("set node %s failed, %s", node, err)
	}
	log.Printf("set node %s success\n", node)
}

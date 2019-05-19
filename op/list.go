package op

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
)

func List(conn *zk.Conn, node string) {
	exists, _, err := conn.Exists(node)
	if err != nil {
		log.Fatalf("check node %s exists failed, %s", node, err)
	}
	if !exists {
		log.Fatalf("node %s does not exists", node)
	}
	fmt.Println(node)
	prefix := "|---"
	printChildren(conn, node, prefix)
}

func printChildren(conn *zk.Conn, node string, prefix string) {
	data, _, err := conn.Children(node)

	if err != nil {
		log.Fatalf("get node %s children failed, %s", node, err)
	}

	for _, v := range data {
		fmt.Printf("%s%s\n", prefix, v)
		var childPath string
		if node == "/" {
			childPath = fmt.Sprintf("%s%s", node, v)
		} else {
			childPath = fmt.Sprintf("%s/%s", node, v)
		}
		printChildren(conn, childPath, "    "+prefix)
	}
}

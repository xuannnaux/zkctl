package op

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"strings"
)

func Create(conn *zk.Conn, node string, data string, force bool, ephemeral bool) {
	exists, _, err := conn.Exists(node)
	if err != nil {
		log.Fatalf("check node %s exists failed, %s", node, err)
	}
	if exists {
		if !force {
			log.Fatalf("node %s exists, use -f flag to force create", node)
		}
	} else {
		paths := strings.Split(node[1:], "/")
		for i := 0; i < len(paths)-1; i++ {
			dir := "/" + strings.Join(paths[:i+1], "/")
			exists, _, err := conn.Exists(dir)
			if err != nil {
				log.Fatalf("check path %s failed", dir)
			}
			if !exists {
				_, err = conn.Create(dir, []byte(""), 0, zk.WorldACL(zk.PermAll))
				if err != nil {
					log.Fatalf("create parent path %s failed, %s", dir, err)
				}
				log.Printf("create parent path %s success\n", dir)
			}
		}
	}

	var nodeFlag int32 = 0
	if ephemeral {
		nodeFlag = zk.FlagEphemeral
	}
	_, err = conn.Create(node, []byte(data), nodeFlag, zk.WorldACL(zk.PermAll))
	log.Printf("create node %s success\n", node)
}

package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"github.com/spf13/cobra"
	"github.com/xuannnaux/zkctl/op"
	"github.com/xuannnaux/zkctl/util"
	"time"
	"log"
)

var (
	version = "0.0.1"

	zkConn           *zk.Conn
	zkSessionTimeout = 3 * time.Second

	serverList []string
	node       string
	data       string
)

func init() {
	log.SetFlags(0)
}

func PreInit() {
	var err error
	zkConn, _, err = zk.Connect(serverList, zkSessionTimeout, zk.WithLogInfo(false))
	if err != nil {
		log.Fatalf("connect to zk failed, %s", err)
		return
	}
}

func main() {
	defer func() {
		if zkConn != nil {
			zkConn.Close()
		}
	}()

	
	// root command and global flag
	cmdRoot := &cobra.Command{
		Version: version,
		Use:     "zkctl",
		Short:   "zookeeper management tool",
		Long:    "a tool to manage zookeeper node",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error

			serverList, err = util.ParseServer(cmd.Flags())
			if err != nil {
				log.Fatal(err)
			}

			node, err = util.ParseNode(cmd.Flags())
			if err != nil {
				log.Fatal(err)
			}

			data, err = util.ParseData(cmd.Flags())
			if err != nil {
				log.Fatal(err)
			}

			PreInit()
		},
	}
	cmdRoot.ValidArgs = []string{"create", "list", "get", "exist", "set", "delete", "meta", "size"}
	cmdRoot.PersistentFlags().StringP("server", "s", "127.0.0.1:2181", "zk server address")
	cmdRoot.PersistentFlags().StringP("node", "n", "", "node path")
	cmdRoot.PersistentFlags().StringP("data", "d", "", "node data")

	// create node
	cmdCreate := &cobra.Command{
		Use:   "create",
		Short: "create a zk node",
		Long:  "create a zk node, default to permanent node",
		Run: func(cmd *cobra.Command, args []string) {
			force, err := cmd.Flags().GetBool("force")
			if err != nil {
				log.Fatal("invalid force flag")
			}

			ephemeral, err := cmd.Flags().GetBool("ephemeral")
			if err != nil {
				log.Fatal("invalid ephemeral flag")
			}

			op.Create(zkConn, node, data, force, ephemeral)
		},
	}
	cmdCreate.Flags().BoolP("force", "f", false, "force replace node if node exists")
	cmdCreate.Flags().BoolP("ephemeral", "e", false, "set node as ephemeral node")

	// list node children
	cmdList := &cobra.Command{
		Use:   "list",
		Short: "list path children",
		Long:  "list path children",
		Run: func(cmd *cobra.Command, args []string) {
			op.List(zkConn, node)
		},
	}

	// get node data
	cmdGet := &cobra.Command{
		Use:   "get",
		Short: "show node data",
		Long:  "show node data",
		Run: func(cmd *cobra.Command, args []string) {
			op.Get(zkConn, node)
		},
	}

	// node exists
	cmdExists := &cobra.Command{
		Use:   "exist",
		Short: "check if node exist or not",
		Long:  "check if node exist or not",
		Run: func(cmd *cobra.Command, args []string) {
			op.Exists(zkConn, node)
		},
	}

	// set node
	cmdSet := &cobra.Command{
		Use:   "set",
		Short: "set node data",
		Long:  "set node data, node must exists",
		Run: func(cmd *cobra.Command, args []string) {
			op.Set(zkConn, node, data)
		},
	}

	// delete node
	cmdDelete := &cobra.Command{
		Use:   "delete",
		Short: "delete node",
		Long:  "delete node",
		Run: func(cmd *cobra.Command, args []string) {
			op.Delete(zkConn, node)
		},
	}

	cmdRoot.AddCommand(cmdCreate, cmdList, cmdGet, cmdExists, cmdSet, cmdDelete)

	cmdRoot.Execute()
}

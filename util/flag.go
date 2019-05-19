package util

import (
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
	"net"
	"strings"
)

var (
	defaultServer = "127.0.0.1:2181"
)

func ParseServer(flags *flag.FlagSet) ([]string, error) {
	server, err := flags.GetString("server")
	if err != nil {
		return nil, errors.New("invalid server flag")
	}
	servers := strings.Split(server, ",")

	serverList := make([]string, 0)
	for _, s := range servers {
		s := strings.TrimSpace(s)
		if s == "" {
			continue
		}
		_, _, err := net.SplitHostPort(s)
		if err != nil {
			return nil, errors.New("invalid server flag")
		}
		serverList = append(serverList, s)
	}

	if len(serverList) == 0 {
		serverList = append(serverList, defaultServer)
	}
	return serverList, nil
}

func ParseNode(flags *flag.FlagSet) (string, error) {
	node, err := flags.GetString("node")
	if err != nil {
		return "", errors.New("invalid path flag")
	}

	node = strings.TrimSpace(node)
	if node == "" || node[0] != "/"[0] {
		return "", errors.New("invalid path flag")
	}
	if node != "/" && node[len(node)-1] == "/"[0] {
		return "", errors.New("path must not ends with '/'")
	}
	return node, nil
}

func ParseData(flags *flag.FlagSet) (string, error) {
	data, err := flags.GetString("data")
	if err != nil {
		return "", errors.New("invalid data flag")
	}

	return data, nil
}

package snowflake

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func Init() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	node = n
}

func Generate() int64 {
	return node.Generate().Int64()
}

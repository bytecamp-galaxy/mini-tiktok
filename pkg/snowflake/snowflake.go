package snowflake

import "github.com/bwmarrin/snowflake"

var generator *snowflake.Node

func Init(node int64) {
	n, err := snowflake.NewNode(node)
	if err != nil {
		panic(err)
	}
	generator = n
}

func Generate() int64 {
	return generator.Generate().Int64()
}

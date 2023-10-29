package flags

import "flag"

type Flags struct {
	DbPath *string
}

func ParseFlags() *Flags {
	flags := Flags{}

	flags.DbPath = flag.String("db", "bolt.db", "Bolt Database path")

	flag.Parse()

	return &flags
}

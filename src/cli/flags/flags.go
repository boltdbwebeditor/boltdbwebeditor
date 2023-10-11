package flags

import "flag"

type Flags struct {
	DB *string
}

func ParseFlags() *Flags {
	flags := Flags{}

	flags.DB = flag.String("db", "bolt.db", "Bolt Database path")

	flag.Parse()

	return &flags
}

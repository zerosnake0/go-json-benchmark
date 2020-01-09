package json_benchmark

import (
	"github.com/zerosnake0/jzon"
)

var (
	jzonFastDecoder = jzon.NewDecoder(&jzon.DecoderOption{
		CaseSensitive: true,
	})
)

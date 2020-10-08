package json_benchmark

import (
	"github.com/zerosnake0/jzon"
)

var (
	jzonFastDecoderConfig = jzon.NewDecoderConfig(&jzon.DecoderOption{
		CaseSensitive: true,
	})
)

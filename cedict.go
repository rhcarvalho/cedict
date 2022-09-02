// Package cedict provides an embedded copy of the [CC-CEDICT] Chinese to
// English word dictionary.
//
// [CC-CEDICT]: https://www.mdbg.net/chinese/dictionary?page=cc-cedict
package cedict

import _ "embed"

// Bytes are the raw bytes of the dictionary.
//
// The format is described on the [CC-CEDICT Wiki].
//
// [CC-CEDICT Wiki]: https://cc-cedict.org/wiki/format:syntax
//
//go:embed cedict_1_0_ts_utf-8_mdbg.txt
var Bytes []byte

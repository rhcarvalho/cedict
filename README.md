# cedict

An always up-to-date copy of the
[CC-CEDICT](https://www.mdbg.net/chinese/dictionary?page=cc-cedict), a Chinese
to English word dictionary with pronunciation in pinyin for the Chinese
characters.

<a href="https://www.mdbg.net/chinese/dictionary"><img src="https://www.mdbg.net/logos/mdbg_dictionary_128x32.png" alt="MDBG Chinese-English dictionary" title="MDBG Chinese-English dictionary" style="border: solid 1px #c0c0c0" border="0" /></a>

The original dictionary is available as a plain text file. In addition to that,
this repository is a Go module, which allows for easy consumption of the
dictionary bytes from a Go program.

```go
package main

import (
	"os"

	"github.com/rhcarvalho/cedict"
)

func main() {
	os.Stdout.Write(append(cedict.Bytes[:200], []byte("...\n")...))
}
```

<!--
| CC-CEDICT Release Date | Number of Entries |
|------------------------|-------------------|
| 2013-03-19T13:58:14Z   | 105565            |
-->

# License

The original [CC-CEDICT](https://www.mdbg.net/chinese/dictionary?page=cc-cedict) by [MDBG](https://www.mdbg.net/) is licensed under [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/).

The remaining content of this repository is licensed under the [Apache License 2.0](LICENSE).

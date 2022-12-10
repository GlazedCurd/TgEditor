*bold \*text*
_italic \*text_

\#hashtag
__underline__
~strikethrough~
||spoiler||
*bold _italic bold ~italic bold strikethrough ||italic bold strikethrough spoiler||~ __underline italic bold___ bold*

[inline URL](http://www.example.com/)

`inline fixed-width code`

*bold \*text*


```
    pre-formatted fixed-width code block
```

__underline__
~strikethrough~


```go
package main

import (
	"fmt"
)

func getCancelButtonData() string {
	return "c"
}

func getFileButtonData(filename string) string {
	return fmt.Sprintf("f:%s", filename)
}

```
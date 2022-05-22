package main

import (
	"bytes"
	"fmt"
)

func writeEvents(events []Event) {
	buf := bytes.Buffer{}
	buf.WriteString(GO_OBS_PACKAGE)
	for _, e := range events {
		buf.WriteString(wrapComment(e.Docs))
		buf.WriteString(fmt.Sprintf("type %sEvent struct {\n", e.Name))
		for _, p := range e.Returns {
			str := fmt.Sprintf(
				"%s%s %s\n",
				wrapComment(p.Docs),
				p.Name,
				p.Type.String(),
			)
			buf.WriteString(str)
		}
		buf.WriteString("}\n\n")
	}
	fmtWrite("./gen_events.go", buf)
}

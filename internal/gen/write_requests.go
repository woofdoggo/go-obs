package main

import (
	"bytes"
	"fmt"
)

func writeRequests(reqs []Request) {
	buf := bytes.Buffer{}
	buf.WriteString(GO_OBS_PACKAGE)

	for _, r := range reqs {
		buf.WriteString(wrapComment(r.Docs))
		if r.Deprecated != "" {
			buf.WriteString("//\n// Deprecated:\n")
			buf.WriteString(wrapComment(r.Deprecated))
		}
		buf.WriteString(fmt.Sprintf("type %sRequest struct {\n", r.Name))
		for _, p := range r.Parameters {
			str := fmt.Sprintf(
				"%s%s %s\n",
				wrapComment(p.Docs),
				p.Name,
				p.Type.String(),
			)
			buf.WriteString(str)
		}
		buf.WriteString("}\n\n")

		buf.WriteString(fmt.Sprintf("type %sResponse struct {\n", r.Name))
		for _, p := range r.Returns {
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
	fmtWrite("./gen_requests.go", buf)
}

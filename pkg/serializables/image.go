package serializables

import (
	"bytes"
	"path"

	"golang.org/x/net/html"
)

func TransformInlineImages(src []byte, images map[string]string) []byte {
	root, err := html.Parse(bytes.NewReader(src))
	if err != nil {
		return src
	}

	replace(root, images)
	transformed := make([]byte, 0)
	buf := bytes.NewBuffer(transformed)
	if err = html.Render(buf, root); err != nil {
		return src
	}
	return buf.Bytes()
}

func replace(n *html.Node, images map[string]string) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			for i, myattr := range n.Attr {
				if myattr.Key == "src" {
					value := path.Base(myattr.Val)
					if data, found := images[value]; found {
						n.Attr[i].Val = data
					}
				}
			}
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		replace(child, images)
	}
}

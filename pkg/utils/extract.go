package utils

import (
	"strings"

	"golang.org/x/net/html"
)

func ExtractTextContent(rawHTML string) string {
	var sb strings.Builder
	doc, _ := html.Parse(strings.NewReader(rawHTML))

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "meta" {
				for _, a := range n.Attr {
					sb.WriteString(a.Key + "=" + a.Val + " ")
				}
				sb.WriteString("\n")
			}
			if n.Data == "script" {
				for _, a := range n.Attr {
					if a.Key == "type" && a.Val == "application/ld+json" {
						if n.FirstChild != nil {
							sb.WriteString(n.FirstChild.Data + "\n")
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)
	return sb.String()
}

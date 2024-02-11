package link

import (
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Useful for debugging if necessary
func HumanReadableNodeType(nodeType html.NodeType) string {

	switch nodeType {
	case 0:
		return "ErrorNode"
	case 1:
		return "TextNode"
	case 2:
		return "DocumentNode"
	case 3:
		return "ElementNode"
	case 4:
		return "CommentNode"
	case 5:
		return "DoctypeNode"
	default:
		return "Unknown"
	}
}

func HtmlBytesFromFile(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	checkErr(err)
	return bytes
}

func ParseHtml(htmlbytes []byte) []Link {
	reader := strings.NewReader(string(htmlbytes))
	document, err := html.Parse(reader)

	checkErr(err)

	var traverseDocument func(*html.Node) []Link
	traverseDocument = func(node *html.Node) []Link {

		links := make([]Link, 0)

		// Check to see if the node is a link <a> tag
		if node.Type == html.ElementNode && node.Data == "a" {
			var path string
			for _, attribute := range node.Attr {
				if attribute.Key == "href" {
					path = attribute.Val
					// fmt.Println("Path:", path)
					break
				}
			}

			// once we've found the href link, let's scan for text
			var text string
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				if child.Type == html.TextNode {
					text = child.Data
					link := Link{path, text}
					links = append(links, link)
				}
				links = append(links, traverseDocument(child)...)
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			links = append(links, traverseDocument(child)...)
		}
		return links
	}
	return traverseDocument(document)
}

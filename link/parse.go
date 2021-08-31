package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

//TODO: Read about public vs private members of a struct
type Link struct {
	href string
	text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := arrOfATagNodes(doc)
	var links []Link
	for _, n := range nodes {
		links = append(links, getLink(n))
	}
	return links, nil
}

// dfs algo, getting slice of nodes with <a> tag for links
func arrOfATagNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}

	//TODO: What is the best practice to create slice
	//Either by using var like below or by using make function
	var arr []*html.Node
	for it := node.FirstChild; it != nil; it = it.NextSibling {
		arr = append(arr, arrOfATagNodes(it)...)
	}

	return arr
}

func getLink(node *html.Node) Link {
	var lnk Link
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			//TODO: Check what is the best practice.
			//To use setter or directly use members
			lnk.href = attr.Val
			break
		}
	}

	//TODO: Same as above
	lnk.text = getText(node)
	return lnk
}

func getText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	var str string
	for it := node.FirstChild; it != nil; it = it.NextSibling {
		str += getText(it) + " "
	}

	return strings.Join(strings.Fields(str), " ")
}

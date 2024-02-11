package link

import (
	"testing"
)

var testCases = []struct {
	description   string
	filename      string
	rawHtml       []byte
	expectedLinks []Link
}{
	{
		description: "simple example",
		filename:    "",
		rawHtml:     []byte("<html><a href=\"/path\">foo</a></html"),
		expectedLinks: []Link{
			{
				Href: "/path",
				Text: "foo",
			},
		},
	},
	{
		description: "simple example from file",
		filename:    "./testFiles/ex1.html",
		rawHtml:     []byte{},
		expectedLinks: []Link{
			{
				Href: "/other-page",
				Text: "A link to another pag",
			},
		},
	},
}

func TestParseHtml(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			// Run ParseHtml
			var links []Link
			if len(testCase.rawHtml) > 0 {
				links = ParseHtml(testCase.rawHtml)
			} else if testCase.filename != "" {
				filebytes := HtmlBytesFromFile(testCase.filename)
				links = ParseHtml(filebytes)
			} else {
				t.Errorf("Invalid test case. Test case must reference an html file or have raw html")
			}

			// Assert Results of Parse match expected
			if len(testCase.expectedLinks) != len(links) {
				t.Errorf("Incorrect number of links - expected %d got %d", len(testCase.expectedLinks), len(links))
			}
			for i, link := range links {
				if link.Href != testCase.expectedLinks[i].Href || link.Text != testCase.expectedLinks[i].Text {
					t.Errorf("Found mismatching links - expected %s got %s", testCase.expectedLinks[i], link)
				}
			}
		})
	}

}

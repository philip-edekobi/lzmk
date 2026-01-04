package codegen

import (
	"fmt"
	"time"

	"github.com/philip-edekobi/lzmk/pkg/parser"
)

func generateHyperTextForNodeType(node *parser.Node) (string, error) {
	switch node.Kind {
	case parser.HeadingNode:
		return "\n\t\t" + generateHeading(node), nil
	case parser.TextNode:
		return "\n\t\t" + generateParagraph(node), nil
	case parser.URLNode:
		return "\n\t\t" + generateImgUrl(node), nil
	default:
		return "", fmt.Errorf("found invalid node type")
	}
}

func generateHeading(node *parser.Node) string {
	return fmt.Sprintf("<h2 class=\"text-xl sm:text-2xl font-semibold tracking-tight\">%s</h2>", node.Value())
}

func generateParagraph(node *parser.Node) string {
	return fmt.Sprintf("<p class=\"text-base leading-7 text-zinc-700 dark:text-zinc-300\">%s</p>", node.Value())
}

func generateImgUrl(node *parser.Node) string {
	return fmt.Sprintf("<img src=\"%s\" alt=\"%s\" class=\"w-full max-w-full rounded-lg border border-zinc-200 dark:border-zinc-800\" />\n", node.URLData.URL, node.URLData.AltText)
}

func convertDateToIsoTime(dateStr string) string {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		fmt.Print(err.Error())
		panic("error parsing date")
	}

	isoTime := t.Format(time.RFC3339)
	return isoTime
}

func getYear(dateStr string) int {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Print(err.Error())
		panic("error parsing date")
	}

	return t.Year()
}

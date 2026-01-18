package codegen

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/philip-edekobi/lzmk/pkg/parser"
)

func generateHyperTextForNodeType(node *parser.Node) (string, error) {
	switch node.Kind {
	case parser.HeadingNode:
		return "\n\t\t" + generateHeading(node), nil
	case parser.TextNode:
		return "\n\t\t" + generateParagraph(node), nil
	case parser.MediaNode:
		mediaMarkup, err := generateMediaUrl(node)
		if err != nil {
			return "", err
		}

		return "\n\t\t" + mediaMarkup, nil
	default:
		return "", fmt.Errorf("found invalid node type")
	}
}

func generateHeading(node *parser.Node) string {
	return fmt.Sprintf("<h2 class=\"text-xl sm:text-2xl font-semibold tracking-tight\">%s</h2>", node.Value())
}

func generateParagraph(node *parser.Node) string {
	return fmt.Sprintf("<p class=\"text-base leading-7 text-zinc-700 dark:text-zinc-300\">%s</p>", parseLinksInText(node.Value()))
}

func generateMediaUrl(node *parser.Node) (string, error) {
	switch strings.Trim(node.MediaData.MediaType, " \n") {
	case "vid":
		return fmt.Sprintf("<video controls class=\"w-full rounded-lg border border-zinc-200 dark:border-zinc-800\">\n\t\t\t<source src=\"%s\" />\n\t\t</video>\n", node.MediaData.URL), nil
	case "img":
		return fmt.Sprintf("<img src=\"%s\" alt=\"%s\" class=\"w-full max-w-full rounded-lg border border-zinc-200 dark:border-zinc-800\" />\n", node.MediaData.URL, node.MediaData.AltText), nil
	default:
		fmt.Println("PROBLEMATIC NODE:", node)
		return "", fmt.Errorf("invalid media type")
	}
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

func isExternalLink(link string) bool {
	u, err := url.Parse(link)
	if err != nil {
		return false
	}

	if u.Host != "" {
		return true
	}

	if strings.HasPrefix(link, "//") {
		return true
	}

	if u.Scheme != "" {
		return true
	}

	return false
}

func parseLinksInText(text string) string {
	re := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 3 {
			return match
		}
		alt := submatches[1]
		url := submatches[2]

		if isExternalLink(url) {
			return fmt.Sprintf("<a target=\"_blank\" rel=\"noopener noreferrer\" class=\"text-blue-600 dark:text-blue-400 underline underline-offset-4 hover:text-blue-700 dark:hover:text-blue-300\" href=\"%s\">%s</a>\n", url, alt)
		}

		return fmt.Sprintf("<a class=\"font-medium text-blue-600 dark:text-blue-400 underline underline-offset-4 hover:text-blue-700 dark:hover:text-blue-300\" href=\"%s\">%s</a>\n", url, alt)
	})
}

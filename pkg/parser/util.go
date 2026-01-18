package parser

import (
	"fmt"
	"strings"

	"github.com/philip-edekobi/lzmk/pkg/lexer"
)

func (p *Parser) parseHeader() *Node {
	_, err := p.consumeToken(lexer.HeaderHash)
	if err != nil {
		return nil
	}

	headerToken, err := p.consumeToken(lexer.String)
	if err != nil {
		return nil
	}

	n := newNode(HeadingNode)
	n.StringValue = headerToken.Value

	return n
}

func (p *Parser) parseString() *Node {
	stringToken, err := p.consumeToken(lexer.String)
	if err != nil {
		return nil
	}

	tails := []*lexer.Token{}
	for {
		_, err := p.consumeToken(lexer.NewLine)
		if err != nil {
			break
		}

		strToken, err := p.consumeToken(lexer.String)
		if err != nil {
			break
		}

		tails = append(tails, strToken)
	}

	n := newNode(TextNode)
	n.StringValue = stringToken.Value

	for _, tailToken := range tails {
		n.StringValue += "<br>\n" + tailToken.Value
	}

	return n
}

func (p *Parser) parseMedia() *Node {
	_, err := p.consumeToken(lexer.HashBang)
	if err != nil {
		return nil
	}

	mediaType, err := p.consumeToken(lexer.String)
	if err != nil {
		return nil
	}

	_, err = p.consumeToken(lexer.LeftBrace)
	if err != nil {
		return nil
	}

	altTextToken, err := p.consumeToken(lexer.String)
	if err != nil {
		return nil
	}

	_, err = p.consumeToken(lexer.RightBrace)
	if err != nil {
		return nil
	}

	_, err = p.consumeToken(lexer.LeftParen)
	if err != nil {
		return nil
	}

	urlToken, err := p.consumeToken(lexer.String)
	if err != nil {
		return nil
	}

	_, err = p.consumeToken(lexer.RightParen)
	if err != nil {
		return nil
	}

	n := newNode(MediaNode)
	n.MediaData = MediaInfo{URL: urlToken.Value, AltText: altTextToken.Value, MediaType: mediaType.Value}

	return n
}

func (p *Parser) parseMetadata() *Node {
	_, err := p.consumeToken(lexer.MetaHash)
	if err != nil {
		return nil
	}

	metaString, err := p.consumeToken(lexer.String)
	if err != nil {
		return nil
	}

	metaSubStrings := strings.SplitN(metaString.Value, " ", 2)

	n := newNode(MetadataNode)
	n.Metadata = MetadataInfo{Key: metaSubStrings[0], Value: metaSubStrings[1]}

	p.MetaHashMap[n.Metadata.Key] = n.Metadata.Value

	return n
}

func (p *Parser) parseTitleNode() (*Node, error) {
	for p.peek().Kind == lexer.NewLine {
		p.advance()
	}

	_, err := p.consumeToken(lexer.TitleHash)
	if err != nil {
		return nil, err
	}

	titleToken, err := p.consumeToken(lexer.String)
	if err != nil {
		return nil, err
	}

	n := newNode(TitleNode)
	n.StringValue = titleToken.Value

	return n, nil
}

func (p *Parser) parseBody() (*Node, error) {
	b := newNode(BodyNode)

	for p.peek().Kind != lexer.EOF {
		for p.peek().Kind == lexer.NewLine {
			p.advance()
		}

		if p.peek().Kind == lexer.EOF {
			break
		}

		headerNode := p.parseHeader()
		if headerNode != nil {
			b.Children = append(b.Children, headerNode)
			continue
		}

		stringNode := p.parseString()
		if stringNode != nil {
			b.Children = append(b.Children, stringNode)
			continue
		}

		mediaNode := p.parseMedia()
		if mediaNode != nil {
			b.Children = append(b.Children, mediaNode)
			continue
		}

		metaNode := p.parseMetadata()
		if metaNode != nil {
			b.Children = append(b.Children, metaNode)
			continue
		}

		return nil, fmt.Errorf(
			"Error occured at %d:%d: failed to find any valid tokens",
			p.peek().Line,
			p.peek().Col,
		)
	}

	return b, nil
}

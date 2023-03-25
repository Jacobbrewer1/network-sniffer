package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Jacobbrewer1/network-sniffer/scraper/src/config"
	"golang.org/x/net/context"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func scrape() ([][]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, config.Cfg.Setup.ApiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", getToken()))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	got, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println(err)
		}
	}(resp.Body)

	p, err := format(string(got))
	if err != nil {
		return nil, err
	}

	return parseTable(strings.NewReader(p)), nil
}

func getToken() string {
	unparsedToken := fmt.Sprintf("%s:%s", config.Cfg.Setup.Username, config.Cfg.Setup.Password)
	t := base64.StdEncoding.EncodeToString([]byte(unparsedToken))
	return t
}

func parseTable(body io.Reader) [][]string {
	z := html.NewTokenizer(body)
	var content [][]string
	var innerContent []string
	counter := 0

	// While have not hit the </html> tag
	for z.Token().Data != "html" {
		tt := z.Next()
		if tt == html.StartTagToken {
			t := z.Token()

			if t.Data == "tr" {
				innerContent = []string{}
				counter = 0
			}

			if t.Data == "td" {
				inner := z.Next()
				if inner == html.TextToken {
					text := (string)(z.Text())
					t := strings.TrimSpace(text)
					innerContent = append(innerContent, t)
					counter++
				}

				// change counter according to your coulumn size
				if counter == 8 {
					content = append(content, innerContent)
				}
			}
		}
	}
	return content
}

func format(page string) (string, error) {
	// Remove everything before table
	s := strings.Split(page, `<table border="1" cellpadding="0" cellspacing="0" width="99%">`)
	if len(s) < 2 {
		return "", errors.New("cannot split html, most likely unauthorised")
	}

	p := s[1]

	p = removeSpan(p)

	// Remove everything after table
	p = removeLastTableClosure(p)
	p = removeLastTableClosure(p)

	p = fmt.Sprintf("<html><body><table>%s</table></body></html>", p)

	return p, nil
}

func removeSpan(text string) string {
	text = strings.ReplaceAll(text, `<span class="thead">`, "")
	text = strings.ReplaceAll(text, `<span class="ttext">`, "")
	text = strings.ReplaceAll(text, `</span>`, "")
	return text
}

func removeLastTableClosure(text string) string {
	pos := strings.LastIndex(text, `</table>`)
	if pos == -1 {
		panic("invalid string")
	}

	text = text[:pos]
	return text
}

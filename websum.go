package websum

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SummarizeWeb returns the summary of web containing
// html versin, title, count headings and links, and whether the page contains login form or not
func SummarizeWeb(url string) (sum Summary, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	sum.HTMLVersion = htmlVersion(doc)
	sum.Title = doc.FindMatcher(goquery.Single("title")).Text()
	sum.HeadingsCount = countHeadings(doc)
	sum.LinksCount = countLinks(doc, url)
	sum.ContainLogin = loginForm(doc)
	return
}

// htmlVersion gets the html version based on HTMLTypes
func htmlVersion(doc *goquery.Document) string {
	raw, err := doc.Html()
	if err != nil {
		return unknown
	}
	for version, matcher := range HTMLTypes {
		if strings.Contains(raw, matcher) {
			return version
		}
	}
	return unknown
}

// countHeadings counts all h1 to h6 html tag
func countHeadings(doc *goquery.Document) map[string]int {
	res := make(map[string]int)
	for i := 1; i <= 6; i++ {
		tag := fmt.Sprintf("h%d", i)
		res[tag] = len(doc.Find(tag).Nodes)
	}
	return res
}

// countLinks counts all a tag html
// directly do http request to href link to count the inaccessable
// if href link contains the same domain as requested url, counts as internal link
func countLinks(doc *goquery.Document, url string) Links {
	regex := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	partURL := regex.FindStringSubmatch(url)
	domain := partURL[1]

	links := Links{}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if !exists {
			return
		}

		// doesn't contain domain in link, means link to internal page or using the same domain
		if len(regex.FindStringSubmatch(link)) == 0 {
			link = partURL[0] + link
		}

		res, err := http.Get(link)
		if err != nil || res.StatusCode != http.StatusOK {
			links.Inaccessable++
		}

		if strings.Contains(link, domain) {
			links.Internal++
			return
		}
		links.External++
	})
	return links
}

// loginForm finds the input type password and login button
// login button whether HTML tag button or input type submit button
func loginForm(doc *goquery.Document) bool {
	passExists, loginBtn := false, false
	doc.Find("input").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("type")
		if val == "password" {
			passExists = true
		}
		// find login button type submit based on loginKeys
		if val == "submit" {
			for _, v := range s.Get(0).Attr {
				for _, key := range loginKeys {
					if strings.Contains(strings.ToLower(v.Val), key) {
						loginBtn = true
					}
				}
			}
		}
	})

	// finding button with keyword from loginKeys
	doc.Find("button").Each(func(i int, s *goquery.Selection) {
		for _, key := range loginKeys {
			if strings.Contains(strings.ToLower(s.Text()), key) {
				loginBtn = true
			}
		}
	})
	return passExists && loginBtn
}

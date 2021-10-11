package websum

var (
	// reference: https://howtocheckversion.com/check-html-version-website/
	HTMLTypes = map[string]string{
		"HTML 4.01 Strict":       `"-//W3C//DTD HTML 4.01//EN"`,
		"HTML 4.01 Transitional": `"-//W3C//DTD HTML 4.01 Transitional//EN"`,
		"HTML 4.01 Frameset":     `"-//W3C//DTD HTML 4.01 Frameset//EN"`,
		"XHTML 1.0 Strict":       `"-//W3C//DTD XHTML 1.0 Strict//EN"`,
		"XHTML 1.0 Transitional": `"-//W3C//DTD XHTML 1.0 Transitional//EN"`,
		"XHTML 1.0 Frameset":     `"-//W3C//DTD XHTML 1.0 Frameset//EN"`,
		"XHTML 1.1":              `"-//W3C//DTD XHTML 1.1//EN"`,
		"HTML 5":                 `<!DOCTYPE html>`,
	}

	loginKeys = []string{
		"login",
		"log in",
		"sign in",
	}
)

const (
	unknown = "UNKNOWN"
)

type Summary struct {
	HTMLVersion   string         `json:"html_version"`
	Title         string         `json:"title"`
	HeadingsCount map[string]int `json:"headings_count"`
	LinksCount    Links          `json:"links_count"`
	ContainLogin  bool           `json:"contain_login"`
}

type Links struct {
	External     int `json:"external"`
	Internal     int `json:"internal"`
	Inaccessable int `json:"inaccessible"`
}

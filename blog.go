package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/tj/front"
	stripmd "github.com/writeas/go-strip-markdown"
)

const (
	wordsPerMinute = 200
)

type Post struct {
	SmallContainer bool
	Slug           string
	WordCount      int
	ReadingTime    int          // in minutes
	Title          string       `yaml:"title"`
	AuthorHandles  string       `yaml:"authors"`
	Authors        []TeamMember `yaml:"ignore"`
	IsDraft        bool         `yaml:"draft"`
	Date           string       `yaml:"date"`
	DateTime       time.Time
	Description    string `yaml:"desc"`
	DescBody       template.HTML
	HTMLBody       template.HTML // html content
}

func (p *Post) FormatDate() string {
	return p.DateTime.Format("January 2, 2006")
}

func createPage(p *Post) error {
	fp := filepath.Join(buildDir, "blog", p.Slug+".html")
	os.Remove(fp)

	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	err = postTmpl.ExecuteTemplate(f, "base", p)
	if err != nil {
		return err
	}

	fmt.Printf("📔  blog/%s.html sucessfully generated.\n", p.Slug)
	return nil
}

func createBlogIndex(posts []*Post) error {
	fp := filepath.Join(buildDir, "blog", "index.html")
	os.Remove(fp)

	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	err = blogTmpl.ExecuteTemplate(f, "base", Page{
		Title: "Blog",
		Posts: posts,
	})
	if err != nil {
		return err
	}

	fmt.Printf("📚  blog/index.html sucessfully generated.\n")
	return nil

}

func newPost(fn string) (*Post, error) {
	var p Post
	p.SmallContainer = true
	p.Slug = strings.Replace(fn, ".md", "", 1)

	path := filepath.Join(blogDir, fn)
	md, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	content, err := front.Unmarshal(md, &p)
	if err != nil {
		return nil, err
	}

	p.DateTime, err = time.Parse("02/01/2006", p.Date)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	// Get author by handle
	authors := []TeamMember{}
	handles := strings.Split(p.AuthorHandles, ",")
	for _, member := range Members {
		if slices.Contains(handles, member.Handle) {
			authors = append(authors, member)
		}
	}
	sortTeamMembers(authors)
	p.Authors = authors

	p.DescBody = template.HTML(p.Description)
	// parser
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.MathJax | parser.Footnotes
	parser := parser.NewWithExtensions(extensions)
	// renderer
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.FootnoteNoHRTag
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	p.HTMLBody = template.HTML(markdown.ToHTML(content, parser, renderer))

	// Word count and reading time
	plaintext := stripmd.Strip(string(content))
	words := strings.Fields(plaintext)
	p.WordCount = len(words)
	p.ReadingTime = p.WordCount/wordsPerMinute + 1

	return &p, nil
}

func genBlog() {
	var posts []*Post

	// Read markdown files
	mdfiles, err := ioutil.ReadDir(blogDir)
	if err != nil {
		log.Fatal("Can't read the posts directory")
	}

	for _, file := range mdfiles {
		fn := file.Name()
		// Create post for each markdown file
		p, err := newPost(fn)
		if err != nil {
			log.Fatal(err)
		}
		// Add it to the list
		posts = append(posts, p)
	}

	subdir := filepath.Join(buildDir, "blog")
	_ = os.Mkdir(subdir, os.ModePerm)

	// Build each post page
	for _, post := range posts {
		err := createPage(post)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create blog index

	err = createBlogIndex(posts)
	if err != nil {
		log.Fatal(err)
	}
}

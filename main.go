package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/exp/slices"
)

const (
	servePort = "8100"
	tmplDir   = "./templates"
	buildDir  = "./public"
	blogDir   = "./blog"

	layoutTmplName   = "layout.html"
	indexTmplName    = "index.html"
	teamTmplName     = "team.html"
	postTmplName     = "post.html"
	clientTmplName   = "client.html"
	researchTmplName = "research.html"
	bridgesTmplName  = "bridges.html"
)

const description string = "Common Prefix is a small team of scientists and software engineers offering blockchain science consulting services."

var layoutPath = filepath.Join(tmplDir, layoutTmplName)
var homeTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, indexTmplName)))
var teamTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, teamTmplName)))
var clientTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, clientTmplName)))
var researchTmpl = template.Must(template.New("").Funcs(template.FuncMap{
	"sub": func(a, b int) int {
		return a - b
	},
	"teamHandle": func(name string) string {
		idx := slices.IndexFunc(team, func(m TeamMember) bool {
			n := m.Name
			n = strings.TrimPrefix(n, "Prof. ")
			n = strings.TrimPrefix(n, "Dr. ")
			return n == name
		})
		if idx == -1 {
			return ""
		}
		member := team[idx]
		return member.Handle
	},
}).ParseFiles(layoutPath, filepath.Join(tmplDir, researchTmplName)))
var bridgesTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, bridgesTmplName)))
var postTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, postTmplName)))

// Data structures

type TeamMember struct {
	Handle         string
	Name           string
	Specialization string
	Desc           template.HTML
	Image          string
}

// john_doe.jpg -> john_doe_w150.jpg
func (m *TeamMember) ImageLow() string {
	bits := strings.Split(m.Image, ".")
	return fmt.Sprintf("%s_w150.%s", bits[0], bits[1])
}

func sortTeamMembers(team []TeamMember) {
	sort.Slice(team, func(i, j int) bool {
		lastname := func(n string) string {
			n = strings.TrimPrefix(n, "Prof. ")
			n = strings.TrimPrefix(n, "Dr. ")
			n = strings.Split(n, " ")[1]
			return n
		}
		n1 := lastname(team[i].Name)
		n2 := lastname(team[j].Name)
		return n1 < n2
	})
}

type Finding struct {
	Url  string
	Name string
	Date string // optional
}

func (f *Finding) Ext() string {
	bits := strings.Split(f.Url, ".")
	return bits[len(bits)-1]
}

type ProjectLink struct {
	Url  string
	Name string
	Date string // optional
}

func (f *ProjectLink) Ext() string {
	bits := strings.Split(f.Url, ".")
	return bits[len(bits)-1]
}

func (f *ProjectLink) IsFile() bool {
	// primitive check
	u, _ := url.Parse(f.Url)
	return strings.Contains(u.Path, ".")
}

type Project struct {
	Title    string
	Desc     template.HTML
	Date     string // example 28/03/2024
	DateTime time.Time
	Findings []Finding
	Links    []ProjectLink
}

type Client struct {
	Handle   string
	Name     string
	Body     template.HTML
	Image    template.HTML
	Projects []Project
	Findings []Finding
	Team     []TeamMember
}

type ResearchPaper struct {
	Handle         string
	Name           string
	Conference     string
	ConferenceYear int
	Authors        []string
	Url            string
	Tags           []Tag
}

type Research struct {
	ResearchPapers []ResearchPaper
	TagToColor     map[Tag]string
	AllAuthors     []string
	AllTags        []Tag
	AllConferences []string
	AllYears       []int
}

type Page struct {
	SmallContainer bool
	Title          string
	Description    string
	Members        []TeamMember
	Clients        []Client
	Research       Research
}

type ClientPage struct {
	SmallContainer bool
	Title          string
	Description    string
	Client         Client
	NextClient     Client
}

type ResearchPage struct {
	Title      string
	TagToColor map[Tag]string
}

var team = []TeamMember{}

func htmlToFormattedString(s template.HTML) string {
	bmp := bluemonday.StripTagsPolicy()
	replacer := strings.NewReplacer("\n", " ", "\t", "")
	return replacer.Replace(strings.TrimSpace(bmp.Sanitize(string(s))))
}

func build() {
	//
	// Build index page
	//
	index := filepath.Join(buildDir, indexTmplName)
	// Remove the old version
	os.Remove(index)
	// Create new file
	f, err := os.Create(index)
	if err != nil {
		log.Fatalf("can't create %s", indexTmplName)
	}
	homeTmpl.ExecuteTemplate(f, "base", Page{SmallContainer: true, Title: "", Description: description, Members: team, Clients: Clients})
	f.Close()
	fmt.Printf("🏠  %s sucessfully generated.\n", indexTmplName)

	//
	// Build team page
	//
	teamPage := filepath.Join(buildDir, teamTmplName)
	// Remove the old version
	os.Remove(teamPage)
	// Create new file
	f, err = os.Create(teamPage)
	if err != nil {
		log.Fatalf("can't create %s", teamTmplName)
	}
	teamTmpl.ExecuteTemplate(f, "base", Page{Title: "Team", Members: team, Description: description, Clients: Clients})
	f.Close()
	fmt.Printf("👫  %s sucessfully generated.\n", teamTmplName)

	//
	// Build clients directory
	//
	psf := filepath.Join(buildDir, "clients")
	_ = os.Mkdir(psf, os.ModePerm)

	// Build clients pages
	for idx, p := range Clients {
		nextP := Clients[(idx+1)%len(Clients)]
		pf := filepath.Join(buildDir, "clients", p.Handle+".html")
		// Remove the old version
		os.Remove(pf)
		// Create new file
		f, err = os.Create(pf)
		if err != nil {
			log.Fatal("can't create clients/" + p.Handle + ".html")
		}

		clientTmpl.ExecuteTemplate(f, "base", ClientPage{SmallContainer: true, Title: p.Name, Description: htmlToFormattedString(p.Body), Client: p, NextClient: nextP})
		f.Close()
		fmt.Printf("📖  clients/%s.html sucessfully generated.\n", p.Handle)
	}

	//
	// Build research page
	//
	researchPage := filepath.Join(buildDir, researchTmplName)
	// Remove the old version
	os.Remove(researchPage)
	// Create new file
	f, err = os.Create(researchPage)
	if err != nil {
		log.Fatalf("can't create %s", researchTmplName)
	}

	allAuthors := []string{}
	for _, tm := range team {
		for _, r := range ResearchPapers {
			name := tm.Name
			name = strings.TrimPrefix(name, "Prof. ")
			name = strings.TrimPrefix(name, "Dr. ")
			if slices.Contains(r.Authors, name) && !slices.Contains(allAuthors, name) {
				allAuthors = append(allAuthors, name)
			}
		}
	}

	allTags := []Tag{}
	for _, r := range ResearchPapers {
		for _, t := range r.Tags {
			if !slices.Contains(allTags, t) {
				allTags = append(allTags, t)
			}
		}
	}

	allConferences := []string{}
	for _, r := range ResearchPapers {
		if !slices.Contains(allConferences, r.Conference) {
			allConferences = append(allConferences, r.Conference)
		}
	}

	allYears := []int{}
	for _, r := range ResearchPapers {
		if r.ConferenceYear != 0 && !slices.Contains(allYears, r.ConferenceYear) {
			allYears = append(allYears, r.ConferenceYear)
		}
	}
	slices.Sort(allYears)
	slices.Reverse(allYears)

	err = researchTmpl.ExecuteTemplate(f, "base", Page{Title: "Research",
		Description: description,
		Research: Research{ResearchPapers: ResearchPapers,
			TagToColor:     TagToColor,
			AllAuthors:     allAuthors,
			AllTags:        allTags,
			AllConferences: allConferences,
			AllYears:       allYears,
		}})
	if err != nil {
		log.Println(err)
		log.Fatalf("can't generate  %s", researchTmplName)
	}

	f.Close()
	fmt.Printf("👫  %s sucessfully generated.\n", researchTmplName)

	//
	// Build blog posts
	//
	genBlog()

	//
	// Build bridges page
	//
	bridgesPage := filepath.Join(buildDir, bridgesTmplName)
	// Remove the old version
	os.Remove(bridgesPage)
	// Create new file
	f, err = os.Create(bridgesPage)
	if err != nil {
		log.Fatalf("can't create %s", bridgesTmplName)
	}
	bridgesTmpl.ExecuteTemplate(f, "base", Page{SmallContainer: true, Title: "Bridges", Description: description})

	f.Close()
	fmt.Printf("🌉  %s sucessfully generated.\n", bridgesTmplName)
}

func main() {
	// build a list of all members
	for _, m := range Members {
		team = append(team, m)
	}
	sortTeamMembers(team)

	serve := flag.Bool("serve", false, "serve mode")
	port := flag.String("p", servePort, "port to serve on")
	flag.Parse()

	build()

	if *serve == true {
		fs := http.FileServer(http.Dir(buildDir))
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// Clean urls
			if strings.HasPrefix(r.URL.Path, "/clients") || strings.HasPrefix(r.URL.Path, "/team") {
				if !strings.HasSuffix(r.URL.Path, ".html") {
					r.URL.Path = r.URL.Path + ".html"
				}
			}

			fs.ServeHTTP(w, r)
		})
		fmt.Printf("🧞‍♂️  Serving on http://localhost:%s\n", *port)
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	}
}

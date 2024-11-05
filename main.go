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
	blogTmplName     = "blog.html"
	postTmplName     = "post.html"
	clientTmplName   = "client.html"
	grantorTmplName  = "grantor.html"
	researchTmplName = "research.html"
	bridgesTmplName  = "bridges.html"
	careersTmplName  = "careers.html"
)

const description string = "Common Prefix is a team of scientists and engineers researching and implementing blockchain protocols."

var layoutPath = filepath.Join(tmplDir, layoutTmplName)
var homeTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, indexTmplName)))
var teamTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, teamTmplName)))
var clientTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, clientTmplName)))
var grantorTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, grantorTmplName)))
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
var blogTmpl = template.Must(template.ParseFiles(layoutPath, filepath.Join(tmplDir, blogTmplName)))
var careersTmpl = template.Must(template.New("").ParseFiles(layoutPath, filepath.Join(tmplDir, careersTmplName)))

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
	IsBridge bool
}

type Client struct {
	Handle      string
	Name        string
	Body        template.HTML
	Image       template.HTML
	Projects    []Project
	Findings    []Finding
	Team        []TeamMember
	HideOnIndex bool
}

type Grantor struct {
	Handle      string
	Name        string
	Body        template.HTML
	Image       template.HTML
	Projects    []Project
	Findings    []Finding
	Team        []TeamMember
	HideOnIndex bool
}

type ResearchPaper struct {
	Handle         string
	Name           string
	Conference     ConferenceAbbreviation
	ConferenceYear int
	Authors        []string
	Url            string
	Tags           []Tag
	Citations      int
}

type Research struct {
	ResearchPapers []ResearchPaper
	TagToColor     map[Tag]string
	AllAuthors     []string
	AllTags        []Tag
	AllConferences []ConferenceAbbreviation
	AllYears       []int
}

type JobOpening struct {
	Name        string
	Url         string
	Location    string
	Type        string
	Description string
}

type Page struct {
	SmallContainer bool
	Title          string
	Description    string
	Members        []TeamMember
	Clients        []Client
	Grants         []Grantor
	Research       Research
	Posts          []*Post
}

type BridgesPage struct {
	SmallContainer bool
	Title          string
	Description    string
	Projects       []Project
}

type ClientPage struct {
	SmallContainer bool
	Title          string
	Description    string
	Client         Client
	NextClient     Client
}

type GrantorPage struct {
	SmallContainer bool
	Title          string
	Description    string
	Grantor        Grantor
	NextGrantor    Grantor
}

type CareersPage struct {
	SmallContainer bool
	Title          string
	Description    string
	Openings       []JobOpening
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
	homeTmpl.ExecuteTemplate(f, "base", Page{SmallContainer: true, Title: "", Description: description, Members: team, Clients: Clients, Grants: Grants})
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
	teamTmpl.ExecuteTemplate(f, "base", Page{Title: "Team", Members: team, Description: description, Clients: Clients, Grants: Grants})
	f.Close()
	fmt.Printf("👫  %s sucessfully generated.\n", teamTmplName)

	//
	// Build clients directory
	//
	psfClients := filepath.Join(buildDir, "clients")
	_ = os.Mkdir(psfClients, os.ModePerm)

	// Build clients pages
	for idx, p := range Clients {
		nextP := Clients[(idx+1)%len(Clients)]
		pfClients := filepath.Join(buildDir, "clients", p.Handle+".html")
		// Remove the old version
		os.Remove(pfClients)
		// Create new file
		f, err = os.Create(pfClients)
		if err != nil {
			log.Fatal("can't create clients/" + p.Handle + ".html")
		}

		clientTmpl.ExecuteTemplate(f, "base", ClientPage{SmallContainer: true, Title: p.Name, Description: htmlToFormattedString(p.Body), Client: p, NextClient: nextP})
		f.Close()
		fmt.Printf("📖  clients/%s.html sucessfully generated.\n", p.Handle)
	}

	//
	// Build grants directory
	//
	psfGrants := filepath.Join(buildDir, "grants")
	_ = os.Mkdir(psfGrants, os.ModePerm)

	// Build grants pages
	for idx, p := range Grants {
		nextP := Grants[(idx+1)%len(Grants)]
		pfGrants := filepath.Join(buildDir, "grants", p.Handle+".html")
		// Remove the old version
		os.Remove(pfGrants)
		// Create new file
		f, err = os.Create(pfGrants)
		if err != nil {
			log.Fatal("can't create grants/" + p.Handle + ".html")
		}

		grantorTmpl.ExecuteTemplate(f, "base", GrantorPage{SmallContainer: true, Title: p.Name, Description: htmlToFormattedString(p.Body), Grantor: p, NextGrantor: nextP})
		f.Close()
		fmt.Printf("🎁  grants/%s.html sucessfully generated.\n", p.Handle)
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

	allConferences := []ConferenceAbbreviation{}
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

	sort.Slice(ResearchPapers, func(i, j int) bool {
		return ResearchPapers[i].Citations > ResearchPapers[j].Citations
	})

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
	bridgeProjects := []Project{}
	for _, c := range Clients {
		for _, p := range c.Projects {
			if p.IsBridge {
				bridgeProjects = append(bridgeProjects, p)
			}
		}
	}

	bridgesPage := filepath.Join(buildDir, bridgesTmplName)
	// Remove the old version
	os.Remove(bridgesPage)
	// Create new file
	f, err = os.Create(bridgesPage)
	if err != nil {
		log.Fatalf("can't create %s", bridgesTmplName)
	}
	bridgesTmpl.ExecuteTemplate(f, "base", BridgesPage{SmallContainer: true, Title: "Bridges", Description: description, Projects: bridgeProjects})

	f.Close()
	fmt.Printf("🌉  %s sucessfully generated.\n", bridgesTmplName)

	//
	// Build careers page
	//
	careersPage := filepath.Join(buildDir, careersTmplName)
	// Remove the old version
	os.Remove(careersPage)
	// Create new file
	f, err = os.Create(careersPage)
	if err != nil {
		log.Fatalf("can't create %s", careersTmplName)
	}
	careersTmpl.ExecuteTemplate(f, "base", CareersPage{SmallContainer: true, Openings: Openings})
	f.Close()
	fmt.Printf("👔  %s successfully generated.\n", careersTmplName)

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
			if strings.HasPrefix(r.URL.Path, "/clients") || strings.HasPrefix(r.URL.Path, "/grants") || strings.HasPrefix(r.URL.Path, "/team") {
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

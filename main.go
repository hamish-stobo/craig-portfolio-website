package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

type PageData struct {
	Title           string
	Name            string
	PersonalSummary string
	Qualifications  []string
	Services        []Service
	CurrentWork     []Position
	PreviousWork    []Position
	Articles        []Article
	PressItems      []PressItem
	Contact         Contact
}

type Service struct {
	Title string
	Image string
	Alt   string
}

type Position struct {
	Company string
	Title   string
	Logo    string
	URL     string
	Alt     string
}

type Article struct {
	Title  string
	Source string
	Date   string
	Image  string
	URL    string
	Alt    string
}

type PressItem struct {
	Title  string
	Source string
	Date   string
	Image  string
	URL    string
	Alt    string
}

type Contact struct {
	LinkedIn string
	Email    string
	Phone    string
}

func main() {
	r := mux.NewRouter()

	// Static file server for assets
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Main route
	r.HandleFunc("/", homeHandler).Methods("GET")

	// HTMX partial routes
	r.HandleFunc("/partials/about/{tab}", aboutHandler).Methods("GET")
	r.HandleFunc("/partials/work/{tab}", workHandler).Methods("GET")
	r.HandleFunc("/partials/articles/{tab}", articlesHandler).Methods("GET")

	// Build static site
	r.HandleFunc("/build", buildStaticSite).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := getPageData()

	// Create template with custom functions
	tmpl := template.New("index.html").Funcs(template.FuncMap{
		"isEven": func(i int) bool {
			return i%2 == 0
		},
		"contains": func(s, substr string) bool {
			return strings.Contains(s, substr)
		},
	})

	var err error
	tmpl, err = tmpl.ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err = tmpl.ParseGlob("templates/partials/*.html")
	if err != nil {
		log.Printf("Error parsing partial templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]

	data := getPageData()

	var tmplFile string
	switch tab {
	case "summary":
		tmplFile = "templates/partials/about-summary.html"
	case "qualifications":
		tmplFile = "templates/partials/about-qualifications.html"
	default:
		tmplFile = "templates/partials/about-summary.html"
	}

	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func workHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]

	data := getPageData()

	var tmplFile string
	switch tab {
	case "current":
		tmplFile = "templates/partials/work-current.html"
	case "previous":
		tmplFile = "templates/partials/work-previous.html"
	default:
		tmplFile = "templates/partials/work-current.html"
	}

	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]

	data := getPageData()

	var tmplFile string
	switch tab {
	case "articles":
		tmplFile = "templates/partials/articles-list.html"
	case "press":
		tmplFile = "templates/partials/press-list.html"
	case "poetry":
		tmplFile = "templates/partials/poetry-list.html"
	default:
		tmplFile = "templates/partials/articles-list.html"
	}

	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func buildStaticSite(w http.ResponseWriter, r *http.Request) {
	data := getPageData()

	// Create dist directory
	err := os.MkdirAll("dist", 0755)
	if err != nil {
		log.Printf("Error creating dist directory: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse and execute main template with custom functions
	tmpl := template.New("index.html").Funcs(template.FuncMap{
		"isEven": func(i int) bool {
			return i%2 == 0
		},
		"contains": func(s, substr string) bool {
			return strings.Contains(s, substr)
		},
	})

	tmpl, err = tmpl.ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err = tmpl.ParseGlob("templates/partials/*.html")
	if err != nil {
		log.Printf("Error parsing partial templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create index.html file
	file, err := os.Create("dist/index.html")
	if err != nil {
		log.Printf("Error creating index.html: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = tmpl.ExecuteTemplate(file, "index.html", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Copy static assets
	err = copyDir("static", "dist/static")
	if err != nil {
		log.Printf("Error copying static assets: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Static site built successfully in ./dist/")
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return copyFile(path, dstPath)
	})
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = srcFile.WriteTo(dstFile)
	return err
}

func getPageData() PageData {
	return PageData{
		Title: "Craig Stobo - Professional Director & Business Owner",
		Name:  "Craig Stobo",
		PersonalSummary: `I'm Craig Stobo. I am a professional director and business owner. I have previously worked 
		as a diplomat, economist, Chief Investment Officer and CEO. Outside of work I have climbed 
		mountains in NZ, Peru, Pakistan, and Nepal. I continue to ski, sail and compete in ocean swims. 
		I also enjoy drama and literature, and I am a long-time supporter of Otago rugby and the 
		Highlanders. Finally, I am Chair of the Waitaki Boys High School Foundation, and I'm a 
		Taumata of Otago University's Business School.`,
		Qualifications: []string{
			"BA (Hons) First Class Economics Otago",
			"Wharton Business School (Advanced Management Programme)",
			"Associate Member CFA Society New Zealand",
			"Chartered Fellow NZ Institute of Directors",
		},
		Services: []Service{
			{Title: "Business Advice", Image: "/static/images/services-photos/photo-1450101499163-c8848c66ca85.jpg", Alt: "Man signing paper"},
			{Title: "Investment Direction", Image: "/static/images/services-photos/photo-1454165804606-c3d57bc86b40.jpg", Alt: "Men discussing sheet of paper"},
			{Title: "Government Advice and Implementation", Image: "/static/images/services-photos/photo-1460925895917-afdab827c52f.jpg", Alt: "Laptop display graph with figures"},
			{Title: "Capital Raising", Image: "/static/images/services-photos/photo-1516321318423-f06f85e504b3.jpg", Alt: "People pointing at laptop screen"},
			{Title: "B2B Consulting", Image: "/static/images/services-photos/photo-1517048676732-d65bc937f952.jpg", Alt: "People sitting around table"},
			{Title: "Directing Companies", Image: "/static/images/services-photos/photo-1523875194681-bedd468c58bf.jpg", Alt: "Chess board"},
		},
		CurrentWork: []Position{
			{Company: "NZ Windfarms Ltd", Title: "Chairman", Logo: "/static/images/company-logos/current/nz_wind_farms.png", URL: "https://www.nzwindfarms.co.nz/", Alt: "NZ Windfarms Logo"},
			{Company: "Financial Markets Authority", Title: "Chairman", Logo: "/static/images/company-logos/current/fma_logo.png", URL: "https://www.fma.govt.nz/", Alt: "Financial Markets Authority Logo"},
			{Company: "Local Government Funding Agency", Title: "Chairman", Logo: "/static/images/company-logos/current/lgfa.png", URL: "https://www.lgfa.co.nz/", Alt: "Local Government Funding Agency Logo"},
			{Company: "Apello Services Ltd", Title: "Director", Logo: "/static/images/company-logos/current/apello-services-logo.png", URL: "https://appelloservices.co.nz/", Alt: "Apello Services Ltd Logo"},
			{Company: "Biomarine Ltd", Title: "Chairman and Shareholder", Logo: "/static/images/company-logos/current/biomarine-logo.png", URL: "http://www.biomarine.co.nz/", Alt: "Biomarine Ltd Logo"},
			{Company: "Cressida Capital", Title: "Board Advisor", Logo: "/static/images/company-logos/current/cressida-capital-logo.png", URL: "https://www.cressida.co.nz/", Alt: "Cressida Capital Logo"},
			{Company: "Urban Partners", Title: "Board Advisor", Logo: "/static/images/company-logos/current/urban-partners.png", URL: "http://www.urbanpartners.co.nz", Alt: "Urban Partners NZ Logo"},
			{Company: "Legend Terrace", Title: "Chairman and Shareholder", Logo: "/static/images/company-logos/current/legend-terrace-logo.png", URL: "", Alt: "Legend Terrace NZ Logo"},
		},
		PreviousWork: []Position{
			{Company: "Saturn Advice Ltd", Title: "Chairman and Shareholder", Logo: "/static/images/company-logos/current/Saturn-Advice-Logo_new 1.png", URL: "https://saturnadvice.co.nz/", Alt: "Saturn Advice Ltd Logo"},
			{Company: "Saturn Invest Ltd", Title: "Chairman and Shareholder", Logo: "/static/images/company-logos/current/saturn-invest-logo.png", URL: "https://www.pssnz.co.nz/", Alt: "Saturn Invest Ltd Logo"},
			{Company: "Elevation Capital", Title: "Chairman and Shareholder", Logo: "/static/images/company-logos/current/elevation-capital-logo.png", URL: "https://www.elevationcapital.co.nz/", Alt: "Elevation Capital Logo"},
			{Company: "AIG Insurance New Zealand", Title: "Chairman", Logo: "/static/images/company-logos/current/AIG-logo.png", URL: "https://www.aig.co.nz/", Alt: "AIG Insurance New Zealand Logo"},
			{Company: "Precinct Properties Ltd", Title: "Chairman", Logo: "/static/images/company-logos/current/precinct-properties-logo.png", URL: "https://www.precinct.co.nz/", Alt: "Precinct Properties Ltd Logo"},
			{Company: "Fliway Group", Title: "Chairman", Logo: "/static/images/company-logos/previous/fliway-group-logo.png", URL: "https://www.fliway.com/", Alt: "Fliway Group Logo"},
			{Company: "Bureau Ltd", Title: "Chairman", Logo: "/static/images/company-logos/previous/bureau-logo.png", URL: "https://www.bureau.co.nz/", Alt: "Bureau Ltd Logo"},
			{Company: "OCG Consulting Ltd", Title: "Chairman", Logo: "/static/images/company-logos/previous/ocg-logo.png", URL: "https://www.ocg.co.nz/", Alt: "OCG Consulting Ltd Logo"},
			{Company: "Industrial Research Ltd", Title: "Director", Logo: "/static/images/company-logos/previous/callaghan-innovation-logo.png", URL: "https://www.callaghaninnovation.govt.nz/", Alt: "Industrial Research Ltd Logo"},
			{Company: "New Zealand Green Investment Finance", Title: "Advisor", Logo: "/static/images/company-logos/previous/NZGIF-logo.png", URL: "https://nzgif.co.nz/", Alt: "NZGIF Logo"},
			{Company: "New Zealand Treasury", Title: "Advisor", Logo: "/static/images/company-logos/previous/treasury-logo.png", URL: "https://treasury.govt.nz/", Alt: "NZ Treasury Logo"},
			{Company: "Local Government New Zealand", Title: "Advisor", Logo: "/static/images/company-logos/previous/lgnz-logo.png", URL: "https://www.lgnz.co.nz/", Alt: "Local Government New Zealand Logo"},
		},
		Articles: []Article{
			{Title: "NZ Local Government Insurance Review", Source: "Local Government New Zealand", Date: "Dec 2013", Image: "/static/images/publications-media/articles/thumbnail_local_government_insurance_market_review.png", URL: "/static/docs/articles/LGNZ140101 Insurance Review Publication FA-4.pdf", Alt: "Thumbnail of Document"},
			{Title: "Exporting Financial Services", Source: "International Financial Services Development Group", Date: "July 2011", Image: "/static/images/publications-media/articles/thumbnail_exporting_financial_services.png", URL: "/static/docs/articles/IFSD Group Presentation_July2011.pdf", Alt: "Thumbnail of Document"},
			{Title: "Towards Consensus on the Taxation of Investment Income", Source: "Inland Revenue Department", Date: "October 2004", Image: "/static/images/publications-media/articles/thumbnail_consensus_taxation_investment_income.png", URL: "/static/docs/articles/2004-other-report-stobo-taxation-investment-income.pdf", Alt: "Thumbnail of Document"},
		},
		PressItems: []PressItem{
			{Title: "Shovel-ready, but not work-ready", Source: "NZ Herald", Date: "9 May 2020", Image: "/static/images/publications-media/media/nz-herald-2020.jpg", URL: "https://www.nzherald.co.nz/business/news/article.cfm?c_id=3&objectid=12330405", Alt: "NZ Herald article thumbnail"},
			{Title: "Q & A with Craig Stobo", Source: "University of Otago Business School", Date: "20 Oct 2017", Image: "/static/images/publications-media/media/UoO-2017-business-school-interview.png", URL: "https://www.youtube.com/watch?v=fDo0HFOe8EU", Alt: "University of Otago interview thumbnail"},
			{Title: "Government assets sheet needs right balance", Source: "NZ Herald", Date: "12 Mar 2015", Image: "/static/images/publications-media/media/nz-herald-2015.jpg", URL: "https://www.nzherald.co.nz/business/news/article.cfm?c_id=3&objectid=11415720", Alt: "NZ Herald article thumbnail"},
		},
		Contact: Contact{
			LinkedIn: "https://nz.linkedin.com/in/craig-stobo-4b40361a2",
			Email:    "mailto:stobo@xtra.co.nz",
			Phone:    "tel:+6421733751",
		},
	}
}

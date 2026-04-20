package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
  "os"
)

type Project struct {
	Title       string
	Description string
	Tags        []string
	Link        string
	Year        string
}

type Experience struct {
	Role        string
	Company     string
	Period      string
	Description string
	Tech        []string
}

type Skill struct {
	Category string
	Items    []string
}

type Portfolio struct {
	Name        string
	Role        string
	Tagline     string
	About       string
	Email       string
	GitHub      string
	LinkedIn    string
	Skills      []Skill
	Experiences []Experience
	Projects    []Project
}

var portfolioData = Portfolio{
	Name:     "Hanshika Sahu",
	Role:     "Web Developer",
	Tagline:  "Crafting digital experiences with clean code & bold design.",
	About:    "I am a passionate web developer with 2+ years of experience building scalable, performant applications. I thrive at the intersection of design and engineering — turning complex problems into elegant, user-friendly solutions. When I am not coding, I am exploring open-source projects and contributing to the dev community.",
	Email:    "hanshikasahu03@gmail.com",
	GitHub:   "https://github.com/HanshikaSahu",
	LinkedIn: "https://www.linkedin.com/in/hanshika-sahu-120b9b301/",
	Skills: []Skill{
		{Category: "Frontend", Items: []string{"HTML5", "CSS3", "JavaScript", "TypeScript", "ReactJS", "Vue.js", "Tailwind CSS"}},
		{Category: "Backend", Items: []string{"Go", "Node.js", "Express.js", "REST APIs", "Django", "Next.js"}},
		{Category: "Tools", Items: []string{"Git", "Vercel", "Render", "VS Code", "Postman", "Supabase", "Firebase"}},
	},
	Experiences: []Experience{
		{
			Role:        "Web Developer",
			Company:     "Paycasso",
			Period:      "June 2025 - Aug 2025",
			Description: "Built features for a secure fintech messaging platform using React, TypeScript, Supabase, and WebSockets, enabling real-time communication and media sharing.",
			Tech:        []string{"Supabase", "Typescript", "Next.js", "WebSocket", "Tailwind CSS"},
		},
		{
			Role:        "Full Stack Developer",
			Company:     "International Institute of SDG's and Public Policy",
			Period:      "July 2025 - Sep 2025",
			Description: "Developed responsive policy research platforms using React.js, improving page performance and content rendering.",
			Tech:        []string{"React.js", "Node.js", "REST APIs", "MongoDB"},
		},
		{
			Role:        "Backend Developer",
			Company:     "First Contact",
			Period:      "July 2025 - Aug 2025",
			Description: "Developed backend functionalities at an LGBTQ+ focused organization using Wix Velo and JavaScript. Built dynamic, data-driven features, managed database interactions, and supported scalable web application performance.",
			Tech:        []string{"WIx", "JavaScript", "Wix Website Builder"},
		},
	},
	Projects: []Project{
		{
			Title:       "ContactBridge",
			Description: "A simple contact website built with Firebase for real-time form submissions.",
			Tags:        []string{"Javascript", "React", "FIrebase"},
			Link:        "https://github.com/HanshikaSahu/ContactBridge",
			Year:        "May-2025",
		},
		{
			Title:       "Gemini Task-Manager",
			Description: "A task management system with secure APIs, structured data handling, and AI-powered task assistance.",
			Tags:        []string{"Next.js", "Typescript", "Drizzle ORM", "Hono", "Clerk", "PostgreSQL", "Gemini API","Docker"},
			Link:        "Frontend on Vercel, Backend on Render for submission purposes",
			Year:        "June 2025",
		},
		{
			Title:       "Cognify",
			Description: "A structured approach to learning, powered by intelligent automation",
			Tags:        []string{"React.js", "Next.js", "Neon", "Razorpay", "Clerk"},
			Link:        "cognify-pi.vercel.app",
			Year:        "Dec 2025",
		},
		{
			Title:       "Distributed-Cache-Backend",
			Description: "Fault-tolerant distributed in-memory cache built with Python and FastAPI, featuring LRU eviction, TTL expiration, replication, metrics, and Dockerized deployment.",
			Tags:        []string{"JavaScript", "Python", "Docker", "FastAPI + Uvicorn"},
			Link:        "https://github.com/HanshikaSahu/Distributed-Cache-Backend",
			Year:        "Jan 2026",
		},
	},
}



func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.Execute(w, portfolioData); err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
	}
}


func main() {
	http.HandleFunc("/", homeHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	fmt.Printf("Portfolio running on port %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
/*
	This file contains the TUI code for Chaker.
*/

package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/olekukonko/ts"
)

// The search bar
var searchBarContent string = ""

// It's search bar time!
var itsSearchBarTime bool = false

// Help line
var help string = lipgloss.NewStyle().
	Faint(true).
	Bold(true).
	Render("↑ - up · ↓ - down · q - quit · ⏎  - open · c - comment section · m - more · p - prev")

// Terminal size
var size, _ = ts.GetSize()

// Style
var style = lipgloss.NewStyle().
	PaddingLeft(4).
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#ff6600")).
	Align(lipgloss.Center).
	Width(size.Col())

// Shortcuts
var render = style.Render

// The Model struct
type Model struct {
	cursor   int    // Which submission is our cursor pointing at
	selected string // Which submission is selected
}

// Open the browser with the URL
func openBrowserWithURL(url string) {
	// To make sure that this works on other platforms, we need to use different commands
	var browserCmd string

	// Detect the OS
	os := runtime.GOOS
	switch os {
	case "windows":
		browserCmd = "explorer"
	case "darwin":
		browserCmd = "open" // Darwin aka Mac
	case "linux":
		browserCmd = "browse"
	default:
		panic("Unknown OS: " + os)
	}

	// Open the default browser with the URL
	cmd := exec.Command(browserCmd, url)
	_, err := cmd.Output()

	// Report the error!
	if err != nil {
		log.Fatal(err)
	}
}

// Make a custom title (and extra information for fainting effect) from the submission
// This is meant to be in main.go, but for the fainting effect, it has to be move here.
func returnCustomTitle(submission Submission) (string, string) {
	// The submission time (Idk how to do something like Hacker New's one)
	submissionTime := time.Unix(int64(submission.Time), 0).Format("15:04 PM")

	if submission.Type == "job" {
		// If the submission is a 'job', then we don't need to print unnecessary information,
		// we will just show the title, how old is it and the URL
		return submission.Title, spf("(%s)", submissionTime)
	}

	return submission.Title, spf(
		"(%d points by %s | %s | %d comments)",
		submission.Score,
		submission.By,
		submissionTime,
		submission.Descendants,
	)
}

// The main function
func tui() {
	// Initial model
	var initialModel Model = Model{
		cursor:   1,
		selected: "",
	}

	p := tea.NewProgram(
		initialModel,
		tea.WithAltScreen(),
	)
	if err := p.Start(); err != nil {
		log.Fatalf("We got an error! %v", err)
		os.Exit(1)
	}
}

// Initialize the app
func (m Model) Init() tea.Cmd {
	return nil
}

// Update the app
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !itsSearchBarTime {
		switch msg := msg.(type) {
		// Handle keyboard event
		case tea.KeyMsg:
			switch msg.String() {
			// It is 'q', quit the program!
			case "q":
				return m, tea.Quit

			// An Up? Fly up!
			case "up":
				if m.cursor > 1 {
					m.cursor--
				}

			// A Down? Fall down!
			case "down":
				if m.cursor < len(submissions)-1 {
					m.cursor++
				}

			// Enter? Enter the web!
			case "enter":
				openBrowserWithURL(submissions[m.cursor].URL)

			// 'c'? Open the comment section of the title in the cursor!
			case "c":
				openBrowserWithURL(spf("https://news.ycombinator.com/item?id=%d", submissions[m.cursor].ID))

			// 'm'? Next page, please!
			case "m":
				pageNum++
				submissions = []Submission{}
				submissions = Scrape(pageNum) // Scrape fresh data

			// Same as 'm', but previous page, and also checks if pageNum is larger than 1 to prevent go to page 0
			case "p":
				if pageNum > 1 {
					pageNum--
					submissions = []Submission{}
					submissions = Scrape(pageNum) // Scrape fresh data
				}

			// The 's'? It's seaarch bar time!
			case "s":
				itsSearchBarTime = !itsSearchBarTime
			}
		}
	} else {
		// This is the second input, in case the program is in the search bar time
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "s":
				itsSearchBarTime = !itsSearchBarTime
			}
		}
	}

	return m, nil
}

// The UI of the app
func (m Model) View() string {
	// Header
	s := render(spf("Today is %s\n", time.Now().Format("Monday, January 2, 2006, at 15:04 PM")))

	for i := range submissions {
		// Is the cursor pointing at this title?
		cursor := " " // No cursor
		if m.cursor == i {
			cursor = ">" // Yes cursor!
		}

		// Render the row
		title, extraInfo := returnCustomTitle(submissions[i]) // Get the title and extra info
		urlHost := parseURLHost(submissions[i].URL)           // Get the host of the URL

		// If the cursor is not pointing to this title, we won't need the extraInfo
		if m.cursor != i {
			extraInfo = ""
		} else {
			extraInfo = lipgloss.NewStyle().Faint(true).Render(spf("%s %s", urlHost, extraInfo))
		}
		s += spf("%s %s %s\n", cursor, title, extraInfo)
	}
	// Footer (the page where the user are in and help)
	s += spf("You are at page %d %s", pageNum, help)

	// Send the UI for rendering
	return s
}

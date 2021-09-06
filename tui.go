/*
	This file contains the TUI code for Hecker.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/olekukonko/ts"
)

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

// The Model
type Model struct {
	submissions   []Submission  // List of submission
	cursor   			int      			// Which submission is our cursor pointing at
	selected 			string   			// Which submission is selected
}

// Open the browse with the URL
func open_browser_with_url(url string) {
	// Trigger xdg-open with the URL
	cmd := exec.Command("browse", url)
	_, err := cmd.Output()

	// Report the error!
	if err != nil {
		log.Fatal(err)
	}
}

// Make a custom title (and extra information for fainting effect) from the submission
func return_custom_title(submission Submission) (string, string) {
	// The submission time (Idk how to do something like Hacker New's one)
	submission_time := time.Unix(time.Now().Unix() - int64(submission.Time), 0).Format("15:04 PM")

	if submission.Type == "job" {
		// If the submission is a 'job', then we don't need to print unnecessary information,
		// we will just show the title, how old is it and the URL
		return submission.Title, spf("(%s)", submission_time)
	} else {
		return submission.Title, spf(
			"(%d points by %s | %s | %d comments)",
			submission.Score,
			submission.By,	
			submission_time,
			submission.Descendants,
		)
	}
}

// The main function
func tui(submissions []Submission) {
	initialModel := Model {
		submissions: submissions,
		selected: "",
	}

	p := tea.NewProgram(
		initialModel,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)
	if err := p.Start(); err != nil {
		log.Fatalf("We got an error! %v", err)
		os.Exit(1)
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Handle keyboard event
	case tea.KeyMsg:
		switch msg.String() {
		// It is 'q', quit the program!
		case "q":
			return m, tea.Quit

		// An Up? Fly up!
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		// A Down? Fall down!
		case "down":
			if m.cursor < len(m.submissions) - 1 {
				m.cursor++
			}

		// Enter? Enter the web!
		case "enter":
			// Take the URL
			open_browser_with_url(m.submissions[m.cursor].URL)
		
		// 'c'? Open the comment section
		case "c":
			open_browser_with_url(spf("https://news.ycombinator.com/item?id=%d", m.submissions[m.cursor].ID))
		}
	}

	return m, nil
}

func (m Model) View() string {
	// Header
	s := render(spf("Today is %s\n", time.Now().Format("Monday, January 2, 2006, at 15:04 PM")))
	s += "\nTop Submissions:\n"

	for i := range m.submissions {
		// Is the cursor pointing at this title?
		cursor := " " // No cursor
		if m.cursor == i {
			cursor = "⏵ " // Yes cursor!
		}

		// Render the row
		title, extra_info := return_custom_title(submissions[i])
		extra_info = lipgloss.NewStyle().Faint(true).Render(" " + extra_info)
		s += fmt.Sprintf("%s %s\n", cursor, title + extra_info)
	}

	// Footer (basically the help part)
	s += lipgloss.NewStyle().
		Faint(true).
		Bold(true).
		Render("\n↑ - up · ↓ - down · q - quit · ⏎  - open · c - comment section")

	// Send the UI for rendering
	return s
}

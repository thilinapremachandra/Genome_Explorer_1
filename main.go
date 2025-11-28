package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	sequence string
	cursor   int
	width    int
	height   int
	stats    Stats
}

type Stats struct {
	Length int
	GC     float64
	A      int
	T      int
	G      int
	C      int
}

func initialModel(seq string) Model {
	return Model{
		sequence: seq,
		cursor:   0,
		stats:    calculateStats(seq),
	}
}

// --------- Bubble Tea Required Methods ----------

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "down", "j":
			if m.cursor+1 < len(m.sequence) {
				m.cursor++
			}

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		}
	}
	return m, nil
}

func (m Model) View() string {

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Render("🧬 Genome Explorer\n")

	stats := fmt.Sprintf(
		"Length: %d | GC: %.2f%% | A:%d T:%d G:%d C:%d\n",
		m.stats.Length, m.stats.GC,
		m.stats.A, m.stats.T, m.stats.G, m.stats.C,
	)

	seqView := renderSequence(m.sequence, m.cursor)

	footer := "\n↑ ↓ to scroll | q to quit"

	return title + stats + "\n" + seqView + footer
}

// --------- Helper Functions ----------

func renderSequence(seq string, cursor int) string {
	var out strings.Builder

	styles := map[rune]lipgloss.Style{
		'A': lipgloss.NewStyle().Foreground(lipgloss.Color("82")),
		'T': lipgloss.NewStyle().Foreground(lipgloss.Color("45")),
		'G': lipgloss.NewStyle().Foreground(lipgloss.Color("214")),
		'C': lipgloss.NewStyle().Foreground(lipgloss.Color("199")),
	}

	start := cursor
	end := cursor + 200
	if end > len(seq) {
		end = len(seq)
	}

	for _, c := range seq[start:end] {
		if style, ok := styles[c]; ok {
			out.WriteString(style.Render(string(c)))
		} else {
			out.WriteRune(c)
		}
	}

	return out.String()
}

func calculateStats(seq string) Stats {
	var s Stats
	s.Length = len(seq)

	for _, c := range seq {
		switch c {
		case 'A':
			s.A++
		case 'T':
			s.T++
		case 'G':
			s.G++
		case 'C':
			s.C++
		}
	}

	gc := s.G + s.C
	s.GC = float64(gc) / float64(s.Length) * 100

	return s
}

func loadFASTA(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic("failed to open FASTA file")
	}
	defer file.Close()

	var seq strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, ">") {
			continue
		}
		seq.WriteString(strings.ToUpper(line))
	}
	return seq.String()
}

// --------- Main ----------

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go sample.fasta")
		return
	}

	sequence := loadFASTA(os.Args[1])
	p := tea.NewProgram(initialModel(sequence))
	_ = p.Start()
}

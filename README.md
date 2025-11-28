# 🧬 Genome Explorer

Genome Explorer is a terminal-based interactive tool built with **Go** and **Bubble Tea**.  
It allows users to explore DNA sequences from FASTA files in a clean, readable interface.  
This project is an early version and serves as a hands-on introduction to building bioinformatics tools with modern terminal UI libraries.

---

## Features (v1)

- **FASTA File Handling**
  - Load DNA sequences from FASTA files
  - Interactive scrolling through sequences inside the terminal

- **Real-time Sequence Statistics**
  - Sequence length
  - GC content (%)
  - Base counts (A, T, G, C)

- **Visual Enhancements**
  - Color-coded nucleotides for better readability

---

## 📂 Project Structure

- `main.go` → Application entry point  
- `go.mod` → Go module dependencies  
- `sample.fasta` → Example DNA sequence file  
- `README.md` → Documentation  

---

## How It Works

- Reads a FASTA file and extracts the DNA sequence  
- Displays the sequence in an interactive terminal view  
- Allows scrolling through the sequence using keyboard controls  
- Calculates and displays sequence statistics in real time  

---

## Installation

- **Prerequisites**
  - Go 1.18 or newer

- **Steps**
  - Clone the repository:
    ```bash
    git clone <your-repository-url>
    cd genome-explorer
    ```
  - Install dependencies:
    ```bash
    go get github.com/charmbracelet/bubbletea
    go get github.com/charmbracelet/lipgloss
    ```

---

## Usage

- Run the application:
  ```bash
  go run main.go sample.fasta


### Controls
- ↑ / k	Scroll up
- ↓ / j	Scroll down
- q	Quit
 
### 📄 Example FASTA File
>Example DNA Sequence
ATGCGGCTTAGCTAGCTAGCTAGGCTTAGCGATCGATCGATCGTAGCTAGCTA

## Motivation

- Most bioinformatics tools rely on raw command-line outputs or heavy graphical software.
- Genome Explorer explores a middle ground: clean, interactive terminal-based exploration of genomic data.

🔮 Future Improvements

- Motif search and highlighting
- Codon and ORF visualization
- Support for multiple sequences
- Better layout and navigation

## License

- MIT License

# Greppo

**Greppo** is a simple and efficient command-line tool written in Go that searches for text patterns in files and directories — similar to Unix `grep`, but with a minimal and clean Go implementation.

Use it for quick pattern matching, recursive search, and filtering lines in files with useful flags like case-insensitivity, line counting, and more.

---

## 🚀 Features

- ✅ Case-insensitive matching (`-i`)
- ✅ Inverted match (exclude pattern) (`-v`)
- ✅ Count of matching lines (`-c`)
- ✅ Display line numbers (`-n`)
- ✅ Recursive search in directories (`-r`)
---

## 🧑‍💻 Planned features

- 🎨 Colored output: Highlight matching terms in output
- 🔍 Regex support: Use regular expressions for advanced matching
- ⚡ Parallel file traversal: Speed up recursive search on large directories

Have ideas? Feel free to suggest or contribute!

## 📦 Installation

You’ll need [Go 1.20+](https://golang.org/dl/) installed.

```bash
git clone https://github.com/alexandru356/greppo.git
cd greppo
go build -o greppo ./cmd

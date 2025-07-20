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

## 📦 Installation

You’ll need [Go 1.20+](https://golang.org/dl/) installed.

```bash
git clone https://github.com/alexandru356/greppo.git
cd greppo
go build -o greppo ./cmd

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)


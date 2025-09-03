# lzmk

`lzmk` is a compiler for **Lazymark** — a minimal, Markdown-inspired grammar for writing blogs.  
It parses Lazymark source files and compiles them into **HTML** or **React components**, making it easy to publish structured posts without the overhead of full Markdown or a CMS.

---

## Features

- **Simple grammar**: Titles, headers, paragraphs, images, and metadata.
- **Two targets**: Output as plain HTML or as React components.
- **CLI tool**: Write, compile, and publish in seconds.
- **Minimalism first**: Lazymark strips down to what’s essential for blogs.

---

## Installation

Clone the repo and build with Go:

```bash
git clone https://github.com/yourname/lzmk.git
cd lzmk
go build -o lzmk ./cmd/lzmk
```

Or install directly:

```bash
go install github.com/yourname/lzmk/cmd/lzmk@latest
```

---

## Usage

```bash
lzmk [options] file...
```

### Options

- `-o, --output=FILE` → Write output to file (default: stdout).  
- `-t, --target=TYPE` → Output type: `html` (default) or `react`.  
- `-w, --watch` → Watch files and recompile on change.  
- `-v, --version` → Show version info.  
- `-h, --help` → Display usage help.  

---

## Example

**Source file (`post.lzmk`):**
```text
# My First Post

## Introduction
This is a Lazymark paragraph.

## An Image
#!(cat.jpg){My cat}
```

**Compile to HTML:**
```bash
lzmk post.lzmk -o post.html
```

**Compile to React:**
```bash
lzmk -t react post.lzmk > Post.jsx
```

---

## Lazymark Grammar

- `# Title` → Document title  
- `## Header` → Section headers  
- Paragraphs → plain text lines  
- `#!(url){alt}` → Images with alt text  
- `key value` → Metadata entries  

---

## Roadmap

- [ ] Syntax highlighting for Lazymark  
- [ ] Plugin system for custom output  
- [ ] Static site generator mode (`lzmk build`)  
- [ ] VSCode extension for Lazymark syntax  

---

## Development

This project follows the usual Go project layout:

- `cmd/lzmk` → CLI entrypoint.  
- `pkg/parser` → Lazymark lexer/parser.  
- `pkg/compiler` → Code generation for HTML/React.  

Run tests:

```bash
go test ./...
```

---

## License

MIT License. See [LICENSE](LICENSE) for details.

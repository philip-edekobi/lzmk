LZMK(1)                 User Commands                 LZMK(1)

NAME
       lzmk - compile Lazymark documents into HTML or React

SYNOPSIS
       lzmk [options] file...

DESCRIPTION
       lzmk is a compiler for Lazymark, a minimal markdown-like
       grammar designed for blogging.

       It parses Lazymark source files and outputs HTML or React
       components suitable for static sites, blogs, or rendering
       pipelines.

OPTIONS
       -o, --output=FILE
              Write output to FILE instead of stdout.

       -t, --target=TYPE
              Specify output target: html (default) or react.

       -w, --watch
              Watch files and recompile on change.

       -v, --version
              Print version info.

       -h, --help
              Show this help message.

EXAMPLES
       lzmk post.lzmk -o post.html
              Compile Lazymark to HTML.

       lzmk -t react post.lzmk > Post.jsx
              Compile Lazymark to a React component.

SEE ALSO
       markdown(1), pandoc(1)

lzmk 0.1.0                August 2025                 LZMK(1)

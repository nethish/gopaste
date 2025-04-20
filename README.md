# gopaste - A utility to paste clipboard contents to a file

## Installation
```bash
go install github.com/nethish/gopaste@v0.0.1
```

## Go Paste
* `gopaste file.txt` will create a new file `file.txt` and paste the contents of the clipboard to it
* If the file already exists it will NOT over write it and instead report it to you

## Motivation
I do lot of experiments and often copy paste code from stack overflow or ChatGPT. Every time I start a new project, create files and paste contents to it.
It was slow and annoying (well for me it was). 
I could've used `pbpaste` or `xclip` but it's fun to use your own tool :evil: (and reinvent the wheel)


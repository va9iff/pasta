# Pasta
## Paste text with markdown

Needed a simple recursive text composition tool with no destructive syntax.
Decided to use markdown links with `pasta: ` prefix in their title to denote
embedding the document directly instead of just having a link to a document.
Right now Pasta only works with relative file pathes, not urls.

## Cool things about Pasta

1. Recursive nature allows us to have flexible composition. `a.md` embeds
   `b.md` and `b.md` embeds `c.md`, then the generated output would include
`a.md`, `b.md` and `c.md` all in the correct places

2. Pasta links are valid markdown links and would have a working link to the
   embedded documents in markdown editors even without Pasta

3. Written in Go, can easily be a part of any build process

4. Doesn't care about filetypes, Markdown is just the syntax it uses. Plain
   text, HTML, CSV or code, anything that requires pasting text from another
file, Pasta got you

## Examples

This link will be replaced with `./subdoc.md`'s content in the generated
output. You have to put the link in a standalone line with no other text to get
it embedded. Don't forget to cd into your generated document's folder before
running pasta.


```md

[pasta: my sub document](./subdoc.md)


```

Not all links are embedded. Links with no `pasta: ` prefix stays untouched

```md

[this link is not embedded](./pasta.go)

```

# Usage

```sh 

pasta input.md

pasta input.md output.md

```

Even if you decide to leave Pasta completely, your markdown will still be
intact. I don't like vendorlock

# Unicode in Go

A simple example that shows common operations on Unicode characters in Go.

Unicode can be complex, but it isn't rocket science. Here's a short [primer](../../docs/unicode.md) to understand the basics of Unicode.

In Go, use escape `\u` followed by a 4 digit unicode code point for basic unicode code point. And use escape `\U` followed by an 8 digit unicode for supplementary code point.

## Setup

1. Run the program.

   ```bash
   $ make run
   ```

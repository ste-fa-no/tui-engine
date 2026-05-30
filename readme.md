## Introduction

A small toy graphics engine for text-based user interface applications.

> Built for fun and the curiosity of understanding how modern TUIs work under the hood.

### Think this project sucks? Should I spend my time better?

[Hire me! 😛]()

### Features
- [x] Cross-platform (Unix + Windows)
- [x] Raw mode + alternate screen
- [x] Event loop with goroutines and channels
- [x] ANSI escape sequence parser
- [x] Support for terminal resize
- [x] Differential rendering
- [x] True Color
- [x] Text formatting
- [x] Composable and declarative widgets
- [x] Focus system
- [x] Layout constraints
- [x] Structured application state management
- [ ] Decorations for core widgets
- [ ] Mouse support

### Available widgets
- [x] Horizontal stacks
- [x] Vertical stacks
- [x] Text blocks
- [x] Lists
- [x] Text fields (single line with horizontal scrolling)
- [x] Text areas (multiple lines with word-wrap and vertical scrolling)
- [x] Scrollbar decorator for vertical-scrolling widgets
- [x] Border decorator for all widgets (with support for titles)
- [ ] Buttons
- [ ] Checkboxes and radio buttons
- [ ] Progress bars
- [ ] Modals

Philosophy

This project does not aim to reinvent the terminal.
My goal was to build a small engine that is readable and reasonably simple to understand, in order to encourage:
- experimentation
- learning terminal rendering
- rapid development of proof of concepts (and maybe something more!)
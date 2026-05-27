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
- [x] Resize events
- [x] Differential rendering
- [x] True Color
- [x] Text formatting
- [x] Composable and declarative widgets
- [x] Context with relative coordinates
- [x] Focus system
- [x] List widget with scrolling
- [x] TextField with horizontal scrolling and placeholder
- [x] Multiline TextArea with word wrap and placeholder
- [x] Visual scrollbar for List
- [ ] Decorations for core widgets
- [ ] Layout constraints
- [ ] Mouse support

Philosophy

This project does not aim to reinvent the terminal.
My goal was to build a small engine that is readable and reasonably simple to understand, in order to encourage:
- experimentation
- learning terminal rendering
- rapid development of proof of concepts (and maybe something more!)
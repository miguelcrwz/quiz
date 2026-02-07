<div align="center">

<pre>
 ┌────────────────────────────────────────────────────┐
 │  user@dev:~/projects$ ./quiz                       │
 │                                                    │
 │         ██████╗   ██╗   ██╗  ██╗  ███████╗         │
 │        ██╔═══██╗  ██║   ██║  ██║  ╚══███╔╝         │
 │        ██║   ██║  ██║   ██║  ██║    ███╔╝          │
 │        ██║▄▄ ██║  ██║   ██║  ██║   ███╔╝           │
 │        ╚██████╔╝  ╚██████╔╝  ██║  ███████╗         │
 │         ╚══▀▀═╝    ╚═════╝   ╚═╝  ╚══════╝         │
 │                                                    │
 │  > formula 1 trivia quiz · go 1.23 stdlib          │
 └────────────────────────────────────────────────────┘
</pre>

![type](https://img.shields.io/badge/type-cli%20game-2ea043?style=for-the-badge&labelColor=0d1117)
![version](https://img.shields.io/badge/version-1.0.0-58a6ff?style=for-the-badge&labelColor=0d1117)
![license](https://img.shields.io/badge/license-mit-8b949e?style=for-the-badge&labelColor=0d1117)
![stack](https://img.shields.io/badge/stack-go%201.23-d29922?style=for-the-badge&labelColor=0d1117)

</div>

---

## `$ cat about.txt`

```
quiz is a formula 1 trivia game for the terminal: it asks your name,
goes through ten multiple-choice questions and prints your score and
how long you took.

the questions are data, not code — encoding/csv reads quiz.csv in one
ReadAll, skips the header row and turns every line into a Questions
value (text, four options, number of the right one). rand.Shuffle
reorders them on each run, and an answer is asked again until it is a
number between 1 and 4, so a stray letter or a 7 never costs a point.

built as a hands-on study of go — structs with pointer-receiver
methods, csv parsing and reading stdin with bufio.
```

## `$ ls features/`

```
[x] 10 formula 1 questions, 4 options each, loaded from quiz.csv
[x] a new question order on every run (rand.Shuffle)
[x] non-numeric or out-of-range answers are asked again, not scored
[x] instant feedback — a miss shows the right option
[x] final score and total time in seconds
[x] colored prompts with ansi escape codes
[ ] quiz.csv path is hardcoded — run from the repo root or it panics
[ ] eof on stdin (ctrl+d / ctrl+z) panics instead of quitting
[ ] piped input panics at question 1 — Init's bufio.Reader buffers it all
[ ] a row under 6 fields, or not as many as the header, panics at startup
[ ] an Answer that is not 1-4 panics once that question is answered
[ ] the clock counts two 350ms pauses per question — 7s of every run
[ ] only the questions are shuffled, never the options
[ ] score is not shown out of the total, and nothing is saved
[ ] no newline after the score — the shell prompt lands on the same line
[ ] module is github.com/miguelcrwz, so go build names it miguelcrwz
[ ] no tests
```

## `$ cat stack.lock`

```
runtime     go 1.23.3 (go.mod), tested with go 1.25.4
framework   none — stdlib only (bufio, encoding/csv, math/rand)
data        quiz.csv, header + 10 rows of 6 columns
ui          terminal, ansi color escape codes
infra       none, runs locally
tests       none yet
tooling     go run / go build
```

## `$ ./install.sh`

```bash
# clone
git clone https://github.com/miguelcrwz/quiz.git
cd quiz

# dependencies (stdlib only — just needs go 1.23.3+)
go version
```

## `$ ./run`

```bash
# development (from the repo root — quiz.csv is opened from the cwd)
go run .

# production — pass -o, the module path would name the binary miguelcrwz
go build -o quiz . && ./quiz
```

```
WELCOME TO F1 QUIZ! ⚒︎
 Enter your NAME: miguel
Hello, miguel! LET'S START!

 1. Who is the youngest driver to win an F1 race?
[1] Max Verstappen
[2] Lewis Hamilton
[3] Fernando Alonso
[4] Sebastian Vettel
 Enter an ALTERNATIVE: 1

✔ CORRECT ANSWER

 2. What is the name of the famous F1 circuit in Brazil?
[1] Vegas Circuit
[2] Silverstone Circuit
[3] Interlagos Circuit
[4] Suzuka Circuit
 Enter an ALTERNATIVE: abc

✘ INVALID VALUE, PLEASE SELECT A VALID OPTION.

 Enter an ALTERNATIVE: 7

✘ INVALID OPTION, PLEASE SELECT A VALID OPTION.

 Enter an ALTERNATIVE: 4

✘ INCORRECT ANSWER / RIGHT ANSWER: [3] Interlagos Circuit

 ...

YOU SCORED 9 POINTS IN 14.22 SECONDS!
```

## `$ describe quiz.csv`

```
Questions   string   question text — row 1 is the header, skipped
Option 1    string   shown as [1]
Option 2    string   shown as [2]
Option 3    string   shown as [3]
Option 4    string   shown as [4]
Answer      int      1-4, number of the right option, not checked on load
```

## `$ cat game.flow`

```
start
  │
  ├─ ProcessCSV
  │    ├─ no quiz.csv ───────────────────────> panic
  │    ├─ field count differs from header ───> panic
  │    ├─ every row under 6 fields ──────────> panic (record[5])
  │    └─ ok ─> skip header ─> []Questions ─> rand.Shuffle
  │
  ├─ Init: read the name
  │    └─ eof ───────────────────────────────> panic
  │
  ├─ Run: for each question, read a line
  │    ├─ eof ───────────────────────────────> panic (slice [:-1])
  │    ├─ not a number ──────────────────────> invalid value, ask again
  │    ├─ outside 1-4 ───────────────────────> invalid option, ask again
  │    ├─ == Answer ─────────────────────────> correct, points++
  │    ├─ Answer not 1-4 ────────────────────> panic (Options[Answer-1])
  │    └─ otherwise ─────────────────────────> incorrect + right option
  │
  └─ points + seconds since Run started ─────> stdout
```

## `$ tree`

```
quiz/
├── main.go          # Questions, GameState — load, shuffle, ask, score
├── quiz.csv         # the 10 formula 1 questions, header row first
├── preview/
│   └── gopher.gif   # the gopher that celebrates exit 0
├── go.mod           # module github.com/miguelcrwz, go 1.23.3
├── .gitignore
├── LICENSE
└── README.md
```

## `$ git log --oneline roadmap`

```
a1b2c3d  share one bufio.Reader for stdin and exit cleanly on eof
e4f5g6h  validate quiz.csv on load instead of panicking mid-game
i7j8k9l  rename the module to github.com/miguelcrwz/quiz
m0n1o2p  show the score out of the total, stop the clock during pauses
q3r4s5t  table-driven tests for toInt and the csv loader
```

## `$ cat LICENSE | head -1`

```
MIT © 2026 Miguel Cruz
```

## `$ whoami`

```
miguel cruz
├── github    https://github.com/miguelcrwz
├── linkedin  https://linkedin.com/in/miguelcrwz
└── email     miguwlcrwz@gmail.com
```

<div align="center">

<pre>
$ exit
[process exited with code 0]
</pre>

<img src="preview/gopher.gif" alt="dancing gopher" width="128">

</div>

# Advent of Code 2025 – Solutions in Go

🎄 Welcome to my repository containing solutions for 
**[Advent of Code 2025](https://adventofcode.com/2025)**, implemented in the 
**Go programming language**. 

Advent of Code is an annual programming event featuring daily puzzles throughout
December, designed to challenge problem-solving and algorithmic thinking.

---

## 📂 Repository Structure

|            |            |
|  ------    |  ------    |
| Day 01  ✔️  ✔️ | Day 07  ✔️  ✔️ |
| Day 02  ✔️  ✔️ | Day 08  ✔️  ✔️ |
| Day 03  ✔️  ✔️ | Day 09  ✔️  ☐  |
| Day 04  ✔️  ✔️ | Day 10  ✔️  ☐  |
| Day 05  ✔️  ✔️ | Day 11  ✔️  ✔️  |
| Day 06  ✔️  ✔️ | Day 12  ☐  ☐  |
|------------|-----------|

- `day01/` – Solution for Day 1 puzzle  
- `day02/` – Solution for Day 2 puzzle
- ... and so on, up to Day 12
- `common/` – Shared utilities, helpers, and reusable components  
- `README.md` – This document

Each day’s folder contains:
- `main.go` – Entry point for the solution  
- `input.txt` – Puzzle input (not included here, see 
[Advent of Code 2025](https://adventofcode.com/2025))  
- `part-1.txt`, `part-2.txt` – Puzzle descriptions (not included here, see 
[Advent of Code 2025](https://adventofcode.com/2025))
- `test.go` – Optional unit tests  

---

## 🚀 Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) version 1.25 or higher

### Running a Solution
Navigate to the desired day’s folder and run:

```bash
go run .
```

If the solution requires an input file, place it in the same folder as `input.txt`.

### Running Tests
Some solutions include unit tests (and benchmarks). Run them in solution's folder  with:

```bash
go test .
```

```bash
go test -bench . -benchmem
```

---

## ✨ Goals

- Practice Go by solving algorithmic challenges  
- Write clean, modular, and reusable code  
- Document approaches and highlight interesting techniques  
- Share solutions with the community 
- Have fun

---

## 📖 Notes

- Puzzle inputs are **not included** in this repository, as per Advent of Code rules.  
- Solutions are written with clarity and readability in mind, sometimes prioritizing simplicity over extreme optimization.  
- Feel free to explore, learn, and adapt these solutions for your own practice.

## Create gif after render to png-images

```bash
ffmpeg -i %3d.png -vf "fps=10,scale=640:-1:flags=lanczos" output.gif
```

and with reverse
```bash
# first a direct pass
ffmpeg -framerate 10 -i %3d.png -vf "scale=640:-1:flags=lanczos" forward.mp4

# then reverse
ffmpeg -i forward.mp4 -vf reverse reverse.mp4

# join
ffmpeg -i forward.mp4 -i reverse.mp4 -filter_complex "[0:v][1:v]concat=n=2:v=1:a=0" output.gif

```

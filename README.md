# Go Quests


![Go Quests banner](https://images.steamusercontent.com/ugc/964242373533560541/F1BD3729A9743D8D062FE780774044B192356454/?imw=5000&imh=5000&ima=fit&impolicy=Letterbox&imcolor=%23000000&letterbox=false)

<p align="center"><strong>⭐ If these quests helped you learn Go, please consider starring the repo.</strong></p>



A beginner-friendly set of small Go exercises (“quests”). Each quest is intentionally small so you can practice one Go concept at a time.

You implement functions in `<topic>.go` and confirm correctness with unit tests in `<topic>_test.go`.

---

### Requirements

- [Golang](https://go.dev/dl/) (install from the official site)
- [Visual Studio Code](https://code.visualstudio.com/download) (install from the official site)
- [VS Code Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [VS Code Readme extension](https://marketplace.visualstudio.com/items?itemName=manishsencha.readme-preview)
- [Git](https://git-scm.com/install/) (optional, but recommended)


Check your Go version:

```sh
go version
```

---

### Get the repo (two ways)

**Option A: Clone with Git (recommended)**

Repo: <https://github.com/lite-quests/go-quests>

1. Open Visual studio code.
2. File → Open Folder → Select the folder where you want the repository to reside.
2. Toolbar → Terminal → New Terminal

```sh
git clone https://github.com/lite-quests/go-quests.git
cd go-quests # This command helps you go into the quest directory from the terminal
```

**Option B: Download ZIP (no Git required)**
1. Open the repo on GitHub: <https://github.com/lite-quests/go-quests>
2. Click **Code** → **Download ZIP**
3. Unzip the file
4. Visual Studio Code → Open Folder → Select the extracted/unzipped folder
---

### Code structure

- `quests/` — exercises you edit (each quest is a Go package)

- `quests/<n.topic>/README.md` — instructions for that quest

- `quests/<n.topic>/<topic>.go` — your implementation (edit this)

- `quests/<n.topic>/<topic>_test.go` — tests (do not change)

- `solutions/` — reference implementations (use to compare after you try)

---

### Quests

| # | Topic | Folder | Difficulty |
|---|-------|--------|------------|
| 1 | Hello Go | [quests/001.hello_go](quests/001.hello_go) | Easy |
| 2 | Values | [quests/002.values](quests/002.values) | Easy |
| 3 | Loops | [quests/003.loops](quests/003.loops) | Easy |
| 4 | Conditions | [quests/004.conditions](quests/004.conditions) | Easy |
| 5 | Slice | [quests/005.slice](quests/005.slice) | Easy |
| 6 | Maps | [quests/006.maps](quests/006.maps) | Easy |
| 7 | Functions | [quests/007.functions](quests/007.functions) | Easy |
| 8 | Pointers | [quests/008.pointers](quests/008.pointers) | Easy |
| 9 | Strings | [quests/009.strings](quests/009.strings) | Easy |
| 10 | Structs | [quests/010.structs](quests/010.structs) | Easy |
| 11 | Interfaces | [quests/011.interfaces](quests/011.interfaces) | Medium |
| 12 | Enum | [quests/012.enum](quests/012.enum) | Medium |
| 13 | Generics | [quests/013.generics](quests/013.generics) | Medium |
| 14 | Error | [quests/014.error](quests/014.error) | Medium |
| 15 | Go Routine | [quests/015.go_routine](quests/015.go_routine) | Medium |
| 16 | Channel | [quests/016.channel](quests/016.channel) | Medium |
| 17 | Select & Timeout | [quests/017.select_timeout](quests/017.select_timeout) | Medium |
| 18 | Tickers | [quests/018.tickers](quests/018.tickers) | Medium |
| 19 | Worker Pool | [quests/019.worker_pool](quests/019.worker_pool) | Hard |
| 20 | Rate Limiting | [quests/020.rate_limiting](quests/020.rate_limiting) | Hard |
| 21 | Mutexes | [quests/021.mutexes](quests/021.mutexes) | Hard |
| 22 | String Formatting | [quests/022.string_formatting](quests/022.string_formatting) | Medium |
| 23 | Regex | [quests/023.regex](quests/023.regex) | Medium |
| 24 | Files | [quests/024.files](quests/024.files) | Medium |
| 25 | CLI | [quests/025.cli](quests/025.cli) | Medium |
| 26 | Env Vars | [quests/026.env_vars](quests/026.env_vars) | Easy |
| 27 | HTTP Server | [quests/027.http_server](quests/027.http_server) | Hard |
| 28 | HTTP Client | [quests/028.http_client](quests/028.http_client) | Hard |
| 29 | Context | [quests/029.context](quests/029.context) | Hard |
| 30 | Logging | [quests/030.logging](quests/030.logging) | Medium |
| 31 | Processes | [quests/031.processes](quests/031.processes) | Hard |
| 32 | Exit | [quests/032.exit](quests/032.exit) | Easy |
| 33 | Signals | [quests/033.signals](quests/033.signals) | Hard |
| 34 | JSON | [quests/034.json](quests/034.json) | Medium |
| 35 | Defer, Panic & Recover | [quests/035.defer_panic_recover](quests/035.defer_panic_recover) | Medium |
| 36 | Sorting | [quests/036.sorting](quests/036.sorting) | Easy |

---

### How to solve a quest (intended workflow)
1. Open the repo folder in VS Code (Or any other IDE of your choice)
2. Pick a quest folder under `quests/` (start with smaller numbers first)
3. Read `quests/<n.topic>/README.md` thoroughly. Review all referenced links and documentation to gain a complete understanding of the underlying concept. If the preview does not render correctly, use **Ctrl + Shift + V** to open the README in preview mode for improved readability.
4. Open the implementation file:
   - Example: `quests/010.structs/structs.go`
5. Implement the required functions/methods.
6. Run the quest tests:

```sh
go test -v ./quests/010.structs
```

6. Repeat until all tests pass. If stuck, inspect the failing test output first.

7. Only after you’ve tried: compare with the reference in `solutions/`.

This repo is designed so that tests teach you the spec: make changes → run tests → refine → repeat.


---

### Tips that save beginners a lot of time

- **Don’t print unless the quest asks**
  Tests usually check return values or exact stdout bytes.

- **Read test failures carefully**: they usually tell you the exact edge case you missed.

- **Avoid changing test files**
  If tests don’t pass, fix `<topic>.go` (the tests define the requirements).

- **Disable AI assistance while solving**
  You’ll learn faster if you struggle a bit, read the errors, and look things up in the Go docs. Use AI only after you’ve made a real attempt.

- **Prefer the intended concept over “test hacks”**
  You might find a workaround that passes the tests, but the goal is to practice the concept the quest is designed to teach (the quest README will usually hint what to use) so please stick to it.


---

### Where are the answers?

- `solutions/` contains reference implementations for each quest.
- Best practice: try first, then compare after you pass (or when you’re stuck).


---

### Troubleshooting

- **“package not found”**: make sure you’re running test command from the correct directory.
- **Tests fail on stdout**: check newline requirements (`fmt.Print` vs `fmt.Println`).
- **Changes don’t seem to affect tests**
  Clear the test cache:
  ```sh
  go clean -testcache
  ```

---

### Contact

For any issues, contact either:
- [Mani Yadla](https://x.com/mani_yadla_)
- [Ananya Pappula](https://x.com/AnanyaPappula)

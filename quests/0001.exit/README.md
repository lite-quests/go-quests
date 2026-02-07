# Go Database Connection Exit Codes

## Concept

In Go, `os.Exit(code)` immediately terminates the current process with the given exit status code. Unlike a normal function return, `os.Exit` stops execution instantly and does not run any deferred functions.

Exit codes are used by command-line programs to signal success (`0`) or failure (`non-zero`). This allows the operating system and other programs to detect whether a program completed successfully.

Understanding `os.Exit` is important when building CLI tools and handling fatal errors.

---

## References

* [https://pkg.go.dev/os#Exit](https://pkg.go.dev/os#Exit)
* [https://gobyexample.com/exit](https://gobyexample.com/exit)
* [https://go.dev/doc/effective_go#errors](https://go.dev/doc/effective_go#errors)

---

## Quest

### Objective

In real-world applications, programs often connect to databases to fetch or store data. When a connection attempt fails, the program must exit with a specific exit code indicating the exact reason for failure.

Exit codes allow operating systems, scripts, and other programs to understand what went wrong and take appropriate action. Using meaningful exit codes is essential when building backend services, CLI tools, and infrastructure systems.

Implement a function that simulates connecting to a database and exits with different exit codes depending on the connection result.

---

### Requirements

#### `ConnectDB`

Function signature:

```go
ConnectDB(host string, port int, credentialsValid bool)
```

The function must exit based on the following conditions:

---

### Exit Conditions
| Case  | Condition             | Input Scenario                                             | Exit Code | Meaning                                        |
| ----- | --------------------- | ---------------------------------------------------------- | --------- | ---------------------------------------------- |
| **1** | Host unreachable      | `host` is empty                                               | `1`       | The database host is empty or unreachable      |
| **2** | Invalid port          | `port` is lesser than or equal to 0 (OR)  `port` is greater than 65535                            | `2`       | The port number is outside the valid TCP range |
| **3** | Authentication failed | `credentialsValid` is false                                | `3`       | Database credentials are incorrect             |
| **4** | Successful connection | `host` is not empty (AND) valid port (AND) `credentialsValid` is true | `0`       | Database connection successful                 |


### Inputs

| Parameter        | Type   | Description                     |
| ---------------- | ------ | ------------------------------- |
| host             | string | Database hostname               |
| port             | int    | Database port number            |
| credentialsValid | bool   | Whether credentials are correct |

---

### Outputs

None (process exits with a status code)

---

### Examples

```go
ConnectDB("localhost", 5432, true)
```

Exit code:

```
0
```

---

```go
ConnectDB("", 5432, true)
```

Exit code:

```
1
```


---

## Testing

To run the tests, execute the following command from the root directory:

```bash
go test -v ./quests/0001.exit
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestConnectDB
--- PASS: TestConnectDB (0.00s)
PASS
```

## Exercises

Total time: 30 minutes. Each exercise is self-contained and uses the chapter order from _Learning Go with Tests_: Install Go, Hello World, Integers.

1. **Verify the Go toolchain and initialise a module (5 min)**  
   Confirm that Go is installed and usable by running `go version` and `go env GOPATH`. Create a new directory for the exercises, initialise a Go module inside it with a name such as `github.com/yourname/ledgerflow-tdd`, and confirm that a `go.mod` file is created. Expected result: all commands succeed, `go.mod` exists, and `go test ./...` reports no test files (or no failures).

   ```sh
   ~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging*
    ❯ ls
    .rw-r--r-- 79 erona  9 out 08:09  go.mod

    ~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging*
    ❯ ls
    .rw-r--r-- 79 erona  9 out 08:09  go.mod
   ```

   See diff bellow

   ````diff
   diff --git a/exercises/ledgerflow-tdd/go.mod b/exercises/ledgerflow-tdd/go.mod
    new file mode 100644
    index 0000000..60dff0c
    --- /dev/null
    +++ b/exercises/ledgerflow-tdd/go.mod
    @@ -0,0 +1,3 @@
    +module github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd
    +
    +go 1.25.3
    diff --git a/exercises/lgwt-1.md b/exercises/lgwt-1.md
    new file mode 100644
    index 0000000..282a534
    --- /dev/null
    +++ b/exercises/lgwt-1.md
    @@ -0,0 +1,48 @@
    +## Exercises
    +
    +Total time: 30 minutes. Each exercise is self-contained and uses the chapter order from _Learning Go with Tests_: Install Go, Hello World, Integers.
    +
    +1. **Verify the Go toolchain and initialise a module (5 min)**
    +   Confirm that Go is installed and usable by running `go version` and `go env GOPATH`. Create a new directory for the exercises, initialise a Go module inside it with a name such as `github.com/yourname/ledgerflow-tdd`, and confirm that a `go.mod` file is created. Expected result: all commands succeed, `go.mod` exists, and `go test ./...` reports no test files (or no failures).
    +
    +   ```sh
    +   ~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging*
    +    ❯ ls
    +    .rw-r--r-- 79 erona  9 out 08:09  go.mod
    +
    +    ~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging*
    +    ❯ ls
    +    .rw-r--r-- 79 erona  9 out 08:09  go.mod
    +   ```
    +
    +2. **Hello World with TDD, constants, if/switch, sub-tests, and closures (10 min)**
    +   In a package named `greeting`, write a failing test first for a function named `Hello` that takes a name and returns a greeting. The function should return `Hello, world!` when the name is empty, and `Hello, <name>!` otherwise. Use table-driven sub-tests with closures to cover both cases. Implement the function using a constant for the default name and an `if` or `switch` statement. Run `go test -v` and confirm that every sub-test passes. Expected result: the test output shows named sub-tests, and the implementation uses a constant rather than a hard-coded string in the return.
    +
    +3. **Integers, documentation, and coverage (15 min)**
    +   Create a package named `integers`. Write a failing test file first for an `Add` function that takes two integers and returns their sum. Use a table-driven test covering positive, negative, and zero cases: `2+2=4`, `-1+1=0`, `0+0=0`, and `-3+-2=-5`. Run the test to confirm it fails, then implement `Add`. Add a Go doc comment directly above `Add` that begins with the function name and describes the parameters and return value. Run `go test -cover` and `go doc . Add`. Expected result: all test cases pass, coverage for `Add` is 100%, and `go doc` prints the comment you wrote.
    +
    +---
    +
    +## Feature Suggestion
    +
    +### Business Justification
    +
    +LedgerFlow explicitly lists observability as a core goal: structured logging, distributed tracing, and dashboards for fintech flows. The `transfer-api` currently has no request-level logging, which makes it difficult to see how the service behaves under real HTTP traffic, debug failed health checks, or prepare for OpenTelemetry instrumentation. Adding a request-logging middleware is a small, high-value first step that improves auditability, troubleshooting, and operational readiness without changing the public API surface.
    +
    +### What Needs to Be Done
    +
    +- Implement an HTTP middleware for `transfer-api` that wraps the existing router.
    +- The middleware must capture the incoming request method, path, response status code, and total handling duration.
    +- The middleware must use the same injected logger instance that `main` already configures, rather than a global logger or `fmt.Println`.
    +- Ensure the middleware passes requests through to the existing `/health` and `/ready` handlers without altering their response bodies or status codes.
    +- Add a test that exercises at least one endpoint through the middleware and confirms the response is still correct.
    +- Update the service README only if the middleware changes how the service is started or configured.
    +
    +### Acceptance Criteria
    +
    +- Every HTTP request to `/health` and `/ready` produces a structured log entry containing the request method, path, response status, and duration.
    +- The middleware uses the logger instance created in `main`; no new global logger is introduced.
    +- `/health` still returns `{"Status":"Ok"}` with HTTP 200, and `/ready` still returns `{"Status":"Ready"}` with HTTP 200.
    +- `go test .` passes, including the new middleware test.
    +- Manual verification with `go run .` shows a log line for each request when the endpoints are called.
    +- No request-logging output uses `fmt.Println`; all request information is logged through the injected structured logger.

   ````

2. **Hello World with TDD, constants, if/switch, sub-tests, and closures (10 min)**  
    In a package named `greeting`, write a failing test first for a function named `Hello` that takes a name and returns a greeting. The function should return `Hello, world!` when the name is empty, and `Hello, <name>!` otherwise. Use table-driven sub-tests with closures to cover both cases. Implement the function using a constant for the default name and an `if` or `switch` statement. Run `go test -v` and confirm that every sub-test passes. Expected result: the test output shows named sub-tests, and the implementation uses a constant rather than a hard-coded string in the return.

   ````sh
   ~/repos/ledgerflow/exercises/ledgerflow-tdd/greeting feat/request-audit-logging*
    ❯ go test -v
    === RUN   TestHelloWorld
    === RUN   TestHelloWorld/When_empty_returns_hello_world
    === RUN   TestHelloWorld/When_have_a_name_returns_hello_name
    --- PASS: TestHelloWorld (0.00s)
        --- PASS: TestHelloWorld/When_empty_returns_hello_world (0.00s)
        --- PASS: TestHelloWorld/When_have_a_name_returns_hello_name (0.00s)
    PASS
    ok      github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/greeting   0.002s
    ```

    ```diff
    diff --git a/exercises/ledgerflow-tdd/greeting/greeting.go b/exercises/ledgerflow-tdd/greeting/greeting.go
   new file mode 100644
   index 0000000..13c3637
   --- /dev/null
   +++ b/exercises/ledgerflow-tdd/greeting/greeting.go
   @@ -0,0 +1,18 @@
   +package greeting
   ````

- +import "fmt"
- +const (
- empty = ""
- messageTemplate = "Hello, %s!"
- defaultFiller = "world"
  +)
- +func Hello(name string) string {
- filler := defaultFiller
- if name != empty {
-     filler = name
- }
-
- return fmt.Sprintf(messageTemplate, filler)
  +}
  diff --git a/exercises/ledgerflow-tdd/greeting/greeting_test.go b/exercises/ledgerflow-tdd/greeting/greeting_test.go
  new file mode 100644
  index 0000000..7221b33
  --- /dev/null
  +++ b/exercises/ledgerflow-tdd/greeting/greeting_test.go
  @@ -0,0 +1,34 @@
  +package greeting
- +import (
- "testing"
  +)
- +func TestHelloWorld(t \*testing.T) {
- table := []struct {
-     test string
-     name string
-     want string
- }{
-     {
-     	test: "When empty returns hello world",
-     	name: "",
-     	want: "Hello, world!",
-     },
-     {
-     	test: "When have a name returns hello name",
-     	name: "Rachel",
-     	want: "Hello, Rachel!",
-     },
- }
-
- for \_, v := range table {
-     t.Run(v.test, func(t *testing.T) {
-     	result := Hello(v.name)
-
-     	if result != v.want {
-     		t.Errorf("want %s, have %s", v.want, result)
-     	}
-     })
- }
  +}

  ```

  ```

3. **Integers, documentation, and coverage (15 min)**  
    Create a package named `integers`. Write a failing test file first for an `Add` function that takes two integers and returns their sum. Use a table-driven test covering positive, negative, and zero cases: `2+2=4`, `-1+1=0`, `0+0=0`, and `-3+-2=-5`. Run the test to confirm it fails, then implement `Add`. Add a Go doc comment directly above `Add` that begins with the function name and describes the parameters and return value. Run `go test -cover` and `go doc . Add`. Expected result: all test cases pass, coverage for `Add` is 100%, and `go doc` prints the comment you wrote.

   ```sh
   ~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging*
   ❯ go test ./integers/
   --- FAIL: TestAdd (0.00s)
    --- FAIL: TestAdd/positive_case (0.00s)
        integers_test.go:43: want '\x04', have '\x00'
    --- FAIL: TestAdd/negative_case (0.00s)
        integers_test.go:43: want '�', have '\x00'
   FAIL
   FAIL    github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers   0.003s
   FAIL
   ```

~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
❯ go test ./integers/
--- FAIL: TestAdd (0.00s)
--- FAIL: TestAdd/positive_case (0.00s)
integers_test.go:43: want 4, have 0
--- FAIL: TestAdd/negative_case (0.00s)
integers_test.go:43: want -1, have 0
FAIL
FAIL github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers 0.003s
FAIL

~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
❯ go test -v ./integers/
=== RUN TestAdd
=== RUN TestAdd/positive_case
integers_test.go:43: want 4, have 0
=== RUN TestAdd/negative_case
integers_test.go:43: want -1, have 0
=== RUN TestAdd/zero_case
=== RUN TestAdd/another_zero_case
--- FAIL: TestAdd (0.00s)
--- FAIL: TestAdd/positive_case (0.00s)
--- FAIL: TestAdd/negative_case (0.00s)
--- PASS: TestAdd/zero_case (0.00s)
--- PASS: TestAdd/another_zero_case (0.00s)
FAIL
FAIL github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers 0.003s
FAIL

~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
❯ go test -v ./integers/
=== RUN TestAdd
=== RUN TestAdd/positive_case
=== RUN TestAdd/negative_case
=== RUN TestAdd/zero_case
=== RUN TestAdd/another_zero_case
--- PASS: TestAdd (0.00s)
--- PASS: TestAdd/positive_case (0.00s)
--- PASS: TestAdd/negative_case (0.00s)
--- PASS: TestAdd/zero_case (0.00s)
--- PASS: TestAdd/another_zero_case (0.00s)
PASS
ok github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers 0.003s

~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
❯ go test -cover

# .

no Go files in /home/erona/repos/ledgerflow/exercises/ledgerflow-tdd
FAIL . [setup failed]

~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
❯ cd integers/

~/repos/ledgerflow/exercises/ledgerflow-tdd/integers feat/request-audit-logging\*
❯ go test -cover
PASS
coverage: 100.0% of statements
ok github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers 0.003s

~/repos/ledgerflow/exercises/ledgerflow-tdd/integers feat/request-audit-logging\*
❯ go doc . Add
package integers // import "github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers"

func Add(a, b int) int
Add takes two parameters of integer type, a and b, and returns it's sum

````

```diff
diff --git a/exercises/ledgerflow-tdd/integers/integers.go b/exercises/ledgerflow-tdd/integers/integers.go
new file mode 100644
index 0000000..52b5e53
--- /dev/null
+++ b/exercises/ledgerflow-tdd/integers/integers.go
@@ -0,0 +1,7 @@
+package integers
+
+// Add takes two parameters of integer type, a and b,
+// and returns it's sum
+func Add(a, b int) int {
+	return a + b
+}
diff --git a/exercises/ledgerflow-tdd/integers/integers_test.go b/exercises/ledgerflow-tdd/integers/integers_test.go
new file mode 100644
index 0000000..26d97bf
--- /dev/null
+++ b/exercises/ledgerflow-tdd/integers/integers_test.go
@@ -0,0 +1,48 @@
+package integers
+
+import "testing"
+
+func TestAdd(t *testing.T) {
+	table := []struct {
+		test           string
+		number1        int
+		number2        int
+		expectedResult int
+	}{
+		{
+			test:           "positive case",
+			number1:        2,
+			number2:        2,
+			expectedResult: 4,
+		},
+		{
+			test:           "negative case",
+			number1:        3,
+			number2:        -4,
+			expectedResult: -1,
+		},
+		{
+			test:           "zero case",
+			number1:        5,
+			number2:        -5,
+			expectedResult: 0,
+		},
+		{
+			test:           "another zero case",
+			number1:        0,
+			number2:        0,
+			expectedResult: 0,
+		},
+	}
+
+	for _, v := range table {
+		t.Run(v.test, func(t *testing.T) {
+			result := Add(v.number1, v.number2)
+
+			if result != v.expectedResult {
+				t.Errorf("want %d, have %d", v.expectedResult, result)
+			}
+		})
+	}
+
+}
diff --git a/exercises/lgwt-1.md b/exercises/lgwt-1.md
index e453147..166a5df 100644
--- a/exercises/lgwt-1.md
+++ b/exercises/lgwt-1.md
@@ -162,12 +162,98 @@ Total time: 30 minutes. Each exercise is self-contained and uses the chapter ord
 -     })
 - }
   +}
+
````

````

3. **Integers, documentation, and coverage (15 min)**
-   Create a package named `integers`. Write a failing test file first for an `Add` function that takes two integers and returns their sum. Use a table-driven test covering positive, negative, and zero cases: `2+2=4`, `-1+1=0`, `0+0=0`, and `-3+-2=-5`. Run the test to confirm it fails, then implement `Add`. Add a Go doc comment directly above `Add` that begins with the function name and describes the parameters and return value. Run `go test -cover` and `go doc . Add`. Expected result: all test cases pass, coverage for `Add` is 100%, and `go doc` prints the comment you wrote.
+    Create a package named `integers`. Write a failing test file first for an `Add` function that takes two integers and returns their sum. Use a table-driven test covering positive, negative, and zero cases: `2+2=4`, `-1+1=0`, `0+0=0`, and `-3+-2=-5`. Run the test to confirm it fails, then implement `Add`. Add a Go doc comment directly above `Add` that begins with the function name and describes the parameters and return value. Run `go test -cover` and `go doc . Add`. Expected result: all test cases pass, coverage for `Add` is 100%, and `go doc` prints the comment you wrote.
+
+   ```sh
+   ~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging*
+   ❯ go test ./integers/
+   --- FAIL: TestAdd (0.00s)
+    --- FAIL: TestAdd/positive_case (0.00s)
+        integers_test.go:43: want '\x04', have '\x00'
+    --- FAIL: TestAdd/negative_case (0.00s)
+        integers_test.go:43: want '�', have '\x00'
+   FAIL
+   FAIL    github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers   0.003s
+   FAIL
+   ```
+
+~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
+❯ go test ./integers/
+--- FAIL: TestAdd (0.00s)
+--- FAIL: TestAdd/positive_case (0.00s)
+integers_test.go:43: want 4, have 0
+--- FAIL: TestAdd/negative_case (0.00s)
+integers_test.go:43: want -1, have 0
+FAIL
+FAIL github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers 0.003s
+FAIL
+
+~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
+❯ go test -v ./integers/
+=== RUN TestAdd
+=== RUN TestAdd/positive_case
+integers_test.go:43: want 4, have 0
+=== RUN TestAdd/negative_case
+integers_test.go:43: want -1, have 0
+=== RUN TestAdd/zero_case
+=== RUN TestAdd/another_zero_case
+--- FAIL: TestAdd (0.00s)
+--- FAIL: TestAdd/positive_case (0.00s)
+--- FAIL: TestAdd/negative_case (0.00s)
+--- PASS: TestAdd/zero_case (0.00s)
+--- PASS: TestAdd/another_zero_case (0.00s)
+FAIL
+FAIL github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers 0.003s
+FAIL
+
+~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
+❯ go test -v ./integers/
+=== RUN TestAdd
+=== RUN TestAdd/positive_case
+=== RUN TestAdd/negative_case
+=== RUN TestAdd/zero_case
+=== RUN TestAdd/another_zero_case
+--- PASS: TestAdd (0.00s)
+--- PASS: TestAdd/positive_case (0.00s)
+--- PASS: TestAdd/negative_case (0.00s)
+--- PASS: TestAdd/zero_case (0.00s)
+--- PASS: TestAdd/another_zero_case (0.00s)
+PASS
+ok github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers 0.003s
+
+~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
+❯ go test -cover
+
+# .
+
+no Go files in /home/erona/repos/ledgerflow/exercises/ledgerflow-tdd
+FAIL . [setup failed]
+
+~/repos/ledgerflow/exercises/ledgerflow-tdd feat/request-audit-logging\*
+❯ cd integers/
+
+~/repos/ledgerflow/exercises/ledgerflow-tdd/integers feat/request-audit-logging\*
+❯ go test -cover
+PASS
+coverage: 100.0% of statements
+ok github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers 0.003s
+
+~/repos/ledgerflow/exercises/ledgerflow-tdd/integers feat/request-audit-logging\*
+❯ go doc . Add
+package integers // import "github.com/EronAlves1996/ledgerflow/exercises/ledgerflow-tdd/integers"
+
+func Add(a, b int) int
+Add takes two parameters of integer type, a and b, and returns it's sum
+
+```
+
+

---

@@ -194,3 +280,4 @@ LedgerFlow explicitly lists observability as a core goal: structured logging, di
- `go test .` passes, including the new middleware test.
- Manual verification with `go run .` shows a log line for each request when the endpoints are called.
- No request-logging output uses `fmt.Println`; all request information is logged through the injected structured logger.
+```
````

---

## Feature Suggestion

### Business Justification

LedgerFlow explicitly lists observability as a core goal: structured logging, distributed tracing, and dashboards for fintech flows. The `transfer-api` currently has no request-level logging, which makes it difficult to see how the service behaves under real HTTP traffic, debug failed health checks, or prepare for OpenTelemetry instrumentation. Adding a request-logging middleware is a small, high-value first step that improves auditability, troubleshooting, and operational readiness without changing the public API surface.

### What Needs to Be Done

- Implement an HTTP middleware for `transfer-api` that wraps the existing router.
- The middleware must capture the incoming request method, path, response status code, and total handling duration.
- The middleware must use the same injected logger instance that `main` already configures, rather than a global logger or `fmt.Println`.
- Ensure the middleware passes requests through to the existing `/health` and `/ready` handlers without altering their response bodies or status codes.
- Add a test that exercises at least one endpoint through the middleware and confirms the response is still correct.
- Update the service README only if the middleware changes how the service is started or configured.

### Acceptance Criteria

- Every HTTP request to `/health` and `/ready` produces a structured log entry containing the request method, path, response status, and duration.
- The middleware uses the logger instance created in `main`; no new global logger is introduced.
- `/health` still returns `{"Status":"Ok"}` with HTTP 200, and `/ready` still returns `{"Status":"Ready"}` with HTTP 200.
- `go test .` passes, including the new middleware test.
- Manual verification with `go run .` shows a log line for each request when the endpoints are called.
- No request-logging output uses `fmt.Println`; all request information is logged through the injected structured logger.

```

```

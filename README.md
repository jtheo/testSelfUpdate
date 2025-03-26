# Testing minio selfupdate for hacking week

- [doc](https://pkg.go.dev/github.com/minio/selfupdate)
- [repo](https://github.com/minio/selfupdate)

## Example

To use this example, you need to run in another terminal:

```
$: cd server
$: go run .
2025/03/26 16:32:16 Serving files from ../dist on 0.0.0.0:8080
```

### Create a new release

```
$: ./build.sh
===== Wed 26 Mar 2025 16:25:18 GMT ==> Running go vet
===== Wed 26 Mar 2025 16:25:18 GMT ==> building version 1...
Building
  - linux amd64 arm64
  - darwin amd64 arm64
```

### Local copy for test

```
$: cp dist/1/testselfupdate-darwin-arm64 testselfupdate
```

### Tests

```
$: ./testselfupdate
Hello, World!

$: ./testselfupdate -v
testselfupdate, version: 1

$: ./testselfupdate -u
You are already at the latest version
```

### New version:

```
$: ./build.sh
===== Wed 26 Mar 2025 16:25:32 GMT ==> Running go vet
===== Wed 26 Mar 2025 16:25:32 GMT ==> building version 2...
Building
  - linux amd64 arm64
  - darwin amd64 arm64
```

### New tests

```
$: ./testselfupdate
Your version of testselfupdate is out of date
You are on version 1, but there's a new version 2

Hello, World!

$: ./testselfupdate -u
Your version of testselfupdate is out of date
You are on version 1, but there's a new version 2

Running the update
Update successful, bye!

$: ./testselfupdate -u
You are already at the latest version
```

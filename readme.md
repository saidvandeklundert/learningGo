Thing I learn about Go.

```
# install go
sudo apt update && sudo apt upgrade
sudo apt install golang-go
wget -c https://golang.org/dl/go1.23.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.23.2.linux-amd64.tar.gz
export  PATH=$PATH:/usr/local/go/bin
go env -w GOPROXY=direct
go install -v golang.org/x/tools/gopls@latest

# make the first project
mkdir helloworld
cd .\helloworld
# make the file
go build main.go
go run .

# format the code
go fmt

# run tests:
go test
go test -v

# check test coverage:
go test -cover .
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# put tests in 'file-name'_test.go in the same folder and
# prefer table tests. Example:
```
```go
func Test_isPrimeTableTests(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
	}

	for _, tc := range primeTests {
		result, msg := isPrime(tc.testNum)
		if result != tc.expected {
			t.Errorf("%s: expected %t got %t", tc.name, tc.expected, result)
		}

		if msg != tc.msg {
			t.Errorf("%s: expected %s got %s", tc.name, tc.msg, msg)
		}
	}
}
```
```
# check for potential errors or suspicious constructs
go vet

# get another package
go get <package name>
go get golang.org/x/sys/unix

# trace allocations:
go build -gcflags "-m -m"

# benchmark:
go test -bench=.

#benchmark the memory as well:
go test -bench=. -benchmem

# run the benchmark for 3 seconds and do it 5 times:
go test -bench=BenchmarkMultiply -benchtime=3s -count=5

go install golang.org/x/perf/cmd/benchstat@latest
# go test -bench=. > old.txt
# go test -bench=. > new.txt
# benchstat old.txt new.txt


# cpu profiling requires changes to the program:
import (
	"runtime/pprof"
)
func main() {
	// ...
	f, err := os.Create("cpuprofile.out")
	if err != nil {
	// Handle error
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// ... (Rest of your code)
}
go build monitor.go
./monitor

# analyze the output:
go tool pprof cpuprofile.out

# create a flame graph:
go tool pprof -web cpuprofile.out

# setup memory profiling:

f, err := os.Create("memprofile.out")
if err != nil {
// Handle error
}
defer f.Close()
runtime.GC()
pprof.WriteHeapProfile(f)



# analyze the memory profiling output
go tool pprof memprofile.out`
go tool pprof -web memprofile.out

# sync workspace code:
go work sync
```
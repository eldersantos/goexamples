package fetchall

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

<<<<<<< HEAD:fetchall.go
// type for urlcounts
type UrlCounts struct {
	Message string
	Error   int
	Success int
}

func main() {
=======
// Main method
func Main() {
	runtime.GOMAXPROCS(4)
>>>>>>> 7272bbdfce59c6670f49700e762ddd6988d3c5fb:fetchall/fetchall.go
	start := time.Now()
	ch := make(chan UrlCounts)

	f, err := ioutil.ReadFile("urls.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(f), "\r\n")

	for _, url := range sliceData {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		go fetch(url, ch) // start a goroutine
	}
	for range sliceData {
		fmt.Println(<-ch.Message) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- UrlCounts) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- UrlCounts{fmt.Sprint(err), ch.Error + 1, ch.Success + 1}
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch.Message <- fmt.Sprintf("while reading %s: %v", url, err)
		ch.Error++
		return
	}
	secs := time.Since(start).Seconds()
<<<<<<< HEAD:fetchall.go
	ch.Message <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	ch.Success++
=======
	ch <- fmt.Sprintf("%.2fs  %7d  %3d  %s", secs, nbytes, resp.StatusCode, url)
>>>>>>> 7272bbdfce59c6670f49700e762ddd6988d3c5fb:fetchall/fetchall.go
}

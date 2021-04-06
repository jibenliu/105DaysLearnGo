package main

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"os"
	"time"
)

func main() {
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", "http://www.golang-book.com/public/pdf/gobook.pdf")

	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("	%v\n", resp.HTTPResponse.Status)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()
Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("	transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress(),
			)
		case <-resp.Done:
			break Loop
		}
	}
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed : %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Download saved to ./%v \n", resp.Filename)
}

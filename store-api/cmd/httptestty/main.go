package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	//client := NewHTTPClient()

	/*prefix := strings.NewReader("Product image data:\n")

	file, err := os.Open("cmd/httptestty/sample.txt")

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer file.Close()

	modcont := io.MultiReader(prefix, file)

	dst, err := os.Create("cmd/httptestty/dest-file.txt")

	if err != nil {

		fmt.Printf("error: %v\n", err)
		return
	}

	defer dst.Close()

	cop, err := io.Copy(dst, modcont)

	if err != nil {

		fmt.Printf("error: %v\n", err)
		return
	}

	size, err := dst.Stat()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("combinedsize: %d\n", cop)

	fmt.Printf("filesize: %d\n", size.Size()) */

	standardTransport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	}

	standardClient := &http.Client{
		Transport: standardTransport,
		Timeout:   10 * time.Second,
	}

	newclient := NewHTTPClient(standardClient)

	file, err := os.Open("cmd/httptestty/sample.txt")

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer file.Close()

	/* currentPosition, err := file.Seek(10, io.SeekCurrent)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	} */

	size, err := file.Stat()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("filesize: %d\n", size.Size())
	//fmt.Printf("current position: %d\n", currentPosition)

	//rem := size.Size() - currentPosition

	buf := make([]byte, 512)

	preview, err := file.Read(buf)

	if err != nil {

		fmt.Printf("zaerror: %v\n", err)
		return
	}

	fmt.Printf("no of bytes previewed: %d\n", preview)

	fmt.Printf("detected content type: %s\n", http.DetectContentType(buf[:preview]))

	_, err = file.Seek(0, io.SeekStart)

	if err != nil {

		fmt.Printf("error: %v\n", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, "https://httpbin.org/status/404", file)

	if err != nil {

		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Println(req.ContentLength)

	resp, err := newclient.Do(req)

	if err != nil {
		fmt.Printf("error: %v\n", err)

		return
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {

		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("body: %d\n", len(b))

	fmt.Printf("User-Agent: %s\n", req.Header.Get("User-Agent"))
	fmt.Printf("Content-Type: %s\n", req.Header.Get("Content-Type"))
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Header: %v\n", resp.Header)
	fmt.Printf("Response Content-Length: %d\n", resp.ContentLength)
	fmt.Printf("accept: %s\n", req.Header.Get("accept")) //testing if it is truly case insensitive

	/*
		b, err := io.ReadAll(limited)

		if err != nil {

			fmt.Printf("error: %v\n", err)
			return
		}

		fmt.Printf("body: %s\n", b)
		fmt.Printf("length: %d\n", len(b))

		rem, err := io.ReadAll(file)
		if err != nil {

			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Printf("body: %s\n", rem)
		fmt.Printf("length: %d\n", len(rem))

		//

		/*



			type Product struct {
				Name     string `json:"name"`
				Category string `json:"category"`
				Price    int    `json:"price"`
			}

			product := &Product{
				Name:     "Classic Boxer",
				Category: "Men's Underwear",
				Price:    5000,
			}

			body, err := json.Marshal(product)
			if err != nil {
				return
			}





			b, err := io.ReadAll(resp.Body)

			if err != nil {
				return
			}

			fmt.Printf("Content-Type: %v\n", resp.Header.Get("Content-Type"))

			fmt.Printf("Status code: %d\n", resp.StatusCode)

			fmt.Printf("body: %s\n", b)
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				fmt.Printf("code in range\n")
			} else {
				fmt.Printf("code not in range")
			}*/

}

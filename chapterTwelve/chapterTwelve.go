package chaptertwelve

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

func SlowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Slow response"))
	}))
	return s
}

func FastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("error") == "true" {

			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	}))
	return s
}

var client = http.Client{}

func CallBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			cancel()
		}
	}()
	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL)
		if err != nil {
			cancel()
		}
	}()
	wg.Wait()
	fmt.Println("done with both")
}

func callServer(ctx context.Context, label string, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "request err:", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "response err:", err)
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(label, "read err:", err)
	}
	result := string(data)
	if result != "" {
		fmt.Println(label, "result:", result)
	}
	if result == "error" {
		fmt.Println("cancelling from", label)
		return errors.New("error happened")
	}
	return nil
}

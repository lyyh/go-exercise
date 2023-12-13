package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		resp, err := http.Get("https://www.baidu.com")
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("unexpected status code %d", resp.StatusCode)
		}
		fmt.Println("baidu response")
		return nil
	})
	g.Go(func() error {
		resp, err := http.Get("https://www.google.com")
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("unexpected status code %d", resp.StatusCode)
		}
		fmt.Println("google response")
		return nil
	})
	err := g.Wait()
	if err != nil {
		fmt.Println("Errors:", err)
	} else {
		fmt.Println("All tasks Completed successfully")
	}
}

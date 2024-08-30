package main

import (
	"context"
	"fmt"
	"time"
    "net/http"
    "io"
)

func QueryFrameworkStats(ctx context.Context, framework string) <-chan string {
    stats := make(chan string)

    go func() {
        repos := "http://api.github.com/repos/" + framework
        req, err := http.NewRequest("GET", repos, nil)
        if err != nil {
            return
        }

        req = req.WithContext(ctx)

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            return
        }

        data, err := io.ReadAll(resp.Body)
        if err != nil {
            return
        }

        defer resp.Body.Close()
        stats <- string(data)
    }()

    return stats
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    framework := "marevilspirit/mlog"

    select {
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    case statsInfo := <-QueryFrameworkStats(ctx, framework):
        fmt.Println(framework, "fork and start info:", statsInfo)
    }
}

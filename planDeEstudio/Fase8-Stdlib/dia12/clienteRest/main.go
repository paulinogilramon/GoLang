package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	// GET - Obtener posts
	fmt.Println("=== GET: Posts ===")
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts?_limit=3")
	if err != nil {
		fmt.Println("Error GET:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var posts []Post
	json.Unmarshal(body, &posts)

	for _, p := range posts {
		fmt.Printf("[%d] %s\n", p.ID, p.Title)
	}

	// GET con ID específico
	fmt.Println("\n=== GET: Post 1 ===")
	resp, _ = client.Get("https://jsonplaceholder.typicode.com/posts/1")
	var post Post
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal(body, &post)
	fmt.Printf("Post: %s\n%s\n", post.Title, post.Body)

	// GET con query params
	fmt.Println("\n=== GET: Todos (pendientes) ===")
	req, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos", nil)
	q := req.URL.Query()
	q.Add("_limit", "5")
	req.URL.RawQuery = q.Encode()

	resp, _ = client.Do(req)
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()

	var todos []Todo
	json.Unmarshal(body, &todos)
	for _, t := range todos {
		status := "✓"
		if !t.Completed {
			status = "✗"
		}
		fmt.Printf("[%s] %s\n", status, t.Title)
	}

	fmt.Println("\nCliente REST finalizado.")
}

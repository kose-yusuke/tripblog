package main

import (
    "github.com/yusukekoseki/gocrud/db"
    "github.com/yusukekoseki/gocrud/server"
)

func main() {
    db.Init()
    server.Init()
    db.Close()
}



// import (
//     "encoding/json"
//     "fmt"
//     "net/http"
// )

// type Article struct {
//     Title   string `json:"Title"`
//     Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// type Articles []Article

// func homePage(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Welcome to the HomePage!")
//     fmt.Println("Endpoint Hit: homePage")
// }

// func handleRequests() error{
//     http.HandleFunc("/", homePage)
//     http.HandleFunc("/articles", returnAllArticles)
//     return http.ListenAndServe(":8080", nil)
// }

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
//     articles := Articles{}
//     for i := 0; i < 10; i++ {
//         title := "Hello_%d"
//         articles = append(
//             articles,
//             Article{Title: fmt.Sprintf(title, i), Desc: "Article Description", Content: "Article Content"})
//     }
//     fmt.Println("Endpoint Hit: returnAllArticles")
//     json.NewEncoder(w).Encode(articles[0])
// }

// func main() {
//     handleRequests()
// }
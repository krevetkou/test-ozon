package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/krevetkou/test-ozon/app/domain/repository/comment"
	"github.com/krevetkou/test-ozon/app/domain/repository/post"
	"github.com/krevetkou/test-ozon/app/generated"
	"github.com/krevetkou/test-ozon/app/infra/db"
	"github.com/krevetkou/test-ozon/app/infra/persistence"
	"github.com/krevetkou/test-ozon/app/interfaces"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	var (
		defaultPort      = "8080"
		databaseUser     = os.Getenv("DATABASE_USER")
		databaseName     = os.Getenv("DATABASE_NAME")
		databaseHost     = os.Getenv("DATABASE_HOST")
		databasePort     = os.Getenv("DATABASE_PORT")
		databasePassword = os.Getenv("DATABASE_PASSWORD")
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, databasePort, databaseUser, databaseName, databasePassword)

	conn := db.OpenDB(dbConn)

	var postService post.PostService
	var commentService comment.CommentService

	postService = persistence.NewPost(conn)
	commentService = persistence.NewAnswer(conn)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		PostService:    postService,
		CommentService: commentService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

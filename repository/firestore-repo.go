package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/at8109/golang-rest-api/entity"
	"google.golang.org/api/iterator"
)

type repo struct{}

//NewFirestoreRepository creates a new repo
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "golang-project-90315"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a FireStore client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (*repo) FindByID(id string) (entity.Post, error) {

	var fullpost entity.Post
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a FireStore client: %v", err)
		return fullpost, err
	}

	defer client.Close()
	var postid string
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return fullpost, err
		}
		post := doc.Data()["ID"].(int64)
		fmt.Print(post)
		if postid == strconv.FormatInt(post, 10) {

			fullPost := entity.Post{
				ID:    doc.Data()["ID"].(int64),
				Title: doc.Data()["Title"].(string),
				Text:  doc.Data()["Text"].(string),
			}

			fullpost = fullPost
		}
	}
	return fullpost, nil
}

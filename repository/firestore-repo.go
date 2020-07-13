package repository

import (
	"context"
	"log"

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
	_, err = client.Collection(collectionName).Doc(post.ID).Set(ctx, post)
	if err != nil {
		log.Fatalf("Failed to set id to the post: %v", err)
		return nil, err
	}

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
			ID:    doc.Data()["ID"].(string),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (*repo) FindByID(id string) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	dsnap, err := client.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	post := &entity.Post{
		ID:    dsnap.Data()["ID"].(string),
		Title: dsnap.Data()["Title"].(string),
		Text:  dsnap.Data()["Text"].(string),
	}
	return post, nil
}

func (*repo) DeleteByID(id string) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a FireStore client: %v", err)
		return err
	}

	defer client.Close()
	_, er := client.Collection("posts").Doc(id).Delete(ctx)
	if er != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", er)
	}
	return er
}

func (*repo) UpdateByID(id string) error {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return err
	}
	_, er := client.Collection("posts").Doc(id).Set(ctx, map[string]interface{}{
		"Title": "change title",
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	return er
}

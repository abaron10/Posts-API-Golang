package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/abaron10/Posts-API-Golang/model"
	"github.com/abaron10/Posts-API-Golang/repository"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

type postRepo struct{}

//Pseudo contructor para crear un objeto que implementa la interfaz
func NewFirestorePostRepository() repository.PostRepository {
	return &postRepo{}
}

const (
	projectId      string = "pragmatic-reviews-101f6"
	collectionName string = "posts"
)

func (*postRepo) Save(post *model.Post) (*model.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":        post.Id,
		"Title":     post.Title,
		"Text":      post.Text,
		"CreatedBy": post.CreatedBy,
		"CreatedOn": time.Now().UTC(),
	})
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*postRepo) FindAll() ([]model.Post, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []model.Post

	itr := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := model.Post{
			Id:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/abaron10/Posts-API-Golang/models"
	"github.com/abaron10/Posts-API-Golang/repository"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

type userRepo struct{}

const (
	projectUserId      string = "pragmatic-reviews-101f6"
	userCollectionName string = "users"
)

//Pseudo contructor para crear un objeto que implementa la interfaz
func NewFirestoreUserRepository() repository.UserRepository {
	return &userRepo{}
}

func (u *userRepo) SignIn(user *models.User) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(userCollectionName).Add(ctx, map[string]interface{}{
		"id":         user.Id,
		"name":       user.Name,
		"last_name":  user.LastName,
		"user_name":  user.UserName,
		"email":      user.Email,
		"password":   user.Password,
		"created_at": time.Now().UTC(),
	})
	if err != nil {
		log.Fatalf("Error Signing in: %v", err)
		return nil, err
	}
	return user, nil
}

func (*userRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	itr := client.Collection(userCollectionName).Where("email", "==", email).Documents(ctx)
	user, err := GetUserByDocumentIterator(itr)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (*userRepo) GetUserById(id string) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	itr := client.Collection(userCollectionName).Where("id", "==", id).Documents(ctx)
	user, err := GetUserByDocumentIterator(itr)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByDocumentIterator(itr *firestore.DocumentIterator) (*models.User, error) {
	var user models.User
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to get user: %v", err)
			return nil, err
		}
		user = models.User{
			Id:        doc.Data()["id"].(string),
			Name:      doc.Data()["name"].(string),
			LastName:  doc.Data()["last_name"].(string),
			UserName:  doc.Data()["user_name"].(string),
			Email:     doc.Data()["email"].(string),
			Password:  doc.Data()["password"].(string),
			CreatedAt: doc.Data()["created_at"].(time.Time),
		}
	}
	return &user, nil
}

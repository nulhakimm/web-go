package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/nulhakimm/web-go/model"
	"google.golang.org/api/iterator"
)

type ProjectRepo interface {
	FindAll(ctx context.Context) ([]*model.Project, error)
}

type ProjectRepoImpl struct {
	Client *firestore.Client
}

func NewProjectRepo(client *firestore.Client) ProjectRepo {
	return &ProjectRepoImpl{
		Client: client,
	}
}

func (repository *ProjectRepoImpl) FindAll(ctx context.Context) ([]*model.Project, error) {

	iter := repository.Client.Collection("projects").Documents(ctx)

	defer iter.Stop()

	var projects []*model.Project

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to iterate through documents: %v", err)
		}

		var project model.Project
		if err := doc.DataTo(&project); err != nil {
			return nil, fmt.Errorf("failed to decode document data: %v", err)
		}

		projects = append(projects, &project)
	}

	return projects, nil

}

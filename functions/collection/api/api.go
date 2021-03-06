package api

import (
	"os"
	"time"

	"github.com/guregu/dynamo"
	"github.com/pkg/errors"

	"github.com/gofrs/uuid"

	authenticator "github.com/myuon/portals-me/functions/authenticator/lib"
	. "github.com/myuon/portals-me/functions/collection/lib"
)

func DoList(
	userID string,
	entityTable dynamo.Table,
) ([]Collection, error) {
	var colDBOs []CollectionDBO
	if err := entityTable.
		Get("sort", "collection##detail").
		Index(os.Getenv("SortIndex")).
		Range("sort_value", dynamo.Equal, userID).
		All(&colDBOs); err != nil {
		return nil, err
	}

	var cols []Collection = make([]Collection, len(colDBOs))
	for index, col := range colDBOs {
		cols[index] = col.FromDBO()
	}

	return cols, nil
}

type GetOutput struct {
	Collection
	OwnerName string `json:"owner_name"`
}

func DoGet(
	collectionID string,
	entityTable dynamo.Table,
) (GetOutput, error) {
	var colDBO CollectionDBO
	if err := entityTable.
		Get("id", "collection##"+collectionID).
		Range("sort", dynamo.Equal, "collection##detail").
		One(&colDBO); err != nil {
		return GetOutput{}, err
	}

	col := colDBO.FromDBO()

	var userDBO authenticator.UserDBO
	if err := entityTable.
		Get("id", col.Owner).
		Range("sort", dynamo.Equal, "user##detail").
		One(&userDBO); err != nil {
		return GetOutput{}, err
	}

	user := userDBO.FromDBO()

	return GetOutput{
		Collection: col,
		OwnerName:  user.Name,
	}, nil
}

type CreateInput struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Cover       map[string]string `json:"cover"`
}

func DoCreate(
	createInput CreateInput,
	userID string,
	entityTable dynamo.Table,
) (string, error) {
	collectionID := uuid.Must(uuid.NewV4()).String()

	col := Collection{
		ID:             collectionID,
		Owner:          userID,
		Title:          createInput.Title,
		Description:    createInput.Description,
		Cover:          createInput.Cover,
		Media:          []string{},
		CommentMembers: []string{userID},
		CommentCount:   0,
		CreatedAt:      time.Now().Unix(),
	}

	if err := entityTable.Put(col.ToDBO()).Run(); err != nil {
		return "", errors.Wrap(err, "putCollection failed")
	}

	return collectionID, nil
}

func DoUpdate(
	collection Collection,
	entityTable dynamo.Table,
) error {
	if err := entityTable.
		Put(collection.ToDBO()).
		If("$ = ?", "sort_value", collection.Owner).
		Run(); err != nil {
		return err
	}

	return nil
}

func DoDelete(
	collectionID string,
	userID string,
	entityTable dynamo.Table,
) error {
	var idList []string
	if err := entityTable.
		Get("id", "collection##"+collectionID).
		Range("sort", dynamo.Equal, userID).
		Project("sort").
		All(&idList); err != nil {
		return err
	}

	var keys = make([]dynamo.Keyed, len(idList))
	for index, id := range idList {
		keys[index] = dynamo.Keys{"collection##" + collectionID, id}
	}
	if _, err := entityTable.
		Batch("id", "sort").
		Write().
		Delete(keys...).
		Run(); err != nil {
		return err
	}

	return nil
}

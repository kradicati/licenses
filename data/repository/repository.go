package repository

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Repository[T interface{}] interface {
	Get(id string, obj *T) error
	Insert(id string, obj *T) error
}

// IteratorToArray probably not the correct place to put this
func iteratorToArray[T interface{}](iter *firestore.DocumentIterator, array *[]T) error {
	defer iter.Stop()

	var object *T

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return err
		}

		err = doc.DataTo(object)

		if err != nil {
			return err
		}

		*array = append(*array, *object)
	}

	return nil
}

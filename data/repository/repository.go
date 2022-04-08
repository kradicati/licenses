package repository

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"google.golang.org/api/iterator"
)

// IteratorToArray probably not the correct place to put this
func iteratorToArray[T interface{}](iter *firestore.DocumentIterator, array *[]T) error {
	defer iter.Stop()

	object := new(T)

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
			fmt.Println("2", err.Error())
			return err
		}

		*array = append(*array, *object)
	}

	return nil
}

package service

import (
	"testing"

	"github.com/prateek69/go-grpc/sample"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	t.Run("LaptopSavedSuccessfully", func(t *testing.T) {
		t.Parallel()
		laptop := sample.GetLaptop()
		store := NewInMemoryLaptopStore()

		err := store.Save(laptop)
		assert.NoError(t, err)
	})
	t.Run("LaptopNotSaved", func(t *testing.T) {
		t.Parallel()
		laptop := sample.GetLaptop()
		store := NewInMemoryLaptopStore()

		err := store.Save(laptop)
		assert.NoError(t, err)

		err = store.Save(laptop)
		assert.Error(t, err, ErrorLaptopAlreadyExists)
	})
}

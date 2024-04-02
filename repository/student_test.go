package repository

import (
	"log"
	"testing"

	"rinto11804/activity_point_backend/storage"
)

func TestFetchStudentTotalPoints(t *testing.T) {
	studentrepo := NewStudentRepository(storage.GetDB())
	got, err := studentrepo.FetchStudentTotalPoints(6, "2023")
	if err != nil {
		t.Errorf("error occures %s", err)
	}
	log.Println(got)
}

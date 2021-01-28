package openapi

import (
	"context"
	"fmt"
	"os"
	"testing"
)

var client *APIClient

const testScheme = "http"

func TestMain(m *testing.M) {
	cfg := NewConfiguration()
	cfg.Scheme = testScheme
	client = NewAPIClient(cfg)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestFetchPosts(t *testing.T) {
	posts, r, err := client.PostsApi.FetchPosts(context.Background()).Execute()
	if err.Error() != "" {
		t.Fatalf("Error while getting pet by id: %v", err)
		t.Log(r)
	} else {

		t.Log(r)

	}

	for _, p := range posts {
		fmt.Printf("PostID: %v UserID: %v Title: %v Body: %v \n", p.Id, *p.UserId, *p.Title, *p.Body)
		fmt.Println("-------------------------------------------------------------------------------")
	}
}
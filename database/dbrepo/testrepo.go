package dbrepo

import (
	"github.com/schattenbrot/mini-blog-api/models"
)

func (m *testDBRepo) InsertPost(p models.Post) (*string, error) {
	return nil, nil
}

func (m *testDBRepo) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post

	return posts, nil
}

func (m *testDBRepo) UpdatePost(p models.Post) error {
	return nil
}

func (m *testDBRepo) DeleteOnePost(id string) error {
	return nil
}

func (m *testDBRepo) GetPostById(id string) (*models.Post, error) {
	return nil, nil
}

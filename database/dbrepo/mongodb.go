package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/schattenbrot/mini-blog-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrorTitleAndTextEmpty = "title and text cannot be empty"
var ErrorDocumentNotFound = "document not found"
var ErrorAlreadyUpToDate = "up to date"
var ErrorNameEmailPasswordEmpty = "either name or email or password cannot be empty"

// Post is the Post type used for communication with the mongo driver.
type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Text      string             `bson:"text,omitempty"`
	Creator   string             `bson:"creator,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

// User is the User type used for communication with the mongo driver.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Email     string             `bson:"email,omitempty" validate:"omitempty,email"`
	Password  string             `bson:"password,omitempty"`
	Roles     []string           `bson:"roles,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}

// toModelPost converts a mongoPost to a models.Post.
func toModelPost(post *Post) models.Post {
	var modelPost models.Post
	modelPost.ID = post.ID.Hex()
	modelPost.Title = post.Title
	modelPost.Text = post.Text
	modelPost.Creator = post.Creator
	modelPost.CreatedAt = post.CreatedAt
	modelPost.UpdatedAt = post.UpdatedAt

	return modelPost
}

// toModelUser converts a mongoUser to a models.User.
func toModelUser(user *User) models.User {
	var modelUser models.User
	modelUser.ID = user.ID.Hex()
	modelUser.Name = user.Name
	modelUser.Email = user.Email
	modelUser.Password = user.Password
	modelUser.Roles = user.Roles
	modelUser.CreatedAt = user.CreatedAt

	return modelUser
}

// InsertPost inserts a given post into the database.
// Returns the post ID of the inserted post and an error if any occurred.
func (m *mongoDBRepo) InsertPost(p models.Post) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var post Post
	post.Title = p.Title
	post.Text = p.Text
	post.Creator = p.Creator
	post.CreatedAt = p.CreatedAt
	post.UpdatedAt = p.UpdatedAt

	collection := m.DB.Collection("posts")

	result, err := collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	oid := result.InsertedID.(primitive.ObjectID).Hex()

	return &oid, nil
}

// GetPostById gets a post from the database by its ID.
// Returns a post and an error if any occurred.
func (m *mongoDBRepo) GetPostById(id string) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var post Post

	collection := m.DB.Collection("posts")

	filter := Post{ID: oid}

	err = collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return nil, err
	}

	modelPost := toModelPost(&post)

	return &modelPost, nil
}

// GetPosts gets a list of posts from the database.
// Returns a list of posts and an error if any occurred.
func (m *mongoDBRepo) GetPosts() ([]*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	posts := []*models.Post{}

	collection := m.DB.Collection("posts")

	filter := bson.D{}

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post Post
		cursor.Decode(&post)

		newPost := toModelPost(&post)

		posts = append(posts, &newPost)
	}

	return posts, nil
}

// GetPostCreator fetches the creator of a post from the database.
// Returns the creator's id and an error if any occurred.
func (m *mongoDBRepo) GetPostCreator(id string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	collection := m.DB.Collection("posts")

	filter := Post{ID: oid}

	proj := bson.M{"creator": 1}
	var options options.FindOneOptions
	options.SetProjection(proj)

	type Result struct {
		Creator primitive.ObjectID `bson:"creator"`
	}
	var result Result
	err = collection.FindOne(ctx, filter, &options).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Creator.Hex(), nil
}

// GetPostsByPage gets a list of posts by page number and page limit.
// Returns a list of posts and an error if any occurred.
func (m *mongoDBRepo) GetPostsByPage(page, limit int) ([]*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	posts := []*models.Post{}

	collection := m.DB.Collection("posts")

	filter := bson.M{}
	findOptions := options.Find()
	findOptions.SetSkip((int64(page) - 1) * int64(limit))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post Post
		cursor.Decode(&post)

		newPost := toModelPost(&post)

		posts = append(posts, &newPost)
	}

	return posts, nil
}

// UpdatePost updates a given post in the database.
// Returns an error if any occurred.
func (m *mongoDBRepo) UpdatePost(p models.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if p.Title == "" && p.Text == "" {
		return errors.New(ErrorTitleAndTextEmpty)
	}

	var post Post
	if p.Title != "" {
		post.Title = p.Title
	}
	if p.Text != "" {
		post.Text = p.Text
	}
	post.UpdatedAt = time.Now()

	collection := m.DB.Collection("posts")

	oid, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		return err
	}

	update := bson.M{"$set": post}

	result, err := collection.UpdateByID(ctx, oid, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		err = errors.New(ErrorDocumentNotFound)
		return err
	}

	if result.ModifiedCount == 0 {
		err = errors.New(ErrorAlreadyUpToDate)
		return err
	}

	return nil
}

// DeleteOnePost deletes one post from the database by its ID.
// Returns an error if any occurred.
func (m *mongoDBRepo) DeleteOnePost(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := m.DB.Collection("posts")

	filter := Post{ID: oid}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		err = errors.New(ErrorDocumentNotFound)
		return err
	}

	return nil
}

// InsertUser inserts a given user into the database.
// Returns the user ID of the inserted user and an error if any occurred.
func (m *mongoDBRepo) InsertUser(u models.User) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := User{
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Roles:     u.Roles,
		CreatedAt: time.Now(),
	}

	collection := m.DB.Collection("users")

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	oid := result.InsertedID.(primitive.ObjectID).Hex()

	return &oid, nil
}

// GetUserById retrieves a user from the database by its ID.
// Returns a user and an error if any occurred.
func (m *mongoDBRepo) GetUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := User{ID: oid}

	collection := m.DB.Collection("users")

	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	fetchedUser := toModelUser(&user)

	return &fetchedUser, nil
}

// GetUserRoles fetches the roles of a user from the database.
// Returns the user's roles and an error if any occurred.
func (m *mongoDBRepo) GetUserRoles(id string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := m.DB.Collection("users")

	filter := User{ID: oid}

	proj := bson.M{"roles": 1}
	var options options.FindOneOptions
	options.SetProjection(proj)

	type Result struct {
		Roles []string `bson:"roles"`
	}
	var result Result

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Roles, err
}

// GetUserByMail retrieves a user from the database by its email.
// Returns a user and an error if any occurred.
func (m *mongoDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User

	filter := User{Email: email}

	collection := m.DB.Collection("users")

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	fetchedUser := toModelUser(&user)

	return &fetchedUser, nil
}

// UpdateUser updates a given user.
// Returns an error if any occurred.
func (m *mongoDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if u.Name == "" && u.Email == "" && u.Password == "" {
		return errors.New(ErrorNameEmailPasswordEmpty)
	}

	var user User
	if u.Name != "" {
		user.Name = u.Name
	}
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.Password != "" {
		user.Password = u.Password
	}

	collection := m.DB.Collection("users")

	oid, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return err
	}

	update := bson.M{"$set": user}

	result, err := collection.UpdateByID(ctx, oid, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		err = errors.New(ErrorDocumentNotFound)
		return err
	}

	if result.ModifiedCount == 0 {
		err = errors.New(ErrorAlreadyUpToDate)
		return err
	}

	return nil
}

// DeleteUser deletes a user from the database by its ID.
// Returns an error if any occurred.
func (m *mongoDBRepo) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := m.DB.Collection("users")

	filter := User{ID: oid}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		err = errors.New(ErrorDocumentNotFound)
		return err
	}

	return nil
}

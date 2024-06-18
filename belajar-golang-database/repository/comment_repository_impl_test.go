package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	comment, err := CommentRepository.FindById(context.Background(), 100)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestDeleteById(t *testing.T) {
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	comment, err := CommentRepository.DeleteById(context.Background(), 41)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment.Comment)
}
func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	comments, err := CommentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

package repository

import (
	belajargolangdatabase "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository 2",
	}

	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comment, err := CommentRepository.FindById(ctx, 33)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comments, err := CommentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

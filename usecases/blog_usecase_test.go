package usecases_test

import (
	"app/entity"
	"app/graphql/model"
	"app/test/mock_repository"
	"app/usecases"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockBlogsRepository *mock_repository.MockBlogsRepository
var blogUsecase usecases.BlogUsecase

func blogTestSetUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlogsRepository = mock_repository.NewMockBlogsRepository(ctrl)

	blogUsecase = usecases.BlogUsecase{
		BlogsRepo: mockBlogsRepository,
	}

	return func() {
		defer ctrl.Finish()
	}
}

func TestBlogList(t *testing.T) {
	blogTestSetUp := blogTestSetUp(t)
	defer blogTestSetUp()

	var blogList []*entity.Blog

	for i := 1; i <= 5; i++ {
		blogList = append(blogList, &entity.Blog{
			ID:                 uint64(i),
			Title:              fmt.Sprintf("Title %s", strconv.Itoa(i)),
			Description:        fmt.Sprintf("Description %s", strconv.Itoa(i)),
			Body:               fmt.Sprintf("Body %s", strconv.Itoa(i)),
			NgihtBody:          fmt.Sprintf("NgihtBody %s", strconv.Itoa(i)),
			MobileBody:         fmt.Sprintf("MobileBody %s", strconv.Itoa(i)),
			ThumbnailImagePath: fmt.Sprintf("ThumbnailImagePath %s", strconv.Itoa(i)),
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		})
	}

	params := &model.PageCondition{PageNo: 0, Limit: 5}

	mockBlogsRepository.EXPECT().FindAll(params).Return(blogList, nil)
	mockBlogsRepository.EXPECT().TotalCount().Return(int64(5), nil)

	blogListConnection, err := blogUsecase.BlogList(params)

	assert.NoError(t, err)

	assert.Equal(t, strconv.Itoa(int(blogList[0].ID)), blogListConnection.Nodes[0].ID)
	assert.Equal(t, blogList[0].Title, blogListConnection.Nodes[0].Title)
	assert.Equal(t, blogList[0].ThumbnailImagePath, blogListConnection.Nodes[0].ThumbnailImagePath)
	assert.Equal(t, blogList[0].CreatedAt.Format("2006-01-02"), blogListConnection.Nodes[0].CreatedAt)
	assert.Equal(t, 5, blogListConnection.PageInfo.TotalCount)
}

func TestRecommendBlogList(t *testing.T) {
	blogTestSetUp := blogTestSetUp(t)
	defer blogTestSetUp()

	var blogList []*entity.Blog

	for i := 1; i <= 5; i++ {
		blogList = append(blogList, &entity.Blog{
			ID:                 uint64(i),
			Title:              fmt.Sprintf("Title %s", strconv.Itoa(i)),
			Description:        fmt.Sprintf("Description %s", strconv.Itoa(i)),
			Body:               fmt.Sprintf("Body %s", strconv.Itoa(i)),
			NgihtBody:          fmt.Sprintf("NgihtBody %s", strconv.Itoa(i)),
			MobileBody:         fmt.Sprintf("MobileBody %s", strconv.Itoa(i)),
			ThumbnailImagePath: fmt.Sprintf("ThumbnailImagePath %s", strconv.Itoa(i)),
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		})
	}

	params := &model.PageCondition{PageNo: 0, Limit: 3}

	mockBlogsRepository.EXPECT().FindAll(params).Return(blogList, nil)
	mockBlogsRepository.EXPECT().TotalCount().Return(int64(3), nil)

	blogListConnection, err := blogUsecase.BlogList(params)

	assert.NoError(t, err)

	assert.Equal(t, strconv.Itoa(int(blogList[0].ID)), blogListConnection.Nodes[0].ID)
	assert.Equal(t, blogList[0].Title, blogListConnection.Nodes[0].Title)
	assert.Equal(t, blogList[0].ThumbnailImagePath, blogListConnection.Nodes[0].ThumbnailImagePath)
	assert.Equal(t, blogList[0].CreatedAt.Format("2006-01-02"), blogListConnection.Nodes[0].CreatedAt)
}

func TestBlog(t *testing.T) {
	blogTestSetUp := blogTestSetUp(t)
	defer blogTestSetUp()

	var blog = &entity.Blog{
		ID:                 uint64(1),
		Title:              "Title",
		Description:        "Description",
		Body:               "Body",
		NgihtBody:          "NgihtBody",
		MobileBody:         "MobileBody",
		ThumbnailImagePath: "ThumbnailImagePath",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	mockBlogsRepository.EXPECT().Find(int(blog.ID)).Return(blog, nil)

	modelBlog, err := blogUsecase.Blog(int(blog.ID))

	assert.NoError(t, err)

	assert.Equal(t, strconv.Itoa(int(blog.ID)), modelBlog.ID)
	assert.Equal(t, blog.Title, modelBlog.Title)
	assert.Equal(t, blog.ThumbnailImagePath, modelBlog.ThumbnailImagePath)
	assert.Equal(t, blog.Body, modelBlog.Body)
	assert.Equal(t, blog.CreatedAt.Format("2006-01-02"), modelBlog.CreatedAt)
	assert.Equal(t, blog.UpdatedAt.Format("2006-01-02"), modelBlog.UpdatedAt)
}

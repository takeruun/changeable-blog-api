package mock_repository

import (
	entity "app/entity"
	"app/graphql/generated"
	"app/graphql/model"
	"app/graphql/resolver"
	"app/interactor"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

var mockBlogsRepository *MockBlogsRepository
var resolvers resolver.Resolver

func setUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlogsRepository = NewMockBlogsRepository(ctrl)

	interactor := interactor.BlogsInteractor{
		BlogsRepo: mockBlogsRepository,
	}

	resolvers = resolver.Resolver{BlogsInteractor: interactor}

	return func() {
		defer ctrl.Finish()
	}
}

func TestBlogList(t *testing.T) {
	setup := setUp(t)
	defer setup()

	var returnFindBlogList []*entity.Blog

	for i := 1; i <= 5; i++ {
		returnFindBlogList = append(returnFindBlogList, &entity.Blog{
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

	mockBlogsRepository.EXPECT().FindAll(params).Return(returnFindBlogList, nil)
	mockBlogsRepository.EXPECT().TotalCount().Return(int64(5), nil)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

	var resp struct {
		BlogList struct {
			Nodes []struct {
				ID, Title, ThumbnailImagePath, CreatedAt string
			}
			PageInfo struct {
				TotalCount int
			}
		}
	}

	query := `
		query {
			blogList(input: {pageNo: 0, limit: 5}) {
				nodes {
					id
					title
					thumbnailImagePath
					createdAt
				}
				pageInfo {
					totalCount
				}
			}
		}
	`

	c.MustPost(query, &resp)
	assert.Equal(t, strconv.Itoa(int(returnFindBlogList[0].ID)), resp.BlogList.Nodes[0].ID)
	assert.Equal(t, returnFindBlogList[0].Title, resp.BlogList.Nodes[0].Title)
	assert.Equal(t, returnFindBlogList[0].ThumbnailImagePath, resp.BlogList.Nodes[0].ThumbnailImagePath)
	assert.Equal(t, returnFindBlogList[0].CreatedAt.Format("2006-01-02"), resp.BlogList.Nodes[0].CreatedAt)
	assert.Equal(t, 5, resp.BlogList.PageInfo.TotalCount)
}

func TestRecommendBlogList(t *testing.T) {
	setup := setUp(t)
	defer setup()

	var returnFindRecommendBlogList []*entity.Blog

	for i := 1; i <= 5; i++ {
		returnFindRecommendBlogList = append(returnFindRecommendBlogList, &entity.Blog{
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

	mockBlogsRepository.EXPECT().FindAll(params).Return(returnFindRecommendBlogList, nil)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

	var resp struct {
		RecommendBlogList struct {
			Nodes []struct {
				ID, Title, ThumbnailImagePath, CreatedAt string
			}
		}
	}

	query := `
		query {
			recommendBlogList {
				nodes {
					id
					title
					thumbnailImagePath
					createdAt
				}
			}
		}
	`

	c.MustPost(query, &resp)
	assert.Equal(t, strconv.Itoa(int(returnFindRecommendBlogList[0].ID)), resp.RecommendBlogList.Nodes[0].ID)
	assert.Equal(t, returnFindRecommendBlogList[0].Title, resp.RecommendBlogList.Nodes[0].Title)
	assert.Equal(t, returnFindRecommendBlogList[0].ThumbnailImagePath, resp.RecommendBlogList.Nodes[0].ThumbnailImagePath)
	assert.Equal(t, returnFindRecommendBlogList[0].CreatedAt.Format("2006-01-02"), resp.RecommendBlogList.Nodes[0].CreatedAt)
}

func TestBlog(t *testing.T) {
	setup := setUp(t)
	defer setup()

	var returnBlog = &entity.Blog{
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

	mockBlogsRepository.EXPECT().Find(int(returnBlog.ID)).Return(returnBlog, nil)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

	var resp struct {
		Blog struct {
			ID, Title, Body, ThumbnailImagePath, CreatedAt, UpdatedAt string
		}
	}

	query := `
		query {
			blog(id: 1) {
				id
				title
				body
				thumbnailImagePath
				createdAt
				updatedAt
			}
		}
	`

	c.MustPost(query, &resp)
	assert.Equal(t, strconv.Itoa(int(returnBlog.ID)), resp.Blog.ID)
	assert.Equal(t, returnBlog.Title, resp.Blog.Title)
	assert.Equal(t, returnBlog.ThumbnailImagePath, resp.Blog.ThumbnailImagePath)
	assert.Equal(t, returnBlog.Body, resp.Blog.Body)
	assert.Equal(t, returnBlog.CreatedAt.Format("2006-01-02"), resp.Blog.CreatedAt)
	assert.Equal(t, returnBlog.UpdatedAt.Format("2006-01-02"), resp.Blog.UpdatedAt)
}

func TestNotFoundBlog(t *testing.T) {
	setup := setUp(t)
	defer setup()

	var returnErr error = gorm.ErrRecordNotFound

	mockBlogsRepository.EXPECT().Find(1).Return(nil, returnErr)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

	var resp struct {
		Blog struct {
			ID, Title, Body, ThumbnailImagePath, CreatedAt, UpdatedAt string
		}
	}

	query := `
		query {
			blog(id: 1) {
				id
				title
				body
				thumbnailImagePath
				createdAt
				updatedAt
			}
		}
	`

	errResp := c.Post(query, &resp)
	assert.Containsf(t, errResp.Error(), "Error record not found", "expected error containing %s", "Error record not found")
	assert.Containsf(t, errResp.Error(), "NOT_FOUND_ERROR", "expected error containing error code %s", "NOT_FOUND_ERROR")
}

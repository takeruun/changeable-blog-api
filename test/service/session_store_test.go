package service

import (
	"app/config"
	"app/middleware"
	"app/service"

	"app/test/test_utils"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

var SSService service.SessionStoreServiceRepository

func SetUp(db *gorm.DB) {
	os.Setenv("GO_MODE", "test")
	store := config.NewSessionStore(db)
	SSService = service.NewSessionStoreService(store)
}

func NewRouter() *chi.Mux {
	r := chi.NewMux()
	r.Use(middleware.InjectHTTPMiddleware())

	return r
}

func TestGetSession(t *testing.T) {
	db := test_utils.NewDB(t)
	SetUp(db)

	r := NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Context().Value(service.HTTPKey("http")).(service.HTTP); !ok {
			panic("missing http context")
		}

		session, err := SSService.GetSession(r.Context(), "session")
		if err != nil {
			t.Errorf("Error get session %s", err.Error())
		}

		if !reflect.DeepEqual(session.Values, map[interface{}]interface{}{}) {
			t.Errorf("Sessoin value not empty \n got: %s", session.Values)
		}
	})

	ts := httptest.NewUnstartedServer(r)
	ts.Start()

	defer ts.Close()

	test_utils.TestRequest(t, ts, "GET", "/", nil)
}

func TestSaveSession(t *testing.T) {
	var (
		key   string = "userId"
		value string = "1"
	)

	db := test_utils.NewDB(t)

	defer func() {
		db.Exec("DELETE FROM sessions")
	}()

	SetUp(db)

	r := NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Context().Value(service.HTTPKey("http")).(service.HTTP); !ok {
			panic("missing http context")
		}

		session, err := SSService.GetSession(r.Context(), "session")
		if err != nil {
			t.Errorf("Error get session %s", err.Error())
		}

		SSService.SaveValue(r.Context(), key, value)

		if !reflect.DeepEqual(session.Values, map[interface{}]interface{}{key: value}) {
			t.Errorf("Sessoin value not %s \n got: %s", map[interface{}]interface{}{key: value}, session.Values)
		}
	})

	ts := httptest.NewUnstartedServer(r)
	ts.Start()

	defer ts.Close()

	test_utils.TestRequest(t, ts, "GET", "/", nil)
}

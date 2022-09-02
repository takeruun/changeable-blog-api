package test

import (
	"app/middleware"
	"app/service"
	"app/test/test_utils"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	return chi.NewMux()
}

// HTTP の context が取得されているかのテスト
// 公式参考 : https://github.com/go-chi/chi/blob/master/mux_test.go#L1697

func TestHttpContext(t *testing.T) {
	r := NewRouter()
	r.Use(middleware.InjectHTTPMiddleware())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Context().Value(http.ServerContextKey).(*http.Server); !ok {
			panic("missing server context")
		}

		if _, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); !ok {
			panic("missing local addr context")
		}

		if _, ok := r.Context().Value(service.HTTPKey("http")).(service.HTTP); !ok {
			panic("missing http context")
		}
	})

	ts := httptest.NewUnstartedServer(r)
	ts.Start()

	defer ts.Close()

	test_utils.TestRequest(t, ts, "GET", "/", nil)
}

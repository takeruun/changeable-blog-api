package test

import (
	"app/middleware"
	"app/service"
	"io"
	"io/ioutil"
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

	testRequest(t, ts, "GET", "/", nil)
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}

package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "router"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloController(t *testing.T) {
	w := httptest.NewRecorder()
	Convey("Test GetHello", t, func() {
		r, _ := http.NewRequest(http.MethodGet, "/hello?name=Lyric", nil)
		http.DefaultServeMux.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, 200)
		So(w.Body.String(), ShouldEqual, "Hello,Lyric")
	})
}

func TestSpec(t *testing.T) {
    // Only pass t into top-level Convey calls
    Convey("Given some integer with a starting value", t, func() {

        Convey("When the integer is incremented", func() {
            So(1, ShouldEqual, 1)
        })
    })
}

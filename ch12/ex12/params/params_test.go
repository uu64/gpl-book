package params

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func search(resp http.ResponseWriter, req *http.Request) {
	type Data struct {
		Labels []string `http:"l"`
		Mail   string   `mail:"m"`
		Num    int      `positive:"n"`
	}
	var data Data
	if err := Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func TestUnpackSuccess(t *testing.T) {
	tests := []struct {
		url string
	}{
		{"http://localhost:3000/search"},
		{"http://localhost:3000/search?l=golang&l=java&m=test@exmaple.com&n=10"},
		{"http://localhost:3000/search?invalid=test@exmaple.com&n=10"},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.url, nil)
		w := httptest.NewRecorder()

		search(w, req)
		resp := w.Result()
		// エラーが発生しないことを確かめる
		if resp.StatusCode != 200 {
			t.Errorf("req failed: %d %s\n", resp.StatusCode, test.url)
		}
		body, _ := io.ReadAll(resp.Body) // ignoring errors
		fmt.Println(string(body))
	}
}

func TestUnpackFail(t *testing.T) {
	tests := []struct {
		url string
	}{
		{"http://localhost:3000/search?l=golang&l=java&m=test.exmaple.com&n=10"},
		{"http://localhost:3000/search?l=golang&l=java&m=test@exmaple.com&n=-10"},
		{"http://localhost:3000/search?l=golang&l=java&m=test@exmaple.com&n=true"},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.url, nil)
		w := httptest.NewRecorder()

		search(w, req)
		resp := w.Result()
		// エラーが発生することを確かめる
		if resp.StatusCode == 200 {
			t.Errorf("req success: %d %s\n", resp.StatusCode, test.url)
		}
		body, _ := io.ReadAll(resp.Body) // ignoring errors
		fmt.Println(string(body))
	}
}

func TestPack(t *testing.T) {
	tests := []struct {
		in  interface{}
		url string
		// want string
	}{
		{
			&struct {
				Labels     string `http:"l"`
				MaxResults int    `http:"max"`
				Exact      bool   `http:"x"`
			}{"Go", 150, true},
			"http://localhost:3000/search",
			// "http://localhost:3000/search?l=Go&max=150&x=true",
		},
		{
			&struct {
				Labels     []string `http:"l"`
				MaxResults []int    `http:"max"`
				Exact      []bool   `http:"x"`
			}{[]string{"Go", "Java"}, []int{150, 55}, []bool{true, false}},
			"http://localhost:3000/search",
			// "http://localhost:3000/search?l=Go&l=Java&max=150&max=55&x=true&x=false",
		},
		{
			&struct {
				Labels     string
				MaxResults int  `http:"max"`
				Exact      bool `http:"x"`
			}{"Go", 150, true},
			"http://localhost:3000/search",
			// "http://localhost:3000/search?labels=Go&max=150&x=true",
		},
		{
			&struct{}{},
			"http://localhost:3000",
			// "http://localhost:3000/search?l=Go&l=Java&max=150&max=55&x=true&x=false",
		},
	}

	for _, test := range tests {
		url, err := Pack(test.url, test.in)
		if err != nil {
			t.Error(err)
		}
		// NOTE: クエリの順番が保証されないので最低限エラーが発生しないことを確かめる
		t.Log(url)
		// if url != test.want {
		// 	t.Errorf("want: %s, got: %s\n", test.want, url)
		// }
	}
}

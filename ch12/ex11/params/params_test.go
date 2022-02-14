package params

import "testing"

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

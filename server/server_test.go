package server

import "testing"

func TestHandling(t *testing.T) {
	testData := []struct {
		name string
		body []byte
		res  string
		err  error
	}{
		{
			name: "good request",
			body: []byte(`{"object":"page","entry":["id":"<PAGE_ID>","messaging":[]]}`),
			res:  "page",
			err:  nil,
		},
		{
			name: "bad request with no object",
			body: []byte(`{"no_object":"page","entry":["id":"<PAGE_ID>","messaging":[]]}`),
			res:  "",
			err:  errNoObjectKey,
		},
		{
			name: "bad request with no entry",
			body: []byte(`{"object":"page","no_entry":["id":"<PAGE_ID>","messaging":[]]}`),
			res:  "",
			err:  errNoEntryKey,
		},
	}

	for _, v := range testData {
		t.Run(v.name, func(t *testing.T) {
			res, err := handleRequest(v.body)
			if err != v.err {
				t.Errorf("handleRequest().error = %v, expected: %v", err, v.err)
			}

			if res != v.res {
				t.Errorf("handleRequest() = %q, expected: %q", res, v.res)
			}
		})
	}
}

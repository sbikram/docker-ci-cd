package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestUnit(t *testing.T) {
	cases := []struct{ question, reply string }{
		{"check", "If you are reading this that means your application is UP"},
		{"api", "API is working"},
	}

	for _, x := range cases {
		r := httptest.NewRequest("GET", "http://dummy/"+x.question, nil)
		w := httptest.NewRecorder()
		index(w, r)

		if body, err := ioutil.ReadAll(w.Body); err != nil {
			t.Error(err)
		} else if string(body) != x.reply {
			t.Error("oops we have a problem: expected reply - ", x.reply, ", but got - ", string(body))
		}
	}
}

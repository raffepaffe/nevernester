package below_four

import (
	"fmt"
	"net/http"
)

type Hero struct {
	Kind string
	Name string
}

type Identity struct {
}

func (i Identity) String() (string, error) {
	return "", nil
}

type source struct {
	Id Identity
}

type pair struct {
	source  *source
	compare *source
}

func findDiffs(pair pair) ([]Hero, error) {
	diffs := []Hero{}

	if pair.compare == nil {
		id, err := pair.source.Id.String()
		if err != nil {
			return []Hero{}, nil
		}
		return []Hero{
			{
				Kind: "kind.Good",
				Name: fmt.Sprintf("no id '%s' found", id),
			},
		}, nil
	}

	return diffs, nil
}

func aHandler() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if ctx == nil {
				w.WriteHeader(500)
				return
			}

		})
}

package json

import (
	"net/http"

	json "github.com/json-iterator/go"
)

// EmptySliceRender implements gin render interface almost as render.JSON,
// except it marshall nil slices as empty array([]).
type EmptySliceRender struct {
	Data interface{}
}

func (r EmptySliceRender) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	jsonBytes, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonBytes)
	return err
}

func (r EmptySliceRender) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
}

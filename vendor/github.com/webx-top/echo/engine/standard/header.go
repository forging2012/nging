package standard

import "net/http"

type (
	Header struct {
		header http.Header
	}
)

func (h *Header) Add(key, val string) {
	h.header.Add(key, val)
}

func (h *Header) Del(key string) {
	h.header.Del(key)
}

func (h *Header) Get(key string) string {
	return h.header.Get(key)
}

func (h *Header) Set(key, val string) {
	h.header.Set(key, val)
}

func (h *Header) reset(hdr http.Header) {
	h.header = hdr
}

func (h *Header) Object() interface{} {
	return h.header
}

func (h *Header) Std() http.Header {
	return h.header
}

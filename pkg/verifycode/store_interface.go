package verifycode

type Store interface {
	Set(id string, code string) bool

	Get(id string, clear bool) string

	Verify(id, answer string, clear bool) bool
}

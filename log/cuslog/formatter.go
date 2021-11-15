package cuslog

const (
	FmtEmptySeparate = ""
)

type Formatter interface {
	Format(entry *Entry) error
}

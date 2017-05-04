package layoutparam

const (
	FILL_PARENT  = -1
	MATCH_PARENT = -1
	WRAP_CONTENT = -2
)

type SizeEntry struct {
	Width  int
	Height int
}

func (sizeEntry *SizeEntry) SetSize(width, height int) {
	sizeEntry.Width = width
	sizeEntry.Height = height
}

func NewLayoutParam(widthNew, heightNew int) *SizeEntry {
	return &SizeEntry{
		Width:  widthNew,
		Height: heightNew,
	}
}

func test() {
	test2 := NewLayoutParam(1, 2)
	test2.SetSize(2, 3)
}

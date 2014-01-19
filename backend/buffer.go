package backend

type Buffer struct {
    data []rune
}

func (b *Buffer) Size() int {
    return len(b.data)
}

func (buf *Buffer) Substr(r Region) string {
    l := len(buf.data)
    a, b := clamp(0, 1, r.Begin()), clamp(0, 1, r.End())
    return string(buf.data[a:b])
}
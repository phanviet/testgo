package backend

type View struct {
    HasSettings
    filename string
    window *window
    buffer *Buffer
}

func (v *View) Window() *Window {
    return v.window
}

func (v *View) Buffer() *Buffer {
    return v.buffer
}
package backend

type Editor struct {
    HasSettings
}

type Window struct {
    HasSettings
    views []View
}

func (w *Window) NewView() *View {
    w.views = append(w.views, View{window: w})
    v := &w.views[len(w.views) - 1]
    v.setBuffer(&Buffer{})
    v.selection.Regions = []Region{{0, 0}}
    return v
}
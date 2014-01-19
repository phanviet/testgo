package sublime

import "lime/backend"

type (
    Region backend.Region
    View Struct {
        bv *backend.View
        scratch bool
    }

    Settings backend.Settings
    Window backend.Window
)

func (v *View) Size() int {
    return v.bv.Buffer().size()
}

func (v *View) IsScratch() bool {
    return v.scratch
}

func (v *View) Settings() *Settings {
    return (*Settings) (v.bv.Settings())
}
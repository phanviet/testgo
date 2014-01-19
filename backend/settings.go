package backend

type HasSettings struct {
    settings Settings
}

func (s *HasSettings) Settings() *Settings {
    if s.settings.data == nil {
        s.settings = NewSettings()
    }
    return &s.settings
}

type OnChangeCallback func()
type settingsMap map[string]interface {}
type Settings struct {
    OnChangeCallbacks map[string]OnChangeCallback
    data settingsMap
}

func NewSettings() Settings {
    return Settings{make(map[string]OnChangeCallback), make(settingsMap)}
}

func (s *Settings) AddOnChange(key string, cb OnChangeCallback) {
    s.OnChangeCallbacks[key] = cb
}

func (s *Settings) ClearOnChange(key string) {
    s.OnChangeCallbacks[key] = nil
}

func (s *Settings) Get(name string, def ...interface{}) interface{} {
    if v, ok := s.data[name]; ok {
        return v
    }
    return def
}

func (s *Settings) Set(name string, val interface{}) {
    s.data[name] = val
    s.onChange()
}

func (s *Settings) onChange() {
    for _, v := range s.OnChangeCallbacks {
        v()
    }
}

func (s *Settings) Erase(name string) {
    s.data[name] = nil
}

func (s *Settings) merge(other SettingsMap) {
    for k, v := range other {
        s.data[k] = v
    }
    s.onChange()
}
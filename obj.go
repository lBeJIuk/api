package api

import "image"

type StreamDeckInfo struct {
	Cols     int    `json:"cols,omitempty"`
	Rows     int    `json:"rows,omitempty"`
	IconSize int    `json:"icon_size,omitempty"`
	Page     int    `json:"page"`
	Profile  string `json:"profile"`
	Serial   string `json:"serial,omitempty"`
}

type Page []Key

type Profile struct {
	Name  string `json:"name"`
	Pages []Page `json:"pages"`
}

type DeprecatedDeck struct {
	Serial string `json:"serial"`
	Pages  []Page `json:"pages"`
}

type Deck struct {
	Serial   string    `json:"serial"`
	Profiles []Profile `json:"profiles"`
}

type Config struct {
	Modules []string `json:"modules,omitempty"`
	Decks   []Deck   `json:"decks"`
}

type DepracatedConfig struct {
	Modules []string `json:"modules,omitempty"`
	Pages   []Page   `json:"pages"`
}

type DepracatedConfig_v2 struct {
	Modules []string         `json:"modules,omitempty"`
	Decks   []DeprecatedDeck `json:"decks"`
}

type Key struct {
	Icon              string            `json:"icon,omitempty"`
	SwitchPage        int               `json:"switch_page,omitempty"`
	Text              string            `json:"text,omitempty"`
	TextSize          int               `json:"text_size,omitempty"`
	TextAlignment     string            `json:"text_alignment,omitempty"`
	Keybind           string            `json:"keybind,omitempty"`
	Command           string            `json:"command,omitempty"`
	Brightness        int               `json:"brightness,omitempty"`
	Url               string            `json:"url,omitempty"`
	IconHandler       string            `json:"icon_handler,omitempty"`
	KeyHandler        string            `json:"key_handler,omitempty"`
	IconHandlerFields map[string]string `json:"icon_handler_fields,omitempty"`
	KeyHandlerFields  map[string]string `json:"key_handler_fields,omitempty"`
	Buff              image.Image       `json:"-"`
	IconHandlerStruct IconHandler       `json:"-"`
	KeyHandlerStruct  KeyHandler        `json:"-"`
}

type Module struct {
	Name       string
	IconFields []Field
	KeyFields  []Field
	IsIcon     bool
	IsKey      bool
}

type Field struct {
	Title     string
	Name      string
	Type      string
	FileTypes []string
	Values    []string
}

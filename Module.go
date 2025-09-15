package api

type Module struct {
	Name             string  `json:"name,omitempty"`
	IconFields       []Field `json:"icon_fields,omitempty"`
	KeyFields        []Field `json:"key_fields,omitempty"`
	IsIcon           bool    `json:"is_icon,omitempty"`
	IsKey            bool    `json:"is_key,omitempty"`
	IsLinkedHandlers bool    `json:"is_linked_handlers,omitempty"`
}

type Field struct {
	Title     string   `json:"title,omitempty"`
	Name      string   `json:"name,omitempty"`
	Type      string   `json:"type,omitempty"`
	FileTypes []string `json:"file_types,omitempty"`
	ListItems []string `json:"list_items,omitempty"`
}

package model

// NodeElement represents a DOM element node.
type NodeElement struct {
	// Name of the DOM element. Available tags: a, aside, b, blockquote, br, code, em, figcaption, figure, h3, h4, hr,
	// i, iframe, img, li, ol, p, pre, s, strong, u, ul, video.
	Tag string `json:"tag"`
	// Attributes of the DOM element. Key of object represents name of attribute, value represents value of attribute.
	// Available attributes: href, src.
	Attrs map[string]string `json:"attrs,omitempty"`
	// List of child nodes for the DOM element.
	Children []Node `json:"children,omitempty"`
}

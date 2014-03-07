package elements

import "encoding/xml"

// Create an (empty) Parsable to parse XML into
type Constructor func() Element

// Maintain a mapping between tag names (and namespaces) and Constructors
type ElementFactory map[string]Constructor

func NewElementFactory() ElementFactory {
	return ElementFactory(make(map[string]Constructor))
}

func (self ElementFactory) AddConstructor(key string, constructor Constructor) {
	self[key] = constructor
}

// Call a constructor for specified key or "*", if defined. Otherwise return an error
func (self ElementFactory) Get(name string) (Element, error) {
	if constructor, ok := self[name]; ok {
		return constructor(), nil
	}

	// This is default constructor if defined
	if constructor, ok := self["*"]; ok {
		return constructor(), nil
	}

	return &InnerXML{}, nil
}

func (self ElementFactory) DecodeElement(d *xml.Decoder, element *xml.StartElement) (interface{}, error) {
	elementObject, err := self.Get(element.Name.Space + " " + element.Name.Local)
	if err != nil {
		return nil, err
	}
	if err := d.DecodeElement(elementObject, element); err != nil {
		return nil, err
	}

	return elementObject, nil
}
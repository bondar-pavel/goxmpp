package auth

import (
	"encoding/xml"

	"github.com/dotdoom/goxmpp/stream"
	"github.com/dotdoom/goxmpp/stream/elements"
	"github.com/dotdoom/goxmpp/stream/elements/features"
)

type mechanismsElement struct {
	XMLName    xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl mechanisms"`
	Mechanisms []string `xml:"mechanism"`
	*features.Container
}

func newMechanismsElement() *mechanismsElement {
	return &mechanismsElement{
		Container: features.NewContainer(),
	}
}

func (self *mechanismsElement) IsRequiredFor(stream *stream.Stream) bool {
	var state *State
	err := stream.State.Get(&state)
	return err != nil || state.UserName == ""
}

func (self *mechanismsElement) CopyIfAvailable(stream *stream.Stream) elements.Element {
	if self.IsRequiredFor(stream) {
		x := newMechanismsElement()
		mechanisms_element.CopyAvailableFeatures(stream, x.Container)
		return x
	}
	return nil
}

var mechanisms_element = newMechanismsElement()

var mechanism_handlers map[string]Handler = make(map[string]Handler)

func AddMechanism(name string, handler Handler) {
	mechanisms_element.AddElement(&MechanismElement{Name: name})
	mechanism_handlers[name] = handler
}

func init() {
	features.Tree.AddElement(mechanisms_element)
}

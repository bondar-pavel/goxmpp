package bind

import (
	"encoding/xml"
	"github.com/dotdoom/goxmpp/stream/connection"
	"github.com/dotdoom/goxmpp/stream/elements/features"
)

type bind struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-bind bind"`
	features.Elements
}

func (self *bind) IsRequiredFor(fs connection.State) bool {
	return fs["bound"] == nil
}

func (self *bind) CopyIfAvailable(fs connection.State) interface{} {
	if self.IsRequiredFor(fs) && fs["authenticated"] != nil {
		return self.CopyAvailableFeatures(fs, new(bind))
	}
	return nil
}

func init() {
	features.List.AddElement(new(bind))
}

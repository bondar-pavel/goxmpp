package compression

import (
	"encoding/xml"

	"github.com/dotdoom/goxmpp/stream"
	"github.com/dotdoom/goxmpp/stream/elements"
)

func init() {
	stream.StreamFactory.AddConstructor(func() elements.Element {
		return NewCompressionHandler()
	})
}

var CompressionFactory = elements.NewFactory()

func NewCompressionHandler() *CompressionHandler {
	return &CompressionHandler{InnerElements: elements.NewInnerElements(CompressionFactory)}
}

type BaseCompression struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl compression"`
}

// This struct is used for marshaling
type CompressionFeature struct {
	BaseCompression
	*elements.InnerElements
}

// This struct is used for unmarshaling and stream handling
type CompressionHandler struct {
	BaseCompression
	*elements.InnerElements
}

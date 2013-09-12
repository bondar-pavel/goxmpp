package goxmpp

import (
	_ "github.com/dotdoom/goxmpp/extensions/features/compression"
	_ "github.com/dotdoom/goxmpp/extensions/features/compression/method"
	_ "github.com/dotdoom/goxmpp/extensions/features/starttls"
	_ "github.com/dotdoom/goxmpp/extensions/iq/disco_info"
	_ "github.com/dotdoom/goxmpp/extensions/iq/disco_items"
	_ "github.com/dotdoom/goxmpp/extensions/iq/last"
	_ "github.com/dotdoom/goxmpp/extensions/iq/ping"
	_ "github.com/dotdoom/goxmpp/extensions/iq/privacy"
	_ "github.com/dotdoom/goxmpp/extensions/iq/query"
	_ "github.com/dotdoom/goxmpp/extensions/iq/query/item"
	_ "github.com/dotdoom/goxmpp/extensions/iq/stats"
	_ "github.com/dotdoom/goxmpp/extensions/iq/time"
	_ "github.com/dotdoom/goxmpp/extensions/iq/version"
	_ "github.com/dotdoom/goxmpp/stream"
)

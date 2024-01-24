package zammad

type TicketState string

const (
	NewTicketState    TicketState = "new"
	OpenTicketState   TicketState = "open"
	ClosedTicketState TicketState = "closed"
)

// TicketSearchResult represents the results from a search
// We currently only care about the assets in which the tickets
// are contained
type TicketSearchResult struct {
	Assets Assets `json:"assets"`
}

// Assets represents the assets in the search result
type Assets struct {
	Tickets map[string]Ticket `json:"Ticket"`
}

// Ticket represents a Zammad Ticket
// We use two custom field attributes for the tickets
// icinga_host and icinga_service to track existing tickets
type Ticket struct {
	ID            int     `json:"id,omitempty"`
	Title         string  `json:"title"`
	Group         string  `json:"group"`
	Customer      string  `json:"customer"`
	IcingaHost    string  `json:"icinga_host"`
	IcingaService string  `json:"icinga_service"`
	Article       Article `json:"article,omitempty"`
}

// Article represents a Zammad Ticket Article
type Article struct {
	TicketID    int    `json:"ticket_id,omitempty"`
	Internal    bool   `json:"internal"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	ContentType string `json:"content_type"`        // "text/html"
	Type        string `json:"type"`                // "phone"
	Sender      string `json:"sender"`              // "Agent"
	TimeUnit    string `json:"time_unit,omitempty"` // "15"
}

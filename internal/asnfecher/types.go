package asnfetcher

// RipeAdressBlock Adrress range returned by Ripe
type RipeAdressBlock struct {
	Resource string `json:"resource"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
}

// RipeResourceResponseBlock the actual content date of the response
type RipeResourceResponseBlock struct {
	Resource       string          `json:"resource"`
	Type           string          `json:"type"`
	QueryStarTtime string          `json:"query_starttime"`
	Block          RipeAdressBlock `json:"block"`
	QueryEndtime   string          `json:"queryEndtime"`
	Holder         string          `json:"holder"`
	Announced      bool            `json:"announced"`
}

// RipeQueryResponse JSON respone message from ripe query
type RipeQueryResponse struct {
	Status         string     `json:"status"`
	ServerID       string     `json:"server_id"`
	StatusCode     int        `json:"status_code"`
	Version        string     `json:"version"`
	Cached         bool       `json:"cached"`
	SeeAlso        []string   `json:"see_also"`
	Time           string     `json:"time"`
	Messages       [][]string `json:"messages"`
	DataCallStatus string     `json:"data_call_status"`
	ProcessTime    int        `json:"process_time"`
	BuildVersion   string     `json:"build_version"`
	QueryID        string     `json:"query_id"`
}

// RipeAsOverviewResponse Response of AS Overview query
type RipeAsOverviewResponse struct {
	RipeQueryResponse `json:",inline"`
	Data              RipeResourceResponseBlock `json:"data"`
}

// RipeAnnouncedPrefixesResponse Reponse of Assigned prefixes by an AS
type RipeAnnouncedPrefixesResponse struct {
	RipeQueryResponse `json:",inline"`
	Data              AnnouncedPrefixesDataBlock `json:"data"`
}

// AnnouncedPrefixesDataBlock The data field of announced prefixed data bloxk
type AnnouncedPrefixesDataBlock struct {
	Resource string            `json:"resource"`
	Prefixes []PrefixAsingment `json:"prefixes"`
}

// PrefixAsingment An assigment of an address range withing a timeline
type PrefixAsingment struct {
	Timelines []RipeTimeline `json:"timelines"`
	Prefix    string         `json:"prefix"`
}

// RipeTimeline time interval
type RipeTimeline struct {
	Endtime   string `json:"endtime"`
	Starttime string `json:"starttime"`
}

package storage

type SMS []*SMSData
type MMS []*MMSData

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type MMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func NewSMS(country, bandwidth, responseTime, provider string) SMSData {
	return SMSData{
		Country:      country,
		Bandwidth:    bandwidth,
		ResponseTime: responseTime,
		Provider:     provider,
	}
}

func New() SMS {
	var s []*SMSData
	return s
}

func (s *SMS) Put(sms *SMSData) {
	*s = append(*s, sms)
}

func (m *MMS) Put(mms *MMSData) {
	*m = append(*m, mms)
}

func (s *SMS) Get() *SMS {
	return s
}

package info

type JobPost struct {
	Name            string   `json:"name"`
	Company         string   `json:"company"`
	Techniques      []string `json:"techniques"`
	Location        string   `json:"location"`
	Career          string   `json:"career"`
	Link            string   `json:"link"`
	RecruitmentSite `json:"recruitmentSite"`
}

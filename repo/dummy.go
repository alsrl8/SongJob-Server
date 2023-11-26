package repo

import "SongJob_Server/info"

func GetDummyJobPost() info.JobPost {
	return info.JobPost{
		Name:            "백엔드 엔지니어",
		Company:         "다이렉트클라우드랩",
		Techniques:      []string{"GCP", "Golang", "AWS"},
		Location:        "Seoul",
		Career:          "3 years",
		Link:            "https://example.com",
		RecruitmentSite: info.Jumpit,
	}
}

func GetDummyJobPosts() []info.JobPost {
	return []info.JobPost{
		{
			Name:            "예제",
			Company:         "어떤 회사",
			Techniques:      []string{"어떤 기술1", "어떤 기술2", "어떤 기술3"},
			Location:        "서울",
			Career:          "몇 년",
			Link:            "https://example.com",
			RecruitmentSite: info.Saramin,
		},
		{
			Name:            "백엔드 엔지니어",
			Company:         "다이렉트클라우드랩",
			Techniques:      []string{"GCP", "Golang", "AWS"},
			Location:        "Seoul",
			Career:          "3 years",
			Link:            "https://www.jumpit.co.kr/position/15779",
			RecruitmentSite: info.Jumpit,
		},
		{
			Name:            "백엔드 엔지니어",
			Company:         "다이렉트클라우드랩",
			Techniques:      []string{"GCP", "Golang", "AWS"},
			Location:        "Seoul",
			Career:          "3 years",
			Link:            "https://www.jumpit.co.kr/position/15779",
			RecruitmentSite: info.Jumpit,
		},
		{
			Name:            "백엔드 엔지니어",
			Company:         "다이렉트클라우드랩",
			Techniques:      []string{"GCP", "Golang", "AWS"},
			Location:        "Seoul",
			Career:          "3 years",
			Link:            "https://www.jumpit.co.kr/position/15779",
			RecruitmentSite: info.Jumpit,
		},
		{
			Name:            "백엔드 엔지니어",
			Company:         "다이렉트클라우드랩",
			Techniques:      []string{"GCP", "Golang", "AWS"},
			Location:        "Seoul",
			Career:          "3 years",
			Link:            "https://www.jumpit.co.kr/position/15779",
			RecruitmentSite: info.Jumpit,
		},
	}
}

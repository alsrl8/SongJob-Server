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
			Link:            "https://www.wikipedia.org",
			RecruitmentSite: info.Saramin,
		},
		{
			Name:            "백엔드 엔지니어",
			Company:         "백엔드 엔지니어를 뽑는 회사",
			Techniques:      []string{"GCP", "Golang", "AWS"},
			Location:        "Seoul",
			Career:          "3 years",
			Link:            "https://www.stackoverflow.com",
			RecruitmentSite: info.Jumpit,
		},
		{
			Name:            "프론트엔드 개발자",
			Company:         "다이나믹 웹사이트",
			Techniques:      []string{"React", "JavaScript", "HTML5"},
			Location:        "부산",
			Career:          "2 years",
			Link:            "https://www.github.com",
			RecruitmentSite: info.JobKorea,
		},
		{
			Name:            "데이터 과학자",
			Company:         "데이터 분석 회사",
			Techniques:      []string{"Python", "R", "SQL"},
			Location:        "대구",
			Career:          "5 years",
			Link:            "https://www.medium.com",
			RecruitmentSite: info.Saramin,
		},
		{
			Name:            "풀스택 개발자",
			Company:         "스타트업 A",
			Techniques:      []string{"Node.js", "React", "MongoDB"},
			Location:        "인천",
			Career:          "4 years",
			Link:            "https://www.khanacademy.org",
			RecruitmentSite: info.Jumpit,
		},
		{
			Name:            "모바일 개발자",
			Company:         "모바일 앱 개발",
			Techniques:      []string{"Swift", "Kotlin", "Flutter"},
			Location:        "광주",
			Career:          "1 year",
			Link:            "https://www.ted.com",
			RecruitmentSite: info.JobKorea,
		},
		{
			Name:            "클라우드 엔지니어",
			Company:         "클라우드 서비스 회사",
			Techniques:      []string{"AWS", "Azure", "Docker"},
			Location:        "대전",
			Career:          "6 years",
			Link:            "https://www.coursera.com",
			RecruitmentSite: info.Saramin,
		},
		{
			Name:            "시스템 관리자",
			Company:         "IT 인프라 회사",
			Techniques:      []string{"Linux", "Networking", "Security"},
			Location:        "울산",
			Career:          "8 years",
			Link:            "https://www.edx.com",
			RecruitmentSite: info.Jumpit,
		},
		{
			Name:            "AI 엔지니어",
			Company:         "인공지능 개발사",
			Techniques:      []string{"TensorFlow", "PyTorch", "Machine Learning"},
			Location:        "제주",
			Career:          "7 years",
			Link:            "https://www.nasa.com",
			RecruitmentSite: info.JobKorea,
		},
		{
			Name:            "UX/UI 디자이너",
			Company:         "디자인 스튜디오",
			Techniques:      []string{"Sketch", "Adobe XD", "Figma"},
			Location:        "서울",
			Career:          "2 years",
			Link:            "https://www.nationalgeographic.com",
			RecruitmentSite: info.Saramin,
		},
	}
}

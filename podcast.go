package main

import "html/template"

type PodcastEpisode struct {
	Number      int
	AffiliationNumber      int
	Title       string
	Guest       string
	GuestURL       string
	Affiliation string
	AffiliationURL string
	Description template.HTML
	Resources   []ResourceLink
	Date        string // optional
	EpisodeURLs EpisodeURLs
}

type ResourceLink struct {
	Url  string
	Name string
}

type Podcast struct {
	Title           string
	Subtitle        string
	Description     template.HTML
	SpotifyUrl      string
	ApplePodcastUrl string
	YouTubeUrl      string
	LastUpdate      string
	Episodes        []PodcastEpisode
}

type EpisodeURLs struct {
	Spotify string
	ApplePodcasts string
	YouTube string
}

var HonestMajorityPodcast = Podcast{
	Title:           "Honest Majority",
	Subtitle:        "a podcast series by Common Prefix",
	Description:     template.HTML(`Honest Majority by Common Prefix explores the people and ideas shaping trust technology. Guests share their journey into the space, the technical problems they are solving, and where they see the next five years of secure, scalable, and interoperable distributed systems headed.`),
	SpotifyUrl:      "https://open.spotify.com/show/3dZtQUzhbcnv9I5qZmaRkF?si=2996b42327d444fa",
	ApplePodcastUrl: "https://podcasts.apple.com/us/podcast/honest-majority-by-common-prefix/id1860881793",
	YouTubeUrl:      "https://www.youtube.com/@CommonPrefix",
	LastUpdate:      "December 12, 2025",
	Episodes: []PodcastEpisode{
		{
			Number:      1,
			AffiliationNumber: 2,
			Title:       "Can Data Availability be Zero-Knowledge",
			Guest:       "Alex Evans",
			GuestURL:    "https://x.com/alexhevans",
			Affiliation: "Bain Capital Crypto",
			AffiliationURL: "https://x.com/baincapcrypto",
			Description: template.HTML(`In the inaugural episode of Honest Majority, we speak with Alex Evans (Bain Capital Crypto) about the evolving relationship between data availability and zero-knowledge proofs in blockchain systems. 
The conversation begins with Alex’s path into crypto and research, before diving into the fundamentals of data availability and recent advances in zero-knowledge data availability constructions. We examine the theoretical guarantees these approaches offer, the trade-offs they introduce in terms of assumptions and complexity, and where current research still falls short. The episode concludes with a discussion of open problems and how these ideas may shape the design of future blockchain protocols.`),
			Resources: []ResourceLink{
				{
					Url:  "https://arxiv.org/pdf/1809.09044",
					Name: "⁠Fraud and Data Availability Proofs: Maximising Light Client Security and Scaling Blockchains with Dishonest Majorities",
				},
				{
					Url:  "https://eprint.iacr.org/2025/034.pdf",
					Name: "ZODA: Zero-Overhead Data Availability",
				},
			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/432Fe1a4yZPYUbZXztzaKo?si=2a8f4dab29364550",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/alex-evans-can-data-availability-be-zero-knowledge/id1860881793?i=1000741007403",
				YouTube: "https://youtu.be/YhChZn2s9Uc?si=1otW6qokJ_vOc8Q5",
			},
		},
	},
}




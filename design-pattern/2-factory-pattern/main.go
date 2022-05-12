package main

import "fmt"

// Content is a Factory Pattern.
type Content interface {
	Play()
	ArtistInfo()
}

// AnimeContent is a concrete class of Content.
type AnimeContent struct {
	Name string
}

func (c *AnimeContent) Play() {
	fmt.Println("Playing", c.Name)
}

func (c *AnimeContent) ArtistInfo() {
	fmt.Println("Artist:", "A.I.M.")
}

func (c *AnimeContent) CustomFuncAnime() {
	fmt.Println("Custom func anime")
}

// KDramaContent is a concrete class of Content.
type KDramaContent struct {
	Name string
}

func (c *KDramaContent) Play() {
	fmt.Println("Playing", c.Name)
}

func (c *KDramaContent) ArtistInfo() {
	fmt.Println("Artist:", "KBS")
}

// Studio is a Factory Pattern.
type Studio interface {
	CreateContent(name string) Content
}

// Netflix is a concrete class of Studio.
type Netflix struct{}

func (n *Netflix) CreateContent(name string) Content {
	return &AnimeContent{Name: name}
}

// Mappa is a concrete class of Studio.
type Mappa struct{}

func (m *Mappa) CreateContent(name string) Content {
	return &KDramaContent{Name: name}
}

func main() {
	n := new(Netflix)

	content1 := n.CreateContent("AoT")
	content1.Play()
	content1.ArtistInfo()

	m := new(Mappa)
	content2 := m.CreateContent("Drama Korea Satu")
	content2.Play()
	content2.ArtistInfo()

	animeC := new(AnimeContent)
	animeC.Name = "Anime"
	animeC.Play()

	kdramaC := new(KDramaContent)
	kdramaC.Name = "KDrama"
	kdramaC.Play()
}

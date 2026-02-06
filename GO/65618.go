package main

import "strings"
import "fmt"
import "sort"


type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() {
	base := float64(p.Goals) + float64(p.Assists)/2
	if p.Misses == 0 {
		p.Rating = base
	} else {
		p.Rating = base / float64(p.Misses)
	}
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}
	p.calculateRating()
	return p
}


func goalsSort(players []Player) []Player {
	sorted := make([]Player, len(players))
	copy(sorted, players)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Goals != sorted[j].Goals {
			return sorted[i].Goals > sorted[j].Goals
		}
		return strings.ToLower(sorted[i].Name) < strings.ToLower(sorted[j].Name)
	})
	return sorted
}

func ratingSort(players []Player) []Player {
	sorted := make([]Player, len(players))
	copy(sorted, players)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Rating != sorted[j].Rating {
			return sorted[i].Rating > sorted[j].Rating
		}
		return strings.ToLower(sorted[i].Name) < strings.ToLower(sorted[j].Name)
	})
	return sorted
}

func gmSort(players []Player) []Player {
	sorted := make([]Player, len(players))
	copy(sorted, players)
	sort.Slice(sorted, func(i, j int) bool {
		var ratioI, ratioJ float64
		if sorted[i].Misses == 0 {
			ratioI = float64(sorted[i].Goals)
		} else {
			ratioI = float64(sorted[i].Goals) / float64(sorted[i].Misses)
		}
		if sorted[j].Misses == 0 {
			ratioJ = float64(sorted[j].Goals)
		} else {
			ratioJ = float64(sorted[j].Goals) / float64(sorted[j].Misses)
		}
		if ratioI != ratioJ {
			return ratioI > ratioJ
		}
		return strings.ToLower(sorted[i].Name) < strings.ToLower(sorted[j].Name)
	})
	return sorted
}



func main() {
	players := []Player{
		NewPlayer("Арсений", 10, 2, 4),
		NewPlayer("Иван", 8, 0, 6),
		NewPlayer("Пётр", 12, 3, 2),
		NewPlayer("Сергей", 7, 1, 5),
	}

	fmt.Println("Сортировка по головам:")
	for _, p := range goalsSort(players) {
		fmt.Printf("%s: %d голов\n", p.Name, p.Goals)
	}

	fmt.Println("\nСортировка по рейтингу:")
	for _, p := range ratingSort(players) {
		fmt.Printf("%s: рейтинг %.2f\n", p.Name, p.Rating)
	}

	fmt.Println("\nСортировка по отношению голов к промахам:")
	for _, p := range gmSort(players) {
		var ratio float64
		if p.Misses == 0 {
			ratio = float64(p.Goals)
		} else {
			ratio = float64(p.Goals) / float64(p.Misses)
		}
		fmt.Printf("%s: %.2f\n", p.Name, ratio)
	}
}

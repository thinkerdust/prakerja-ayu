package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score []int
}

func (s *Student) AddScore(score int) {
	s.Score = append(s.Score, score)
}

func (s *Student) CalculateAverageScore() float64 {
	sum := 0
	for _, score := range s.Score {
		sum += score
	}
	average := float64(sum) / float64(len(s.Score))
	return average
}

func (s *Student) GetMinimumScore() int {
	minScore := math.MaxInt64
	for _, score := range s.Score {
		if score < minScore {
			minScore = score
		}
	}
	return minScore
}

func (s *Student) GetMaximumScore() int {
	maxScore := math.MinInt64
	for _, score := range s.Score {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	students := make([]Student, 0)

	// Memasukkan data siswa sebanyak 5 siswa
	for i := 0; i < 5; i++ {
		var name string
		var scores []int
		fmt.Printf("Masukkan nama siswa ke-%d: ", i+1)
		fmt.Scanln(&name)
		fmt.Println("Masukkan skor siswa (pisahkan dengan spasi):")
		for j := 0; j < 5; j++ {
			var score int
			fmt.Scan(&score)
			scores = append(scores, score)
		}
		student := Student{
			Name:  name,
			Score: scores,
		}
		students = append(students, student)
	}

	// Menampilkan skor rata-rata, siswa dengan skor minimum, dan siswa dengan skor maksimum
	for i, student := range students {
		averageScore := student.CalculateAverageScore()
		minScore := student.GetMinimumScore()
		maxScore := student.GetMaximumScore()

		fmt.Printf("Siswa %s:\n", student.Name)
		fmt.Printf("Skor rata-rata: %.2f\n", averageScore)
		fmt.Printf("Skor minimum: %d\n", minScore)
		fmt.Printf("Skor maksimum: %d\n", maxScore)

		// Pemisah antara setiap siswa
		if i != len(students)-1 {
			fmt.Println("---------------------")
		}
	}
}

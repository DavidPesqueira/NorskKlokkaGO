package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	for {
		// Generate random hours and minutes
		hours := rand.Intn(24)
		minutes := rand.Intn(60)

		// Format the correct answer in Norwegian-style time-telling
		correctAnswer := "Klokka er " + formatNorwegianTime(hours, minutes)
		acceptedAnswers := generateAcceptedAnswers(hours, minutes)

		// Display the random 24-hour time
		fmt.Printf("Hva er klokken? %02d:%02d\n", hours, minutes)

		var userInput string
		correct := false
		for !correct {
			fmt.Print("Skriv inn tid (f.eks., 'ti over to'): ")
			userInput, _ = reader.ReadString('\n')
			userInput = strings.TrimSpace(userInput)

			// Check if the user wants to see the correct answer
			if strings.ToLower(userInput) == "s" {
				fmt.Printf("Riktig svar: %s\n", correctAnswer)
				fmt.Printf("Andre aksepterte svar: %s\n", strings.Join(acceptedAnswers, ", "))
				break
			}

			// Display the user's input
			fmt.Println("Du svarte:", userInput)

			// Check if the user's input matches any accepted answers
			if isAnswerAccepted(userInput, correctAnswer, acceptedAnswers) {
				fmt.Printf("Riktig! %s\n", correctAnswer)
				correct = true
			} else {
				fmt.Println("Feil! Prøv igjen eller trykk 's' for å se riktig svar.")
			}
		}

		// Ask if the user wants to try again or exit
		fmt.Println("Prøv igjen eller avslutte?")
		fmt.Println("1. Igjen")
		fmt.Println("2. Avslutt")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "2" {
			fmt.Println("Ha det bra!")
			break
		}
	}
}

// formatNorwegianTime converts hours and minutes to a Norwegian-style time format
func formatNorwegianTime(hours, minutes int) string {
	hours = hours % 12
	if hours == 0 {
		hours = 12
	}

	nextHour := (hours % 12) + 1
	if nextHour > 12 {
		nextHour = 1
	}

	if minutes == 0 {
		return fmt.Sprintf("%s", numberToNorwegian(hours))
	} else if minutes == 30 {
		return fmt.Sprintf("halv %s", numberToNorwegian(nextHour))
	} else if minutes == 15 {
		return fmt.Sprintf("kvart over %s", numberToNorwegian(hours))
	} else if minutes == 45 {
		return fmt.Sprintf("kvart på %s", numberToNorwegian(nextHour))
	} else if minutes < 30 {
		return handleMinutesLessThanHalfPast(hours, minutes)
	} else {
		return handleMinutesMoreThanHalfPast(nextHour, minutes)
	}
}

func handleMinutesLessThanHalfPast(hours, minutes int) string {
	switch minutes {
	case 1:
		return fmt.Sprintf("ett minutt over %s", numberToNorwegian(hours))
	case 2:
		return fmt.Sprintf("to minutter over %s", numberToNorwegian(hours))
	case 3:
		return fmt.Sprintf("tre minutter over %s", numberToNorwegian(hours))
	case 4:
		return fmt.Sprintf("fire minutter over %s", numberToNorwegian(hours))
	case 5:
		return fmt.Sprintf("fem over %s", numberToNorwegian(hours))
	case 10:
		return fmt.Sprintf("ti over %s", numberToNorwegian(hours))
	case 20:
		return fmt.Sprintf("tjue over %s", numberToNorwegian(hours))
	default:
		return fmt.Sprintf("%s minutter over %s", numberToNorwegian(minutes), numberToNorwegian(hours))
	}
}

func handleMinutesMoreThanHalfPast(nextHour, minutes int) string {
	switch minutes {
	case 31:
		return fmt.Sprintf("ett minutt på %s", numberToNorwegian(nextHour))
	case 32:
		return fmt.Sprintf("to minutter på %s", numberToNorwegian(nextHour))
	case 33:
		return fmt.Sprintf("tre minutter på %s", numberToNorwegian(nextHour))
	case 34:
		return fmt.Sprintf("fire minutter på %s", numberToNorwegian(nextHour))
	case 35:
		return fmt.Sprintf("fem over halv %s", numberToNorwegian(nextHour))
	case 40:
		return fmt.Sprintf("ti over halv %s", numberToNorwegian(nextHour))
	case 50:
		return fmt.Sprintf("ti på %s", numberToNorwegian(nextHour))
	case 55:
		return fmt.Sprintf("fem på %s", numberToNorwegian(nextHour))
	default:
		return fmt.Sprintf("%s minutter på %s", numberToNorwegian(60-minutes), numberToNorwegian(nextHour))
	}
}

// numberToNorwegian converts an integer to its Norwegian text representation
func numberToNorwegian(n int) string {
	numbers := []string{
		"null", "ett", "to", "tre", "fire", "fem", "seks", "sju", "åtte", "ni",
		"ti", "elleve", "tolv", "tretten", "fjorten", "femten", "seksten", "sytten",
		"atten", "nitten", "tjue", "tjueen", "tjueto", "tjuetre", "tjuefire", "tjuefem",
		"tjueseks", "tjueni", "trettien", "trettito", "trettitre", "trettifire", "trettiseks",
		"trettifem", "trettisju", "førti", "femti", "seksti", "sytti", "åtti", "nitti",
	}
	if n < 0 || n >= len(numbers) {
		return ""
	}
	return numbers[n]
}

// generateAcceptedAnswers generates a list of accepted answers for a given hour and minute
func generateAcceptedAnswers(hours, minutes int) []string {
	var answers []string
	answers = append(answers, formatNorwegianTime(hours, minutes))

	if minutes == 0 {
		return answers
	}

	nextHour := (hours + 1) % 24
	if minutes < 30 {
		answers = append(answers, fmt.Sprintf("klokka er %s minutter over %s", numberToNorwegian(minutes), numberToNorwegian(hours)))
	} else if minutes == 30 {
		answers = append(answers, fmt.Sprintf("klokka er halv %s", numberToNorwegian(nextHour)))
	} else {
		answers = append(answers, fmt.Sprintf("klokka er %s minutter på %s", numberToNorwegian(60-minutes), numberToNorwegian(nextHour)))
	}

	if minutes >= 15 && minutes < 30 {
		answers = append(answers, fmt.Sprintf("klokka er kvart over %s", numberToNorwegian(hours)))
	} else if minutes >= 30 && minutes < 45 {
		answers = append(answers, fmt.Sprintf("klokka er kvart på %s", numberToNorwegian(nextHour)))
	}

	return answers
}

// isAnswerAccepted checks if the user input matches any accepted answers
func isAnswerAccepted(userInput, correctAnswer string, acceptedAnswers []string) bool {
	// Normalize the user input for comparison
	userInput = strings.ToLower(userInput)
	correctAnswer = strings.ToLower(correctAnswer)

	// Check if the user's input matches the correct answer or any accepted alternatives
	if userInput == correctAnswer {
		return true
	}
	for _, answer := range acceptedAnswers {
		if userInput == answer {
			return true
		}
	}
	return false
}

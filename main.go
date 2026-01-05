// Vladislav Ebert
// 241RDB316
package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Bus struct {
	Sākuma_pietura string `json:"Sākuma_pietura"`
	Gala_pietura   string `json:"Gala_pietura"`
	Nedēļas_diena  string `json:"Nedēļas_diena"`
	Laiks          string `json:"Laiks"`
	Biļetes_cena   string `json:"Biļetes_cena"`
}

var Nedēļas_dienas = map[string]bool{"Pr": true,
	"Ot": true,
	"Tr": true,
	"Ce": true,
	"Pt": true,
	"St": true,
	"Sv": true}

func CLaiks(t string) bool {
	if len(t) != 5 || t[2] != ':' {
		return false
	}
	for i := 0; i < 5; i++ {
		if i != 2 && (t[i] < '0' || t[i] > '9') {
			return false
		}
	}
	return true
}

func CBiļetes_cena(p string) bool {
	parts := strings.Split(p, ".")
	if len(parts) != 2 || len(parts[1]) != 2 {
		return false
	}
	for _, c := range parts[0] + parts[1] {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func processFile() ([]Bus, []string) {
	file, _ := os.Open("db.csv")
	defer file.Close()

	var valid []Bus
	var errors []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		r, _ := csv.NewReader(strings.NewReader(line)).ReadAll()
		if len(r) != 1 || len(r[0]) != 5 {
			errors = append(errors, line)
			continue
		}

		area := make([]string, 5)
		for i, f := range r[0] {
			area[i] = strings.TrimSpace(f)
		}

		Sākuma_pietura, Gp, Nedēļas_diena, Laiks, Biļetes_cena := area[0], area[1], area[2], area[3], area[4]
		if !Nedēļas_dienas[Nedēļas_diena] || !CLaiks(Laiks) || !CBiļetes_cena(Biļetes_cena) {
			errors = append(errors, line)
			continue
		}

		valid = append(valid, Bus{Sākuma_pietura, Gp, Nedēļas_diena, Laiks, Biļetes_cena})
	}
	return valid, errors
}

func SaglJson(data []Bus) {
	file, _ := os.Create("bus.json")
	defer file.Close()
	json.NewEncoder(file).Encode(data)
}

func SaglKļūd(errors []string) {
	file, _ := os.Create("err.txt")
	defer file.Close()
	for _, e := range errors {
		fmt.Fprintln(file, e)
	}
}

func ParbMaršrutus(maršruti []Bus) {
	fmt.Println("result:")
	for _, r := range maršruti {
		fmt.Printf("%s %s %s %s %s\n", r.Sākuma_pietura,
			r.Gala_pietura,
			r.Nedēļas_diena,
			r.Laiks,
			r.Biļetes_cena)
	}
}

func main() {
	maršruti, errs := processFile()
	SaglJson(maršruti)
	SaglKļūd(errs)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		cmd := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch cmd {
		case "print":
			scanner.Scan()
			Sākuma_pietura := strings.TrimSpace(scanner.Text())
			scanner.Scan()
			Gp := strings.TrimSpace(scanner.Text())
			var matches []Bus
			for _, r := range maršruti {
				if strings.EqualFold(r.Sākuma_pietura, Sākuma_pietura) && strings.EqualFold(r.Gala_pietura, Gp) {
					matches = append(matches, r)
				}
			}
			ParbMaršrutus(matches)

		case "find":
			scanner.Scan()
			Nedēļas_diena := strings.TrimSpace(scanner.Text())
			var matches []Bus
			for _, r := range maršruti {
				if r.Nedēļas_diena == Nedēļas_diena {
					matches = append(matches, r)
				}
			}
			ParbMaršrutus(matches)

		case "error":
			data, _ := os.ReadFile("err.txt")
			fmt.Println("result:")
			fmt.Print(string(data))

		case "json":
			data, _ := os.ReadFile("bus.json")
			fmt.Println("result:")
			fmt.Print(string(data))

		case "exit":
			return
		}
	}
}

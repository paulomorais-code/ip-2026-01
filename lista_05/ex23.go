package main

import f "fmt"

func main() {
	corredor := []int{0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	janelas := []int{0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var disponivel []int
	var sel string

	f.Println("Escolha uma área, Janela ou Corredor")
	f.Scan(&sel)

	if sel == "corredor" || sel == "Corredor" || sel == "CORREDOR" {
		for i, c := range corredor {
			if c == 0 {
				disponivel = append(disponivel, i+1)
			}
		}
		if len(disponivel) == 0 {
			f.Println("Não há poltronas disponiveis")
		} else {
			f.Printf("As poltronas %d estão disponiveis\n", disponivel)
		}
	}

	if sel == "corredor" || sel == "Corredor" || sel == "CORREDOR" {
		for i, c := range corredor {
			if c == 0 {
				disponivel = append(disponivel, i+1)
			}
		}
		if len(disponivel) == 0 {
			f.Println("Não há poltronas disponiveis")
		} else {
			f.Printf("As poltronas %d estão disponiveis\n", disponivel)
		}
	} else if sel == "janela" || sel == "Janela" || sel == "JANELA" {
		for i, c := range janelas {
			if c == 0 {
				disponivel = append(disponivel, i+1)
			}
		}
		if len(disponivel) == 0 {
			f.Println("Não há poltronas disponiveis")
		} else {
			f.Printf("As poltronas %d estão disponiveis\n", disponivel)
		}
	} else {
		f.Println("Erro")
	}

}

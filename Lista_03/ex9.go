package main

import f "fmt"

type alunos struct {
	nome string

	nota1 float64

	nota2 float64
}

func main() {
	var pessoa string
	var nota1, nota2, media float64
	var al []alunos
	soma := 0.0
	ContA := 0
	ContR := 0
	ContE := 0

	for {
		f.Println("Escreva o nome do aluno ou FIM para sair:")
		f.Scan(&pessoa)

		if pessoa == "FIM" {
			break
		}
		f.Println("Escreva a primeira nota:")
		f.Scan(&nota1)

		f.Println("Escreva a segunda nota:")
		f.Scan(&nota2)

		novoaluno := alunos{
			nome:  pessoa,
			nota1: nota1,
			nota2: nota2,
		}

		al = append(al, novoaluno)
	}

	for _, p := range al {
		media = (p.nota1 + p.nota2) / 2
		soma = soma + media
		if media < 3.0 {
			f.Printf("O aluno %s de média %.2f está reprovado\n", p.nome, media)
			ContR++
		} else if media >= 7.0 {
			f.Printf("O aluno %s de média %.2f está aprovado\n", p.nome, media)
			ContA++
		} else {
			f.Printf("O aluno %s de média %.2f está de exame\n", p.nome, media)
			ContE++
		}
	}
	mediaG := soma / float64(len(al))
	f.Printf("%d Alunos foram aprovados\n", ContA)
	f.Printf("%d Alunos foram reprovados\n", ContR)
	f.Printf("%d Alunos estão de exame\n", ContE)
	f.Println(" A média geral foi:", mediaG)
}

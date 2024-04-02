package main

import "fmt"
import "strings"
import "time"

type cliente  struct { 
	nome string
	idade int
	empregado bool
	tempo_empregado float64
	salario float64
}

func main(){

	//exercicio 1
	fmt.Println("\n \n Exercicio 1\n")
	palavra := "palavrao"
	fmt.Printf("A palavra contem: %v letras \n", strings.Count(palavra, ""))

	letras := strings.Split(palavra, "")

	for i := 0; i < len(letras); i++{
		fmt.Println(letras[i])
	}
	

	//exercicio 2
	fmt.Println("\n \n Exercicio 2\n")
	cliente1 := cliente{nome: "Cesar", idade: 36, empregado: true, tempo_empregado: 12, salario: 1000}
	cliente2 := cliente{nome: "Augusto", idade: 36, empregado: true, tempo_empregado: 3, salario: 200000}
	cliente3 := cliente{nome: "Oliveira", idade: 21, empregado: true, tempo_empregado: 3, salario: 200000}
	cliente4 := cliente{nome: "Silva", idade: 23, empregado: false, tempo_empregado: 3, salario: 200000}
	cliente5 := cliente{nome: "Adriano", idade: 24, empregado: true, tempo_empregado: 1, salario: 200000}
	cliente6 := cliente{nome: "Rocha", idade: 24, empregado: true, tempo_empregado: 2, salario: 2000}
	
	var clientes []cliente

	clientes = append(clientes, cliente1)
	clientes = append(clientes, cliente2)
	clientes = append(clientes, cliente3)
	clientes = append(clientes, cliente4)
	clientes = append(clientes, cliente5)
	clientes = append(clientes, cliente6)

	for j := 0; j < len(clientes); j++{
		fmt.Println("Cliente sendo verificado: ")
		fmt.Println(clientes[j])
		if (clientes[j].idade > 22 && clientes[j].empregado  && clientes[j].tempo_empregado > 1){
			
			fmt.Println("Pode conceder emprestimo ao cliente ",clientes[j].nome)
			if (clientes[j].salario > 100000){
				fmt.Println("Cliente com salario maior que 100.000 emprestimo sem juros ")
			}
		} else {
			fmt.Println("Cliente fora dos criterios de emprestimo")
		}

		
	}

	fmt.Println("Cliente: ", cliente1)

	//exercicio 3
	fmt.Println("\n \n Exercicio 3\n")

	t := time.Now()
	//var mes string
	//mes = t.Format(("January"))

	fmt.Printf("Hoje é dia %v de %v \n",t.Day(), t.Month())
	fmt.Printf("Numero do Mes: %02d", t.Month())


	// exercicio 4
	fmt.Println("\n \n Exercicio 4\n")
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	//De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.
	fmt.Println("Idade de Benjamin", employees["Benjamin"])
	// Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
	employees["Federico"] = 25
	// Excluir Pedro do mapa.
	delete(employees, "Pedro")

	// Saiba quantos de seus funcionários têm mais de 21 anos.

	for key, value := range employees {

		if value > 25 { 
		fmt.Printf("%q possui a idade de %v\n", key, value)
		}

	}

}

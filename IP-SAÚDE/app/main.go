package main

// Importa os pacotes necessários para o funcionamento do programa
import (
	"fmt"                       // Usado para imprimir mensagens no terminal
	"log"                       // Usado para registrar mensagens de erro ou log
	"net"                       // Usado para obter informações de rede, como IPs
	"net/http"                  // Usado para criar o servidor HTTP
	"servidorHTTP/app/handlers" // Importa os handlers definidos na aplicação
	"servidorHTTP/app/utils"    // Importa utilitários como conexão ao banco de dados
)

func main() {
	// Conecta ao banco de dados utilizando a função definida em utils
	utils.ConnectToDB()

	// Garante que a tabela de pacientes exista antes de subir a aplicação
	if err := utils.EnsurePatientTable(); err != nil {
		log.Fatal(err)
	}

	// Define a rota raiz e a página inicial como conteúdo dinâmico
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/index.html", handlers.HomeHandler)

	// Serve os estilos estáticos da aplicação
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./static/styles"))))

	// Define a rota "/registerPatient" e associa ao handler de cadastro de paciente
	http.HandleFunc("/registerPatient", handlers.RegisterPatientHandler)

	// Define a rota "/updatePatient" e associa ao handler de atualização de paciente
	http.HandleFunc("/updatePatient", handlers.UpdatePatientHandler)

	// Define a rota "/dischargePatient" e associa ao handler de saída de paciente
	http.HandleFunc("/dischargePatient", handlers.DischargePatientHandler)

	// Obtém os endereços de rede disponíveis na máquina
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		// Encerra o programa caso ocorra um erro ao obter os endereços
		log.Fatal(err)
	}

	var localIP string
	// Itera sobre os endereços de rede para encontrar o IP local
	for _, addr := range addrs {
		// Verifica se o endereço é um IP válido (não é loopback e é IPv4)
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			// Armazena o IP local encontrado
			localIP = ipNet.IP.String()
			break
		}
	}

	// Define a porta padrão do servidor como "3000"
	port := "3000"

	// Caso nenhum IP local seja encontrado, utiliza "127.0.0.1" como padrão
	if localIP == "" {
		localIP = "127.0.0.1"
	}

	// Imprime no terminal o endereço e a porta onde o servidor está rodando
	fmt.Printf("Servidor rodando em: http://%s:%s/\n", localIP, port)

	// Inicia o servidor HTTP na porta especificada e escuta conexões
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		// Encerra o programa caso ocorra um erro ao iniciar o servidor
		log.Fatal(err)
	}
}

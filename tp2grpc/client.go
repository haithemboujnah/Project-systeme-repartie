package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"tp2grpc/factorization_pb" // Généré à partir du fichier .proto
)

func main() {
	// Connexion au serveur
	conn, err := grpc.Dial("172.20.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := factorization_pb.NewPrimeFactorizationServiceClient(conn)

	// Demander à l'utilisateur de saisir un nombre
	var number int64
	fmt.Print("Enter a number to factorize: ")
	_, err = fmt.Scanln(&number)
	if err != nil {
		log.Fatalf("could not read input: %v", err)
	}

	// Envoi de la requête pour factoriser le nombre
	req := &factorization_pb.NumberRequest{Number: number}
	resp, err := client.GetPrimeFactors(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get prime factors: %v", err)
	}

	// Affichage du résultat
	fmt.Printf("Prime factors: %v\n", resp.Factors)
}

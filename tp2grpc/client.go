package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"tp2grpc/factorization_pb"
	"bufio"
	"os"
)

func main() {
	conn, err := grpc.Dial("172.20.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := factorization_pb.NewPrimeFactorizationServiceClient(conn)

	for {
		fmt.Println("\nChoose the operation:")
		fmt.Println("1. Request/Response")
		fmt.Println("2. Bidirectional Streaming")
		fmt.Println("0. Exit")
		fmt.Print("Enter choice (0, 1 or 2): ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			log.Fatalf("could not read input: %v", err)
		}

		switch choice {
		case 1:
			handleUnaryRequestResponse(client)
		case 2:
			handleBidirectionalStreaming(client)
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose 0, 1, or 2.")
		}
	}
}


func handleUnaryRequestResponse(client factorization_pb.PrimeFactorizationServiceClient) {
	var number int64
	fmt.Print("Enter a number to factorize: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		log.Fatalf("could not read input: %v", err)
	}

	req := &factorization_pb.NumberRequest{Number: number}
	resp, err := client.GetPrimeFactors(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get prime factors: %v", err)
	}

	
	fmt.Printf("Prime factors: %v\n", resp.Factors)
}


func handleBidirectionalStreaming(client factorization_pb.PrimeFactorizationServiceClient) {
	stream, err := client.GetPrimeFactorsStream(context.Background())
	if err != nil {
		log.Fatalf("could not open stream: %v", err)
	}

	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a number to factorize (or type '0' to exit): ")
		scanner.Scan()
		input := scanner.Text()

		
		if input == "0" {
			break
		}
		
		var number int64
		_, err := fmt.Sscanf(input, "%d", &number)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		
		req := &factorization_pb.NumberRequest{Number: number}
		err = stream.Send(req)
		if err != nil {
			log.Fatalf("could not send request: %v", err)
		}

		
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("could not receive response: %v", err)
		}
		fmt.Printf("Prime factors: %v\n", resp.Factors)
	}

	
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("could not close stream: %v", err)
	}
}

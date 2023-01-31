package main

import (
	"blockchainEtheriumGRPCclient/pkg/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("gRPC Client")
	createWallet("five", 21000)
	getWallet("five")
	getWallet("one")
	sendMoney("five", "one", 50)
	getWallet("five")
	getWallet("one")
}

func getWallet(walletName string) {
	conn := connectGrpcServer()
	defer conn.Close()
	c := service.NewGrpcServiceClient(conn)

	response, err := c.GetWallet(context.Background(), &service.Message{Body: walletName})
	if err != nil {
		log.Fatalf("Error when calling GetWallet: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}

func createWallet(walletName string, balance int64) {
	conn := connectGrpcServer()
	defer conn.Close()
	c := service.NewGrpcServiceClient(conn)

	response, err := c.CreateWallet(context.Background(), &service.WalletCreate{WalletName: walletName, Balance: balance})
	if err != nil {
		log.Fatalf("Error when calling GetWallet: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}

func sendMoney(walletSender string, wallletRecipient string, sendmoney int64) {
	conn := connectGrpcServer()
	defer conn.Close()
	c := service.NewGrpcServiceClient(conn)

	response, err := c.SendMoneyWallet(context.Background(), &service.SendMoney{WalletSender: walletSender, WalletRecipient: wallletRecipient, SendMoney: sendmoney})
	if err != nil {
		log.Fatalf("Error when calling GetWallet: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}

func connectGrpcServer() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	return conn
}

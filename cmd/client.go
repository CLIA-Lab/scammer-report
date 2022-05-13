/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/spf13/cobra"

	pb "github.com/CLIA-Lab/scammer-report/pkg/report"
)

const (
	address     = "localhost:8000"
	defaultName = "dr-who"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Send report",
	Long:  `Send report to the server. The report contains user public key, transaction hash and an optional description.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		client := pb.NewReportClient(conn)

		var user pb.UserReport
		user.Send = timestamppb.Now()
		// Contact the server and print out its response.
		// name := defaultName
		if len(os.Args) > 2 {
			user.PublicKey, user.HashTransaction, user.Description = os.Args[2], os.Args[3], os.Args[4]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.AddReport(ctx, &pb.AddReportRequest{UserReport: &user})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Response: %s", r.GetMessage())
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

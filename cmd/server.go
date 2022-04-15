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
	"net"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	pb "github.com/CLIA-Lab/scammer-report/pkg/report"
	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type Server struct {
	pb.UnimplementedReportServer
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()

		// Register services
		pb.RegisterReportServer(grpcServer, &Server{})

		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

// GetGopher implements gopher.GopherServer
func (s *Server) GetReport(ctx context.Context, req *pb.ReportRequest) (*pb.ReportReply, error) {
	res := &pb.ReportReply{}

	// Check request
	if req == nil {
		fmt.Println("request must not be nil")
		return res, xerrors.Errorf("request must not be nil")
	}

	if req.UserReport == nil {
		fmt.Println("public key must not be nil in the request")
		return res, xerrors.Errorf("user report must not be empty in the request")
	}

	if req.UserReport.PublicKey == "" {
		fmt.Println("public key must not be empty in the request")
		return res, xerrors.Errorf("public key must not be empty in the request")
	}

	if req.UserReport.HashTransaction == "" {
		fmt.Println("transaction hash must not be empty in the request")
		return res, xerrors.Errorf("transaction hash must not be empty in the request")
	}

	if req.UserReport.Description == nil {
		fmt.Println("description of the report must not be nil in the request")
		return res, xerrors.Errorf("description of the report must not be empty in the request")
	}

	log.Printf("Received: %v", strings.Join(req.UserReport.Description, " "))

	res.Message = "OK reported: user = " + req.UserReport.PublicKey + ", transaction = " + req.UserReport.HashTransaction

	return res, nil
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

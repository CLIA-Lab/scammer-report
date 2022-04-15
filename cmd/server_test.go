package cmd_test

import (
	"context"
	"testing"

	cmd "github.com/CLIA-Lab/scammer-report/cmd"
	pb "github.com/CLIA-Lab/scammer-report/pkg/report"
	"google.golang.org/protobuf/types/known/timestamppb"

	. "github.com/onsi/gomega"
)

func TestGetReport(t *testing.T) {
	s := cmd.Server{}

	testCases := []struct {
		name        string
		req         *pb.ReportRequest
		message     string
		expectedErr bool
	}{
		{
			name:        "req ok",
			req:         &pb.ReportRequest{UserReport: &pb.UserReport{PublicKey: "9201948012402748", HashTransaction: "938642983659283659", Description: []string{"hola", "mundo"}, Send: timestamppb.Now()}},
			message:     "OK reported: user = 9201948012402748, transaction = 938642983659283659",
			expectedErr: false,
		},
		{
			name:        "req ok",
			req:         &pb.ReportRequest{UserReport: &pb.UserReport{PublicKey: "9201948402748", HashTransaction: "93864298365921183659", Send: timestamppb.Now()}},
			expectedErr: true,
		},
		{
			name:        "req with empty request",
			req:         &pb.ReportRequest{},
			expectedErr: true,
		},
		{
			name:        "req with nil user",
			req:         &pb.ReportRequest{UserReport: nil},
			expectedErr: true,
		},
		{
			name:        "req with empty name",
			req:         &pb.ReportRequest{UserReport: &pb.UserReport{}},
			expectedErr: true,
		},
		{
			name:        "nil request",
			req:         nil,
			expectedErr: true,
		},
	}
	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctx := context.Background()

			// call
			response, err := s.GetReport(ctx, testCase.req)

			t.Log("Got : ", response)

			// assert results expectations
			if testCase.expectedErr {
				g.Expect(response).ToNot(BeNil(), "Result should be nil")
				g.Expect(err).ToNot(BeNil(), "Result should be nil")
			} else {
				g.Expect(response.Message).To(Equal(testCase.message))
			}
		})
	}
}

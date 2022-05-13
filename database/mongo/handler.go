package mongo

import (
	"context"
	"fmt"

	model "github.com/CLIA-Lab/scammer-report/models"
	pb "github.com/CLIA-Lab/scammer-report/pkg/report"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddReport(ctx context.Context, req *pb.AddReportRequest, reportdb *mongo.Collection, mongoCtx context.Context) (*pb.AddReportReply, error) {
	// Essentially doing req.UserReport to access the struct with a nil check
	user := req.GetUserReport()
	// Now we have to convert this into a Report type to convert into BSON
	data := model.Report{
		ID:              [12]byte{},
		PublicKey:       user.PublicKey,
		TransactionHash: user.HashTransaction,
		Description:     req.UserReport.Description,
		Time:            primitive.Timestamp{T: uint32(user.GetSend().AsTime().Unix()), I: 0},
	}

	// Insert the data into the database, result contains the newly generated Object ID for the new document
	result, err := reportdb.InsertOne(mongoCtx, data)
	// check for potential errors
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// add the id to report, first cast the "generic type" (go doesn't have real generics yet) to an Object ID.
	oid := result.InsertedID.(primitive.ObjectID)
	// Convert the object id to it's string counterpart
	user.Id = oid.Hex()
	// return the blog in a report type
	return &pb.AddReportReply{Message: "Id: " + user.GetId()}, nil
}

func GetReport(ctx context.Context, req *pb.GetReportRequest, reportdb *mongo.Collection) (*pb.GetReportReply, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetPublicKey())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to user public key: %v", err))
	}
	result := reportdb.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty report to write our decode result to
	data := model.Report{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find user with public key %s: %v", req.GetPublicKey(), err))
	}
	// Cast to get report type
	response := &pb.GetReportReply{
		UserReport: &pb.UserReport{
			Id:              oid.Hex(),
			PublicKey:       data.PublicKey,
			HashTransaction: data.TransactionHash,
			Send:            &timestamppb.Timestamp{Seconds: int64(data.Time.T)},
			Description:     data.Description,
		},
	}
	return response, nil
}

func ListReport(req *pb.ListReportRequest, stream pb.Report_ListReportsServer, reportdb *mongo.Collection) error {
	// Initiate a report type to write decoded data to
	data := &model.Report{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := reportdb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send report over stream
		stream.Send(&pb.ListReportReply{
			UserReport: &pb.UserReport{
				Id:              data.ID.Hex(),
				PublicKey:       data.PublicKey,
				HashTransaction: data.TransactionHash,
				Send:            &timestamppb.Timestamp{Seconds: int64(data.Time.T)},
				Description:     data.Description,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/cybersamx/go-recipes/grpc/test/genproto"
)

const (
	port = ":50051"
)

func newService() *service {
	serv := service{
		notes: make([]*pb.Note, 0),
	}

	if err := serv.readCatalogFile(); err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	return &serv
}

type service struct {
	notes []*pb.Note
}

func (s *service) readCatalogFile() error {
	catalogJSON, err := os.ReadFile("notes.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(catalogJSON, &s.notes); err != nil {
		return err
	}

	return nil
}

func (s *service) AddNote(_ context.Context, in *pb.Note) (*pb.NoteID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create a in ID: %v", err)
	}

	in.Id = id.String()
	s.notes = append(s.notes, in)

	return &pb.NoteID{Value: in.Id}, nil
}

func (s *service) GetNote(_ context.Context, in *pb.NoteID) (*pb.Note, error) {
	for _, item := range s.notes {
		if item.Id == in.Value {
			return item, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Failed to get note with ID %s", in.Value)
}

func (s *service) ListNotes(_ context.Context, _ *pb.Empty) (*pb.Notes, error) {
	res := pb.Notes{
		Notes: s.notes,
	}

	return &res, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	service := newService()
	pb.RegisterNoteServiceServer(srv, service)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to launch: %v", err)
	}
}

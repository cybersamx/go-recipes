package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"

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
		notes: make(map[string]*pb.Note),
	}

	if err := serv.readCatalogFile(); err != nil {
		log.Fatalf("failed to read catalog: %v", err)
	}

	return &serv
}

type service struct {
	notes map[string]*pb.Note
}

func (s *service) readCatalogFile() error {
	catalogJSON, err := ioutil.ReadFile("notes.json")
	if err != nil {
		return err
	}

	var items []*pb.Note
	if err := json.Unmarshal(catalogJSON, &items); err != nil {
		return err
	}

	for _, item := range items {
		s.notes[item.Id] = item
	}

	return nil
}

func (s *service) AddNote(_ context.Context, note *pb.Note) (*pb.NoteID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create a note ID: %v", err)
	}

	note.Id = id.String()
	s.notes[note.Id] = note

	res := pb.NoteID{
		Value: note.Id,
	}

	return &res, nil
}
func (s *service) GetNote(_ context.Context, noteID *pb.NoteID) (*pb.Note, error) {
	found := s.notes[noteID.Value]
	if found == nil {
		return nil, status.Errorf(codes.NotFound, "Failed to get note with ID %s", noteID.Value)
	}

	return found, nil
}
func (s *service) ListNotes(_ context.Context, _ *pb.Empty) (*pb.Notes, error) {
	notes := make([]*pb.Note, 0)

	for _, note := range s.notes {
		notes = append(notes, note)
	}

	res := pb.Notes{
		Notes: notes,
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

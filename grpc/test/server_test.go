package main

import (
	"context"
	"net"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/cybersamx/go-recipes/grpc/test/genproto"
)

func newGRPCServer(t *testing.T, service *service) *grpc.Server {
	srv := grpc.NewServer()
	pb.RegisterNoteServiceServer(srv, service)

	return srv
}

func newGRPCClient(t *testing.T, addr string) (pb.NoteServiceClient, func()) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)

	client := pb.NewNoteServiceClient(conn)

	return client, func() {
		_ = conn.Close()
	}
}

func TestNoteService_AddNote(t *testing.T) {
	listen, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	service := newService()
	srv := newGRPCServer(t, service)
	defer srv.GracefulStop()
	go func() {
		err = srv.Serve(listen)
		require.NoError(t, err)
	}()

	client, teardown := newGRPCClient(t, listen.Addr().String())
	defer teardown()

	count := len(service.notes)

	newNote := pb.Note{
		Title:   "Agenda",
		Content: "1. Discuss priorities, 2. Questions",
	}
	res, err := client.AddNote(context.Background(), &newNote)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Value)
	assert.Len(t, service.notes, count+1)

	var foundNote *pb.Note
	for _, item := range service.notes {
		if item.Id == res.Value {
			foundNote = item
		}
	}
	if foundNote != nil {
		assert.Equal(t, newNote.Title, foundNote.Title)
		assert.Equal(t, newNote.Content, foundNote.Content)
	} else {
		assert.Failf(t, "nil value", "can't find note %s in service", res.Value)
	}
}

func TestNoteService_GetNote(t *testing.T) {
	listen, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	service := newService()
	srv := newGRPCServer(t, service)
	defer srv.GracefulStop()
	go func() {
		err = srv.Serve(listen)
		require.NoError(t, err)
	}()

	client, teardown := newGRPCClient(t, listen.Addr().String())
	defer teardown()

	id := "315dcbef-9684-4930-bca4-5d446fd4a4fa"
	noteID := pb.NoteID{
		Value: id,
	}
	res, err := client.GetNote(context.Background(), &noteID)
	assert.NoError(t, err)

	var foundNote *pb.Note
	for _, item := range service.notes {
		if item.Id == res.Id {
			foundNote = item
		}
	}
	if foundNote != nil {
		assert.Equal(t, foundNote.Id, res.Id)
		assert.Equal(t, foundNote.Title, res.Title)
		assert.Equal(t, foundNote.Content, res.Content)
	} else {
		assert.Failf(t, "nil value", "can't find note %s in service", res.Id)
	}
}

func TestNoteService_ListNotes(t *testing.T) {
	listen, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	service := newService()
	srv := newGRPCServer(t, service)
	defer srv.GracefulStop()
	go func() {
		err = srv.Serve(listen)
		require.NoError(t, err)
	}()

	client, teardown := newGRPCClient(t, listen.Addr().String())
	defer teardown()

	res, err := client.ListNotes(context.Background(), &pb.Empty{})
	assert.NoError(t, err)
	assert.Len(t, res.Notes, len(service.notes))
	diff := cmp.Diff(res.Notes, service.notes, cmp.Comparer(proto.Equal))
	assert.Empty(t, diff)
}

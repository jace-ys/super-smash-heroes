package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/jace-ys/super-smash-heroes/libraries/go/errors"
	"github.com/jace-ys/super-smash-heroes/libraries/go/response"
	"github.com/jace-ys/super-smash-heroes/libraries/go/utils"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

type SuperheroServiceClient struct {
	conn *grpc.ClientConn
}

func (c *SuperheroServiceClient) GetAllSuperheroes(w http.ResponseWriter, r *http.Request) {
	client := pb.NewSuperheroServiceClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	stream, err := client.GetAllSuperheroes(ctx, &pb.Empty{})
	if err != nil {
		response.HandleGrpcError(w, err)
		return
	}

	var superheroes []response.PbJSON
	for {
		s, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			response.HandleGrpcError(w, err)
			return
		}
		superheroes = append(superheroes, response.EncodePbToJSON(s))
	}

	response.SendJSON(w, http.StatusOK, map[string]interface{}{"superheroes": superheroes})
}

func (c *SuperheroServiceClient) GetOneSuperhero(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	id := utils.Atoi(strId)
	if id == -1 {
		response.SendError(w, http.StatusBadRequest, errors.InvalidIDFormat.Error())
		return
	}

	client := pb.NewSuperheroServiceClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	superhero, err := client.GetOneSuperhero(ctx, &pb.SuperheroIdRequest{
		Val: int32(id),
	})
	if err != nil {
		response.HandleGrpcError(w, err)
		return
	}

	response.SendJSON(w, http.StatusOK, response.EncodePbToJSON(superhero))
}

type searchRequest struct {
	FullName string `json:"fullName"`
	AlterEgo string `json:"alterEgo"`
}

func (c *SuperheroServiceClient) AddSuperhero(w http.ResponseWriter, r *http.Request) {
	var s searchRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&s)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	client := pb.NewSuperheroServiceClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	superhero, err := client.AddSuperhero(ctx, &pb.SearchRequest{
		FullName: s.FullName,
		AlterEgo: s.AlterEgo,
	})
	if err != nil {
		response.HandleGrpcError(w, err)
		return
	}

	response.SendJSON(w, http.StatusOK, response.EncodePbToJSON(superhero))
}

func (c *SuperheroServiceClient) DeleteOneSuperhero(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, errors.InvalidIDFormat.Error())
		return
	}

	client := pb.NewSuperheroServiceClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err = client.DeleteOneSuperhero(ctx, &pb.SuperheroIdRequest{
		Val: int32(id),
	})
	if err != nil {
		response.HandleGrpcError(w, err)
		return
	}

	response.SendJSON(w, http.StatusNoContent, nil)
}

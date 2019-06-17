package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

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
	id := mux.Vars(r)["id"]

	client := pb.NewSuperheroServiceClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	superhero, err := client.GetOneSuperhero(ctx, &pb.SuperheroIdRequest{
		Val: id,
	})
	if err != nil {
		response.HandleGrpcError(w, err)
		return
	}

	response.SendJSON(w, http.StatusOK, response.EncodePbToJSON(superhero))
}

func (c *SuperheroServiceClient) AddSuperhero(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var s map[string]string
	err := decoder.Decode(&s)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(s) > 1 {
		response.SendError(w, http.StatusBadRequest, errInvalidRequest.Error())
		return
	}

	var key int32
	var val string
	for k, v := range s {
		k = strings.ToUpper(utils.CamelToSnakeCase(k))
		k, ok := pb.SearchRequest_Key_value[k]
		if !ok {
			response.SendError(w, http.StatusBadRequest, errInvalidRequest.Error())
			return
		}
		key = k
		val = v
	}

	client := pb.NewSuperheroServiceClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	superhero, err := client.AddSuperhero(ctx, &pb.SearchRequest{
		Key: pb.SearchRequest_Key(key),
		Val: val,
	})
	if err != nil {
		response.HandleGrpcError(w, err)
		return
	}

	response.SendJSON(w, http.StatusOK, response.EncodePbToJSON(superhero))
}

func (c *SuperheroServiceClient) DeleteOneSuperhero(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	client := pb.NewSuperheroServiceClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err := client.DeleteOneSuperhero(ctx, &pb.SuperheroIdRequest{
		Val: id,
	})
	if err != nil {
		response.HandleGrpcError(w, err)
		return
	}

	response.SendJSON(w, http.StatusNoContent, nil)
}

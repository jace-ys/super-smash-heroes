package superhero

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log/level"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"

	pb "github.com/jace-ys/super-smash-heroes/services/superhero/api/superhero"
)

func (s *SuperheroService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	level.Info(s.logger).Log("event", "superhero.list.started")
	defer level.Info(s.logger).Log("event", "superhero.list.finished")

	superheroes, err := s.list(ctx)
	if err != nil {
		level.Error(s.logger).Log("event", "superhero.list.failure", "msg", err)
		return nil, s.Error(codes.Internal, err)
	}

	level.Info(s.logger).Log("event", "superhero.list.success")
	return &pb.ListResponse{
		Superheroes: superheroes,
	}, nil
}

func (s *SuperheroService) list(ctx context.Context) ([]*pb.Superhero, error) {
	var superheroes []*pb.Superhero
	err := s.db.Transact(ctx, func(tx *sqlx.Tx) error {
		query := `
		SELECT s.id, s.full_name, s.alter_ego, s.image_url, s.intelligence, s.strength, s.speed, s.durability, s.power, s.combat
		FROM superheroes AS s
		`
		rows, err := tx.QueryxContext(ctx, query)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var superhero pb.Superhero
			if err := rows.StructScan(&superhero); err != nil {
				return err
			}
			superheroes = append(superheroes, &superhero)
		}
		return rows.Err()
	})
	if err != nil {
		return nil, err
	}
	return superheroes, nil
}

func (s *SuperheroService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	level.Info(s.logger).Log("event", "superhero.get.started")
	defer level.Info(s.logger).Log("event", "superhero.get.finished")

	superhero, err := s.get(ctx, req.Id)
	if err != nil {
		level.Error(s.logger).Log("event", "superhero.get.failure", "msg", err)
		switch {
		case errors.Is(err, ErrSuperheroNotFound):
			return nil, s.Error(codes.NotFound, err)
		default:
			return nil, s.Error(codes.Internal, err)
		}
	}

	level.Info(s.logger).Log("event", "superhero.get.success")
	return &pb.GetResponse{
		Superheroes: superhero,
	}, nil
}

func (s *SuperheroService) get(ctx context.Context, id int32) (*pb.Superhero, error) {
	var superhero pb.Superhero
	err := s.db.Transact(ctx, func(tx *sqlx.Tx) error {
		query := `
		SELECT s.id, s.full_name, s.alter_ego, s.image_url, s.intelligence, s.strength, s.speed, s.durability, s.power, s.combat
		FROM superheroes AS s
		WHERE s.id=$1
		`
		row := tx.QueryRowxContext(ctx, query, id)
		return row.StructScan(&superhero)
	})
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrSuperheroNotFound
		default:
			return nil, err
		}
	}
	return &superhero, nil
}

func (s *SuperheroService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	level.Info(s.logger).Log("event", "superhero.create.started")
	defer level.Info(s.logger).Log("event", "superhero.create.finished")

	err := s.validateCreateRequest(req.FullName, req.AlterEgo)
	if err != nil {
		level.Error(s.logger).Log("event", "superhero.create.failure", "msg", err)
		return nil, s.Error(codes.InvalidArgument, err)
	}

	superhero, err := s.registry.Find(req.FullName, req.AlterEgo)
	if err != nil {
		level.Error(s.logger).Log("event", "superhero.create.failure", "msg", err)
		return nil, s.Error(codes.NotFound, err)
	}

	id, err := s.create(ctx, superhero)
	if err != nil {
		level.Error(s.logger).Log("event", "superhero.create.failure", "msg", err)
		switch {
		case errors.Is(err, ErrSuperheroExists):
			return nil, s.Error(codes.AlreadyExists, err)
		default:
			return nil, s.Error(codes.Internal, err)
		}
	}

	level.Info(s.logger).Log("event", "superhero.create.success")
	return &pb.CreateResponse{
		Id: id,
	}, nil
}

func (s *SuperheroService) validateCreateRequest(fullName, alterEgo string) error {
	switch {
	case fullName == "":
		return ErrRequestInvalid
	case alterEgo == "":
		return ErrRequestInvalid
	}
	return nil
}

func (s *SuperheroService) create(ctx context.Context, superhero *pb.Superhero) (int32, error) {
	var id int32
	err := s.db.Transact(ctx, func(tx *sqlx.Tx) error {
		query := `
		INSERT INTO superheroes (full_name, alter_ego, image_url, intelligence, strength, speed, durability, power, combat)
		VALUES (:full_name, :alter_ego, :image_url, :intelligence, :strength, :speed, :durability, :power, :combat)
		RETURNING id
		`
		stmt, err := tx.PrepareNamedContext(ctx, query)
		if err != nil {
			return err
		}
		return stmt.QueryRowxContext(ctx, superhero).Scan(&id)
	})
	if err != nil {
		var pqErr *pq.Error
		switch {
		case errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation":
			return 0, ErrSuperheroExists
		default:
			return 0, err
		}
	}
	return id, nil
}

func (s *SuperheroService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	level.Info(s.logger).Log("event", "superhero.delete.started")
	defer level.Info(s.logger).Log("event", "superhero.delete.finished")

	err := s.delete(ctx, req.Id)
	if err != nil {
		level.Error(s.logger).Log("event", "superhero.delete.failure", "msg", err)
		switch {
		case errors.Is(err, ErrSuperheroNotFound):
			return nil, s.Error(codes.NotFound, err)
		default:
			return nil, s.Error(codes.Internal, err)
		}
	}

	level.Info(s.logger).Log("event", "superhero.delete.success")
	return &pb.DeleteResponse{}, nil
}

func (s *SuperheroService) delete(ctx context.Context, id int32) error {
	err := s.db.Transact(ctx, func(tx *sqlx.Tx) error {
		query := `
		DELETE FROM superheroes
		WHERE id=$1
		`
		res, err := tx.ExecContext(ctx, query, id)
		if err != nil {
			return err
		}
		count, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if count == 0 {
			return sql.ErrNoRows
		}
		return nil
	})
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrSuperheroNotFound
		default:
			return err
		}
	}
	return nil
}

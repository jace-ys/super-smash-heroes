package integration

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jace-ys/super-smash-heroes/test/api/superhero"
)

var _ = Describe("SuperheroService", func() {
	Context("List", func() {
		It("returns OK", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			req := &superhero.ListRequest{}

			resp, err := superheroService.List(ctx, req)

			Expect(status.Code(err).String()).To(Equal(codes.OK.String()))
			Expect(resp).NotTo(BeNil())
			Expect(resp.Superheroes).To(HaveLen((2)))
		})
	})

	Context("Get", func() {
		It("returns OK", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			req := &superhero.GetRequest{
				Id: 1,
			}

			resp, err := superheroService.Get(ctx, req)

			Expect(status.Code(err).String()).To(Equal(codes.OK.String()))
			Expect(resp).NotTo(BeNil())
			Expect(resp.Superheroes.FullName).To(Equal("Oliver Queen"))
			Expect(resp.Superheroes.AlterEgo).To(Equal("Green Arrow"))
		})

		Context("Non-existent superhero", func() {
			It("returns NotFound", func() {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				req := &superhero.GetRequest{
					Id: 0,
				}

				resp, err := superheroService.Get(ctx, req)

				Expect(status.Code(err).String()).To(Equal(codes.NotFound.String()))
				Expect(resp).To(BeNil())
			})
		})
	})

	Context("Create", func() {
		It("returns OK", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			req := &superhero.CreateRequest{
				FullName: "Thor Odinson",
				AlterEgo: "Thor",
			}

			resp, err := superheroService.Create(ctx, req)

			Expect(status.Code(err).String()).To(Equal(codes.OK.String()))
			Expect(resp).NotTo(BeNil())
			Expect(int(resp.Id)).To(Equal(3))
		})

		Context("Missing field", func() {
			It("returns InvalidArgument", func() {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				req := &superhero.CreateRequest{}

				resp, err := superheroService.Create(ctx, req)

				Expect(status.Code(err).String()).To(Equal(codes.InvalidArgument.String()))
				Expect(resp).To(BeNil())
			})
		})

		Context("Unregistered superhero", func() {
			It("returns NotFound", func() {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				req := &superhero.CreateRequest{
					FullName: "Superhero",
					AlterEgo: "Superhero",
				}

				resp, err := superheroService.Create(ctx, req)

				Expect(status.Code(err).String()).To(Equal(codes.NotFound.String()))
				Expect(resp).To(BeNil())
			})
		})

		Context("Existing superhero", func() {
			It("returns AlreadyExists", func() {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				req := &superhero.CreateRequest{
					FullName: "Oliver Queen",
					AlterEgo: "Green Arrow",
				}

				resp, err := superheroService.Create(ctx, req)

				Expect(status.Code(err).String()).To(Equal(codes.AlreadyExists.String()))
				Expect(resp).To(BeNil())
			})
		})
	})

	Context("Delete", func() {
		It("returns OK", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			req := &superhero.DeleteRequest{
				Id: 1,
			}

			resp, err := superheroService.Delete(ctx, req)

			Expect(status.Code(err).String()).To(Equal(codes.OK.String()))
			Expect(resp).NotTo(BeNil())
		})

		Context("Non-existent superhero", func() {
			It("returns NotFound", func() {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				req := &superhero.DeleteRequest{
					Id: 0,
				}

				resp, err := superheroService.Delete(ctx, req)

				Expect(status.Code(err).String()).To(Equal(codes.NotFound.String()))
				Expect(resp).To(BeNil())
			})
		})
	})
})

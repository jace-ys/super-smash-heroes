package integration

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jace-ys/super-smash-heroes/test/api/battle"
)

var _ = Describe("BattleService", func() {
	Context("GetResult", func() {
		It("returns OK", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			req := &battle.GetResultRequest{
				PlayerOne: &battle.Player{},
				PlayerTwo: &battle.Player{},
			}

			resp, err := battleService.GetResult(ctx, req)

			Expect(status.Code(err)).To(Equal(codes.OK))
			Expect(resp).NotTo(BeNil())
			Expect(int(resp.Winner)).Should(Or(Equal(1), Equal(2)))
		})

		Context("Missing player", func() {
			It("returns InvalidArgument", func() {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				req := &battle.GetResultRequest{
					PlayerOne: &battle.Player{},
				}

				resp, err := battleService.GetResult(ctx, req)

				Expect(status.Code(err).String()).To(Equal(codes.InvalidArgument.String()))
				Expect(resp).To(BeNil())
			})
		})
	})
})

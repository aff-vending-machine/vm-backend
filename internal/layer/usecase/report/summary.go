package report

import (
	"context"

	"vm-backend/internal/layer/usecase/report/request"
	"vm-backend/internal/layer/usecase/report/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Summary(ctx context.Context, req *request.Summary) ([]response.Machine, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToMachineQuery()
	machines, err := uc.machineRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find many machines")
		return nil, errors.Wrap(err, "unable to find many machines")
	}

	query = req.ToChannelQuery()
	channels, err := uc.channelRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find many channels")
		return nil, errors.Wrap(err, "unable to find many channels")
	}

	res := make([]response.Machine, len(machines))
	mapMachineIndex := map[uint]int{}
	mapChannelIndex := map[uint]string{}
	for i, machine := range machines {
		res[i] = response.Machine{
			ID:            machine.ID,
			Name:          machine.Name,
			SerialNumber:  machine.SerialNumber,
			TotalPayments: make(map[string]float64, 0),
			Total:         0,
		}
		for _, channel := range channels {
			res[i].TotalPayments[channel.Channel] = 0
			mapChannelIndex[channel.ID] = channel.Channel
		}

		mapMachineIndex[machine.ID] = i
	}

	query = req.ToTransactionQuery()
	entities, err := uc.transactionRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find many transactions")
		return nil, errors.Wrap(err, "unable to find many transactions")
	}

	for _, entity := range entities {
		if i, ok := mapMachineIndex[entity.MachineID]; ok {
			if channel, ok := mapChannelIndex[entity.ChannelID]; ok {
				res[i].TotalPayments[channel] += entity.PaidPrice
			}
			res[i].Total += entity.PaidPrice
		}
	}

	return res, nil
}

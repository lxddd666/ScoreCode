package tg

import (
	"context"
	"hotgo/internal/model/input/tgin"
)

func (s *sTgArts) GetUserChannels(ctx context.Context, inp *tgin.GetUserChannelsInp) (res []*tgin.TgDialogModel, err error) {
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	list, err := s.TgGetDialogs(ctx, inp.Account)
	if err != nil {
		return
	}
	for _, dialog := range list {
		if dialog.Type == 3 {
			res = append(res, dialog)
		}
	}
	return
}

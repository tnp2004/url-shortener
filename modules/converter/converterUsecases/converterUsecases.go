package converterUsecases

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/tnp2004/url-shortener/config"
	"github.com/tnp2004/url-shortener/modules/converter"
	"github.com/tnp2004/url-shortener/modules/converter/converterRepositories"
	"github.com/tnp2004/url-shortener/pkg/utils"
)

type (
	IConverterUsecases interface {
		GetShortUrl(pctx context.Context, req *converter.ConverterReq) (*converter.ConverterRes, error)
		SearchDestination(pctx context.Context, req *converter.SearchShortIdReq) (string, error)
	}
	converterUsecases struct {
		converterRepository converterRepositories.IConverterRepository
		cfg                 *config.App
	}
)

func NewConverterUsecases(repository converterRepositories.IConverterRepository, cfg *config.App) IConverterUsecases {
	return &converterUsecases{
		converterRepository: repository,
		cfg:                 cfg,
	}
}

func (u *converterUsecases) GetShortUrl(pctx context.Context, req *converter.ConverterReq) (*converter.ConverterRes, error) {
	// Check if the destination URL already exists
	existDoc, _ := u.converterRepository.FindOneDestination(pctx, req.URL)
	if existDoc != nil {
		// Update expiration time  when destination is already exist
		if err := u.converterRepository.UpdateOneExpiration(pctx, existDoc.Id, time.Now().Add(86400*time.Second)); err != nil {
			return nil, err
		}

		return &converter.ConverterRes{
			ShortenedURL: fmt.Sprintf("%s/%s/%s", u.cfg.Url, u.cfg.Version, existDoc.ShortId),
		}, nil
	}

	// Insert Short id
	shortId := utils.GenerateShortId(6)

	insertRequest := &converter.Url{
		ShortId:     shortId,
		Destination: req.URL,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(86400 * time.Second),
	}

	_, err := u.converterRepository.InsertUrl(pctx, insertRequest)
	if err != nil {
		log.Fatalf("Error: InsertUrl failed: %s", err.Error())
		return nil, errors.New("error: get short url failed")
	}

	return &converter.ConverterRes{
		ShortenedURL: fmt.Sprintf("%s/%s/%s", u.cfg.Url, u.cfg.Version, shortId),
	}, nil
}

func (u *converterUsecases) SearchDestination(pctx context.Context, req *converter.SearchShortIdReq) (string, error) {
	des, err := u.converterRepository.FindOneDestinationByShortId(pctx, req.ShortId)
	if err != nil {
		return "", errors.New("error: search destination failed")
	}

	return des, nil
}

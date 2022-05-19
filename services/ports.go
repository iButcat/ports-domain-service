package services

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"strings"

	"port-domain-service/domain/entity"
	"port-domain-service/domain/repository"
	"port-domain-service/pkg/utils"
)

type (
	PortDomainService interface {
		Create(ctx context.Context, file multipart.File) error
		Update(ctx context.Context, file multipart.File) error
	}

	portDomainService struct {
		repo repository.PortRepository
	}
)

func NewPortDomainService(repo repository.PortRepository) PortDomainService {
	return &portDomainService{
		repo: repo,
	}
}

// Probably not the most efficient way to do this, but it works.
func (s *portDomainService) Create(ctx context.Context, file multipart.File) error {
	var ports []entity.Port
	var port entity.Port
	fileContent, err := utils.ReadFileByChunks(file, 1024)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(strings.NewReader(string(fileContent)))
	for {
		err := dec.Decode(&port)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		ports = append(ports, port)
	}

	for _, port := range ports {
		if err := s.repo.Create(ctx, &port); err != nil {
			return err
		}
	}
	return nil
}

func (s *portDomainService) Update(ctx context.Context, file multipart.File) error {
	var ports []entity.Port
	fileContent, err := utils.ReadFileByChunks(file, 1024)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(fileContent, &ports); err != nil {
		return err
	}

	for _, port := range ports {
		if err := s.repo.Update(ctx, &port); err != nil {
			return err
		}
	}

	return nil
}

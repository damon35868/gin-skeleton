package service

import "fmt"

type IndexService struct{}

func NewIndexService() *IndexService {
	return &IndexService{}
}

func (indexService *IndexService) Index(way string) {
	fmt.Printf("[Service] Form: %s\n", way)
}

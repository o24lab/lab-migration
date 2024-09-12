package esx

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"pkg/errorx"
)

type Es[T any] interface {
	GetOne(query elastic.Query) (*T, errorx.Error)
	GetList(query elastic.Query, options ...EsOption) ([]T, errorx.Error)
}

type es[T any] struct {
	tableIndex string
}

func Define[T any](tableIndex string) Es[T] {
	return &es[T]{tableIndex: tableIndex}
}

func (e es[T]) GetOne(query elastic.Query) (*T, errorx.Error) {
	client := GetClient()
	result, err := client.Search(e.tableIndex).Query(query).Do(context.Background())
	if err != nil {
		return nil, errorx.WithStack(err)
	}
	if len(result.Hits.Hits) == 0 {
		return nil, errorx.ErrDataNotFound.WithStack()
	}
	var data T
	err1 := json.Unmarshal(result.Hits.Hits[0].Source, &data)
	if err1 != nil {
		return nil, errorx.WithStack(err1)
	}
	return &data, nil
}

func (e es[T]) GetList(query elastic.Query, options ...EsOption) ([]T, errorx.Error) {
	client := GetClient()
	q := client.Scroll(e.tableIndex).Query(query).Size(1000).FetchSource(true)
	for _, o := range options {
		o(q)
	}
	result, err := q.Do(context.Background())
	if err != nil {
		return nil, errorx.WithStack(err)
	}
	if len(result.Hits.Hits) == 0 {
		return nil, errorx.ErrDataNotFound.WithStack()
	}
	var out []T
	for _, hit := range result.Hits.Hits {
		var data T
		err1 := json.Unmarshal(hit.Source, &data)
		if err1 != nil {
			return nil, errorx.WithStack(err1)
		}
		out = append(out, data)
	}
	if result.ScrollId != "" {
		list, err1 := e.GetList(query, WithScrollId(result.ScrollId))
		if err1 != nil {
			return nil, errorx.WithStack(err1)
		}
		out = append(out, list...)
	}
	return out, nil
}

type EsOption func(s *elastic.ScrollService)

func WithScrollId(scrollId string) EsOption {
	return func(s *elastic.ScrollService) {
		s.ScrollId(scrollId)
	}
}

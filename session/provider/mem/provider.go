package mem

import (
	"context"
	"time"

	"github.com/gopub/errors"
	"github.com/gopub/wine/session"
	"github.com/patrickmn/go-cache"
)

type Provider struct {
	cache *cache.Cache
}

var _ session.Provider = (*Provider)(nil)

func NewProvider() *Provider {
	p := new(Provider)
	p.cache = cache.New(session.DefaultOptions().TTL, session.DefaultOptions().TTL*50)
	return p
}

func (p *Provider) Get(ctx context.Context, id string) (session.Session, error) {
	v, ok := p.cache.Get(id)
	if !ok {
		return nil, errors.NotExist
	}
	return v.(*Session), nil
}

func (p *Provider) Create(ctx context.Context, id string, ttl time.Duration) (session.Session, error) {
	s := &Session{
		id:          id,
		sharedCache: p.cache,
	}
	p.cache.Set(id, s, ttl)
	return s, nil
}

func (p *Provider) Delete(ctx context.Context, id string) error {
	p.cache.Delete(id)
	return nil
}

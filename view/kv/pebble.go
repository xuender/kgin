package kv

import (
	"errors"

	"github.com/cockroachdb/pebble"
	"github.com/retailnext/hllpp"
	"github.com/xuender/kgin/view"
	"github.com/xuender/kit/times"
)

type Pebble struct {
	db *pebble.DB
}

func NewPebble(path string) (*Pebble, error) {
	pdb, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		return nil, err
	}

	return &Pebble{db: pdb}, nil
}

func (p *Pebble) Close() error {
	return p.db.Close()
}

func (p *Pebble) PV(key uint64, day times.IntDay) uint64 {
	return p.get(view.PVKey(key, day))
}

func (p *Pebble) UV(key uint64, day times.IntDay) uint64 {
	keyBytes := view.UVKey(key, day)

	value, closer, err := p.db.Get(keyBytes)
	if err != nil {
		return 0
	}

	if closer != nil {
		closer.Close()
	}

	if numMax := 8; len(value) > numMax {
		now := times.Now2IntDay()

		hll, err := hllpp.Unmarshal(value)
		if err != nil {
			if day < now {
				_ = p.set(keyBytes, 0)
			}

			return 0
		}

		count := hll.Count()

		if day < now {
			_ = p.set(keyBytes, count)
		}

		return count
	}

	return view.ToUint64(value)
}

func (p *Pebble) TV(key uint64) uint64 {
	return p.get(view.TVKey(key))
}

func (p *Pebble) View(key uint64, remoteIP string) error {
	day := times.Now2IntDay()

	if err := p.incr(view.TVKey(key)); err != nil {
		return err
	}

	if err := p.incr(view.PVKey(key, day)); err != nil {
		return err
	}

	uvKey := view.UVKey(key, day)

	return p.uv(uvKey, []byte(remoteIP))
}

func (p *Pebble) get(key []byte) uint64 {
	value, closer, err := p.db.Get(key)
	if closer != nil {
		closer.Close()
	}

	if err != nil {
		return 0
	}

	return view.ToUint64(value)
}

func (p *Pebble) incr(key []byte) error {
	var count uint64

	value, closer, err := p.db.Get(key)

	switch {
	case err == nil:
		count = view.ToUint64(value) + 1

		if closer != nil {
			closer.Close()
		}
	case errors.Is(err, pebble.ErrNotFound):
		count = 1
	default:
		return err
	}

	return p.set(key, count)
}

func (p *Pebble) uv(key, remoteIP []byte) error {
	value, closer, err := p.db.Get(key)
	if closer != nil {
		closer.Close()
	}

	var hll *hllpp.HLLPP

	switch {
	case err == nil:
		hll, err = hllpp.Unmarshal(value)
		if err != nil {
			return err
		}

	case errors.Is(err, pebble.ErrNotFound):
		hll = hllpp.New()
	default:
		return err
	}

	hll.Add(remoteIP)

	return p.db.Set(key, hll.Marshal(), pebble.Sync)
}

func (p *Pebble) set(key []byte, count uint64) error {
	return p.db.Set(key, view.ToBytes(count), pebble.Sync)
}

package pokedex

import (
	"sync"

	"github.com/DuckyDuckDo/bootdev-pokedex/internal/pokeapi"
)

type Pokemon struct {
	name string
	data pokeapi.RespPokemonInfo
}

type Pokedex struct {
	entries map[string]Pokemon // stores pokemon and its information
	mu      sync.Mutex
}

// initializes a pokedex
func NewPokedex() *Pokedex {
	p := &Pokedex{
		entries: make(map[string]Pokemon),
	}

	return p
}

func (p *Pokedex) Add(pokemon string, pokemonInfo pokeapi.RespPokemonInfo) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entry := Pokemon{
		name: pokemon,
		data: pokemonInfo,
	}
	p.entries[pokemon] = entry
}

func (p *Pokedex) Get(pokemon string) (pokeapi.RespPokemonInfo, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.entries[pokemon]; !ok {
		return pokeapi.RespPokemonInfo{}, false
	}
	return p.entries[pokemon].data, true
}

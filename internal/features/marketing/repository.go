package marketing

type Repository interface{}

type StubRepository struct {
	tiers map[string]PricingTier
}

func NewRepository() Repository {
	return &StubRepository{
		tiers: make(map[string]PricingTier),
	}
}

func (r *StubRepository) GetAll() ([]PricingTier, error) {
	tiers := make([]PricingTier, 0, len(r.tiers))

	return tiers, nil
}

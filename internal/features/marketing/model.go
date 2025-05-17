package marketing

type PricingTier struct {
	Name        string
	Price       float32
	Description string
	Features    map[string]string
}

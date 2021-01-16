package database

//CollectionMongoType ...
func CollectionMongoType() struct {
	Business    string
	Activity    string
	Price       string
	PriceGroup  string
	PriceUnit   string
	Sample      string
	SampleGroup string
	Contact     string
	Location    string
	Member      string
	Sites       string

	Project string
} {
	return struct {
		Business    string
		Activity    string
		Price       string
		PriceGroup  string
		PriceUnit   string
		Sample      string
		SampleGroup string
		Contact     string
		Location    string
		Member      string
		Sites       string

		Project string
	}{
		"businesses",
		"activities",
		"price",
		"price_group",
		"price_unit",
		"samples",
		"sample_group",
		"contact",
		"location",
		"member",
		"sites",
		"projects",
	}
}

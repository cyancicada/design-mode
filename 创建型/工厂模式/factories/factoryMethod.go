package factories

/***
 *	Factory Method
 */
type Factory interface {
	Create() Food
}

type MeatFactory struct {
}

func (mf MeatFactory) Create() Food {
	m := Meat{}
	return m
}

type HambergerFactory struct {
}

func (hf HambergerFactory) Create() Food {
	h := Hamberger{}
	return h
}

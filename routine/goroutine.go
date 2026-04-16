package routine

import "github.com/google/uuid"

type Compony struct {
	id             string
	name           string
	founder        string
	avg_make_money float64
	country        []Country
}

type Country struct {
	id   string
	name string
}

func newCompony(name string, name_of_compony string, founder string, avg_make_money float64, my_country Country) Compony {

	return Compony{
		id:             uuid.New().String(),
		name:           name_of_compony,
		founder:        founder,
		avg_make_money: avg_make_money,
		country:        []Country{my_country},
	}

}

// we have table of facts and измерения

package entity

type Response1 struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

//	{
//	    "cep": "56909494",
//	    "state": "PE",
//	    "city": "Serra Talhada",
//	    "neighborhood": "Nossa Senhora de Fátima",
//	    "street": "Rua Professor Eduardo Lopes de Pádua",
//	    "service": "correios"
//	}

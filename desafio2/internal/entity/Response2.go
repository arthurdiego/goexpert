package entity

type Response2 struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

//	{
//	    "cep": "56909-494",
//	    "logradouro": "Rua Professor Eduardo Lopes de Pádua",
//	    "complemento": "",
//	    "bairro": "Nossa Senhora de Fátima",
//	    "localidade": "Serra Talhada",
//	    "uf": "PE",
//	    "ibge": "2613909",
//	    "gia": "",
//	    "ddd": "87",
//	    "siafi": "2577"
//	}

package entity

type Product struct {
    id int
    name string
    value int
}

func NewProduct(id int, name string, value int) Product {
    return Product{id, name, value}
}

func (p Product) GetId() int {
    return p.id
}

func (p *Product) SetId(id int) {
    p.id = id
}

func (p Product) GetName() string {
    return p.name
}

func (p *Product) SetName(name string) {
    p.name = name
}

func (p Product) GetValue() int {
    return p.value
}

func (p *Product) SetValue(value int) {
    p.value = value
}

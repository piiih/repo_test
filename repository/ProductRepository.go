package repository

import "github.com/piiih/repo_test/entity"
import "github.com/piiih/repo_test/dal"

type ProductRepository struct {
    Db dal.DbConnect
}

func (pr ProductRepository) FindById(id int) entity.Product {
    products := pr.Db.GetResultOf("select * from produtos where id ="+toString(id))

    newProduct := entity.NewProduct(
        toInt(products[0]["id"]),
        products[0]["name"],
        toInt(products[0]["value"]),
    )

    return newProduct
}

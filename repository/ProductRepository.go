package repository

import "github.com/piiih/repo_test/entity"
import "github.com/piiih/repo_test/dal"
import c "github.com/piiih/repo_test/conversions"

type ProductRepository struct {
    Db dal.DbConnect
}

func (pr ProductRepository) FindById(id int) entity.Product {
    products := pr.Db.GetResultOf("select * from produtos where id ="+c.ToString(id))

    newProduct := entity.NewProduct(
        c.ToInt(products[0]["id"]),
        products[0]["name"],
        c.ToInt(products[0]["value"]),
    )

    return newProduct
}

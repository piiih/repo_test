package main

import "fmt"
import "github.com/piiih/repo_test/dal"
import "github.com/piiih/repo_test/repository"
import "github.com/piiih/repo_test/entity"

func main() {
    type repositoryInterface interface {
        FindById(int) entity.Product
    }

    dbConnect := dal.DbConnect{}

    dbConnect.Connect("master:mastersecret@(11.11.11.11:3306)/teste")

    var repository repositoryInterface = repository.ProductRepository{dbConnect}

    product := repository.FindById(1)
    fmt.Printf("%+v\n", product.GetId())
}

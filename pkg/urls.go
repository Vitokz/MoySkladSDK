package pkg

import "fmt"

const (
	// Методы
	Get    = "GET"
	Post   = "POST"
	Put    = "PUT"
	Delete = "DELETE"

	// Товары
	Product           = "/entity/product"
	ProductDeleteList = Product + "/delete"
	ProductAttribute  = Product + "/metadata/attributes/"
)

//TakeProduct Возвращает метод и эндпоинт для того чтобы достать определенный товар
func TakeProduct(id string) (string, string) {
	return Get, fmt.Sprintf("%s/%s", Product, id)
}

//TakeProductList Возвращает метод и эндпоинт для того чтобы достать все товары
func TakeProductList() (string, string) {
	return Get, Product
}

//TakeProductAttribute Возвращает метод и эндпоинт для того чтобы достать отдельное поле товара
func TakeProductAttribute(id string) (string, string) {
	return Get, fmt.Sprintf("%s/%s", ProductAttribute, id)
}

//CreateProduct Возвращает метод и эндпоинт для того чтобы создать товар
func CreateProduct() (string, string) {
	return Post, Product
}

//DeleteProduct Возвращает метод и эндпоинт для того чтобы удалить определенный товар
func DeleteProduct(id string) (string, string) {
	return Delete, fmt.Sprintf("%s/%s", Product, id)
}

//DeleteProductList Возвращает метод и эндпоинт для удаления несколько товаров
func DeleteProductList() (string, string) {
	return Post, ProductDeleteList
}

//UpdateProduct Возвращает метод и эндпоинт для изменения товара
func UpdateProduct(id string) (string, string) {
	return Put, fmt.Sprintf("%s/%s", Product, id)
}

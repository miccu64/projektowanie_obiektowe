const categories = [
    {id: 1, name: "Book"},
    {id: 2, name: "Groceries"},
    {id: 3, name: "Electronics"}
];

const products = [
    {id: 1, name: "Cooking book", categoryId: 1, price: 22.22},
    {id: 2, name: "Clean Code", categoryId: 1, price: 111.11},
    {id: 3, name: "Apple", categoryId: 2, price: 1},
    {id: 4, name: "Bottle of juice", categoryId: 2, price: 3.3},
    {id: 5, name: "Smartwatch", categoryId: 3, price: 220},
    {id: 6, name: "PC", categoryId: 3, price: 3333},
];

const cart = {};

export {
    categories, products, cart
}

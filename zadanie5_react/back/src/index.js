const express = require('express');
const PORT = 3001;
const app = express();
const ProductModel = require('./models/ProductModel')

let products = [new ProductModel(1, "Apple", 2.22),
    new ProductModel(2, "Car", 99999),
    new ProductModel(3, "Book", 20)];

let payments = {};

app.use(express.json());       // to support JSON-encoded bodies
app.use(express.urlencoded()); // to support URL-encoded bodies

app.use(function (req, res, next) {
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Cache-Control', 'no-cache');
    res.setHeader('Content-Type', 'application/json');
    next();
});

app.get('/products', function (req, res) {
    res.json(products);
});

app.post('/payments', function (req, res) {
    for (const key of Object.keys(req.body)) {
        let val = payments[key];
        if (val) {
            val = Number(val) + Number(req.body[key]);
        } else {
            val = Number(req.body[key]);
        }
        payments[key] = val;
    }
    console.log(payments)
    res.sendStatus(200);
});

app.get('/cart', function (req, res) {
    const productsCopy = products.slice().map(p => {
        p.qty = payments[p.id];
        if (!p.qty) {
            p.qty = 0;
        }
        return p;
    })
    res.json(productsCopy);
});

app.post('/cart', function (req, res) {
    payments = {};
    res.sendStatus(200);
});

app.listen(PORT, () => {
    console.log(`Server listening on ${PORT}`);
});

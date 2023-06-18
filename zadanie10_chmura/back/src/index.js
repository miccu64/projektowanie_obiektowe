const express = require('express');
const PORT = 80;
const app = express();
const ProductModel = require('./models/ProductModel')

let products = [new ProductModel(1, "Apple", 2.22),
    new ProductModel(2, "Car", 99999),
    new ProductModel(3, "Book", 20)];

let payments = {};

app.use(express.json());
app.use(express.urlencoded());

app.use(function (req, res, next) {
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Cache-Control', 'no-cache');
    res.setHeader('Content-Type', 'application/json');
    next();
});

app.get('/api/products', function (req, res) {
    res.json(products);
});

app.post('/api/payments', function (req, res) {
    for (const key of Object.keys(req.body)) {
        let val = payments[key];
        if (val) {
            val = Number(val) + Number(req.body[key]);
        } else {
            val = Number(req.body[key]);
        }
        payments[key] = val;
    }
    res.sendStatus(200);
});

app.get('/api/cart', function (req, res) {
    const productsCopy = products.slice().map(p => {
        p.qty = payments[p.id];
        if (!p.qty) {
            p.qty = 0;
        }
        return p;
    })
    res.json(productsCopy);
});

app.post('/api/cart', function (req, res) {
    payments = {};
    res.sendStatus(200);
});

app.listen(PORT, '0.0.0.0', () => {
    console.log(`Server listening on ${PORT}`);
});

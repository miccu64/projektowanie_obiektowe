import React, {useState} from "react";

function Products() {
    const [products, setProducts] = useState(null);

    React.useEffect(() => {
        fetch('/products')
            .then(response => response.json())
            .then(data => setProducts(data));
    }, []);

    return <div>
        {products && products.map(product => {
            return <div style={{marginBottom: "10px"}}>
                <h5>Product name: {product.name}</h5>
                <h5>Price: {product.price}</h5>
                <select name={product.id} id={product.id}>{[...Array(10).keys()].map(value => {
                    return <option value={value}>{value}</option>
                })}</select>
            </div>
        })}
    </div>;
}

export default Products;
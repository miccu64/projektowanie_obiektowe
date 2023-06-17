import React, {useState} from "react";

function Products() {
    const [products, setProducts] = useState(null);

    React.useEffect(() => {
        const fetchData = async () => {
            const data = await fetch('/products');
            setProducts(await data.json());
        }
        fetchData().catch(console.error);
    }, []);

    return <div>
        {products?.map(product => {
            return <div key={product.id} style={{marginBottom: "10px"}}>
                <h5>Product name: {product.name}</h5>
                <h5>Price: {product.price}</h5>
                <select name={product.id} id={product.id}>{[...Array(10).keys()].map(value => {
                    return <option key={product.id.toString() + value.toString()} value={value}>{value}</option>
                })}</select>
            </div>
        })}
    </div>;
}

export default Products;
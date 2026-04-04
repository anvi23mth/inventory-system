import React from "react";
import { Product as ProductType } from "../types/product";
import Product from "./Product";

interface Props {
    products: ProductType[];
}

const ProductList: React.FC<Props> = ({ products }) => {
    if (products.length === 0) {
        return <p>No products found.</p>;
    }

    return (
        <div className="product-list">
            {products.map((product) => (
                <Product key={product.id} product={product} />
            ))}
        </div>
    );
};

export default ProductList;
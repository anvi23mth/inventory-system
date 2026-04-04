import React, { useState } from "react";
import { Product as ProductType } from "../types/product";

interface Props {
    product: ProductType;
}

const Product: React.FC<Props> = ({ product }) => {
    const [expanded, setExpanded] = useState(false);

    return (
        <div className="product-tile" onClick={() => setExpanded(!expanded)}>
            <h2 className="product-name">{product.name}</h2>
            <p className="product-price">${product.price}</p>
            <span className="product-category">
                {product.category || "Uncategorized"}
            </span>

            {/* Expands when clicked - Week 7 interactivity goal */}
            {expanded && (
                <div className="product-details">
                    <p><strong>Brand:</strong> {product.brand}</p>
                    <p><strong>Description:</strong> {product.description || "No description"}</p>
                    <p><strong>Quantity:</strong> {product.quantity}</p>
                </div>
            )}
        </div>
    );
};

export default Product;
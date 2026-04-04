import React, { useEffect, useState } from "react";
import ProductList from "./components/ProductList";
import { Product } from "./types/product";
import "./App.css";

const API_BASE = "http://localhost:8080";
const ADMIN_KEY = "my-super-secret-key";

const App: React.FC = () => {
    const [allProducts, setAllProducts] = useState<Product[]>([]);
    const [filteredProducts, setFilteredProducts] = useState<Product[]>([]);
    const [categories, setCategories] = useState<string[]>([]);
    const [selectedCategory, setSelectedCategory] = useState<string>("all");
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        fetch(`${API_BASE}/products/`, {
            headers: { "X-Admin-Key": ADMIN_KEY }
        })
            .then(res => {
                if (!res.ok) throw new Error(`HTTP error: ${res.status}`);
                return res.json();
            })
            .then(data => {
                const products = data || [];
                setAllProducts(products);
                setFilteredProducts(products);

                // Build unique categories list
const cats: string[] = [];
products.forEach((p: Product) => {
    if (p.category && p.category !== "" && !cats.includes(p.category)) {
        cats.push(p.category);
    }
});
setCategories(cats);
setLoading(false);
                setCategories(cats);
                setLoading(false);
            })
            .catch(err => {
                setError(err.message);
                setLoading(false);
            });
    }, []);

    const handleCategoryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const selected = e.target.value;
        setSelectedCategory(selected);
        if (selected === "all") {
            setFilteredProducts(allProducts);
        } else {
            setFilteredProducts(allProducts.filter(p => p.category === selected));
        }
    };

    return (
        <div className="app">
            <header className="app-header">
                <h1>Inventory Management</h1>
                <nav className="nav-bar">
                    <a href="/">Products</a>
                    <a href="/">Categories</a>
                </nav>
            </header>

            <main>
                {/* Filter Bar */}
                <div className="filter-bar">
                    <label htmlFor="category-filter">Filter by Category:</label>
                    <select
                        id="category-filter"
                        value={selectedCategory}
                        onChange={handleCategoryChange}
                    >
                        <option value="all">All Categories</option>
                        {categories.map(cat => (
                            <option key={cat} value={cat}>{cat}</option>
                        ))}
                    </select>
                </div>

                {loading && <p>Loading products...</p>}
                {error && <p className="error">Error: {error}</p>}
                {!loading && !error && <ProductList products={filteredProducts} />}
            </main>
        </div>
    );
};

export default App;
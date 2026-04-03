console.log("app.js loaded!");

const API_BASE = "http://localhost:8080";
const ADMIN_KEY = "my-super-secret-key";

// Store all products globally so we can filter without re-fetching
let allProducts = [];

async function fetchProducts() {
    try {
        const response = await fetch(`${API_BASE}/products/`, {
            headers: {
                "X-Admin-Key": ADMIN_KEY
            }
        });

        if (!response.ok) {
            throw new Error(`HTTP error: ${response.status}`);
        }

        const products = await response.json();
        console.log("Products fetched:", products);

        if (!products || products.length === 0) {
            document.getElementById("product-list").innerHTML = "<p>No products found.</p>";
            return;
        }

        allProducts = products;
        populateCategoryFilter(products);
        renderProducts(products);

    } catch (error) {
        console.error("Failed to fetch products:", error);
    }
}

// Build the category dropdown from real product data
function populateCategoryFilter(products) {
    const select = document.getElementById("category-filter");

    // Get unique categories from products
    const categories = [...new Set(products
        .map(p => p.category)
        .filter(c => c && c !== "")
    )];

    categories.forEach(category => {
        const option = document.createElement("option");
        option.value = category;
        option.textContent = category;
        select.appendChild(option);
    });

    // Listen for filter changes
    select.addEventListener("change", function() {
        const selected = this.value;
        if (selected === "all") {
            renderProducts(allProducts);
        } else {
            const filtered = allProducts.filter(p => p.category === selected);
            renderProducts(filtered);
        }
    });
}

function renderProducts(products) {
    const container = document.getElementById("product-list");
    container.innerHTML = "";

    if (products.length === 0) {
        container.innerHTML = "<p>No products in this category.</p>";
        return;
    }

    products.forEach(product => {
        const tile = document.createElement("div");
        tile.className = "product-tile";
        tile.innerHTML = `
            <h2 class="product-name">${product.name}</h2>
            <p class="product-price">$${product.price}</p>
            <p class="product-description">${product.description || "No description"}</p>
            <span class="product-category">${product.category || "Uncategorized"}</span>
        `;
        container.appendChild(tile);
    });
}

fetchProducts();
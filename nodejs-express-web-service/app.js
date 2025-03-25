const express = require('express');

const app = express();
const PORT = 8080;

// Middleware to log requests
app.use((req, res, next) => {
    console.log(`Received request: ${req.method} ${req.url}`);
    next();
});

// Function to fetch cocktails using fetch
const fetchCocktails = async (letter, res) => {
    try {
        const url = `https://www.thecocktaildb.com/api/json/v1/1/search.php?f=${letter}`;
        const response = await fetch(url);
        const data = await response.json();
        res.json(data);
    } catch (error) {
        console.error('Error fetching data:', error);
        res.status(500).json({ error: 'Failed to fetch data' });
    }
};

// Route to get cocktails starting with 'a'
app.get('/cocktails', (req, res) => fetchCocktails('a', res));

// Route to get cocktails by letter
app.get('/cocktails/:letter', (req, res) => fetchCocktails(req.params.letter, res));

// Start the server
app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});


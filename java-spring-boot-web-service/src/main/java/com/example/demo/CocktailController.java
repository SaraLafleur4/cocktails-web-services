package com.example.demo;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

@RestController
@RequestMapping("/cocktails")
public class CocktailController {

    private static final String API_URL = "https://www.thecocktaildb.com/api/json/v1/1/search.php?f=";

    @GetMapping
    public ResponseEntity<String> getCocktails() {
        return fetchCocktails("a"); // Default letter 'a'
    }

    @GetMapping("/{letter}")
    public ResponseEntity<String> getCocktailsByLetter(@PathVariable String letter) {
        return fetchCocktails(letter);
    }

    private ResponseEntity<String> fetchCocktails(String letter) {
        RestTemplate restTemplate = new RestTemplate();
        String url = API_URL + letter;

        try {
            ResponseEntity<String> response = restTemplate.getForEntity(url, String.class);
            return ResponseEntity.ok(response.getBody());
        } catch (Exception e) {
            return ResponseEntity.internalServerError().body("{\"error\": \"Failed to fetch data\"}");
        }
    }
}

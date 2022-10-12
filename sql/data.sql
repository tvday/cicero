PRAGMA foreign_keys = ON;

INSERT INTO cocktails(name, cocktailid, history, recipe)
VALUES
    ('Old Fashioned', 1, 'This is the history of Old Fashioned', 'This is the recipe of Old Fashioned'),
    ('Sazerac', 2, 'This is the history of Sazerac', 'This is the recipe of Sazerac'),
    ('Margarita', 3, 'This is the history of Margarita', 'This is the recipe of Margarita'),
    ('Moscow Mule', 4, 'This is the history of Moscow Mule', 'This is the recipe of Moscow Mule'),
    ('Mojito', 5, 'This is the history of Mojito', 'This is the recipe of Mojito'),
    ('Long Island Iced Tea', 6, 'This is the history of Long Island Iced Tea', 'This is the recipe of Long Island Iced Tea'),
    ('Negroni', 7, 'This is the history of Negroni', 'This is the recipe of Negroni'),
    ('Whiskey Sour', 8, 'This is the history of Whiskey Sour', 'This is the recipe of Whiskey Sour'),
    ('Martini', 9, 'This is the history of Martini', 'This is the recipe of Martini');

INSERT INTO ingredients(ingredientid, category, type, subtype) {
    (1, 'spirit', 'whiskey', 'bourbon'),
    (2, 'spirit', 'whiskey', 'Scotch'),
    (3, 'spirit', 'whiskey', 'Irish'),
    (4, 'spirit', 'vodka'),
    (5, 'bitters', 'Angostura'),
    (6, 'other', 'sugar cube'),
    (7, 'other', 'water'),
    (8, 'garnish', 'oragne twist'),
    (9, 'garnish', 'cocktail cherry'),
    (10, 'bitters', 'Peychaud'),
    (11, 'spirit', 'whiskey', 'rye'),
    (12, 'spirit', 'brandy', 'cognac'),
    (13, 'spirit', 'absinthe'),
    (14, 'garnish', 'lemon twist'),
    ;

}

INSERT INTO branded_ingredients(name, brand, flavor, proof)
VALUES
    ('Jim Beam', 'Jim Beam', 'This is flavor the of this', 80),
    ('Vanilla', 'Jim Beam', 'This is the flavor of this', 70),
    ('Single Barrel', 'Jim Beam', 'This is flavor of this', 90),
    ('Four Roses', 'Four Roses', 'This is flavor of this', 80),
    ;


INSERT INTO cocktail_contents(cocktailid, ingredientid, amount, measurement)
VALUES
    -- Old Fahsioned
    (1, 1, 1.5, 'oz'), -- bourbon
    (1, 5, 2, 'dash'), -- bitters
    (1, 6, 1, 'count'), -- sugar cube
    (1, 7, 1, 'dash'), -- water
    (1, 8, 1, 'count'), -- orange twist
    (1, 9, 1, 'count'), -- cocktail cherry
    -- Sazerac
    (2, 11, 1.5, 'oz'), -- rye
    (2, 12, 1, 'oz'), -- cognac
    (2, 13, 1, 'teaspoon'), -- absinthe
    (2, 10, 3, 'dashes'), -- bitters
    (2, 6, 1, 'count'), -- sugar cube
    (2, 14, 1, 'count'); -- lemon twist

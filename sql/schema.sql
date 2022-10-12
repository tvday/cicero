PRAGMA foreign_keys = ON;

CREATE TABLE cocktails (
  cocktailid INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(40) NOT NULL,
  history VARCHAR(1024) NOT NULL,
  recipe VARCHAR(64) NOT NULL,
  created DATETIME DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE branded_ingredients (
  brandedid INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(40) NOT NULL,
  brand VARCHAR(40),
  flavor VARCHAR(512),
  proof INTEGER,
  created DATETIME DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE ingredients (
  ingredientid INTEGER PRIMARY KEY AUTOINCREMENT,
  category VARCHAR(40) NOT NULL,
  type VARCHAR(40),
  subtype VARCHAR(40),
  created DATETIME DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE cocktail_contents (
  cocktailid INTEGER NOT NULL,
  ingredientid INTEGER NOT NULL,
  amount REAL,
  measurement VARCHAR(40),
  created DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (cocktailid, ingredientid),
  FOREIGN KEY (cocktailid) REFERENCES cocktails(cocktailid)
    ON DELETE CASCADE,
  FOREIGN KEY (ingredientid) REFERENCES ingredients(ingredientid)
      ON DELETE CASCADE,
);

CREATE TABLE isa (
    brandedid INTEGER NOT NULL,
    ingredientid INTEGER NOT NULL,
    created DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (brandedid, ingredientid),
    FOREIGN KEY (brandedid) REFERENCES brandeds(brandedid)
      ON DELETE CASCADE,
    FOREIGN KEY (ingredientid) REFERENCES ingredients(ingredientid)
        ON DELETE CASCADE,
);

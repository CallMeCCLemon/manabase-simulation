CREATE TABLE cards
(
    ID   CHAR(36) PRIMARY KEY,
    Name VARCHAR(255),
    Colors TEXT CHECK (colors IN ('White', 'Blue', 'Green', 'Red', 'Black', 'Colorless')),
    EntersTapped BOOLEAN,
    Types TEXT CHECK (types IN ('Plains', 'Island', 'Swamp', 'Mountain', 'Forest', 'Wastes'))
);


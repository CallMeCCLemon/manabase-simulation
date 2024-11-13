# Usage

This page will show you how to run this simulation tool on your own custom decklists and objectives.

## Examples
Please consult the [fixtures](../fixtures) directory for working examples of each data structure. 

## Requirements
* go 1.22
* Decklist
* Test Objective
* Game Configuration

## Defining a Decklist
This simulation primarily works on the concept of "lands" which produce mana and "nonlands" which do not. 
MTG has many different nonlands (artifacts, creatures, etc) which produce mana but these are currently "out of scope" for the current release.
Using this simulation in its current form requires a user to define a JSON decklist with the lands and non-lands separated. 
These will then be converted into an instance of a Deck during the simulation and tested against. 

Sample Lotus Field decklist can be found [here](../fixtures/lotus-field-deck.json).

### Lands
The current land structure is documented [here](package/model/decklist.go).

#### Sample
```json
{
  "name": "Forest",
  "colors":  ["green"],
  "entersTapped":  false,
  "activationCost":  [],
  "quantity":  1
}
```

### Non-Lands
The current non-land structure is documented [here](package/model/decklist.go).

```json
{
  "name": "Llanowar Elves",
  "castingCost": [
    {
      "colorRequirements": ["white", "white", "white"],
      "genericCost": 1
    }
  ],
  "quantity":  34
}
```

## Defining a Test Objective
The current test objective structure is documented [here](package/model/objective.go).

### Example
```json
{
  "targetTurn": 3,
  "manaCosts": [
    {
      "colorRequirements": ["white", "white"],
      "genericCost": 1
    }
  ]
}
```

## Defining a Game Configuration
The current test objective structure is documented [here](package/model/config.go).

### Example
```json
{
  "initialHandSize": 7,
  "cardsDrawnPerTurn": 1,
  "onThePlay": true
}
```

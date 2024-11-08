# Manabase Simulation

This is a simulation tool for testing deck list manabases using a monte-carlo style simulation to compute the frequency
a player will have the mana to cast a spell by a target turn.

## Background

I recently played in the North American Regional Championship (RC) for Magic: The Gathering (MTG) in Ocotober 2024. The
format was Pioneer and I decided to play the Lotus field combo deck. With Rakdos Prowess running loose and resulting in
MANY turn-3
kills, I wanted to understand how consistently Lotus field could power out Nine Lives (WW1 cost) given the mana base
popular lists were running. I consistently felt like it was not able to cast the spell even if it was in hand due to the
restrictive mana cost and wanted better data than what I could generate gold-fishing. Nine Lives was far from a good
answer to prowess, but dropping it on T3 could buy you a couple of turns to try and go off before they ticked it to nine
counters. The project unfortunately did not complete before the RC, but it has been proving an interesting project for
me to work with and formalize my thoughts on how I play MTG.

## Overview

This project was largely an attempt to statistically determine how consistently Lotus field could power out Nine Lives.
As someone who isn't much of a statistician, but rather someone familiar with jamming out some code, I knew I could
write
a monte-carlo simulation to run hundreds of thousands of "gold-fished" hands which end with either nine lives
being cast-able or not and then determine the frequency of success.

## Simulation Workflow

A single iteration of the simulation is as follows:

1. Deal a hand of cards to the player based on the configuration (usually 7).
2. For each turn (including the target turn), play the best land from the hand based on the mana costs of the objective
   and the existing lands on the board.
3. If it is the target turn, validate the objective and return the result; true if success, otherwise false.

### Success Evaluation Formula

The formula for the success evaluation of the simulation is a simple ratio:

$$
p = \frac{s}{i} \times 100 \%
$$

where:

- $p$ is the percentage of success.
- $success$ is the number of times Nine Lives is cast-able by the target turn.
- $iterations$ is the total number of hands simulated.

## Playing Lands

Deciding how to play lands in this simulation is a bit more complicated than just arbitrarily picking a land to play.
Experienced players know how important it is to plan future turns and account which lands remain in the deck. In MTG,
lands can produce one or more colors (sometimes colorless) which means there are different reasons to choose
which land to play on a given turn. Additionally, lands can enter play tapped or untapped depending on the land. Lands
which are tapped aren't usable until the following turn after they've "untapped."

### Meeting an Objective

The approach I decided to take was to first start with evaluating a given board state to determine if the objective was
already met. If it was, then we are all good!
If it wasn't, then we can determine the possible combinations of colors/costs which would be needed to meet the
objective. This is where players tend to develop an intuition for how to tap their lands, but codifying it was a bit of
a challenge.
I used a breadth-first search algorithm to identify which different combinations of mana costs could be
leftover after using the current land for each of its colors individually.

#### Rules and Examples:

If a land could produce 2 colors and both were part of the objective, then we would need to keep track of both
possible remaining costs and use those in future calculations.

```
Land -> (can produce) U (blue) / R (red)
Objective (Mana Cost): pay [U+R+G (green)],
 
play(Land) "playing the land" => two new mana costs: [U+G, R+G]
```

If tapping the land for one or more of its colors wouldn't simplify the objective, those would be pruned from the list
of possible costs keeping us at a minimal set of costs.

```
Land -> U / R
Objective: pay [U+G]
play(Land) => one new mana cost: [G]
```

Similarly, if a land's produced colors couldn't be used for explicit color pips, they could still be used for generic
costs and would not be pruned.

```
Land -> U / R
Objective: pay [W (white)+G+2]
play(Land) => one new mana cost: [W+G+1]
```

### Land Selection

Land selection presented an interesting challenge and is an area I plan to improve upon in the future. Currently, land
selection is done using a scoring function to score each land in the hand and playing the land with the highest score.
I wanted to ensure lands which could produce multiple colors were prioritized over lands which produced fewer.
Additionally,
I wanted to ensure untapped lands were ALWAYS prioritized on the target turn when the mana cost needed to be met. I had
to make a conscious decision to simplify the simulation for now to get something functional out and so the current
framework does not allow tapped lands to be played on the target turn to simplify the calculation (as they'd be unusable
anyway).

#### Scoring

Land scoring is as follows:

$$
score = \sum_{i=1}^{n} (10 \times \sum_{j=1}^{m_i} m_i[j] \in l_c) + l_c 
$$

where:

- $score$ is the score for the land.
- $n$ is the number of mana costs.
- $m$ is the set of all mana costs.
- $n_c$ is the _unique_ number of colors in the mana cost n.
- $l_c$ is the set of colors which can be produced by the land.
- $m_i[j] \in l_c$ is 1 if the current color in the mana cost ($m_i[j]$) can be produced by the land.

This scoring function was a useful mechanism to prioritize lands which would meet the existing mana costs available 
while using the number of colors as a tie-breaker. There is more room for improvement here, but this was a functional start.

## Conclusion

Overall, this project felt like a great success. After a million iterations of goldfishing, I see a convergence at ~52% 
success for the lotus field list provided in this repo. It assumes you ALWAYS have 9 lives in hand which doesn't reflect 
reality, but was a good ballpark estimate which matched my experiences while practicing the deck.

## Future Work

- Allow tapped lands to be played on the target turn.
- Allow all non-lands to simultaneously be the objective for the deck.
- Update the simulation success to be equal to the number of castable non-lands in the hands by target turn.
- Allow for more complex lands to be used (filter lands, verges, fetchlands, etc.)

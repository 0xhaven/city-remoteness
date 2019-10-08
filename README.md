# City Remoteness

### Determine how far cities are from other people.


1. Parses ~10,000 cities' location and population data from the [SimpleMaps World Cities Basic Database](https://simplemaps.com/data/world-cities) ([Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/))

2. Assigns each city a score that is a the sum of every other city' population times an inverse exponential of their pairwise ditance.

3. Generates a CSV than I loaded in [Google Sheets](https://docs.google.com/spreadsheets/d/1Y3PI0Lhc9iSK4U2jkoPOzjBqGLZYaZFYgF8Qi5vR2Hk/edit?usp=sharing) to create a pretty chart.

![City Remoteness](assets/City%20Remoteness.svg "City Remoteness")

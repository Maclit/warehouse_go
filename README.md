# WAREHOUSE GO - TEAM AMAZON

## COMPILATION AND EXECUTION

To compile the project, you should run :

```
go build .
```

It will compile an executable for the platform you currently are working on.

To execute the project, you can use the command :

```
go run .
```

The program will first ask you to write as an input the path to the configuration file, wich contains all the map informations.

## PROJECT ORGANISATION

The project in composed of only 1 package : main. All its files can be found in the root directory.
The main part of the program is contained in the warehouse_graph_*.go files, it is where the logic is being run.

## STRATEGY

To complete this project, our team choosed to use a fairly simple strategy :

 1 . Transporters wich are empty will move toward the nearest box to pick it up.

 2 . Transporters carrying a box will go the the nearest truck to drop it off.

 3 . If the truck is too full to accept another box, it will depart the required amount of turns.


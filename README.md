# WAREHOUSE GO - TEAM AMAZON

## COMPILATION AND EXECUTION

To compile the project, you should run :

```
go build .
```

It will compile an executable for the platform you currently are working on.

To execute the project, you can use the command :

```
go run . [map_file_path]
```

map_file_path corresponds to the path of the file containing the map configuration.

To run unit tests, use :

```
go test
```


## PROJECT ORGANISATION

The project in composed of only 1 package : main. All its files can be found in the root directory.

In the warehouse_graph_*.go files, all function related to the graph representing the map can be found.

This functions are used in the game_logic_main_lopp.go files, where all the main loop logic is found.


## STRATEGY

To complete this project, our team choosed to use a fairly simple strategy :

 1 . Transporters wich are empty will move toward the nearest box to pick it up.

 2 . Transporters carrying a box will go the the nearest truck to drop it off.

 3 . If the truck is too full to accept another box, it will depart the required amount of turns.


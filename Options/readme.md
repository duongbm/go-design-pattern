# Options Pattern

### The problems:
- A complex object that has many optional parameters. 
And there is a way to implemented it that is make a constructor
takes all of these parameters and provides the optional values. But there drawbacks of this solution
1. hard to remember the order of the parameters
2. hard to know whether a parameter is optional or not.
3. constructor becomes very long and is not readable.

### The solutions:
- Options pattern can be used to create objects with many optional parameters.
- Define a struct and provide method to set those parameters.

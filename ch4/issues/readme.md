##usage


```sh
# ./issues repo:roylee0704/react-flexbox-grid is:open
```

## formatter

- a digit in front of a formatter: d, s, f indicate character's width.
- a digit after a decimal of a formatter: s indicate only allow n number of characters to be displayed. If unfilled, show all characters. (only works in `string`)


### Example
1. `%9.5s`  shows that there are 9 characters wide spaces provided, and from the given string value, only show up to 5 characters.

Hence, fmt.Printf("%5.3s", "Lee Siong Tai") outputs: `__Lee` with underscore indicate empty space.

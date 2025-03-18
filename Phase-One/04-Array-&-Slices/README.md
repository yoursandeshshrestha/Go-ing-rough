# Phase One: Arrays and Slices

## 04-Array-&-Slices

### JavaScript vs Go Comparison

#### 1. Arrays

**JavaScript:**

```javascript
// Fixed-size array (not really, JS arrays are dynamic)
const numbers = [1, 2, 3];
// or
const numbers = new Array(3);
numbers[0] = 1;
numbers[1] = 2;
numbers[2] = 3;
```

**Go:**

```go
// Fixed-size array
var numbers [3]int
numbers[0] = 1
numbers[1] = 2
numbers[2] = 3
```

Key Differences:

- JavaScript arrays are always dynamic (can grow/shrink)
- Go arrays are fixed-size (like JavaScript's `TypedArray`)
- Go requires type declaration (`[3]int` means array of 3 integers)

#### 2. Dynamic Arrays (Slices in Go)

**JavaScript:**

```javascript
// Dynamic array
const cities = ["New York", "Paris", "London"];
// Adding elements
cities.push("Tokyo");
// Removing elements
cities.splice(2, 1); // Remove 1 element at index 2
```

**Go:**

```go
// Slice (dynamic array)
var cities = []string{"New York", "Paris", "London"}
// Adding elements
cities = append(cities, "Tokyo")
// Removing elements
cities = append(cities[:2], cities[3:]...)
```

Key Differences:

- JavaScript uses `push()` and `splice()`
- Go uses `append()` function
- Go slices are more explicit about their underlying array

#### 3. Copying Arrays

**JavaScript:**

```javascript
// Shallow copy
const newCities = [...cities];
// or
const newCities = cities.slice();
```

**Go:**

```go
// Copy slice
var newCities = make([]string, len(cities))
copy(newCities, cities)
```

Key Differences:

- JavaScript has multiple ways to copy arrays
- Go requires explicit allocation with `make()` and `copy()`

#### 4. Iteration

**JavaScript:**

```javascript
// Using forEach
cities.forEach((city, index) => {
  console.log(`City ${index}: ${city}`);
});

// Using for...of
for (const [index, city] of cities.entries()) {
  console.log(`City ${index}: ${city}`);
}
```

**Go:**

```go
// Using range
for i, city := range cities {
    fmt.Printf("City %d: %s\n", i, city)
}
```

Key Differences:

- JavaScript has multiple iteration methods
- Go uses the `range` keyword for iteration
- Go's `range` provides both index and value automatically

### Key Concepts:

1. **Arrays vs Slices**:

   - Arrays: Fixed-size collections (like JavaScript's TypedArray)
   - Slices: Dynamic arrays (like JavaScript's regular arrays)

2. **Memory Management**:

   - JavaScript: Automatic memory management
   - Go: More explicit memory management with slices

3. **Type Safety**:

   - JavaScript: Dynamic typing (can mix types in arrays)
   - Go: Static typing (all elements must be same type)

4. **Performance**:
   - Go arrays/slices are generally more performant
   - JavaScript arrays are more flexible but less predictable in performance

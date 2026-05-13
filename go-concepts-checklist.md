# Go Concepts Self-Test Checklist

Use this checklist to verify your understanding of the Go concepts covered in
this workspace. For each section, mark the items you can explain and implement
confidently.

---

## 1. Hello Go

- [ ] I can create a Go module and Go file.
- [ ] I can write and run a simple `main` package.
- [ ] I understand the `package main` and `func main()` entry point.
- [ ] I can import standard library packages like `fmt`.
- [ ] I can use `fmt.Println()` to print text.
- [ ] I know how to build and run a Go program (`go run`, `go build`).

---

## 1.1 Packages & Modules

- [ ] I can initialize a module with `go mod init`.
- [ ] I understand the purpose of `go.mod`.
- [ ] I know how package names relate to folder and file structure.
- [ ] I can use imported package functions from another package.
- [ ] I know how to run code in a package with `go run ./...` or
      `go test ./...`.

---

## 2. Simple Values

- [ ] I can declare variables with explicit types.
- [ ] I understand zero values for basic types.
- [ ] I can use `fmt.Printf` with formatting verbs.
- [ ] I know how to use `const` for constants.
- [ ] I can initialize values with shorthand syntax.

---

## 3. Variables

- [ ] I understand the difference between `var` declaration and short assignment
      `:=`.
- [ ] I know how to declare multiple variables at once.
- [ ] I can use typed and untyped variable declarations.
- [ ] I know when to use `var` outside functions vs `:=` inside functions.
- [ ] I understand variable scope in Go.
- [ ] I can use the blank identifier `_` to ignore a value.

---

## 4. Constants

- [ ] I can define named constants.
- [ ] I understand how constant expressions work.
- [ ] I know what `iota` is and how it is used.
- [ ] I can create related constant groups with `iota`.
- [ ] I understand constant types versus literal values.

---

## 5. For Loop

- [ ] I know the basic `for` loop syntax.
- [ ] I can write a loop with initialization, condition, and post statements.
- [ ] I can write a loop with only a condition.
- [ ] I can write an infinite loop and stop it with `break`.
- [ ] I understand using `continue` inside loops.

---

## 6. If / Else

- [ ] I can write a simple `if` statement.
- [ ] I know how to add an `else` branch.
- [ ] I understand `if` with an initialization statement.
- [ ] I can compare values using Go operators.
- [ ] I know how to check multiple conditions.

---

## 7. Switch

- [ ] I can write a basic `switch` statement.
- [ ] I understand how `case` values work.
- [ ] I know how to use `switch` without an expression.
- [ ] I understand fallthrough behavior and the default case.
- [ ] I can use multiple expressions in a single `case`.

---

## 8. Arrays

- [ ] I know how to declare an array with a fixed length.
- [ ] I can set and get values from an array.
- [ ] I understand the difference between length and capacity.
- [ ] I know how to iterate over an array.
- [ ] I understand arrays are value types in Go.

---

## 9. Slices

- [ ] I can declare and initialize a slice.
- [ ] I know how to append items to a slice.
- [ ] I understand the relationship between slices and arrays.
- [ ] I know a nil slice is a valid zero value.
- [ ] I can slice a slice with `[:]` notation.
- [ ] I know the difference between length and capacity of a slice.
- [ ] I understand how `append` may allocate a new underlying array.
- [ ] I know to to `copy` slice from source slice to destination slice.
- [ ] I know how to `Declare 2D Slice`.
- [ ] I Know how about `slices` built in functions for equality, comparision,
      sorting.

---

## 10. Maps

- [ ] I can declare and initialize a map.
- [ ] I know how to add, update, and delete map entries.
- [ ] I can read a value safely from a map using the comma-ok idiom.
- [ ] I understand that maps are reference types.
- [ ] I can iterate over a map with `for range`.
- [ ] I know how to check if a key exists in a map.
- [ ] I understand maps are unordered and iteration order is not fixed.

---

## 11. Range

- [ ] I know how to use `range` with arrays.
- [ ] I can use `range` with slices.
- [ ] I can use `range` with maps.
- [ ] I understand the two-value form: `index, value`.
- [ ] I know how to ignore either index or value with `_`.
- [ ] I can use `range` to iterate over strings and runes safely.
- [ ] I understand `range` over a string iterates UTF-8 code points.

---

## 12. Functions

- [ ] I can define a function with parameters.
- [ ] I know how to return a value from a function.
- [ ] I can define a function returning multiple values.
- [ ] I understand how function arguments are passed by value.
- [ ] I can call a function from `main()`.
- [ ] I know how to use named return values.

---

## 13. Variadic Functions

- [ ] I know what a variadic function is.
- [ ] I can declare a function with `...T` parameters.
- [ ] I can call a variadic function with multiple arguments.
- [ ] I can call a variadic function by expanding a slice with `...`.
- [ ] I understand variadic parameters are received as a slice.

---

## Pointers

- [ ] I understand what a pointer is in Go.
- [ ] I can get the address of a variable using `&`.
- [ ] I can dereference a pointer using `*`.
- [ ] I understand the difference between pointer and value types.
- [ ] I can pass a pointer to a function to modify the original value.
- [ ] I know how to use pointers with composite types where appropriate.
- [ ] I understand a nil pointer and how to avoid dereferencing it.

---

## Self-Test Instructions

1. Read each checklist item.
2. Try to write a small example program or code snippet for each item without
   looking up the answer.
3. If you can explain the concept and code it, mark the item done.
4. If you cannot, review the corresponding folder and `main.go` file in the
   workspace.
5. Repeat the checklist until all items are checked.

---

## Notes

- Use `go run <folder>/main.go` to run examples in each folder.
- Use `go test` only after you add tests if you want to verify behavior
  automatically.
- This checklist is built from the workspace topic folders and is meant to cover
  the core concepts here.

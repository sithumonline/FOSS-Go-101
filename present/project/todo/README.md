# Mini Project: Todo List

Build a small in-memory Todo List using basic Go.

This folder is a **starter classroom version**: core methods in `todo.go`
contain TODO steps, and students implement them level by level.

## Learning goals

- Define and use `struct` types
- Write methods with pointer receivers
- Use slices for storage
- Handle errors cleanly
- Verify behavior with unit tests

## What students build

Implement these behaviors in `todo.go`:

- add a task
- mark a task as done
- remove a task
- list all tasks
- count pending tasks

## Run tests

From the repository root:

```bash
go test ./present/project/todo
```

## Level-by-level test progression

Students can run one level at a time:

```bash
go test ./present/project/todo -run '^TestLevel1'
go test ./present/project/todo -run '^TestLevel2'
go test ./present/project/todo -run '^TestLevel3'
go test ./present/project/todo -run '^TestLevel4'
```

Then run everything:

```bash
go test ./present/project/todo
```

Expected classroom flow:

- Before implementing a level, that level's tests should fail.
- After completing the TODOs for that level, re-run only that level's tests.
- Move to the next level when current level passes.

## Run the sample CLI main

```bash
go run ./present/project/todo/cmd/todo
```

Try:

- `add Learn Go slices`
- `add Build todo app`
- `list`
- `done 1`
- `pending`
- `remove 2`
- `list`

## Classroom idea

In your classroom starter repo, keep tests as-is and replace some method bodies
with TODO stubs so students must implement them to make tests pass.

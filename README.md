# task-tracker

Task tracker is a CLI tool that allows you to add tasks and track them.

***

## Features

CRUD Operations:
- Add a task
- Delete a task
- Update task status to: todo, in-progress, done
- List tasks by status

## Usage

```
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

```
# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1
```

```
# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1
```

```
# Listing all tasks
task-cli list
```

```
# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```


### Installations

Clone the repository

```
git clone https://github.com/baldeosinghm/task-tracker.git
```

Run the executable

```
go run .
```

### Dependencies

None; the tool doesn't make use of any external libraries.
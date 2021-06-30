# Gows - A Golang Workflow Scheduler

<img src="assets/gows.png" width="15%"/>

[![Go](https://github.com/Software-Craft-Factory/Gows/actions/workflows/go.yml/badge.svg)](https://github.com/Software-Craft-Factory/Gows/actions/workflows/go.yml)


<img src="./assets/gows.gif" />


Gows is an easy and super lightweight workflow management tool.
There is no need for setting up databases and web applications, Gows only takes a couples of json configuration files and that is it.
It is all that it takes to define your workflow DAGs.

## Dag configuration
Gows's dag are describe using a json file : 
```json
{
   "dagName":"example-2",
   "schedule":"*/3 * * * * *",
   "tasks":[
      {
         "name":"bash-task-1",
         "type":"bash",
         "parameters":{
            "cmd":"ls -ah"
         }
      },
      {
         "name":"bash-task-2",
         "type":"bash",
         "parameters":{
            "cmd":"echo 'toto'"
         }
      },
      {
         "name":"bash-task-3",
         "type":"bash",
         "parameters":{
            "cmd":"sleep 2"
         },
         "dependencies":[
            "bash-task-1",
            "bash-task-2"
         ]
      }
   ]
}
```

## Package
Gows core package can be used without gows cli : 
 
```go get -u com.github/Fszta/Gows```


```golang
dag, _ := CreateDag("dag-example")

operator1 := operators.CreateBashOperator()
operator1.SetCmd("ls -lah")
task1, _ := task.CreateTask(operator1, "task1")
dag.AddTask(task1)

operator2 := operators.CreateBashOperator()
operator2.SetCmd("sleep 1 && tree")
task2, _ := task.CreateTask(operator2, "task2")
dag.AddTask(task2)
dag.SetDependency(task2, task1)

operator3 := operators.CreateBashOperator()
task3, _ := task.CreateTask(operator3, "task3")
operator3.SetCmd("ps")
dag.AddTask(task3)
dag.SetDependency(task3, task2)

dag.Run()
```

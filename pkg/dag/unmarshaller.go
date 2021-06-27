package dag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"com.github/Fszta/gows/pkg/operators"
	"com.github/Fszta/gows/pkg/task"
)

type DagConfig struct {
	DagName  string `json:"dagName"`
	Schedule string `json:"schedule"`
	Tasks    []struct {
		Name       string `json:"name"`
		Type       string `json:"type"`
		Parameters struct {
			Cmd       string `json:"cmd,omitempty"`
			Src       string `json:"src,omitempty"`
			Arguments []struct {
				Arg   string `json:"arg"`
				Value string `json:"value"`
			} `json:"arguments,omitempty"`
		} `json:"parameters"`
	} `json:"tasks"`
}

func UnmarshalDagConfig(configData []byte) (*DagConfig, error) {
	dagConfig := DagConfig{}

	err := json.Unmarshal(configData, &dagConfig)

	if err != nil {
		return nil, err
	}

	return &dagConfig, nil
}

func GetDagFromConfig(config *DagConfig) (*Dag, error) {

	if len(config.Tasks) == 0 {
		return nil, fmt.Errorf("configuration json must contains at least one task")
	}

	dag, err := CreateDag(config.DagName)

	if err != nil {
		return nil, err
	}

	for _, taskConfig := range config.Tasks {

		if taskConfig.Type == "python" {
			operator := operators.CreatePythonOperator()

			if taskConfig.Parameters.Src == "" {
				return nil, fmt.Errorf("src field can't be empty, task %v use python operator", taskConfig.Name)
			}

			operator.SetSrc(taskConfig.Parameters.Src)

			newTask, err := task.CreateTask(operator, taskConfig.Name)
			fmt.Printf("Add task %v to dag %v \n", taskConfig.Name, config.DagName)

			if err != nil {
				return nil, err
			}
			dag.AddTask(newTask)

		} else if taskConfig.Type == "bash" {
			operator := operators.CreateBashOperator()

			if taskConfig.Parameters.Cmd == "" {
				return nil, fmt.Errorf("cmd field can't be empty, task %v use bash operator", taskConfig.Name)
			}

			operator.SetCmd(taskConfig.Parameters.Cmd)

			newTask, err := task.CreateTask(operator, taskConfig.Name)
			fmt.Printf("Add task %v to dag %v \n", taskConfig.Name, config.DagName)

			if err != nil {
				return nil, err
			}
			dag.AddTask(newTask)

		}

	}

	dag.SetScheduler(config.Schedule)

	return dag, nil
}

func ReadDagConfig(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}
	return data, nil
}

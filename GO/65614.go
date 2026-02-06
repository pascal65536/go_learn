package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)


func SortNames(names []string) {
	sort.Strings(names)
}

func SortNums(nums []uint) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}




type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Company struct {
	workers []worker
}

type worker struct {
	name       string
	position   string
	salary     uint
	experience uint
}

var positionRank = map[string]int{
	"директор":       5,
	"зам. директора": 4,
	"начальник цеха": 3,
	"мастер":         2,
	"рабочий":        1,
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	position = strings.ToLower(strings.TrimSpace(position))
	if _, ok := positionRank[position]; !ok {
		return errors.New("недопустимая должность: " + position)
	}
	c.workers = append(c.workers, worker{
		name:       name,
		position:   position,
		salary:     salary,
		experience: experience,
	})
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	if len(c.workers) == 0 {
		return nil, errors.New("нет данных о сотрудниках")
	}
	sort.SliceStable(c.workers, func(i, j int) bool {
		incomeI := c.workers[i].salary * c.workers[i].experience
		incomeJ := c.workers[j].salary * c.workers[j].experience
		if incomeI != incomeJ {
			return incomeI > incomeJ
		}
		return positionRank[c.workers[i].position] > positionRank[c.workers[j].position]
	})
	result := make([]string, len(c.workers))
	for i, w := range c.workers {
		income := w.salary * w.experience
		result[i] = fmt.Sprintf("%s — %d — %s", w.name, income, w.position)
	}
	return result, nil
}

func main() {
	company := &Company{}
	_ = company.AddWorkerInfo("Михаил", "директор", 1000, 12)
	_ = company.AddWorkerInfo("Андрей", "мастер", 590, 20)
	_ = company.AddWorkerInfo("Игорь", "зам. директора", 1100, 10)
	_ = company.AddWorkerInfo("Олег", "рабочий", 500, 30)
	_ = company.AddWorkerInfo("Виктор", "начальник цеха", 900, 11)
	sorted, err := company.SortWorkers()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	for _, line := range sorted {
		fmt.Println(line)
	}
}





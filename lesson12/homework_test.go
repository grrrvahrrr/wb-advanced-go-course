package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Task struct {
	Identifier int
	Priority   int
}

type Scheduler struct {
	heap []Task
}

func NewScheduler() Scheduler {
	return Scheduler{}
}

func (s *Scheduler) AddTask(task Task) {
	s.heap = append(s.heap, task)
	s.fixHeapAfterAdd(len(s.heap) - 1)
}

func (s *Scheduler) fixHeapAfterAdd(index int) {
	if index == 0 || s.isHigherPriority((index-1)/2, index) {
		return
	}
	s.swap((index-1)/2, index)
	s.fixHeapAfterAdd((index - 1) / 2)
}

func (s *Scheduler) swap(a, b int) {
	s.heap[a], s.heap[b] = s.heap[b], s.heap[a]
}

func (s *Scheduler) isHigherPriority(a, b int) bool {
	return s.heap[a].Priority > s.heap[b].Priority
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	for i := range s.heap {
		if s.heap[i].Identifier == taskID {
			s.heap[i].Priority = newPriority
			break
		}
	}
	s.heapify(len(s.heap) / 2)
}

func (s *Scheduler) heapify(index int) {
	if index < 0 {
		return
	}
	max, left, right := index, index*2+1, index*2+2
	if left < len(s.heap) && s.isHigherPriority(left, max) {
		max = left
	}
	if right < len(s.heap) && s.isHigherPriority(right, max) {
		max = right
	}
	if max != index {
		s.swap(max, index)
	}
	s.heapify(index - 1)
}

func (s *Scheduler) GetTask() Task {
	task := s.heap[0]
	s.heap = s.heap[1:]
	s.heapify(0)
	return task
}

func TestTrace(t *testing.T) {
	task1 := Task{Identifier: 1, Priority: 10}
	task2 := Task{Identifier: 2, Priority: 20}
	task3 := Task{Identifier: 3, Priority: 30}
	task4 := Task{Identifier: 4, Priority: 40}
	task5 := Task{Identifier: 5, Priority: 50}

	scheduler := NewScheduler()
	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)
	scheduler.AddTask(task4)
	scheduler.AddTask(task5)

	task := scheduler.GetTask()
	assert.Equal(t, task5, task)

	task = scheduler.GetTask()
	assert.Equal(t, task4, task)

	scheduler.ChangeTaskPriority(1, 100)
	task1.Priority = 100

	task = scheduler.GetTask()
	assert.Equal(t, task1, task)

	task = scheduler.GetTask()
	assert.Equal(t, task3, task)
}

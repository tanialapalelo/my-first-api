package todo

type Service struct {
	todos []TodoItem
}

type TodoItem struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
}

func NewService(todos []TodoItem) *Service {
	return &Service{
		todos: make([]TodoItem, 0),
	}
}

func (svc *Service) Add(todo TodoItem) {
	svc.todos = append(svc.todos, todo)
}

func (svc *Service) GetAll() []TodoItem {
	return svc.todos
}

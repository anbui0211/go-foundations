package worker

type EmployeeResponse struct {
	Status  string     `json:"status"`
	Data    []Employee `json:"data"`
	Message string     `json:"message"`
}

type Employee struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

type SalaryAgeRatio struct {
	WorkerID   int     `json:"worker_id"`
	EmployeeID int     `json:"employee_id"`
	Ratio      float32 `json:"ratio"`
}

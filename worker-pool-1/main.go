package worker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

/*
1. Gọi API https://dummy.restapiexample.com/api/v1/employees và trả về 1 slice các struct
	có thông tin trong đấy
2. Viết 1 worker pool dựa trên ví dụ ở https://gobyexample.com/worker-pools,
	nhưng job là struct về người gồm các thông tin ở câu trên, job trả về lương chia cho tuổi của mỗi người
*/

/*
	Thực thi

start := time.Now()
elapsed := time.Since(start)
worker.DemoWorker("worker/data.json")
fmt.Println("Thời gian thực thi:", elapsed)
*/
func DemoWorker(url string) {
	// Get employee
	employees, err := getEmployee(url)
	if err != nil {
		log.Fatal(err)
	}

	var (
		numWorker = 5
		numJobs   = len(employees)
		jobs      = make(chan Employee, numJobs)
		results   = make(chan SalaryAgeRatio, numJobs)
	)

	fmt.Println(numJobs)

	// Init worker
	for i := 1; i <= numWorker; i++ {
		go worker(i, jobs, results)
	}

	// Gửi dữ liệu vào channel
	for _, employee := range employees {
		jobs <- employee
	}
	close(jobs)

	// Lấy dữ liệu từ channel
	for i := 0; i < numJobs; i++ {
		res := <-results
		fmt.Println("workerID:", res.WorkerID, "Ratio: ", res.Ratio)
	}

}

// getByJSONFile use to get data from a JSON file
func getByJSONFile(url string) (data EmployeeResponse, err error) {
	plan, _ := ioutil.ReadFile(url) // filename is the JSON file to read
	if err = json.Unmarshal(plan, &data); err != nil {
		return
	}
	return
}

func getEmployee(url string) (data []Employee, err error) {
	employeeRes, err := getByJSONFile(url)
	if err != nil {
		return
	}
	return employeeRes.Data, nil
}

// Initial worker
func worker(workerId int, jobs <-chan Employee, results chan<- SalaryAgeRatio) {
	for job := range jobs {
		// time.Sleep(time.Second)
		results <- SalaryAgeRatio{
			WorkerID:   workerId,
			EmployeeID: job.ID, // EmployeeID
			Ratio:      float32(job.EmployeeSalary / job.EmployeeAge),
		}
	}
}





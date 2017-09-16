package main

import (
    "github.com/prometheus/client_golang/api/prometheus/v1"
    "github.com/prometheus/client_golang/api"
    "time"
    "context"
    "fmt"
    "os"
    "github.com/prometheus/common/model"
    "math"
    "strconv"
)

func main() {
    os.Clearenv()

    c := api.Config{
        Address: `http://localhost:30900`,
    }
    apiClient, _ := api.NewClient(c)
    httpApi := v1.NewAPI(apiClient)


    q := `sum by(instance)(rate(container_cpu_usage_seconds_total{job="kubernetes-cadvisor",id="/"}[5m]))`

    result, _ := httpApi.Query(context.Background(), q,  time.Now())

    fmt.Println(result)
}


func parseTime(s string) (model.Time, error) {
    if t, err := strconv.ParseFloat(s, 64); err == nil {
        ts := t * float64(time.Second)
        if ts > float64(math.MaxInt64) || ts < float64(math.MinInt64) {
            return 0, fmt.Errorf("cannot parse %q to a valid timestamp. It overflows int64", s)
        }
        return model.TimeFromUnixNano(int64(ts)), nil
    }
    if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
        return model.TimeFromUnixNano(t.UnixNano()), nil
    }
    return 0, fmt.Errorf("cannot parse %q to a valid timestamp", s)
}

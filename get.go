package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type total struct {
	Confirmed int `json:"confirmed"`
	Recovered int `json:"recovered"`
	Deaths    int `json:"deaths"`
	Active    int `json:"active"`
}

type state_total struct {
	State     string `json:"state"`
	Confirmed int    `json:"confirmed"`
	Recovered int    `json:"recovered"`
	Deaths    int    `json:"deaths"`
	Active    int    `json:"active"`
}
type dat struct {
	Source        string        `json:"source"`
	LastRefreshed string        `json:"lastRefreshed"`
	Total         total         `json:"total"`
	Statewise     []state_total `json:"statewise"`
}

type ret struct {
	Success          bool   `json:"success"`
	Data             dat    `json:"data"`
	LastRefreshed    string `json:"lastRefreshed"`
	LastOriginUpdate string `json:"lastOriginUpdate"`
}

func check_file(file string) bool {
	_, err := os.Open(file)
	if err != nil {
		return false
	} else {
		return true
	}
}

func reader(file string) [][]string {
	record, err := os.Open(file)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	csvreader := csv.NewReader(record)
	prev_records, _ := csvreader.ReadAll()
	record.Close()
	return prev_records
}

func main() {
	// resp, _ := http.Get("https://api.rootnet.in/covid19-in/stats/latest")

	// change the colorscheme for the ouput by uncommenting the required lines
	colorReset := "\033[0m"
	// colorRed    := "\033[31m"
	// colorGreen  := "\033[32m"
	colorYellow := "\033[33m"
	// colorBlue   := "\033[34m"
	// colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[0;37m"
	heading := colorCyan
	row0 := colorReset
	row1 := colorWhite
	lines := colorYellow

	resp, _ := http.Get("https://api.rootnet.in/covid19-in/unofficial/covid19india.org/statewise")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body_string := string(body)
	bod := ret{}
	err := json.Unmarshal([]byte(body_string), &bod)
	if err != nil {
		fmt.Println(err)
	}

	filename := "cases.csv"
	var new_name string // change cases.csv to date.csv after reading
	var a, c, d, r int
	prev_file := check_file(filename)
	if prev_file {
		prev_records := reader(filename)
		new_name = prev_records[0][0] + ".csv"
		a, _ = strconv.Atoi(prev_records[1][1])
		c, _ = strconv.Atoi(prev_records[1][2])
		d, _ = strconv.Atoi(prev_records[1][3])
		r, _ = strconv.Atoi(prev_records[1][4])
	} else {
		a = 0
		c = 0
		d = 0
		r = 0
	}

	cmd := exec.Command("figlet", "-f", "slant", "Corona Term")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	ta := bod.Data.Total.Active - a
	tc := bod.Data.Total.Confirmed - c
	td := bod.Data.Total.Deaths - d
	tr := bod.Data.Total.Recovered - r

	fmt.Println(string(stdout))
	fmt.Printf("%s%d%s\n", "Active    - ", bod.Data.Total.Active, " ("+strconv.Itoa(ta)+")")
	fmt.Printf("%s%d%s\n", "Confirmed - ", bod.Data.Total.Confirmed, " ("+strconv.Itoa(tc)+")")
	fmt.Printf("%s%d%s\n", "Deaths    - ", bod.Data.Total.Deaths, " ("+strconv.Itoa(td)+")")
	fmt.Printf("%s%d%s\n", "Recovered - ", bod.Data.Total.Recovered, " ("+strconv.Itoa(tr)+")")
	fmt.Printf("%s%s\n", string(lines), "------------------------------------------------------------------------------------------")
	fmt.Printf("%s%-45s%-s\t%-s\t%-s\t%-s\n", string(heading), "STATE", "ACTIVE", "CONFIRMED", "DEATHS", "RECOVERED")
	fmt.Printf("%s%s\n", string(lines), "------------------------------------------------------------------------------------------")

	var curr_records [][]int
	var c_state_seq []string
	for i := 0; i < len(bod.Data.Statewise); i++ {
		var row []int
		row = append(row, bod.Data.Statewise[i].Active, bod.Data.Statewise[i].Confirmed, bod.Data.Statewise[i].Deaths, bod.Data.Statewise[i].Recovered)
		curr_records = append(curr_records, row)
		c_state_seq = append(c_state_seq, bod.Data.Statewise[i].State)
	}

	if prev_file {
		for i := 0; i < len(bod.Data.Statewise); i++ {
			if i%2 == 1 {
				fmt.Printf("%s%-45s%6d%s%6d%s%6d%s%6d\n", string(row1), c_state_seq[i], curr_records[i][0], "      ", curr_records[i][1], "      ", curr_records[i][2], "      ", curr_records[i][3])

			} else {
				fmt.Printf("%s%-45s%6d%s%6d%s%6d%s%6d\n", string(row0), c_state_seq[i], curr_records[i][0], "      ", curr_records[i][1], "      ", curr_records[i][2], "      ", curr_records[i][3])
			}
		}
	} else {
		for i := 0; i < len(bod.Data.Statewise); i++ {
			if i%2 == 1 {
				fmt.Printf("%s%-45s%6d%s%6d%s%6d%s%6d\n", string(row1), c_state_seq[i], curr_records[i][0], "      ", curr_records[i][1], "      ", curr_records[i][2], "      ", curr_records[i][3])

			} else {
				fmt.Printf("%s%-45s%6d%s%6d%s%6d%s%6d\n", string(row0), c_state_seq[i], curr_records[i][0], "      ", curr_records[i][1], "      ", curr_records[i][2], "      ", curr_records[i][3])
			}
		}
	}

	// adding date and details to csv file
	var rows [][]string
	var rowo []string
	rowo = append(rowo, bod.LastRefreshed[:10], bod.LastRefreshed[11:19], "", "", "")
	rows = append(rows, rowo)
	var rowc []string
	rowc = append(rowc, "India", strconv.Itoa(bod.Data.Total.Active), strconv.Itoa(bod.Data.Total.Confirmed), strconv.Itoa(bod.Data.Total.Deaths), strconv.Itoa(bod.Data.Total.Recovered))
	rows = append(rows, rowc)

	// converting the previous file to that date.
	os.Rename("cases.csv", new_name)

	csvfile, err := os.Create("cases.csv")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(bod.Data.Statewise); i++ {
		var row []string
		row = append(row, bod.Data.Statewise[i].State, strconv.Itoa(bod.Data.Statewise[i].Active), strconv.Itoa(bod.Data.Statewise[i].Confirmed), strconv.Itoa(bod.Data.Statewise[i].Deaths), strconv.Itoa(bod.Data.Statewise[i].Recovered))
		rows = append(rows, row)
	}

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range rows {
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()
	csvfile.Close()
}

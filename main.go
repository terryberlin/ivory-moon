package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func main() {
	http.HandleFunc("/getemployees", GetEmployees)
	http.HandleFunc("/postemployees", PostEmployees)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Println(err)
	}
}

type page struct {
	FieldName    string
	GetEmployees []GetEmployee
}

//GetEmployee : GetEmployee is a structure for employee data.
type (
	GetEmployee struct {
		Empl_ID  *string `db:"Empl_ID" json:"Empl_ID"`
		Store_no *string `db:"Store_no" json:"Store_no"`
		Emp_no   *string `db:"Emp_no" json:"Emp_no"`
	}
)

//PostEmployee : PostEmployee is a structure for posting employee data.
type (
	PostEmployee struct {
		Empl_ID  *string `db:"Empl_ID" json:"Empl_ID"`
		Store_no *string `db:"Store_no" json:"Store_no"`
		Emp_no   *string `db:"Emp_no" json:"Emp_no"`
	}
)

//GetEmployees : GetEmployees is a function for requesting employee data.  The url accepts arguments for begin date, end date, and employee units.
func GetEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content Type", "text/html")
	templates := template.New("template")
	templates.New("Body").Parse(doc)
	templates.New("List").Parse(docList)

	copyfromemployee := r.URL.Query().Get("copyfromemployee")

	//sql := `Exec get_employee_attributes $1`
	sql := `select Empl_ID, Store_no, Emp_no from dev_Empls where Emp_no=$1`
	getemployees := []GetEmployee{}
	err := DB().Select(&getemployees, sql, copyfromemployee)
	if err != nil {
		log.Println(err)
	}

	// json, err := json.Marshal(getemployees)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Fprintf(w, string(json))

	page := page{FieldName: "Get Employees", GetEmployees: getemployees}
	templates.Lookup("Body").Execute(w, page)
}

//PostEmployees : PostEmployees is a function for posting employee data.  The url accepts serveral arguments for employee demographic information.
func PostEmployees(w http.ResponseWriter, r *http.Request) {

	// unitid := r.URL.Query().Get("unitid")
	// identifier := r.URL.Query().Get("identifier")

	empl_id := r.URL.Query().Get("empl_id")
	store_no := r.URL.Query().Get("store_no")
	emp_no := r.URL.Query().Get("emp_no")

	//sql := `Exec post_employee_attributes $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43`
	sql := `Exec post_employee_attributes $1, $2, $3`
	postemployees := []PostEmployee{}

	//fmt.Println(employee_from, employee, org, inventory_group, company, market, dm, street, city, state, zip, employee_name, phone, concept, consent, region, fbc, coop, dc, plu_category, labor_guide, import_group, pos, error_guide, computer, account, vendor, facility, quickcook, new_construct, reader, longitude, latitude, golive, fiscal_seed, tax, state_tax, prep_guide, groups_plu, batch_meat, active, upload, refill)

	//err := DB().Select(&postemployees, sql, employee_from, employee, org, inventory_group, company, market, dm, street, city, state, zip, employee_name, phone, concept, consent, region, fbc, coop, dc, plu_category, labor_guide, import_group, pos, error_guide, computer, account, vendor, facility, quickcook, new_construct, reader, longitude, latitude, golive, fiscal_seed, tax, state_tax, prep_guide, groups_plu, batch_meat, active, upload, refill)
	err := DB().Select(&postemployees, sql, empl_id, store_no, emp_no)
	if err != nil {
		log.Println(err)
	}

	json, err := json.Marshal(postemployees)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(json))
}

//DB : DB is a function that connects to SQL server.
func DB() *sqlx.DB {
	serv := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	database := os.Getenv("DB_DATABASE")

	db, err := sqlx.Connect("mssql", fmt.Sprintf(`server=%s;user id=%s;password=%s;database=%s;log64;encrypt=disable`, serv, user, pass, database))

	if err != nil {
		log.Println(err)
	}
	return db
}

const doc = `
 <!DOCTYPE html>
 <html>
	 <head>
		 <title>{{.FieldName}}</title>
		 <meta charset="utf-8">
		 <meta name="viewport" content="width=device-width, initial-scale=1">
		 <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
		 <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		 <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
	 </head>
     <body>
		 <h3>Employee Edits</h3>
		 	 {{template "List" .GetEmployees}}
     </body>
 </html>
 `
const docList = `
<div>
  <form action="/getemployees">
    <label for="copyfromemployee">Get Employee</label>
	<input type="text" maxlength="5" size="4" id="copyfromemployee" name="copyfromemployee" placeholder="get employee..">
    <input type="submit" value="Get">
  </form>
</div>
<div>

<form action="/postemployees"> 
	<!--Post Changes <input type="submit" value="Post">-->
	<ul >
     {{range .}}
		<ul>Employee (copy from): <input type="text" maxlength="5" size="4" name="emp_no" value={{ .Emp_no}}></input></ul>
		<br>
		<br>
		Post Changes <input type="submit" value="Post">
		<ul>Employee Number (post to): <input type="text" maxlength="5" size="4" name="employee" value={{ .Emp_no}}></input></ul>
		<ul>Employee ID: <input type="text" maxlength="20" size="15" name="Empl_ID" value={{ .Empl_ID}}></input></ul>
		<ul>Store: <input type="text" maxlength="20" size="15" name="Store_no" value={{ .Store_no}}></input></ul>
     {{end}}
	 </ul>
</form>
</div>
 `

/*
http://localhost:5000/getemployees?copyfromemployee=1005
http://localhost:5000/postemployees?employee=1005&org=TJCorp&inventory_group=TJCorpQSPOS2&company=TJCorp&market=TJCorp&dm=Kevin+Kreutzer&street=101+South+Greeley+Highway&city=Cheyenne&state=WY&zip=82007&employee_name=S+Greeley&phone=3076341555&concept=TacoJohn&consent=1&region=Corp&fbc=FBC+B+Davis&coop=CoOp+Cheyenne&dc=FSA+Loveland&plu_category=REGLR&labor_guide=TJC&import_group=TJSynchAll&pos=QSP11&error_guide=2&computer=TjCorp&account=TJCorp&vendor=TJCor&facility=&quickcook=Y&new_construct=&reader=0&longitude=-104.8031900&latitude=41.1163400&golive=2012-04-24&fiscal_seed=2012-12-31&tax=0.06&state_tax=0.06000
*/

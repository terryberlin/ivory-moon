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
		EmployeeID            *string `db:"Empl_ID" json:"employeeID"`
		Unit                  *string `db:"Store_no" json:"unit"`
		EmployeeNumber        *string `db:"Emp_no" json:"employeeNumber"`
		SSN                   *string `db:"SSN" json:"ssn"`
		LastName              *string `db:"Last_Name" json:"lastName"`
		FirstName             *string `db:"First_Name" json:"firstName"`
		Address               *string `db:"Address" json:"address"`
		City                  *string `db:"City" json:"city"`
		State                 *string `db:"State" json:"state"`
		Zip                   *string `db:"Zip" json:"zip"`
		Phone                 *string `db:"Phone" json:"phone"`
		BirthDate             *string `db:"Birth_Dt" json:"birthDate"`
		HireDate              *string `db:"Hire_Date" json:"hireDate"`
		PreviousSSN           *string `db:"Prev_SSN" json:"previousSsn"`
		W4Status              *string `db:"W4Status" json:"w4Status"`
		Allowances            *string `db:"Allowances" json:"allowances"`
		EffectDate            *string `db:"Effect_Dt" json:"effectDate"`
		CurrentRate           *string `db:"Curr_Rate" json:"currentRate"`
		PreviousRate          *string `db:"Prev_Rate" json:"previousRate"`
		Jobcode               *string `db:"Jobcode" json:"jobCode"`
		CrewRate              *string `db:"CrewRate" json:"crewRate"`
		OtherRate             *string `db:"OtherRate" json:"otherRate"`
		PermanentSupervisor   *string `db:"PermSupervisor" json:"permanentSupervisor"`
		SalariedManager       *string `db:"SalariedMgr" json:"salariedManager"`
		Terminated            *string `db:"Term" json:"terminated"`
		Fired                 *string `db:"Fired" json:"fired"`
		ReHire                *string `db:"ReHire" json:"rehire"`
		TerminationReason     *string `db:"Term_Reason" json:"terminationReason"`
		TerminatedDate        *string `db:"DateTermed" json:"terminatedDate"`
		Jobcode2              *string `db:"Jobcode2" json:"jobCode2"`
		PayRate2              *string `db:"PayRate2" json:"payRate2"`
		EffectiveDate2        *string `db:"Effec_Dt2" json:"effecDate2"`
		OldRate2              *string `db:"OldRate2" json:"oldRate2"`
		Jobcode3              *string `db:"Jobcode3" json:"jobCode3"`
		PayRate3              *string `db:"PayRate3" json:"payRate3"`
		EffectiveDate3        *string `db:"Effec_Dt3" json:"effecDate3"`
		OldRate3              *string `db:"OldRate3" json:"oldRate3"`
		Jobcode4              *string `db:"Jobcode4" json:"jobCode4"`
		PayRate4              *string `db:"PayRate4" json:"payRate4"`
		EffectiveDate4        *string `db:"Effec_Dt4" json:"effecDate4"`
		OldRate4              *string `db:"OldRate4" json:"oldrate4"`
		Jobcode5              *string `db:"Jobcode5" json:"jobcode5"`
		PayRate5              *string `db:"PayRate5" json:"payrate5"`
		EffectiveDate5        *string `db:"Effec_Dt5" json:"effecDate5"`
		OldRate5              *string `db:"OldRate5" json:"oldrate5"`
		Overtime1             *string `db:"Overtime1" json:"overtime1"`
		Overtime2             *string `db:"Overtime2" json:"overtime2"`
		Overtime3             *string `db:"Overtime3" json:"overtime3"`
		Overtime4             *string `db:"Overtime4" json:"overtime4"`
		Overtime5             *string `db:"Overtime5" json:"overtime5"`
		Gender                *string `db:"Gender" json:"gender"`
		PayrollID             *string `db:"PayrollID" json:"payrollid"`
		StateAllowances       *string `db:"StateAllowances" json:"stateAllowances"`
		TransferToStore       *string `db:"TransferToStore" json:"transferTostore"`
		MiddleInitial         *string `db:"MiddleInitial" json:"middleInitial"`
		AdditionalWithholding *string `db:"AddlWithholding" json:"additionalWithholding"`
		TaxExempt             *string `db:"TaxExempt" json:"taxExempt"`
		MaidenName            *string `db:"MaidenName" json:"maidenName"`
		ImportDate            *string `db:"Import_Date" json:"importDate"`
		Race                  *string `db:"Race" json:"race"`
		BorrowedEmployee      *string `db:"BorrowedEmpl" json:"borrowedEmployee"`
		HomeStore             *string `db:"HomeStore" json:"homeStore"`
		I9Code                *string `db:"I9Code" json:"i9Code"`
		InfoSent              *string `db:"Info_Sent" json:"infoSent"`
		DateSent              *string `db:"Date_Sent" json:"dateSent"`
		BonusEligible         *string `db:"BonusEligible" json:"bonusEligible"`
		BonusRate             *string `db:"BonusRate" json:"bonusRate"`
		SalaryClockIn         *string `db:"SalaryClockIn" json:"salaryClockin"`
		InSchool              *string `db:"InSchool" json:"inSchool"`
		VacationSchedule      *string `db:"VacationSchedule" json:"vacationSchedule"`
		HRDocuments           *string `db:"HRDocuments" json:"hrDocuments"`
		ApartmentNumber       *string `db:"AptNumber" json:"apartmentNumber"`
		DepositType           *string `db:"DepositType" json:"depositType"`
		RoutingNumber         *string `db:"RoutingNumber" json:"routingNumber"`
		AccountType           *string `db:"AccountType" json:"accountType"`
		AccountNumber         *string `db:"AccountNumber" json:"accountNumber"`
		EmailAddress          *string `db:"EmailAddress" json:"emailAddress"`
		ManagerNote           *string `db:"ManagerNote" json:"managerNote"`
		InactiveEmployee      *string `db:"InactiveEmpl" json:"inactiveEmployee"`
		IDNumber              *string `db:"IDNumber" json:"idNumber"`
		PreviousIDNumber      *string `db:"PrevIDNumber" json:"previousIdnumber"`
	}
)

//PostEmployee : PostEmployee is a structure for posting employee data.
type (
	PostEmployee struct {
		EmployeeID            *string `db:"Empl_ID" json:"employeeID"`
		Unit                  *string `db:"Store_no" json:"unit"`
		EmployeeNumber        *string `db:"Emp_no" json:"employeeNumber"`
		SSN                   *string `db:"SSN" json:"ssn"`
		LastName              *string `db:"Last_Name" json:"lastName"`
		FirstName             *string `db:"First_Name" json:"firstName"`
		Address               *string `db:"Address" json:"address"`
		City                  *string `db:"City" json:"city"`
		State                 *string `db:"State" json:"state"`
		Zip                   *string `db:"Zip" json:"zip"`
		Phone                 *string `db:"Phone" json:"phone"`
		BirthDate             *string `db:"Birth_Dt" json:"birthDate"`
		HireDate              *string `db:"Hire_Date" json:"hireDate"`
		PreviousSSN           *string `db:"Prev_SSN" json:"previousSsn"`
		W4Status              *string `db:"W4Status" json:"w4Status"`
		Allowances            *string `db:"Allowances" json:"allowances"`
		EffectDate            *string `db:"Effect_Dt" json:"effectDate"`
		CurrentRate           *string `db:"Curr_Rate" json:"currentRate"`
		PreviousRate          *string `db:"Prev_Rate" json:"previousRate"`
		JobCode               *string `db:"Jobcode" json:"jobCode"`
		CrewRate              *string `db:"CrewRate" json:"crewRate"`
		OtherRate             *string `db:"OtherRate" json:"otherRate"`
		PermanentSupervisor   *string `db:"PermSupervisor" json:"permanentSupervisor"`
		SalariedManager       *string `db:"SalariedMgr" json:"salariedManager"`
		Terminated            *string `db:"Term" json:"terminated"`
		Fired                 *string `db:"Fired" json:"fired"`
		ReHire                *string `db:"ReHire" json:"rehire"`
		TerminationReason     *string `db:"Term_Reason" json:"terminationReason"`
		TerminatedDate        *string `db:"DateTermed" json:"terminatedDate"`
		Jobcode2              *string `db:"Jobcode2" json:"jobCode2"`
		PayRate2              *string `db:"PayRate2" json:"payRate2"`
		EffectiveDate2        *string `db:"Effec_Dt2" json:"effecDate2"`
		OldRate2              *string `db:"OldRate2" json:"oldRate2"`
		Jobcode3              *string `db:"Jobcode3" json:"jobCode3"`
		PayRate3              *string `db:"PayRate3" json:"payRate3"`
		EffectiveDate3        *string `db:"Effec_Dt3" json:"effecDate3"`
		OldRate3              *string `db:"OldRate3" json:"oldRate3"`
		Jobcode4              *string `db:"Jobcode4" json:"jobCode4"`
		PayRate4              *string `db:"PayRate4" json:"payRate4"`
		EffectiveDate4        *string `db:"Effec_Dt4" json:"effecDate4"`
		OldRate4              *string `db:"OldRate4" json:"oldrate4"`
		Jobcode5              *string `db:"Jobcode5" json:"jobcode5"`
		PayRate5              *string `db:"PayRate5" json:"payrate5"`
		EffectiveDate5        *string `db:"Effec_Dt5" json:"effecDate5"`
		OldRate5              *string `db:"OldRate5" json:"oldrate5"`
		Overtime1             *string `db:"Overtime1" json:"overtime1"`
		Overtime2             *string `db:"Overtime2" json:"overtime2"`
		Overtime3             *string `db:"Overtime3" json:"overtime3"`
		Overtime4             *string `db:"Overtime4" json:"overtime4"`
		Overtime5             *string `db:"Overtime5" json:"overtime5"`
		Gender                *string `db:"Gender" json:"gender"`
		PayrollID             *string `db:"PayrollID" json:"payrollid"`
		StateAllowances       *string `db:"StateAllowances" json:"stateAllowances"`
		TransferToStore       *string `db:"TransferToStore" json:"transferTostore"`
		MiddleInitial         *string `db:"MiddleInitial" json:"middleInitial"`
		AdditionalWithholding *string `db:"AddlWithholding" json:"additionalWithholding"`
		TaxExempt             *string `db:"TaxExempt" json:"taxExempt"`
		MaidenName            *string `db:"MaidenName" json:"maidenName"`
		ImportDate            *string `db:"Import_Date" json:"importDate"`
		Race                  *string `db:"Race" json:"race"`
		BorrowedEmployee      *string `db:"BorrowedEmpl" json:"borrowedEmployee"`
		HomeStore             *string `db:"HomeStore" json:"homeStore"`
		I9Code                *string `db:"I9Code" json:"i9Code"`
		InfoSent              *string `db:"Info_Sent" json:"infoSent"`
		DateSent              *string `db:"Date_Sent" json:"dateSent"`
		BonusEligible         *string `db:"BonusEligible" json:"bonusEligible"`
		BonusRate             *string `db:"BonusRate" json:"bonusRate"`
		SalaryClockIn         *string `db:"SalaryClockIn" json:"salaryClockin"`
		InSchool              *string `db:"InSchool" json:"inSchool"`
		VacationSchedule      *string `db:"VacationSchedule" json:"vacationSchedule"`
		HRDocuments           *string `db:"HRDocuments" json:"hrDocuments"`
		ApartmentNumber       *string `db:"AptNumber" json:"apartmentNumber"`
		DepositType           *string `db:"DepositType" json:"depositType"`
		RoutingNumber         *string `db:"RoutingNumber" json:"routingNumber"`
		AccountType           *string `db:"AccountType" json:"accountType"`
		AccountNumber         *string `db:"AccountNumber" json:"accountNumber"`
		EmailAddress          *string `db:"EmailAddress" json:"emailAddress"`
		ManagerNote           *string `db:"ManagerNote" json:"managerNote"`
		InactiveEmployee      *string `db:"InactiveEmpl" json:"inactiveEmployee"`
		IDNumber              *string `db:"IDNumber" json:"idNumber"`
		PreviousIDNumber      *string `db:"PrevIDNumber" json:"previousIdnumber"`
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
	sql := `select 
		  e.[Empl_ID]
		, e.[Store_no]
		, e.[Emp_no]
		, e.[SSN]
		, e.[Last_Name]
		, e.[First_Name]
		, e.[Address]
		, e.[City]
		, e.[State]
		, e.[Zip]
		, e.[Phone]
		, e.[Birth_Dt]
		, e.[Hire_Date]
		, e.[Prev_SSN]
		, e.[W4Status]
		, e.[Allowances]
		, e.[Effect_Dt]
		, e.[Curr_Rate]
		, e.[Prev_Rate]
		, e.[Jobcode]
		, e.[CrewRate]
		, e.[OtherRate]
		, e.[PermSupervisor]
		, e.[SalariedMgr]
		, e.[Term]
		, e.[Fired]
		, e.[ReHire]
		, e.[Term_Reason]
		, e.[DateTermed]
		, e.[Jobcode2]
		, e.[PayRate2]
		, e.[Effec_Dt2]
		, e.[OldRate2]
		, e.[Jobcode3]
		, e.[PayRate3]
		, e.[Effec_Dt3]
		, e.[OldRate3]
		, e.[Jobcode4]
		, e.[PayRate4]
		, e.[Effec_Dt4]
		, e.[OldRate4]
		, e.[Jobcode5]
		, e.[PayRate5]
		, e.[Effec_Dt5]
		, e.[OldRate5]
		, e.[Overtime1]
		, e.[Overtime2]
		, e.[Overtime3]
		, e.[Overtime4]
		, e.[Overtime5]
		, e.[Gender]
		, e.[PayrollID]
		, e.[StateAllowances]
		, e.[TransferToStore]
		, e.[MiddleInitial]
		, e.[AddlWithholding]
		, e.[TaxExempt]
		, e.[MaidenName]
		, e.[Import_Date]
		, e.[Race]
		, e.[BorrowedEmpl]
		, e.[HomeStore]
		, e.[I9Code]
		, e.[Info_Sent]
		, e.[Date_Sent]
		, e.[BonusEligible]
		, e.[BonusRate]
		, e.[SalaryClockIn]
		, e.[InSchool]
		, e.[VacationSchedule]
		, e.[HRDocuments]
		, e.[AptNumber]
		, e.[DepositType]
		, e.[RoutingNumber]
		, e.[AccountType]
		, e.[AccountNumber]
		, e.[EmailAddress]
		, e.[ManagerNote]
		, e.[InactiveEmpl]
		, e.[IDNumber]
		, e.[PrevIDNumber] 
		from dev_Empls e where Emp_no=$1`
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

	employeeID := r.URL.Query().Get("employeeID")
	unit := r.URL.Query().Get("unit")
	employeeNumber := r.URL.Query().Get("employeeNumber")

	//sql := `Exec post_employee_attributes $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43`
	sql := `Exec post_employee_attributes $1, $2, $3`
	postemployees := []PostEmployee{}

	//fmt.Println(employee_from, employee, org, inventory_group, company, market, dm, street, city, state, zip, employee_name, phone, concept, consent, region, fbc, coop, dc, plu_category, labor_guide, import_group, pos, error_guide, computer, account, vendor, facility, quickcook, new_construct, reader, longitude, latitude, golive, fiscal_seed, tax, state_tax, prep_guide, groups_plu, batch_meat, active, upload, refill)

	//err := DB().Select(&postemployees, sql, employee_from, employee, org, inventory_group, company, market, dm, street, city, state, zip, employee_name, phone, concept, consent, region, fbc, coop, dc, plu_category, labor_guide, import_group, pos, error_guide, computer, account, vendor, facility, quickcook, new_construct, reader, longitude, latitude, golive, fiscal_seed, tax, state_tax, prep_guide, groups_plu, batch_meat, active, upload, refill)
	err := DB().Select(&postemployees, sql, employeeID, unit, employeeNumber)
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
		<!--ul>Employee (copy from): <input type="text" maxlength="5" size="4" name="emp_no" value={{ .EmployeeNumber }}></input></ul-->
		<br>
		<br>
		Post Changes <input type="submit" value="Post">
		<ul>Employee Number: <input type="text" maxlength="5" size="4" name="EmployeeNumber" value={{ .EmployeeNumber }}></input></ul>
		<ul>Employee ID: <input type="text" maxlength="20" size="15" name="EmployeeID" value={{ .EmployeeID }}></input></ul>
		<ul>Unit: <input type="text" maxlength="20" size="15" name="Unit" value={{ .Unit }}></input></ul>
		<ul>SSN: <input type="password" maxlength="20" size="15" name="SSN" value={{ .SSN }}></input></ul>
		<ul>First Name: <input type="text" maxlength="20" size="15" name="FirstName" value={{ .FirstName }}></input></ul>
		<ul>Last Name: <input type="text" maxlength="20" size="15" name="LastName" value={{ .LastName }}></input></ul>
		<ul>Address: <input type="text" maxlength="20" size="15" name="Address" value={{ .Address }}></input></ul>
		<ul>City: <input type="text" maxlength="20" size="15" name="City" value={{ .City }}></input></ul>
		<ul>State: <input type="text" maxlength="20" size="15" name="State" value={{ .State }}></input></ul>
		<ul>Zip: <input type="text" maxlength="20" size="15" name="Zip" value={{ .Zip }}></input></ul>
		<ul>Phone: <input type="text" maxlength="20" size="15" name="Phone" value={{ .Phone }}></input></ul>
		<ul>BirthDate: <input type="text" maxlength="20" size="15" name="BirthDate" value={{ .BirthDate }}></input></ul>
		<ul>HireDate: <input type="text" maxlength="20" size="15" name="HireDate" value={{ .HireDate }}></input></ul>
		<ul>PreviousSSN: <input type="text" maxlength="20" size="15" name="PreviousSSN" value={{ .PreviousSSN }}></input></ul>
		<ul>W4Status: <input type="text" maxlength="20" size="15" name="W4Status" value={{ .W4Status }}></input></ul>
		<ul>Allowances: <input type="text" maxlength="20" size="15" name="Allowances" value={{ .Allowances }}></input></ul>
		<ul>EffectDate: <input type="text" maxlength="20" size="15" name="EffectDate" value={{ .EffectDate }}></input></ul>
		<ul>CurrentRate: <input type="text" maxlength="20" size="15" name="CurrentRate" value={{ .CurrentRate }}></input></ul>
		<ul>PreviousRate: <input type="text" maxlength="20" size="15" name="PreviousRate" value={{ .PreviousRate }}></input></ul>
		<ul>Jobcode: <input type="text" maxlength="20" size="15" name="Jobcode" value={{ .Jobcode }}></input></ul>
		<ul>CrewRate: <input type="text" maxlength="20" size="15" name="CrewRate" value={{ .CrewRate }}></input></ul>
		<ul>OtherRate: <input type="text" maxlength="20" size="15" name="OtherRate" value={{ .OtherRate }}></input></ul>
		<ul>PermanentSupervisor: <input type="text" maxlength="20" size="15" name="PermanentSupervisor" value={{ .PermanentSupervisor }}></input></ul>
		<ul>SalariedManager: <input type="text" maxlength="20" size="15" name="SalariedManager" value={{ .SalariedManager }}></input></ul>
		<ul>Terminated: <input type="text" maxlength="20" size="15" name="Terminated" value={{ .Terminated }}></input></ul>
		<ul>Fired: <input type="text" maxlength="20" size="15" name="Fired" value={{ .Fired }}></input></ul>
		<ul>ReHire: <input type="text" maxlength="20" size="15" name="ReHire" value={{ .ReHire }}></input></ul>
		<ul>TerminationReason: <input type="text" maxlength="20" size="15" name="TerminationReason" value={{ .TerminationReason }}></input></ul>
		<ul>TerminatedDate: <input type="text" maxlength="20" size="15" name="TerminatedDate" value={{ .TerminatedDate }}></input></ul>
		<ul>Jobcode2: <input type="text" maxlength="20" size="15" name="Jobcode2" value={{ .Jobcode2 }}></input></ul>
		<ul>PayRate2: <input type="text" maxlength="20" size="15" name="PayRate2" value={{ .PayRate2 }}></input></ul>
		<ul>EffectiveDate2: <input type="text" maxlength="20" size="15" name="EffectiveDate2" value={{ .EffectiveDate2 }}></input></ul>
		<ul>OldRate2: <input type="text" maxlength="20" size="15" name="OldRate2" value={{ .OldRate2 }}></input></ul>
		<ul>Jobcode3: <input type="text" maxlength="20" size="15" name="Jobcode3" value={{ .Jobcode3 }}></input></ul>
		<ul>PayRate3: <input type="text" maxlength="20" size="15" name="PayRate3" value={{ .PayRate3 }}></input></ul>
		<ul>EffectiveDate3: <input type="text" maxlength="20" size="15" name="EffectiveDate3" value={{ .EffectiveDate3 }}></input></ul>
		<ul>OldRate3: <input type="text" maxlength="20" size="15" name="OldRate3" value={{ .OldRate3 }}></input></ul>
		<ul>Jobcode4: <input type="text" maxlength="20" size="15" name="Jobcode4" value={{ .Jobcode4 }}></input></ul>
		<ul>PayRate4: <input type="text" maxlength="20" size="15" name="PayRate4" value={{ .PayRate4 }}></input></ul>
		<ul>EffectiveDate4: <input type="text" maxlength="20" size="15" name="EffectiveDate4" value={{ .EffectiveDate4 }}></input></ul>
		<ul>OldRate4: <input type="text" maxlength="20" size="15" name="OldRate4" value={{ .OldRate4 }}></input></ul>
		<ul>Jobcode5: <input type="text" maxlength="20" size="15" name="Jobcode5" value={{ .Jobcode5 }}></input></ul>
		<ul>PayRate5: <input type="text" maxlength="20" size="15" name="PayRate5" value={{ .PayRate5 }}></input></ul>
		<ul>EffectiveDate5: <input type="text" maxlength="20" size="15" name="EffectiveDate5" value={{ .EffectiveDate5 }}></input></ul>
		<ul>OldRate5: <input type="text" maxlength="20" size="15" name="OldRate5" value={{ .OldRate5 }}></input></ul>
		<ul>Overtime1: <input type="text" maxlength="20" size="15" name="Overtime1" value={{ .Overtime1 }}></input></ul>
		<ul>Overtime2: <input type="text" maxlength="20" size="15" name="Overtime2" value={{ .Overtime2 }}></input></ul>
		<ul>Overtime3: <input type="text" maxlength="20" size="15" name="Overtime3" value={{ .Overtime3 }}></input></ul>
		<ul>Overtime4: <input type="text" maxlength="20" size="15" name="Overtime4" value={{ .Overtime4 }}></input></ul>
		<ul>Overtime5: <input type="text" maxlength="20" size="15" name="Overtime5" value={{ .Overtime5 }}></input></ul>
		<ul>Gender: <input type="text" maxlength="20" size="15" name="Gender" value={{ .Gender }}></input></ul>
		<ul>PayrollID: <input type="text" maxlength="20" size="15" name="PayrollID" value={{ .PayrollID }}></input></ul>
		<ul>StateAllowances: <input type="text" maxlength="20" size="15" name="StateAllowances" value={{ .StateAllowances }}></input></ul>
		<ul>TransferToStore: <input type="text" maxlength="20" size="15" name="TransferToStore" value={{ .TransferToStore }}></input></ul>
		<ul>MiddleInitial: <input type="text" maxlength="20" size="15" name="MiddleInitial" value={{ .MiddleInitial }}></input></ul>
		<ul>AdditionalWithholding: <input type="text" maxlength="20" size="15" name="AdditionalWithholding" value={{ .AdditionalWithholding }}></input></ul>
		<ul>TaxExempt: <input type="text" maxlength="20" size="15" name="TaxExempt" value={{ .TaxExempt }}></input></ul>
		<ul>MaidenName: <input type="text" maxlength="20" size="15" name="MaidenName" value={{ .MaidenName }}></input></ul>
		<ul>ImportDate: <input type="text" maxlength="20" size="15" name="ImportDate" value={{ .ImportDate }}></input></ul>
		<ul>Race: <input type="text" maxlength="20" size="15" name="Race" value={{ .Race }}></input></ul>
		<ul>BorrowedEmployee: <input type="text" maxlength="20" size="15" name="BorrowedEmployee" value={{ .BorrowedEmployee }}></input></ul>
		<ul>HomeStore: <input type="text" maxlength="20" size="15" name="HomeStore" value={{ .HomeStore }}></input></ul>
		<ul>I9Code: <input type="text" maxlength="20" size="15" name="I9Code" value={{ .I9Code }}></input></ul>
		<ul>InfoSent: <input type="text" maxlength="20" size="15" name="InfoSent" value={{ .InfoSent }}></input></ul>
		<ul>DateSent: <input type="text" maxlength="20" size="15" name="DateSent" value={{ .DateSent }}></input></ul>
		<ul>BonusEligible: <input type="text" maxlength="20" size="15" name="BonusEligible" value={{ .BonusEligible }}></input></ul>
		<ul>BonusRate: <input type="text" maxlength="20" size="15" name="BonusRate" value={{ .BonusRate }}></input></ul>
		<ul>SalaryClockIn: <input type="text" maxlength="20" size="15" name="SalaryClockIn" value={{ .SalaryClockIn }}></input></ul>
		<ul>InSchool: <input type="text" maxlength="20" size="15" name="InSchool" value={{ .InSchool }}></input></ul>
		<ul>VacationSchedule: <input type="text" maxlength="20" size="15" name="VacationSchedule" value={{ .VacationSchedule }}></input></ul>
		<ul>HRDocuments: <input type="text" maxlength="20" size="15" name="HRDocuments" value={{ .HRDocuments }}></input></ul>
		<ul>ApartmentNumber: <input type="text" maxlength="20" size="15" name="ApartmentNumber" value={{ .ApartmentNumber }}></input></ul>
		<ul>DepositType: <input type="text" maxlength="20" size="15" name="DepositType" value={{ .DepositType }}></input></ul>
		<ul>RoutingNumber: <input type="text" maxlength="20" size="15" name="RoutingNumber" value={{ .RoutingNumber }}></input></ul>
		<ul>AccountType: <input type="text" maxlength="20" size="15" name="AccountType" value={{ .AccountType }}></input></ul>
		<ul>AccountNumber: <input type="text" maxlength="20" size="15" name="AccountNumber" value={{ .AccountNumber }}></input></ul>
		<ul>EmailAddress: <input type="text" maxlength="20" size="15" name="EmailAddress" value={{ .EmailAddress }}></input></ul>
		<ul>ManagerNote: <input type="text" maxlength="20" size="15" name="ManagerNote" value={{ .ManagerNote }}></input></ul>
		<ul>InactiveEmployee: <input type="text" maxlength="20" size="15" name="InactiveEmployee" value={{ .InactiveEmployee }}></input></ul>
		<ul>IDNumber: <input type="text" maxlength="20" size="15" name="IDNumber" value={{ .IDNumber }}></input></ul>
		<ul>PreviousIDNumber: <input type="text" maxlength="20" size="15" name="PreviousIDNumber" value={{ .PreviousIDNumber }}></input></ul>		
     {{end}}
	 </ul>
</form>
</div>
 `

/*
http://localhost:5000/getemployees?copyfromemployee=1005
http://localhost:5000/postemployees?employee=1005&org=TJCorp&inventory_group=TJCorpQSPOS2&company=TJCorp&market=TJCorp&dm=Kevin+Kreutzer&street=101+South+Greeley+Highway&city=Cheyenne&state=WY&zip=82007&employee_name=S+Greeley&phone=3076341555&concept=TacoJohn&consent=1&region=Corp&fbc=FBC+B+Davis&coop=CoOp+Cheyenne&dc=FSA+Loveland&plu_category=REGLR&labor_guide=TJC&import_group=TJSynchAll&pos=QSP11&error_guide=2&computer=TjCorp&account=TJCorp&vendor=TJCor&facility=&quickcook=Y&new_construct=&reader=0&longitude=-104.8031900&latitude=41.1163400&golive=2012-04-24&fiscal_seed=2012-12-31&tax=0.06&state_tax=0.06000
*/

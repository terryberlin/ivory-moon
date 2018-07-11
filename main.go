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
		EmployeeID            *string `db:"employeeID" json:"employeeID"`
		Unit                  *string `db:"unit" json:"unit"`
		EmployeeNumber        *string `db:"employeeNumber" json:"employeeNumber"`
		SSN                   *string `db:"ssn" json:"ssn"`
		LastName              *string `db:"lastName" json:"lastName"`
		FirstName             *string `db:"firstName" json:"firstName"`
		Address               *string `db:"address" json:"address"`
		City                  *string `db:"city" json:"city"`
		State                 *string `db:"state" json:"state"`
		Zip                   *string `db:"zip" json:"zip"`
		Phone                 *string `db:"phone" json:"phone"`
		BirthDate             *string `db:"birthDate" json:"birthDate"`
		HireDate              *string `db:"hireDate" json:"hireDate"`
		PreviousSSN           *string `db:"previousSsn" json:"previousSsn"`
		W4Status              *string `db:"w4Status" json:"w4Status"`
		Allowances            *string `db:"allowances" json:"allowances"`
		EffectDate            *string `db:"effectiveDate" json:"effectiveDate"`
		CurrentRate           *string `db:"currentRate" json:"currentRate"`
		PreviousRate          *string `db:"previousRate" json:"previousRate"`
		JobCode               *string `db:"jobCode" json:"jobCode"`
		CrewRate              *string `db:"crewRate" json:"crewRate"`
		OtherRate             *string `db:"otherRate" json:"otherRate"`
		PermanentSupervisor   *string `db:"permanentSupervisor" json:"permanentSupervisor"`
		SalariedManager       *string `db:"salariedManager" json:"salariedManager"`
		Terminated            *string `db:"terminated" json:"terminated"`
		Fired                 *string `db:"fired" json:"fired"`
		ReHire                *string `db:"rehire" json:"rehire"`
		TerminationReason     *string `db:"terminationReason" json:"terminationReason"`
		TerminatedDate        *string `db:"terminatedDate" json:"terminatedDate"`
		Jobcode2              *string `db:"jobCode2" json:"jobCode2"`
		PayRate2              *string `db:"payRate2" json:"payRate2"`
		EffectiveDate2        *string `db:"effectiveDate2" json:"effectiveDate2"`
		OldRate2              *string `db:"oldRate2" json:"oldRate2"`
		Jobcode3              *string `db:"jobCode3" json:"jobCode3"`
		PayRate3              *string `db:"payRate3" json:"payRate3"`
		EffectiveDate3        *string `db:"effectiveDate3" json:"effectiveDate3"`
		OldRate3              *string `db:"oldRate3" json:"oldRate3"`
		Jobcode4              *string `db:"jobCode4" json:"jobCode4"`
		PayRate4              *string `db:"payRate4" json:"payRate4"`
		EffectiveDate4        *string `db:"effectiveDate4" json:"effectiveDate4"`
		OldRate4              *string `db:"oldRate4" json:"oldRate4"`
		Jobcode5              *string `db:"jobCode5" json:"jobCode5"`
		PayRate5              *string `db:"payRate5" json:"payRate5"`
		EffectiveDate5        *string `db:"effectiveDate5" json:"effectiveDate5"`
		OldRate5              *string `db:"oldRate5" json:"oldRate5"`
		Overtime1             *string `db:"overtime1" json:"overtime1"`
		Overtime2             *string `db:"overtime2" json:"overtime2"`
		Overtime3             *string `db:"overtime3" json:"overtime3"`
		Overtime4             *string `db:"overtime4" json:"overtime4"`
		Overtime5             *string `db:"overtime5" json:"overtime5"`
		Gender                *string `db:"gender" json:"gender"`
		PayrollID             *string `db:"payrollId" json:"payrollId"`
		StateAllowances       *string `db:"stateAllowances" json:"stateAllowances"`
		TransferToStore       *string `db:"transferTostore" json:"transferTostore"`
		MiddleInitial         *string `db:"middleInitial" json:"middleInitial"`
		AdditionalWithholding *string `db:"additionalWithholding" json:"additionalWithholding"`
		TaxExempt             *string `db:"taxExempt" json:"taxExempt"`
		MaidenName            *string `db:"maidenName" json:"maidenName"`
		ImportDate            *string `db:"importDate" json:"importDate"`
		Race                  *string `db:"race" json:"race"`
		BorrowedEmployee      *string `db:"borrowedEmployee" json:"borrowedEmployee"`
		HomeStore             *string `db:"homeStore" json:"homeStore"`
		I9Code                *string `db:"i9Code" json:"i9Code"`
		InfoSent              *string `db:"infoSent" json:"infoSent"`
		DateSent              *string `db:"dateSent" json:"dateSent"`
		BonusEligible         *string `db:"bonusEligible" json:"bonusEligible"`
		BonusRate             *string `db:"bonusRate" json:"bonusRate"`
		SalaryClockIn         *string `db:"salaryClockin" json:"salaryClockin"`
		InSchool              *string `db:"inSchool" json:"inSchool"`
		VacationSchedule      *string `db:"vacationSchedule" json:"vacationSchedule"`
		HRDocuments           *string `db:"hrDocuments" json:"hrDocuments"`
		ApartmentNumber       *string `db:"apartmentNumber" json:"apartmentNumber"`
		DepositType           *string `db:"depositType" json:"depositType"`
		RoutingNumber         *string `db:"routingNumber" json:"routingNumber"`
		AccountType           *string `db:"accountType" json:"accountType"`
		AccountNumber         *string `db:"accountNumber" json:"accountNumber"`
		EmailAddress          *string `db:"emailAddress" json:"emailAddress"`
		ManagerNote           *string `db:"managerNote" json:"managerNote"`
		InactiveEmployee      *string `db:"inactiveEmployee" json:"inactiveEmployee"`
		IDNumber              *string `db:"idNumber" json:"idNumber"`
		PreviousIDNumber      *string `db:"previousIdnumber" json:"previousIdnumber"`
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
			rtrim(e.[Empl_ID]) as Empl_ID
			, rtrim(e.[Store_no]) as Store_no
			, rtrim(e.[Emp_no]) as Emp_no
			, rtrim(e.[SSN]) as SSN
			, rtrim(e.[Last_Name]) as Last_Name
			, rtrim(e.[First_Name]) as First_Name
			, rtrim(e.[Address]) as Address
			, rtrim(e.[City]) as City
			, rtrim(e.[State]) as State
			, rtrim(e.[Zip]) as Zip
			, rtrim(e.[Phone]) as Phone
			, e.[Birth_Dt]
			, e.[Hire_Date]
			, rtrim(e.[Prev_SSN]) as Prev_SSN
			, rtrim(e.[W4Status]) as W4Status
			, rtrim(e.[Allowances]) as Allowances
			, e.[Effect_Dt]
			, e.[Curr_Rate]
			, e.[Prev_Rate]
			, rtrim(e.[Jobcode]) as Jobcode
			, e.[CrewRate]
			, e.[OtherRate]
			, e.[PermSupervisor]
			, e.[SalariedMgr]
			, e.[Term]
			, e.[Fired]
			, e.[ReHire]
			, rtrim(e.[Term_Reason]) as Term_Reason
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
			, rtrim(e.[Gender]) as Gender
			, rtrim(e.[PayrollID]) as PayrollID
			, rtrim(e.[StateAllowances]) as StateAllowances
			, rtrim(e.[TransferToStore]) as TransferToStore
			, rtrim(e.[MiddleInitial]) as MiddleInitial
			, e.[AddlWithholding]
			, e.[TaxExempt]
			, rtrim(e.[MaidenName]) as MaidenName
			, isnull(e.[Import_Date],'1/1/1900') as Import_Date
			, rtrim(e.[Race]) as Race
			, e.[BorrowedEmpl]
			, rtrim(isnull(e.[HomeStore],'')) as HomeStore
			, isnull(e.[I9Code],'') as I9Code
			, e.[Info_Sent]
			, isnull(e.[Date_Sent],'1/1/1900') as Date_Sent
			, e.[BonusEligible]
			, isnull(e.[BonusRate],0) as BonusRate
			, e.[SalaryClockIn]
			, e.[InSchool]
			, isnull(e.[VacationSchedule],0) as VacationSchedule
			, isnull(e.[HRDocuments],'') as HRDocuments
			, isnull(e.[AptNumber],'') as AptNumber
			, isnull(e.[DepositType],'') as DepositType
			, isnull(e.[RoutingNumber],'') as RoutingNumber
			, isnull(e.[AccountType],'') as AccountType
			, isnull(e.[AccountNumber],'') as AccountNumber
			, isnull(e.[EmailAddress],'') as EmailAddress
			, isnull(e.[ManagerNote],'') as ManagerNote
			, isnull(e.[InactiveEmpl],'') as InactiveEmpl
			, isnull(e.[IDNumber],'') as IDNumber
			, isnull(e.[PrevIDNumber],'') as PrevIDNumber 
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

	employeeID := r.URL.Query().Get("employeeID")
	unit := r.URL.Query().Get("unit")
	employeeNumber := r.URL.Query().Get("employeeNumber")
	ssn := r.URL.Query().Get("ssn")
	firstName := r.URL.Query().Get("firstName")
	lastName := r.URL.Query().Get("lastName")
	address := r.URL.Query().Get("address")
	city := r.URL.Query().Get("city")
	state := r.URL.Query().Get("state")
	zip := r.URL.Query().Get("zip")
	phone := r.URL.Query().Get("phone")
	birthDate := r.URL.Query().Get("birthDate")
	hireDate := r.URL.Query().Get("hireDate")
	previousSsn := r.URL.Query().Get("previousSsn")
	w4Status := r.URL.Query().Get("w4Status")
	allowances := r.URL.Query().Get("allowances")
	effectiveDate := r.URL.Query().Get("effectiveDate")
	currentRate := r.URL.Query().Get("currentRate")
	previousRate := r.URL.Query().Get("previousRate")
	jobCode := r.URL.Query().Get("jobCode")
	crewRate := r.URL.Query().Get("crewRate")
	otherRate := r.URL.Query().Get("otherRate")
	permanentSupervisor := r.URL.Query().Get("permanentSupervisor")
	salariedManager := r.URL.Query().Get("salariedManager")
	terminated := r.URL.Query().Get("terminated")
	fired := r.URL.Query().Get("fired")
	rehire := r.URL.Query().Get("rehire")
	terminationReason := r.URL.Query().Get("terminationReason")
	terminatedDate := r.URL.Query().Get("terminatedDate")
	jobCode2 := r.URL.Query().Get("jobCode2")
	payRate2 := r.URL.Query().Get("payRate2")
	effectiveDatet2 := r.URL.Query().Get("effectiveDatet2")
	oldRate2 := r.URL.Query().Get("oldRate2")
	jobCode3 := r.URL.Query().Get("jobCode3")
	payRate3 := r.URL.Query().Get("payRate3")
	effectiveDatet3 := r.URL.Query().Get("effectiveDatet3")
	oldRate3 := r.URL.Query().Get("oldRate3")
	jobCode4 := r.URL.Query().Get("jobCode4")
	payRate4 := r.URL.Query().Get("payRate4")
	effectiveDatet4 := r.URL.Query().Get("effectiveDatet4")
	oldRate4 := r.URL.Query().Get("oldRate4")
	jobCode5 := r.URL.Query().Get("jobCode5")
	payRate5 := r.URL.Query().Get("payRate5")
	effectiveDatet5 := r.URL.Query().Get("effectiveDatet5")
	oldRate5 := r.URL.Query().Get("oldRate5")
	overtime1 := r.URL.Query().Get("overtime1")
	overtime2 := r.URL.Query().Get("overtime2")
	overtime3 := r.URL.Query().Get("overtime3")
	overtime4 := r.URL.Query().Get("overtime4")
	overtime5 := r.URL.Query().Get("overtime5")
	gender := r.URL.Query().Get("gender")
	payrollId := r.URL.Query().Get("payrollId")
	stateAllowances := r.URL.Query().Get("stateAllowances")
	transferTostore := r.URL.Query().Get("transferTostore")
	middleInitial := r.URL.Query().Get("middleInitial")
	additionalWithholding := r.URL.Query().Get("additionalWithholding")
	taxExempt := r.URL.Query().Get("taxExempt")
	maidenName := r.URL.Query().Get("maidenName")
	importDate := r.URL.Query().Get("importDate")
	race := r.URL.Query().Get("race")
	borrowedEmployee := r.URL.Query().Get("borrowedEmployee")
	homeStore := r.URL.Query().Get("homeStore")
	i9Code := r.URL.Query().Get("i9Code")
	infoSent := r.URL.Query().Get("infoSent")
	dateSent := r.URL.Query().Get("dateSent")
	bonusEligible := r.URL.Query().Get("bonusEligible")
	bonusRate := r.URL.Query().Get("bonusRate")
	salaryClockin := r.URL.Query().Get("salaryClockin")
	inSchool := r.URL.Query().Get("inSchool")
	vacationSchedule := r.URL.Query().Get("vacationSchedule")
	hrDocuments := r.URL.Query().Get("hrDocuments")
	apartmentNumber := r.URL.Query().Get("apartmentNumber")
	depositType := r.URL.Query().Get("depositType")
	routingNumber := r.URL.Query().Get("routingNumber")
	accountType := r.URL.Query().Get("accountType")
	accountNumber := r.URL.Query().Get("accountNumber")
	emailAddress := r.URL.Query().Get("emailAddress")
	managerNote := r.URL.Query().Get("managerNote")
	inactiveEmployee := r.URL.Query().Get("inactiveEmployee")
	idNumber := r.URL.Query().Get("idNumber")
	previousIdnumber := r.URL.Query().Get("previousIdnumber")

	sql := `Exec post_employee_attributes $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81`
	postemployees := []PostEmployee{}

	//fmt.Println(employeeID, unit, employeeNumber, ssn)

	err := DB().Select(&postemployees, sql, employeeID, unit, employeeNumber, ssn, firstName, lastName, address, city, state, zip, phone, birthDate, hireDate, previousSsn, w4Status, allowances, effectiveDate, currentRate, previousRate, jobCode, crewRate, otherRate, permanentSupervisor, salariedManager, terminated, fired, rehire, terminationReason, terminatedDate, jobCode2, payRate2, effectiveDatet2, oldRate2, jobCode3, payRate3, effectiveDatet3, oldRate3, jobCode4, payRate4, effectiveDatet4, oldRate4, jobCode5, payRate5, effectiveDatet5, oldRate5, overtime1, overtime2, overtime3, overtime4, overtime5, gender, payrollId, stateAllowances, transferTostore, middleInitial, additionalWithholding, taxExempt, maidenName, importDate, race, borrowedEmployee, homeStore, i9Code, infoSent, dateSent, bonusEligible, bonusRate, salaryClockin, inSchool, vacationSchedule, hrDocuments, apartmentNumber, depositType, routingNumber, accountType, accountNumber, emailAddress, managerNote, inactiveEmployee, idNumber, previousIdnumber)
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
		<ul>Employee Number: <input type="text" maxlength="5" size="4" name="employeeNumber" value={{ .EmployeeNumber }}></input></ul>
		<ul>Employee ID: <input type="text" maxlength="20" size="15" name="employeeID" value={{ .EmployeeID }}></input></ul>
		<ul>Unit: <input type="text" maxlength="20" size="15" name="unit" value={{ .Unit }}></input></ul>
		<ul>SSN: <input type="password" maxlength="20" size="15" name="ssn" value={{ .SSN }}></input></ul>
		<ul>First Name: <input type="text" maxlength="20" size="15" name="firstName" value={{ .FirstName }}></input></ul>
		<ul>Last Name: <input type="text" maxlength="20" size="15" name="lastName" value={{ .LastName }}></input></ul>
		<ul>Address: <input type="text" maxlength="20" size="15" name="address" value={{ .Address }}></input></ul>
		<ul>City: <input type="text" maxlength="20" size="15" name="city" value={{ .City }}></input></ul>
		<ul>State: <input type="text" maxlength="20" size="15" name="state" value={{ .State }}></input></ul>
		<ul>Zip: <input type="text" maxlength="20" size="15" name="zip" value={{ .Zip }}></input></ul>
		<ul>Phone: <input type="text" maxlength="20" size="15" name="phone" value={{ .Phone }}></input></ul>
		<ul>BirthDate: <input type="text" maxlength="20" size="15" name="birthDate" value={{ .BirthDate }}></input></ul>
		<ul>HireDate: <input type="text" maxlength="20" size="15" name="hireDate" value={{ .HireDate }}></input></ul>
		<ul>PreviousSSN: <input type="text" maxlength="20" size="15" name="previousSSN" value={{ .PreviousSSN }}></input></ul>
		<ul>W4Status: <input type="text" maxlength="20" size="15" name="w4Status" value={{ .W4Status }}></input></ul>
		<ul>Allowances: <input type="text" maxlength="20" size="15" name="allowances" value={{ .Allowances }}></input></ul>
		<ul>EffectDate: <input type="text" maxlength="20" size="15" name="effectiveDate" value={{ .EffectDate }}></input></ul>
		<ul>CurrentRate: <input type="text" maxlength="20" size="15" name="currentRate" value={{ .CurrentRate }}></input></ul>
		<ul>PreviousRate: <input type="text" maxlength="20" size="15" name="previousRate" value={{ .PreviousRate }}></input></ul>
		<ul>Jobcode: <input type="text" maxlength="20" size="15" name="jobCode" value={{ .Jobcode }}></input></ul>
		<ul>CrewRate: <input type="text" maxlength="20" size="15" name="crewRate" value={{ .CrewRate }}></input></ul>
		<ul>OtherRate: <input type="text" maxlength="20" size="15" name="otherRate" value={{ .OtherRate }}></input></ul>
		<ul>PermanentSupervisor: <input type="text" maxlength="20" size="15" name="permanentSupervisor" value={{ .PermanentSupervisor }}></input></ul>
		<ul>SalariedManager: <input type="text" maxlength="20" size="15" name="salariedManager" value={{ .SalariedManager }}></input></ul>
		<ul>Terminated: <input type="text" maxlength="20" size="15" name="terminated" value={{ .Terminated }}></input></ul>
		<ul>Fired: <input type="text" maxlength="20" size="15" name="fired" value={{ .Fired }}></input></ul>
		<ul>ReHire: <input type="text" maxlength="20" size="15" name="reHire" value={{ .ReHire }}></input></ul>
		<ul>TerminationReason: <input type="text" maxlength="20" size="15" name="terminationReason" value={{ .TerminationReason }}></input></ul>
		<ul>TerminatedDate: <input type="text" maxlength="20" size="15" name="terminatedDate" value={{ .TerminatedDate }}></input></ul>
		<ul>Jobcode2: <input type="text" maxlength="20" size="15" name="jobCode2" value={{ .Jobcode2 }}></input></ul>
		<ul>PayRate2: <input type="text" maxlength="20" size="15" name="payRate2" value={{ .PayRate2 }}></input></ul>
		<ul>EffectiveDate2: <input type="text" maxlength="20" size="15" name="effectiveDate2" value={{ .EffectiveDate2 }}></input></ul>
		<ul>OldRate2: <input type="text" maxlength="20" size="15" name="oldRate2" value={{ .OldRate2 }}></input></ul>
		<ul>Jobcode3: <input type="text" maxlength="20" size="15" name="jobCode3" value={{ .Jobcode3 }}></input></ul>
		<ul>PayRate3: <input type="text" maxlength="20" size="15" name="payRate3" value={{ .PayRate3 }}></input></ul>
		<ul>EffectiveDate3: <input type="text" maxlength="20" size="15" name="effectiveDate3" value={{ .EffectiveDate3 }}></input></ul>
		<ul>OldRate3: <input type="text" maxlength="20" size="15" name="oldRate3" value={{ .OldRate3 }}></input></ul>
		<ul>Jobcode4: <input type="text" maxlength="20" size="15" name="jobCode4" value={{ .Jobcode4 }}></input></ul>
		<ul>PayRate4: <input type="text" maxlength="20" size="15" name="payRate4" value={{ .PayRate4 }}></input></ul>
		<ul>EffectiveDate4: <input type="text" maxlength="20" size="15" name="effectiveDate4" value={{ .EffectiveDate4 }}></input></ul>
		<ul>OldRate4: <input type="text" maxlength="20" size="15" name="oldRate4" value={{ .OldRate4 }}></input></ul>
		<ul>Jobcode5: <input type="text" maxlength="20" size="15" name="jobCode5" value={{ .Jobcode5 }}></input></ul>
		<ul>PayRate5: <input type="text" maxlength="20" size="15" name="payRate5" value={{ .PayRate5 }}></input></ul>
		<ul>EffectiveDate5: <input type="text" maxlength="20" size="15" name="effectiveDate5" value={{ .EffectiveDate5 }}></input></ul>
		<ul>OldRate5: <input type="text" maxlength="20" size="15" name="oldRate5" value={{ .OldRate5 }}></input></ul>
		<ul>Overtime1: <input type="text" maxlength="20" size="15" name="overtime1" value={{ .Overtime1 }}></input></ul>
		<ul>Overtime2: <input type="text" maxlength="20" size="15" name="overtime2" value={{ .Overtime2 }}></input></ul>
		<ul>Overtime3: <input type="text" maxlength="20" size="15" name="overtime3" value={{ .Overtime3 }}></input></ul>
		<ul>Overtime4: <input type="text" maxlength="20" size="15" name="overtime4" value={{ .Overtime4 }}></input></ul>
		<ul>Overtime5: <input type="text" maxlength="20" size="15" name="overtime5" value={{ .Overtime5 }}></input></ul>
		<ul>Gender: <input type="text" maxlength="20" size="15" name="gender" value={{ .Gender }}></input></ul>
		<ul>PayrollID: <input type="text" maxlength="20" size="15" name="payrollID" value={{ .PayrollID }}></input></ul>
		<ul>StateAllowances: <input type="text" maxlength="20" size="15" name="stateAllowances" value={{ .StateAllowances }}></input></ul>
		<ul>TransferToStore: <input type="text" maxlength="20" size="15" name="transferTostore" value={{ .TransferToStore }}></input></ul>
		<ul>MiddleInitial: <input type="text" maxlength="20" size="15" name="middleInitial" value={{ .MiddleInitial }}></input></ul>
		<ul>AdditionalWithholding: <input type="text" maxlength="20" size="15" name="additionalWithholding" value={{ .AdditionalWithholding }}></input></ul>
		<ul>TaxExempt: <input type="text" maxlength="20" size="15" name="TaxExempt" value={{ .TaxExempt }}></input></ul>
		<ul>MaidenName: <input type="text" maxlength="20" size="15" name="maidenName" value={{ .MaidenName }}></input></ul>
		<ul>ImportDate: <input type="text" maxlength="20" size="15" name="importDate" value={{ .ImportDate }}></input></ul>
		<ul>Race: <input type="text" maxlength="20" size="15" name="race" value={{ .Race }}></input></ul>
		<ul>BorrowedEmployee: <input type="text" maxlength="20" size="15" name="borrowedEmployee" value={{ .BorrowedEmployee }}></input></ul>
		<ul>HomeStore: <input type="text" maxlength="20" size="15" name="homeStore" value={{ .HomeStore }}></input></ul>
		<ul>I9Code: <input type="text" maxlength="20" size="15" name="i9Code" value={{ .I9Code }}></input></ul>
		<ul>InfoSent: <input type="text" maxlength="20" size="15" name="infoSent" value={{ .InfoSent }}></input></ul>
		<ul>DateSent: <input type="text" maxlength="20" size="15" name="dateSent" value={{ .DateSent }}></input></ul>
		<ul>BonusEligible: <input type="text" maxlength="20" size="15" name="bonusEligible" value={{ .BonusEligible }}></input></ul>
		<ul>BonusRate: <input type="text" maxlength="20" size="15" name="bonusRate" value={{ .BonusRate }}></input></ul>
		<ul>SalaryClockIn: <input type="text" maxlength="20" size="15" name="salaryClockIn" value={{ .SalaryClockIn }}></input></ul>
		<ul>InSchool: <input type="text" maxlength="20" size="15" name="inSchool" value={{ .InSchool }}></input></ul>
		<ul>VacationSchedule: <input type="text" maxlength="20" size="15" name="vacationSchedule" value={{ .VacationSchedule }}></input></ul>
		<ul>HRDocuments: <input type="text" maxlength="20" size="15" name="hrDocuments" value={{ .HRDocuments }}></input></ul>
		<ul>ApartmentNumber: <input type="text" maxlength="20" size="15" name="apartmentNumber" value={{ .ApartmentNumber }}></input></ul>
		<ul>DepositType: <input type="text" maxlength="20" size="15" name="depositType" value={{ .DepositType }}></input></ul>
		<ul>RoutingNumber: <input type="text" maxlength="20" size="15" name="routingNumber" value={{ .RoutingNumber }}></input></ul>
		<ul>AccountType: <input type="text" maxlength="20" size="15" name="accountType" value={{ .AccountType }}></input></ul>
		<ul>AccountNumber: <input type="text" maxlength="20" size="15" name="accountNumber" value={{ .AccountNumber }}></input></ul>
		<ul>EmailAddress: <input type="text" maxlength="20" size="15" name="emailAddress" value={{ .EmailAddress }}></input></ul>
		<ul>ManagerNote: <input type="text" maxlength="20" size="15" name="managerNote" value={{ .ManagerNote }}></input></ul>
		<ul>InactiveEmployee: <input type="text" maxlength="20" size="15" name="inactiveEmployee" value={{ .InactiveEmployee }}></input></ul>
		<ul>IDNumber: <input type="text" maxlength="20" size="15" name="idNumber" value={{ .IDNumber }}></input></ul>
		<ul>PreviousIDNumber: <input type="text" maxlength="20" size="15" name="previousIdnumber" value={{ .PreviousIDNumber }}></input></ul>		
     {{end}}
	 </ul>
</form>
</div>
 `

/*
http://localhost:5000/getemployees?copyfromemployee=85
http://localhost:5000/postemployees?employeeID=010030000000065&unit=1003&employeeNumber=21602&ssn=555555555&lastName=Buskohl&firstName=Kathy
http://localhost:5000/postemployees?employeeNumber=21602&employeeID=010030000000065&unit=1003&ssn=524118880&firstName=Kathy&lastName=Buskohl
*/

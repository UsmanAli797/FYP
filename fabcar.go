/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"

	"fmt"

	// "strconv"
	// "github.com/hyperledger/fabric/core/chaincode/shim"
	// sc "github.com/hyperledger/fabric/protos/peer"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}
type ChargeSheet struct {
	ChargesheetID         string `json:"chargesheetid"`
	SectionOfLaws         string `json:"sectionoflaws"`
	InvestigatingOfficers string `json:"investigatingofficer"`
	AccusedPerson         string `json:"accusedperson"`
	BriefReport           string `json:"briefReport"`
	ChargedPersons        string `json:"chargedPersons"`
	Plea                  string `json:"plea"`
	Sentence              string `json:"sentence"`
	Judgement             string `json:"judgement"`
	Type	string `json:"type"`

}
type Judgement struct {
	CaseID       string `json:"caseid"`
	Sanction     string `json:"sanction"`
	FinalVerdict string `json:"finalverdict"`
	Type	string `json:"type"`

}

type User struct {
	Fname   string `json:"fname"`
	Email   string `json:"email"`
	IdCard  string `json:"idCard"`
	Phone   string `json:"phone"`
	City    string `json:"city"`
	Address string `json:"address"`
	Pass    string `json:"pass"`
	Type	string `json:"type"`

}
type FIR struct {
	Namee            string `json:"namee"`
	CNIC             string `json:"cnic"`
	PhoneNumber      string `json:"phonenumber"`
	GuardianName     string `json:"guardianname"`
	GuardianPhone    string `json:"guardianphone"`
	PlaceOfOccurence string `json:"placeofoccurence"`
	Addres           string `json:"addres"`
	Against          string `json:"against"`
	Tareekh          string `json:"tareekh"`
	Complaint        string `json:"complaint"`
	WitnessName      string `json:"witnessname"`
	WitnessDetails   string `json:"witnessdetails"`
	Offense          string `json:"offense"`
	EviReport        string `json:"evireport"`
	Type	string `json:"type"`

}

// Definition of the Investigation Arrests structure
type Investigation struct {
	DateTim       string `json:"datetim"`
	CitizenName   string `json:"citizenname"`
	Cause         string `json:"cause"`
	Evidenc       string `json:"evidenc"`
	AccusedName   string `json:"accusedname"`
	InvSummary    string `json:"invsummary"`
	WitnesName    string `json: "witnesname"`
	WitnesSummary string `json: "witnessummary"`
	Type	string `json:"type"`

}

type Complaint struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	EmailId string `json:"emailid"`
	Phonee  string `json:"phonee"`
	Date    string `json:"date"`
	Casee   string `json:"casee"`
	Desc    string `json:"desc"`
	Type	string `json:"type"`
}

type Evidence struct {
	Caseid string `json:"caseid"`
	Path   string `json:"path"`
	Time   string `json:"time"`
	Type	string `json:"type"`

}
type EvidenceReport struct {
	Evidenceid string `json:"evidenceid"`
	Brief      string `json:"brief"`
	Report     string `json:"report"`
	Type	string `json:"type"`

}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	return nil
}

//Create Judgement Report
func (s *SmartContract) CreateJudgement(ctx contractapi.TransactionContextInterface, judgementNumber string, caseid string, sanction string, finalverdict string) error {
	judgement := Judgement{
		CaseID:       caseid,
		Sanction:     sanction,
		FinalVerdict: finalverdict,
		Type:	"Judgement",

	}

	judgementAsBytes, _ := json.Marshal(judgement)

	return ctx.GetStub().PutState(judgementNumber, judgementAsBytes)
}

// QueryJudgement returns the report stored in the world state with given id
func (s *SmartContract) QueryJudgement(ctx contractapi.TransactionContextInterface, Caseid string) (*Judgement, error) {
	judgementAsBytes, err := ctx.GetStub().GetState(Caseid)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if judgementAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", Caseid)
	}

	judgement := new(Judgement)
	_ = json.Unmarshal(judgementAsBytes, judgement)

	return judgement, nil
}

// QueryInvestigation returns the report stored in the world state with given id
func (s *SmartContract) QueryInvestigation(ctx contractapi.TransactionContextInterface, DateTim string) (*Investigation, error) {
	investigationAsBytes, err := ctx.GetStub().GetState(DateTim)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if investigationAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", DateTim)
	}

	investigation := new(Investigation)
	_ = json.Unmarshal(investigationAsBytes, investigation)

	return investigation, nil
}

// EvidenceReportResult structure used for handling result of Evidence
type JudgementResult struct {
	Key    string `json:"Key"`
	Record *Judgement
}

// QueryAllEvidenceReports returns all evidence found in world state
func (s *SmartContract) QueryAllJudgement(ctx contractapi.TransactionContextInterface) ([]JudgementResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []JudgementResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		judgement := new(Judgement)
		_ = json.Unmarshal(queryResponse.Value, judgement)

		judgementResult := JudgementResult{Key: queryResponse.Key, Record: judgement}
		results = append(results, judgementResult)
	}

	return results, nil
}

//Create Evidence Report generated by forensic analyst

func (s *SmartContract) CreateEvidenceReport(ctx contractapi.TransactionContextInterface, evidencereportNumber string, evidenceid string, brief string, report string) error {
	evidencereport := EvidenceReport{
		Evidenceid: evidenceid,
		Brief:      brief,
		Report:     report,
		Type:	"EvidenceReport",

	}

	evidencereportAsBytes, _ := json.Marshal(evidencereport)

	return ctx.GetStub().PutState(evidencereportNumber, evidencereportAsBytes)
}

// QueryEvidenceReport returns the report stored in the world state with given id
func (s *SmartContract) QueryEvidenceReport(ctx contractapi.TransactionContextInterface, Evidenceid string) (*EvidenceReport, error) {
	evidencereportAsBytes, err := ctx.GetStub().GetState(Evidenceid)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if evidencereportAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", Evidenceid)
	}

	evidencereport := new(EvidenceReport)
	_ = json.Unmarshal(evidencereportAsBytes, evidencereport)

	return evidencereport, nil
}

// EvidenceReportResult structure used for handling result of Evidence
type EvidenceReportResult struct {
	Key    string `json:"Key"`
	Record *EvidenceReport
}

// QueryAllEvidenceReports returns all evidence found in world state
func (s *SmartContract) QueryAllEvidenceReports(ctx contractapi.TransactionContextInterface) ([]EvidenceReportResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []EvidenceReportResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		evidencereport := new(EvidenceReport)
		_ = json.Unmarshal(queryResponse.Value, evidencereport)

		evidencereportResult := EvidenceReportResult{Key: queryResponse.Key, Record: evidencereport}
		results = append(results, evidencereportResult)
	}

	return results, nil
}

//create evidence will store evidence path in ledger

func (s *SmartContract) CreateEvidence(ctx contractapi.TransactionContextInterface, evidenceNumber string, caseid string, path string, time string) error {
	evidence := Evidence{
		Caseid: caseid,
		Path:   path,
		Time:   time,
		Type:	"Evidence",

	}

	evidenceAsBytes, _ := json.Marshal(evidence)

	return ctx.GetStub().PutState(evidenceNumber, evidenceAsBytes)
}

// QueryEvidence returns the report stored in the world state with given id
func (s *SmartContract) QueryEvidence(ctx contractapi.TransactionContextInterface, Caseid string) (*Evidence, error) {
	evidenceAsBytes, err := ctx.GetStub().GetState(Caseid)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if evidenceAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", Caseid)
	}

	evidence := new(Evidence)
	_ = json.Unmarshal(evidenceAsBytes, evidence)

	return evidence, nil
}

// EvidenceResult structure used for handling result of Evidence
type EvidenceResult struct {
	Key    string `json:"Key"`
	Record *Evidence
}

// QueryAllEvidence returns all evidence found in world state
func (s *SmartContract) QueryAllEvidence(ctx contractapi.TransactionContextInterface) ([]EvidenceResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []EvidenceResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		evidence := new(Evidence)
		_ = json.Unmarshal(queryResponse.Value, evidence)

		evidenceResult := EvidenceResult{Key: queryResponse.Key, Record: evidence}
		results = append(results, evidenceResult)
	}

	return results, nil
}

//create User will store User in ledger

func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, userNumber string, fname string, email string, idCard string, phone string, city string, address string, pass string) error {
	user := User{
		Fname:   fname,
		Email:   email,
		IdCard:  idCard,
		Phone:   phone,
		City:    city,
		Address: address,
		Pass:    pass,
		Type:	"User",

	}

	userAsBytes, _ := json.Marshal(user)
	return ctx.GetStub().PutState(userNumber, userAsBytes)
}

// QueryUser returns the user stored in the world state with given id
func (s *SmartContract) QueryUser(ctx contractapi.TransactionContextInterface, IdCard string) (*User, error) {
	userAsBytes, err := ctx.GetStub().GetState(IdCard)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if userAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", IdCard)
	}

	user := new(User)
	_ = json.Unmarshal(userAsBytes, user)

	return user, nil
}

//Create Complaint
func (s *SmartContract) CreateComplaint(ctx contractapi.TransactionContextInterface, complaintNumber string, id string, name string, emailid string, phonee string, date string, casee string, desc string) error {
	complaint := Complaint{
		Name:    name,
		ID:      id,
		EmailId: emailid,
		Phonee:  phonee,
		Date:    date,
		Casee:   casee,
		Desc:    desc,
		Type:	"Complaint",
	}

	complaintAsBytes, _ := json.Marshal(complaint)

	return ctx.GetStub().PutState(complaintNumber, complaintAsBytes)
}

// QueryComplaint returns the report stored in the world state with given id
func (s *SmartContract) QueryComplaint(ctx contractapi.TransactionContextInterface, ID string) (*Complaint, error) {
	complaintAsBytes, err := ctx.GetStub().GetState(ID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if complaintAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", ID)
	}

	complaint := new(Complaint)
	_ = json.Unmarshal(complaintAsBytes, complaint)

	return complaint, nil
}

// ComplaintResult structure used for handling result of complaint
type ComplaintResult struct {
	Key    string `json:"Key"`
	Record *Complaint
}

// QueryAllComplaints returns all complaints found in world state
func (s *SmartContract) QueryAllComplaint(ctx contractapi.TransactionContextInterface) ([]ComplaintResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []ComplaintResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		complaint := new(Complaint)
		_ = json.Unmarshal(queryResponse.Value, complaint)

		complaintResult := ComplaintResult{Key: queryResponse.Key, Record: complaint}
		results = append(results, complaintResult)
	}

	return results, nil
}

//create fir will store fir in ledger

func (s *SmartContract) CreateFIR(ctx contractapi.TransactionContextInterface, firNumber string, namee string, cnic string, phonenumber string, guardianname string, guardianphone string, placeofoccurence string, addres string, against string, tareekh string, complaint string, witnessname string, witnessdetails string, offense string, evireport string) error {
	fir := FIR{
		Namee:            namee,
		CNIC:             cnic,
		PhoneNumber:      phonenumber,
		GuardianName:     guardianname,
		GuardianPhone:    guardianphone,
		PlaceOfOccurence: placeofoccurence,
		Addres:           addres,
		Against:          against,
		Tareekh:          tareekh,
		Complaint:        complaint,
		WitnessName:      witnessname,
		WitnessDetails:   witnessdetails,
		Offense:          offense,
		EviReport:        evireport,
		Type:	"FIR",

	}

	firAsBytes, _ := json.Marshal(fir)

	return ctx.GetStub().PutState(firNumber, firAsBytes)
}

// QueryEFIR returns the fir stored in the world state with given id
func (s *SmartContract) QueryFIR(ctx contractapi.TransactionContextInterface, CNIC string) (*FIR, error) {
	firAsBytes, err := ctx.GetStub().GetState(CNIC)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if firAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", CNIC)
	}

	fir := new(FIR)
	_ = json.Unmarshal(firAsBytes, fir)

	return fir, nil
}

// FIRResult structure used for handling result of FIR
type FIRResult struct {
	Key    string `json:"Key"`
	Record *FIR
}

// QueryAllFIR returns all fir found in world state
func (s *SmartContract) QueryAllFIR(ctx contractapi.TransactionContextInterface) ([]FIRResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []FIRResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		fir := new(FIR)
		_ = json.Unmarshal(queryResponse.Value, fir)

		firResult := FIRResult{Key: queryResponse.Key, Record: fir}
		results = append(results, firResult)
	}

	return results, nil
}

//create chargessheet will store chargesheet path in ledger

func (s *SmartContract) CreateChargesheet(ctx contractapi.TransactionContextInterface, csNumber string, chargesheetid string, sectionoflaws string, investigatingofficer string, accusedperson string, briefReport string, chargedPersons string, plea string, sentence string, judgement string) error {
	chargesheet := ChargeSheet{
		ChargesheetID:         chargesheetid,
		SectionOfLaws:         sectionoflaws,
		InvestigatingOfficers: investigatingofficer,
		AccusedPerson:         accusedperson,
		BriefReport:           briefReport,
		ChargedPersons:        chargedPersons,
		Plea:                  plea,
		Sentence:              sentence,
		Judgement:             judgement,
		Type:	"ChargeSheet",

	}

	chargesheetAsBytes, _ := json.Marshal(chargesheet)

	return ctx.GetStub().PutState(csNumber, chargesheetAsBytes)
}

// Querychargesheet returns the chargesheet stored in the world state with given id
func (s *SmartContract) QueryChargeSheet(ctx contractapi.TransactionContextInterface, ChargessheetID string) (*ChargeSheet, error) {
	chargesheetAsBytes, err := ctx.GetStub().GetState(ChargessheetID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if chargesheetAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", ChargessheetID)
	}

	chargesheet := new(ChargeSheet)
	_ = json.Unmarshal(chargesheetAsBytes, chargesheet)

	return chargesheet, nil
}

// ChargeSheetResult structure used for handling result of Evidence
type ChargeSheetResult struct {
	Key    string `json:"Key"`
	Record *ChargeSheet
}

// QueryAllChargesheet returns all evidence found in world state
func (s *SmartContract) QueryAllChargeSheet(ctx contractapi.TransactionContextInterface) ([]ChargeSheetResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []ChargeSheetResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		chargesheet := new(ChargeSheet)
		_ = json.Unmarshal(queryResponse.Value, chargesheet)

		chargesheetResult := ChargeSheetResult{Key: queryResponse.Key, Record: chargesheet}
		results = append(results, chargesheetResult)
	}

	return results, nil
}

//create investigationreport will store investigationreport path in ledger

func (s *SmartContract) CreateInvestigationReport(ctx contractapi.TransactionContextInterface, invNumber string, datetim string, citizenname string, cause string, evidenc string, accusedname string, invsummary string, witnesname string, witnessummary string) error {
	investigation := Investigation{
		DateTim:       datetim,
		CitizenName:   citizenname,
		Cause:         cause,
		Evidenc:       evidenc,
		AccusedName:   accusedname,
		InvSummary:    invsummary,
		WitnesName:    witnesname,
		WitnesSummary: witnessummary,
		Type:	"InvestigationReport",

	}

	investigationAsBytes, _ := json.Marshal(investigation)

	return ctx.GetStub().PutState(invNumber, investigationAsBytes)
}

// ChargeSheetResult structure used for handling result of Evidence
type InvestigationResult struct {
	Key    string `json:"Key"`
	Record *Investigation
}

// QueryAllInvestigation returns all investigation reports found in world state
func (s *SmartContract) QueryAllInvestigation(ctx contractapi.TransactionContextInterface) ([]InvestigationResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []InvestigationResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		investigation := new(Investigation)
		_ = json.Unmarshal(queryResponse.Value, investigation)

		investigationResult := InvestigationResult{Key: queryResponse.Key, Record: investigation}
		results = append(results, investigationResult)
	}

	return results, nil
}

// Chaincode for login
func (s *SmartContract) Login(ctx contractapi.TransactionContextInterface, Email string, pass string) bool {
	user, err := s.QueryUser(ctx, Email)
	if err != nil {
		return false
	}

	// if email == user.Email {
	// 	return true
	// }

	if user.Pass == pass {
		fmt.Printf("Successfull Login: %s", Email)
		return true
	}

	return false
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}

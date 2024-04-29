package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// data struct for one record
type Record struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// blockchain block struct
type Block struct {
	Index     int
	Timestamp string
	Record    Record
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%s", block.Index, block.Timestamp, block.Record.ID, block.Record.Name, block.PrevHash)
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock Block, record Record) Block {
	var newBlock Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Record = record
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func isBlockValid(newBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != newBlock.Index {
		return false
	}
	if prevBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func initBlockchain() {
	genesisBlock := Block{}
	genesisBlock.Index = 0
	genesisBlock.Timestamp = time.Now().String()
	genesisBlock.Record = Record{ID: "0", Name: "Genesis Block"}
	genesisBlock.PrevHash = ""
	genesisBlock.Hash = calculateHash(genesisBlock)

	Blockchain = append(Blockchain, genesisBlock) // Add genesis block to the blockchain
}

func main() {
	jsonFile := flag.String("json", "", "Path tp JSON file")
	csvFile := flag.String("csv", "", "Path tp CSV file ")
	flag.Parse()

	initBlockchain()

	if *jsonFile != "" {
		processJSON(*jsonFile)
	}
	if *csvFile != "" {
		processCSV(*csvFile)
	}
	if *jsonFile == "" && *csvFile == "" {
		processManualInput()
	}
	fmt.Println("Blockchain:")
	spew.Dump(Blockchain)
}
func processJSON(filename string) {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	var records []Record
	if err := json.Unmarshal(jsonData, &records); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	for _, record := range records {
		prevBlock := Blockchain[len(Blockchain)-1]
		newBlock := generateBlock(prevBlock, record)

		if isBlockValid(newBlock, prevBlock) {
			Blockchain = append(Blockchain, newBlock)
			fmt.Println("New Block added to the blockchain.")
			spew.Dump(newBlock)
		}
	}
}
func processCSV(filename string) {
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}
	for _, row := range csvData {
		record := Record{
			ID:    row[0],
			Name:  row[1],
			Value: row[2],
		}
		prevBlock := Blockchain[len(Blockchain)-1]
		newBlock := generateBlock(prevBlock, record)

		if isBlockValid(newBlock, prevBlock) {
			Blockchain = append(Blockchain, newBlock)
			fmt.Println("New Block added to the blockchain.")
			spew.Dump(newBlock)
		}
	}
}
func processManualInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter data manually:")
	fmt.Print("ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Value: ")
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	record := Record{
		ID:    id,
		Name:  name,
		Value: value,
	}
	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := generateBlock(prevBlock, record)

	if isBlockValid(newBlock, prevBlock) {
		Blockchain = append(Blockchain, newBlock)
		fmt.Println("New Block added tp the blockchain.")
		spew.Dump(newBlock)
	}
}

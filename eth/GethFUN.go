package eth

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/go.net/context"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type UserInit struct {
	Key      string
	Password string
}

type VoteList struct {
	Ip          common.Address
	VName       string
	Chairperson common.Address
	VoteCount   *big.Int
	LimitTime   *big.Int
}

type BankList struct {
	Ip              common.Address
	OperationAmount *big.Int
	Operation       string
	OperationTime   *big.Int
}

type PerList struct {
	Ip   common.Address
	Fun  string
	Time *big.Int
	Stu  bool
}

//var test = UserInit{Key: `{"address":"0728fb9ca555568ecdc3e064eccbe77fd6337600","crypto":{"cipher":"aes-128-ctr","ciphertext":"6806499658d053764fe6a7ef7134e9ce073f4e2a0aa96849d2f21c857054c0bb","cipherparams":{"iv":"d9c1d8e1928f5ef764d02388fd1976e8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ff81faa317f91f3be717a9154ef8aedb01b02cf82bf46f7d18ff1a3b2a0a1aac"},"mac":"b93affed22ccadc9b063b2eb0a630b856d0e2437578178238921d43e7ef8e0be"},"id":"77b8fba2-5cdb-41ef-8bc0-270bab2e73ef","version":3}`,
//	Password: "1234"}

//var Key = `{"address":"25773d413b266f068dc8dd737b73a5baa363ecf2","crypto":{"cipher":"aes-128-ctr","ciphertext":"abf3912fcbbdeab084da798f18a2f09f3a1fe1cf8c2050b370eb64aeb3892018","cipherparams":{"iv":"559bb76888974a2e731b59cec76db810"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1980384aefcdbad0fbf31b16bd7fb67ef7bee9be58ab7bf379d1ec54818751a7"},"mac":"894c58b65967d4dee33c7ab2d2c8edf89aaf0d439ca152e299e3cb7dca25a787"},"id":"b58b8f1c-5e22-4d18-af01-4643019bea22","version":3}`

const address = "0xcABB6D1E6e45604f045DE21c6470c0217FEf3cd1"

func geth() *Abi {
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
	}
	defer conn.Close()

	token, err := NewAbi(common.HexToAddress(address), conn)
	if err != nil {
		log.Fatalf("could not newAbi: %v", err)
	}
	return token
}

func (u *UserInit) Auth() *bind.TransactOpts {
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(u.Key), u.Password, big.NewInt(1337))
	if err != nil {
		log.Fatalf("could not create auth: %v", err)
	}
	return auth
}

func Listen() []byte {
	client, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	_, err = client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if len(vLog.Data) >= 96 {
				return vLog.Data
			}
		}
	}
	return nil
}

func InitName(name string, auth *bind.TransactOpts) *types.Transaction {
	initialization, err := geth().Initialization(auth, name)
	if err != nil {
		log.Fatalf("InitName is err%v", err)
	}
	initialization.Hash()
	return initialization
}

func InitNameStu(addr string) bool {
	stu, err := geth().InitializationSTU(&bind.CallOpts{From: common.HexToAddress(addr)})
	if err != nil {
		log.Fatalf("InitNameStu is err%v", err)
	}
	return stu
}

func TakePart(fun string, ZHNum int, JLNum int, auth *bind.TransactOpts) float64 {
	_, err := geth().TakeAPart(auth, fun, big.NewInt(int64(ZHNum)), big.NewInt(int64(JLNum)))
	if err != nil {
		log.Fatalf("TakePart is err%v", err)
	}
	var r float64
	var k []string
	for _, v := range Listen()[32:64] {
		if v > 0 {
			k = append(k, hex.EncodeToString([]byte{v}))
		}
	}
	for i, v := range k {
		n, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			fmt.Println(err)
		}
		s := float64(n)
		r += s * math.Pow(256, float64(len(k)-1-i))
	}
	return r
}

func Report(addr string, reason string, auth *bind.TransactOpts) float64 {
	_, err := geth().Report(auth, common.HexToAddress(addr), reason)
	if err != nil {
		log.Fatalf("Report is err%v", err)
	}
	var r float64
	var k []string
	for _, v := range Listen()[32:64] {
		if v > 0 {
			k = append(k, hex.EncodeToString([]byte{v}))
		}
	}
	for i, v := range k {
		n, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			fmt.Println(err)
		}
		s := float64(n)
		r += s * math.Pow(256, float64(len(k)-1-i))
	}
	return r
}

func Finish(pid int, reword int, massID int, auth *bind.TransactOpts) {
	_, err := geth().Finish(auth, big.NewInt(int64(pid)), big.NewInt(int64(reword)), big.NewInt(int64(massID)))
	if err != nil {
		log.Fatalf("Finish is err%v", err)
	}
}

func ReleaseNeutralize(year int, neutralizeTotal int, auth *bind.TransactOpts) {
	_, err := geth().ReleaseNeutralize(auth, big.NewInt(int64(year)), big.NewInt(int64(neutralizeTotal)))
	if err != nil {
		log.Fatalf("ReleaseNeutralize is err%v", err)
	}
}

func Exchange(amount int, auth *bind.TransactOpts) {
	_, err := geth().Exchange(auth, big.NewInt(int64(amount)))
	if err != nil {
		log.Fatalf("Exchange is err%v", err)
	}
}

func Bank(amount int, year int, interest int, auth *bind.TransactOpts) (string, string) {
	_, err := geth().Bank(auth, big.NewInt(int64(amount)), big.NewInt(int64(year)), big.NewInt(int64(interest)))
	if err != nil {
		log.Fatalf("Exchange is err%v", err)
	}

	a := Listen()[32:96]
	return hex.EncodeToString(a[0:32]), hex.EncodeToString(a[32:64])
}

func GetInterest(hash string, auth *bind.TransactOpts) {
	var a []string
	for i := 0; i < len(hash); i += 2 {
		a = append(a, string(hash[i])+string(hash[i+1]))
	}
	var b []byte
	for _, v := range a {
		parseInt, _ := strconv.ParseInt(v, 16, 32)
		b = append(b, byte(parseInt))
	}
	key := [32]byte{}
	copy(key[:], b)
	_, err := geth().GetInterest(auth, key)
	if err != nil {
		log.Fatalf("GetInterest is err%v", err)
	}
}

func GETBankVoucherHash(hash string) struct {
	Participant   common.Address
	PartakeAmount *big.Int
	Deadline      *big.Int
	Interest      *big.Int
	TakeOut       bool
} {
	var a []string
	for i := 0; i < len(hash); i += 2 {
		a = append(a, string(hash[i])+string(hash[i+1]))
	}
	var b []byte
	for _, v := range a {
		parseInt, _ := strconv.ParseInt(v, 16, 32)
		b = append(b, byte(parseInt))
	}
	key := [32]byte{}
	copy(key[:], b)
	voucherHash, err := geth().BankVoucherHash(nil, key)
	if err != nil {
		log.Fatalf("GETBankVoucherHash is err%v", err)
	}
	return voucherHash
}

func GETVote(pid int) VoteList {
	vote, err := geth().Vote(nil, big.NewInt(int64(pid)))
	if err != nil {
		log.Fatalf("GETVote is err%v", err)
	}
	return vote
}

func Voting(pid int, auth *bind.TransactOpts) {
	_, err := geth().Voting(auth, big.NewInt(int64(pid)))
	if err != nil {
		log.Fatalf("Voting is err%v", err)
	}
}

func VotingTF(pid int, addr string) string {
	tf, err := geth().VotingTF(nil, big.NewInt(int64(pid)), common.HexToAddress(addr))
	if err != nil {
		log.Fatalf("VotingTF is err%v", err)
	}
	tm := time.Unix(tf.Int64(), 10)
	dateString := fmt.Sprintf("%d年%d月%d日 %d:%d:%d\n", tm.Year()-1, tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	return dateString
}

func GETNewPersonal(addr string) struct {
	PersonalName      string
	PersonalAmount    *big.Int
	NeutralizedAmount *big.Int
	RegistrationTime  *big.Int
	NumberOfReported  *big.Int
} {
	personal, err := geth().NewPersonal(nil, common.HexToAddress(addr))
	if err != nil {
		log.Fatalf("GETNewPersonal is err%v", err)
	}
	return personal
}

func GetVoting() []VoteList {
	rvnum, err := geth().RVNUM(nil)
	if err != nil {
		log.Fatalf("RVNUM is err%v", err)
	}
	var voteList []VoteList
	for i := 1; i <= int(rvnum.Int64()); i++ {
		vote, err := geth().Vote(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Fatalf("GETVote is err%v", err)
		}
		voteList = append(voteList, vote)
	}
	return voteList
}

func GetYearData() struct {
	Publisher           common.Address
	NeutralizeTotal     *big.Int
	NeutralizePersonals *big.Int
	PeopleNUM           *big.Int
	RecordTime          *big.Int
	UpToStandard        bool
	CurrentQuota        *big.Int
} {
	year, err := geth().CurrentYear(nil)
	if err != nil {
		log.Fatalf("CurrentYear is err%v", err)
	}
	neutralizeYear, err := geth().NeutralizeYear(nil, year)
	if err != nil {
		log.Fatalf("NeutralizeYear is err%v", err)
	}
	return neutralizeYear
}

func GETYear() *big.Int {
	year, err := geth().CurrentYear(nil)
	if err != nil {
		log.Fatalf("CurrentYear is err%v", err)
	}
	return year
}

func TakeOut(amount int, operation string, auth *bind.TransactOpts) {
	_, err := geth().TakeOutBank(auth, big.NewInt(int64(amount)), operation)
	if err != nil {
		log.Fatalf("TakeOutBank is err%v", err)
	}
}

func Deposit(amount int, operation string, auth *bind.TransactOpts) {
	_, err := geth().Deposit(auth, big.NewInt(int64(amount)), operation)
	if err != nil {
		log.Fatalf("Deposit is err%v", err)
	}
}

func GetBankList() []BankList {
	num, err := geth().BankAdminOperationNum(nil)
	if err != nil {
		log.Fatalf("BankAdminOperationNum is err%v", err)
	}
	var list []BankList
	for i := 0; i < int(num.Int64()); i++ {
		lists, err := geth().BankAdminList(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Fatalf("BankAdminList is err%v", err)
		}
		list = append(list, lists)
	}
	return list
}

func Approve(ip string, amount int, auth *bind.TransactOpts) {
	_, err := geth().Approve(auth, common.HexToAddress(ip), big.NewInt(int64(amount)))
	if err != nil {
		log.Fatalf("Approve is err%v", err)
	}
}

func INApprove(ip string, amount int, auth *bind.TransactOpts) {
	_, err := geth().IncreaseAllowance(auth, common.HexToAddress(ip), big.NewInt(int64(amount)))
	if err != nil {
		log.Fatalf("IncreaseAllowance is err%v", err)
	}
}

func DEApprove(ip string, amount int, auth *bind.TransactOpts) {
	_, err := geth().DecreaseAllowance(auth, common.HexToAddress(ip), big.NewInt(int64(amount)))
	if err != nil {
		log.Fatalf("DecreaseAllowance is err%v", err)
	}
}

func SelectApprove(OWNip string, ip string) *big.Int {
	allowance, err := geth().Allowance(nil, common.HexToAddress(OWNip), common.HexToAddress(ip))
	if err != nil {
		log.Fatalf("Allowance is err%v", err)
	}
	return allowance
}

func Transfer(ip string, amount int, auth *bind.TransactOpts) {
	_, err := geth().Transfes(auth, common.HexToAddress(ip), big.NewInt(int64(amount)))
	if err != nil {
		log.Fatalf("Transfes is err%v", err)
	}
	_, err = geth().SubERC(auth, common.HexToAddress(ip), big.NewInt(int64(amount)))
	if err != nil {
		log.Fatalf("SubERC is err%v", err)
	}
}

func TransferForm(ip string, OWNip string, amount int, auth *bind.TransactOpts) {
	_, err := geth().TransferFrom(auth, common.HexToAddress(ip), common.HexToAddress(OWNip), big.NewInt(int64(amount)))
	if err != nil {
		log.Fatalf("Transfes is err%v", err)
	}
	_, err = geth().SubERC(auth, common.HexToAddress(ip), big.NewInt(int64(amount)))
	if err != nil {
		log.Fatalf("SubERC is err%v", err)
	}
}

func ERCInfo() string {
	name, err := geth().Name(nil)
	if err != nil {
		log.Fatalf("Name is err%v", err)
	}
	return name
}

func ERCLogo() string {
	name, err := geth().Symbol(nil)
	if err != nil {
		log.Fatalf("Name is err%v", err)
	}
	return name
}

func ERCAmount() *big.Int {
	name, err := geth().TotalSupply(nil)
	if err != nil {
		log.Fatalf("Name is err%v", err)
	}
	return name
}

func GETNUM() []PerList {
	num, err := geth().NUM(nil)
	if err != nil {
		log.Fatalf("NUM is err%v", err)
	}
	var personalList []PerList
	for i := 0; i < int(num.Int64()); i++ {
		id, err := geth().MessageID(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Fatalf("MessageID is err%v", err)
		}
		personalList = append(personalList, id)
	}

	return personalList
}

func GETUserNum() *big.Int {
	num, err := geth().UserNum(nil)
	if err != nil {
		log.Fatalf("GETUserNum is err%v", err)
	}
	return num
}

func NUM() *big.Int {
	num, err := geth().NUM(nil)
	if err != nil {
		log.Fatalf("NUM is err%v", err)
	}
	return num
}

func GetRVNUM() *big.Int {
	rvnum, err := geth().RVNUM(nil)
	if err != nil {
		log.Fatalf("RVNUM is err%v", err)
	}
	return rvnum
}

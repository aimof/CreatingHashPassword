package main
import (
	"../hash"
	"../db"
	"fmt"
	"sort"
)

func test_HashSum() bool {
	teststring := "teststring"
	var teststring_hash map[string]string = map[string]string{}
	teststring_hash["sha256"]="3c8727e019a42b444667a587b6001251becadabbb36bfed8087a92c18882d111"
	teststring_hash["sha384"]="edd2ab3262b6c0121d706087045b60d51d6dc2f7419987ba12c983053a70c1057f15f58608ee07e1225266df36ba2c9c"
	teststring_hash["sha512"]="6253b39071e5df8b5098f59202d414c37a17d6a38a875ef5f8c7d89b0212b028692d3d2090ce03ae1de66c862fa8a561e57ed9eb7935ce627344f742c0931d72"
	teststring_hash["ripemd160"]="cd566972b5e50104011a92b59fa8e0b1234851ae"
	teststring_hash["sha3_224"]="17fce94d5da9ed5ce89fb6aaa884ecc94daf3ec35946cc39cbefbda6"
	teststring_hash["sha3_256"]="a52ef463edb03558be0598fb9a3ae67e0fc41c44bd0b1d4392b354d1e817fea8"
	teststring_hash["sha3_384"]="39660f7772dcc83fc8b5b67dc25537f23bb1bd34967d1e7ba8ff7585967553404b3114ac74ff6ecf7219200a1533c20c"
	teststring_hash["sha3_512"]="80938d92fe5d216884786ab8c23b87dc1c572e603c5b5d84d2454fa9bebbc84afa1d7a33ed35ed0196a023043e29c8f67b452102571e0f1944b8381b213a81fb"
	teststring_hash["sha512_224"]="da91a185e8c947f031419df9b9fc76accd56bb62ef8024c21046c364"
	teststring_hash["sha512_256"]="107b9a6123f289a3b14d7ce109c9b94feda782bf5bcaa09ce9e6f00773e21430"
	teststring_hash["black2s_256"]="73fb0bb92d259f9c9b13398aec62c73f204279661ea51b37ab72c11acb27aa3a"
	teststring_hash["black2s_384"]="9078c57a4ba97ea1b1bbbdc94df0f0cfd64246d154d7bdd730520737f20c970570e4e87ca17799242fa85d0fc63f56a3"
	teststring_hash["black2s_512"]="eb2c1252ac1d7d684f1b474b207610aa58c62d2b335b9adeac269d4eab5bd2e7bfe2f6a7a51be61b4f55e5c7dfced6922ca66af2f5eddf8d539f4ca1a28c1232"
	teststring_hash["md5"]="d67c5cbf5b01c9f91932e3b8def5e5f8"
	teststring_hash["other"]=""

	for alg, result := range teststring_hash {
		res := hash.HashSum(alg, teststring)
		if result != res {
			fmt.Printf("%s\n", res)
			fmt.Printf("%s\n", result)
			fmt.Printf("Failed to do hash %s\n", alg)
			return false
		}
	}
	return true
}

func test_AlgorithmList() bool {
	var result []string = []string{"sha256", "sha384", "sha512", "sha3_256", "sha3_384", "sha3_512", "sha512_256", "black2s_256", "black2s_384", "black2s_512"}
	sort.SliceStable(result, func(i, j int) bool { return result[i] < result[j] })
	res := hash.AlgorithmList()
	if(len(result) != len(res)) {
		fmt.Printf("Failed to do hash algorithm %d %d\n", len(result), len(res))
		return false
	}

	for i := range res {
		if res[i] != result[i] {
			fmt.Printf("Failed to do hash algorithm %s and %s\n", res[i], result[i])
			return false
		}
	}

	return true
}

func test_GetAccount() bool {
	res,err := db.GetAccount("test")
	if res.Username !=  "test" || res.Passphrase != "testpass" || err != nil {
		fmt.Printf("Failed to get GetAccount with test\n")
		return false
	}

	res, err = db.GetAccount("test2")
	if res.Username !=  "test2" || res.Passphrase != "testpass2" || err != nil{
		fmt.Printf("Failed to get GetAccount with test2\n")
		return false
	}

	res, err = db.GetAccount("nodata")
	if res.Username !=  "" || res.Passphrase != "" || err == nil{
		fmt.Printf("Failed to get GetAccount without data\n")
		return false
	}
	return true
}

func test_GetAllPassphrase() bool {
	res,err := db.GetAllPassphrase("test")
	if len(res) != 2 || err != nil{
		fmt.Printf("Failed to get GetAllPassphrase without data\n")
		return false
	}

	if res[0].Algorithm != "sha256" {
		result1 := (res[0].Username != "test" || res[0].Title != "testtitle2" ||  res[0].Algorithm != "sha512" || res[0].Seed != "seed2" || res[0].Length != "32" || res[0].DisableSymbol != "false" )
		result2 := (res[1].Username != "test" || res[1].Title != "testtitle1" ||  res[1].Algorithm != "sha256" || res[1].Seed != "seed1" || res[1].Length !="64" || res[1].DisableSymbol != "true")
		if result1 && result2 {
			fmt.Printf("Failed to get GetAllPassphrase without data\n")
			return false
		}
	} else {
		result1 := (res[1].Username != "test" || res[1].Title != "testtitle2" ||  res[1].Algorithm != "sha512" || res[1].Seed != "seed2" || res[1].Length != "32" || res[1].DisableSymbol != "false")
		result2 := (res[0].Username != "test" || res[0].Title != "testtitle1" ||  res[0].Algorithm != "sha256" || res[0].Seed != "seed1" || res[0].Length !="64" || res[0].DisableSymbol != "true" )
		if result1 && result2 {
			fmt.Printf("Failed to get GetAllPassphrase without data\n")
			return false
		}
	}
	return true
}

func test_GetPassphraseInfo() bool {
	// no info
	res, err := db.GetPassphraseInfo("test", "testtitle")
	if res.Title != "" || err == nil {
		fmt.Printf("Failed to get GetPassphraseInfo\n")
		return false
	}

	res, err = db.GetPassphraseInfo("test", "testtitle1")
	result1 := (res.Username != "test" || res.Title != "testtitle1" ||  res.Algorithm != "sha256" || res.Seed != "seed1" || res.Length != "64" || res.DisableSymbol != "true")
	if result1 || err != nil{
		fmt.Printf("Failed to get GetPassphraseInfo with testtitle1\n")
		return false
	}

	return true
}

func test_SetAccount() bool {
	//add
	account := db.Account{"newuser", "newpassphrase"}
	if !db.SetAccount(account) {
		fmt.Printf("Failed to set New account newuser\n")
		return false
	}

	res, err := db.GetAccount(account.Username)
	if res.Username != account.Username || res.Passphrase != account.Passphrase || err != nil {
		fmt.Printf("Failed to check New account newuser\n")
		return false
	}

	//update
	account.Passphrase = "newpassphrase2"
	if !db.SetAccount(account) {
		fmt.Printf("Failed to set New account newuser\n")
		return false
	}

	res, err = db.GetAccount(account.Username)
	if res.Username != account.Username || res.Passphrase != account.Passphrase || err != nil {
		fmt.Printf("Failed to check New account newuser\n")
		return false
	}
	return true
}

func test_DeleteAccount() bool{
	if !db.DeleteAccount("newuser") {
		fmt.Printf("Failed to remove account newuser\n")
		return false
	}

	res, err := db.GetAccount("newuser")
	if res.Username != "" || err == nil {
		fmt.Printf("Failed to remove account newuser 2\n")
		return false
	}

	/*true or false?*/
	if !db.DeleteAccount("otheruser") {
		return false
	}
	return true
}

func test_SetPassphraseInfo() bool {
	//add
	passphrase := db.PassphraseInfo{"test", "newalgorithm", "newseed", "newtitle", "10", "true"}
	if !db.SetPassphraseInfo(passphrase) {
		fmt.Printf("Failed to set New passphraseInfo\n")
		return false
	}

	res, err := db.GetPassphraseInfo(passphrase.Username, passphrase.Title)
	if res.Username != passphrase.Username || res.Title != passphrase.Title || res.Algorithm != passphrase.Algorithm || res.Seed != passphrase.Seed || res.Length != passphrase.Length || res.DisableSymbol != passphrase.DisableSymbol || err != nil {
		fmt.Printf("Failed to check New passphrase info (newtitle)\n")
		return false
	}

	//update
	passphrase.Seed = "newseed2"
	if !db.SetPassphraseInfo(passphrase) {
		fmt.Printf("Failed to set New passphrase info (newtitle)\n")
		return false
	}

	res, err = db.GetPassphraseInfo(passphrase.Username, passphrase.Title)
	if res.Username != passphrase.Username || res.Title != passphrase.Title || res.Algorithm != passphrase.Algorithm || res.Seed != passphrase.Seed || res.Length != passphrase.Length || res.DisableSymbol != passphrase.DisableSymbol || err != nil{
		fmt.Printf("Failed to check New passphrase info (newtitle)\n")
		return false
	}

	//no insert data, no user
	passphrase.Username = "unknown_user"
	if db.SetPassphraseInfo(passphrase) {
		fmt.Printf("Not failed to add unknown user\n")
		return false
	}
	return true
}

func test_DeletePassphraseInfo() bool{
	if !db.DeletePassphraseInfo("test", "newtitle") {
		fmt.Printf("Failed to remove account newtitle\n")
		return false
	}

	res, err := db.GetPassphraseInfo("test", "newtitle")
	if res.Username != "" || err == nil {
		fmt.Printf("Failed to remove account newtitle 2\n")
		return false
	}

	/*true or false?*/
	if !db.DeleteAccount("otheruser") {
		return false
	}
	return true
}

func main() {
	if !test_HashSum() {
		fmt.Printf("Failed to do HashSum\n")
		return
	}
	if !test_AlgorithmList() {
		fmt.Printf("Failed to do AlgorithmList\n")
		return
	}

	if !test_GetAccount() {
		fmt.Printf("Failed to do GetAccount\n")
		return
	}

	if !test_GetAllPassphrase() {
		fmt.Printf("Failed to do GetAllPassphrase\n")
		return
	}

	if !test_GetPassphraseInfo() {
		fmt.Printf("Failed to do GetPassphraseInfo\n")
		return
	}

	if !test_SetAccount(){
		fmt.Printf("Failed to do SetAccount\n")
		return
	}

	if !test_DeleteAccount(){
		fmt.Printf("Failed to do DeleteAccount\n")
		return
	}

	if !test_SetPassphraseInfo() {
		fmt.Printf("Failed to do SetPassphraseInfo\n")
		return
	}
	if !test_DeletePassphraseInfo() {
		fmt.Printf("Failed to do DeletePassphraseInfo\n")
		return
	}

	fmt.Printf("Finish to do all test\n")
}

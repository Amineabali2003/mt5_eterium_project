package services

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/idir-44/ethereum/internal/model"
)

func getTransactions(user model.User) (model.TransactionResponse, error) {
	apiKey := os.Getenv("ETH_API_KEY")
	if apiKey == "" {
		return model.TransactionResponse{}, fmt.Errorf("api key not set")
	}

	url := fmt.Sprintf("https://api.etherscan.io/v2/api?chainid=1&module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=10000&sort=asc&apikey=%s", user.WalletAddress, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return model.TransactionResponse{}, fmt.Errorf("erreur lors de la requête HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.TransactionResponse{}, fmt.Errorf("erreur lors de la lecture de la réponse: %v", err)
	}

	var transactionResp model.TransactionResponse
	if err := json.Unmarshal(body, &transactionResp); err != nil {
		return model.TransactionResponse{}, fmt.Errorf("erreur lors du parsing JSON: %v", err)
	}

	return transactionResp, nil
}

func (s service) GetWalletData(userID string) ([]model.WalletDataResponse, error) {
	user, err := s.repository.GetUser(userID)
	if err != nil {
		return []model.WalletDataResponse{}, err
	}

	if user.WalletAddress == "" {
		return []model.WalletDataResponse{}, fmt.Errorf("user doesn't have a wallet")
	}

	transactions, err := getTransactions(user)
	if err != nil {
		return []model.WalletDataResponse{}, err
	}

	dailyBalances := make(map[string]float64)
	currentBalance := 0.0

	for _, tx := range transactions.Result {
		timestamp, _ := strconv.ParseInt(tx.TimeStamp, 10, 64)
		date := time.Unix(timestamp, 0).Format("2006-01-02")

		wei, _ := new(big.Float).SetString(tx.Value)
		eth := new(big.Float).Quo(wei, new(big.Float).SetFloat64(1e18))
		ethValue, _ := eth.Float64()

		var change float64

		if strings.EqualFold(tx.From, user.WalletAddress) {
			gasPrice, _ := new(big.Float).SetString(tx.GasPrice)
			gasUsed, _ := new(big.Float).SetString(tx.GasUsed)
			gasCost := new(big.Float).Mul(gasPrice, gasUsed)
			gasCostEth := new(big.Float).Quo(gasCost, new(big.Float).SetFloat64(1e18))
			gasCostEthValue, _ := gasCostEth.Float64()

			change = -(ethValue + gasCostEthValue)
		} else if strings.EqualFold(tx.To, user.WalletAddress) {
			change = ethValue
		}

		currentBalance += change
		dailyBalances[date] = currentBalance
	}

	chartData := []model.WalletDataResponse{}
	for date, balance := range dailyBalances {

		chartData = append(chartData, model.WalletDataResponse{Date: date, Price: balance})
	}

	return chartData, nil

}

func (s service) GetWalletTransactions(userID string) (model.TransactionResponse, error) {
	user, err := s.repository.GetUser(userID)
	if err != nil {
		return model.TransactionResponse{}, err
	}
	if user.WalletAddress == "" {
		return model.TransactionResponse{}, fmt.Errorf("user doesn't have a wallet")
	}

	transactions, err := getTransactions(user)
	if err != nil {
		return model.TransactionResponse{}, err
	}

	return transactions, nil
}

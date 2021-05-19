package cmd

import (
	"fmt"
	"strconv"

	"github.com/echovl/cardano-wallet/db"
	"github.com/echovl/cardano-wallet/wallet"
	"github.com/spf13/cobra"
)

// listAddressCmd represents the listAddress command
var listAddressCmd = &cobra.Command{
	Use:     "list-address [wallet-id]",
	Short:   "Print a list of known wallet's addresses",
	Aliases: []string{"lsa"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		useTestnet, err := cmd.Flags().GetBool("testnet")
		network := wallet.Mainnet
		if useTestnet {
			network = wallet.Testnet
		}

		walletID := wallet.WalletID(args[0])
		bdb := db.NewBadgerDB()
		defer bdb.Close()

		w, err := wallet.GetWallet(walletID, bdb)
		if err != nil {
			return err
		}
		addresses, err := w.Addresses(network)
		fmt.Printf("%-25v %-9v\n", "PATH", "ADDRESS")
		for i, addr := range addresses {
			fmt.Printf("%-25v %-9v\n", "m/1852'/1815'/0'/0/"+strconv.Itoa(i), addr)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listAddressCmd)
	listAddressCmd.Flags().Bool("testnet", false, "Use testnet network")
}

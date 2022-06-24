package locales

var EnUS = map[string]string{
	"Language":                       "en-US",
	"Markets":                        "Markets",
	"Wallets":                        "Wallets",
	"Notifications":                  "Notifications",
	"Recent Activity":                "Recent Activity",
	"Sign Out":                       "Sign Out",
	"Order History":                  "Order History",
	"load from file":                 "load from file",
	"loaded from file":               "loaded from file",
	"defaults":                       "defaults",
	"Wallet Password":                "Wallet Password",
	"w_password_helper":              "This is the password you have configured with your wallet software.",
	"w_password_tooltip":             "Leave the password empty if there is no password required for the wallet.",
	"App Password":                   "App Password",
	"app_password_helper":            "Your app password is always required when performing sensitive wallet operations.",
	"Add":                            "Add",
	"Unlock":                         "Unlock",
	"Rescan":                         "Rescan",
	"Wallet":                         "Wallet",
	"app_password_reminder":          "Your app password is always required when performing sensitive wallet operations.",
	"DEX Address":                    "DEX Address",
	"TLS Certificate":                "TLS Certificate",
	"remove":                         "remove",
	"add a file":                     "add a file",
	"Submit":                         "Submit",
	"Confirm Registration":           "Confirm Registration",
	"app_pw_reg":                     "Enter your app password to confirm DEX registration.",
	"reg_confirm_submit":             `When you submit this form, funds will be spent from your wallet to pay the registration fee.`,
	"provided_markets":               "This DEX provides the following markets:",
	"accepted_fee_assets":            "This DEX accepts the following fees:",
	"base_header":                    "Base",
	"quote_header":                   "Quote",
	"lot_size_header":                "Lot Size",
	"lot_size_headsup":               "All trades are in multiples of the lot size.",
	"Password":                       "Password",
	"Register":                       "Register",
	"Authorize Export":               "Authorize Export",
	"export_app_pw_msg":              "Enter your app password to confirm Account export for",
	"Disable Account":                "Disable Account",
	"disable_app_pw_msg":             "Enter your app password to disable account",
	"disable_irreversible":           `<span class="red">Note:</span> This action is irreversible - once an account is disabled it can't be re-enabled.`,
	"Authorize Import":               "Authorize Import",
	"app_pw_import_msg":              "Enter your app password to confirm Account import",
	"Account File":                   "Account File",
	"Change Application Password":    "Change Application Password",
	"Current Password":               "Current Password",
	"New Password":                   "New Password",
	"Confirm New Password":           "Confirm New Password",
	"Cancel Order":                   "Cancel Order",
	"cancel_pw":                      "Enter your password to submit a cancel order for the remaining",
	"cancel_no_pw":                   "Submit a cancel order for the remaining",
	"cancel_remain":                  "The remaining amount may change before the cancel order is matched.",
	"Log In":                         "Log In",
	"epoch":                          "epoch",
	"price":                          "price",
	"volume":                         "volume",
	"buys":                           "buys",
	"sells":                          "sells",
	"Buy Orders":                     "Buy Orders",
	"Quantity":                       "Quantity",
	"Rate":                           "Rate",
	"Epoch":                          "Epoch",
	"Limit Order":                    "Limit Order",
	"Market Order":                   "Market Order",
	"reg_status_msg":                 `In order to trade at <span id="regStatusDex" class="text-break"></span>, the registration fee payment needs <span id="confReq"></span> confirmations.`,
	"Buy":                            "Buy",
	"Sell":                           "Sell",
	"Lot Size":                       "Lot Size",
	"Rate Step":                      "Rate Step",
	"Max":                            "Max",
	"lot":                            "lot",
	"Price":                          "Price",
	"Lots":                           "Lots",
	"min trade is about":             "min trade is about",
	"immediate_explanation":          "If the order doesn't fully match during the next match cycle, any unmatched quantity will not be booked or matched again. Taker-only order.",
	"Immediate or cancel":            "Immediate or cancel",
	"Balances":                       "Balances",
	"outdated_tooltip":               "Balance may be outdated. Connect to the wallet to refresh.",
	"available":                      "available",
	"connect_refresh_tooltip":        "Click to connect and refresh",
	"add_a_base_wallet":              `Add a<br><span data-unit="base"></span><br>wallet`,
	"add_a_quote_wallet":             `Add a<br><span data-unit="quote"></span><br>wallet`,
	"locked":                         "locked",
	"immature":                       "immature",
	"Sell Orders":                    "Sell Orders",
	"Your Orders":                    "Your Orders",
	"Type":                           "Type",
	"Side":                           "Side",
	"Age":                            "Age",
	"Filled":                         "Filled",
	"Settled":                        "Settled",
	"Status":                         "Status",
	"view order history":             "view order history",
	"cancel order":                   "cancel order",
	"order details":                  "order details",
	"verify_order":                   `Verify <span id="vSideHeader"></span>  Order`,
	"You are submitting an order to": "You are submitting an order to",
	"at a rate of":                   "at a rate of",
	"for a total of":                 "for a total of",
	"verify_market":                  "This is a market order and will match the best available order(s) on the book. Based on the current market mid-gap rate, you might receive about",
	"auth_order_app_pw":              "Authorize this order with your app password.",
	"lots":                           "lots",
	"order_disclaimer": `<span class="red">IMPORTANT</span>: Trades take time to settle, and you cannot turn off the DEX client software,
		or the <span data-unit="quote"></span> or <span data-unit="base"></span> blockchain and/or wallet software, until
		settlement is complete. Settlement can complete as quickly as a few minutes or take as long as a few hours.`,
	"Order":                     "Order",
	"see all orders":            "see all orders",
	"Exchange":                  "Exchange",
	"Market":                    "Market",
	"Offering":                  "Offering",
	"Asking":                    "Asking",
	"Fees":                      "Fees",
	"order_fees_tooltip":        "On-chain transaction fees, typically collected by the miner. Decred DEX collects no trading fees.",
	"Matches":                   "Matches",
	"Match ID":                  "Match ID",
	"Time":                      "Time",
	"ago":                       "ago",
	"Cancellation":              "Cancellation",
	"Order Portion":             "Order Portion",
	"you":                       "you",
	"them":                      "them",
	"Redemption":                "Redemption",
	"Refund":                    "Refund",
	"Funding Coins":             "Funding Coins",
	"Exchanges":                 "Exchanges",
	"apply":                     "apply",
	"Assets":                    "Assets",
	"Trade":                     "Trade",
	"Set App Password":          "Set App Password",
	"reg_set_app_pw_msg":        "Set your app password. This password will protect your DEX account keys and connected wallets.",
	"Password Again":            "Password Again",
	"Add a DEX":                 "Add a DEX",
	"Pick a server":             "Pick a server",
	"reg_ssl_needed":            "Looks like we don't have an SSL certificate for this DEX. Add the server's certificate to continue.",
	"Dark Mode":                 "Dark Mode",
	"Show pop-up notifications": "Show pop-up notifications",
	"Account ID":                "Account ID",
	"Export Account":            "Export Account",
	"simultaneous_servers_msg":  "The Decred DEX Client supports simultaneous use of any number of DEX servers.",
	"Change App Password":       "Change App Password",
	"Build ID":                  "Build ID",
	"Connect":                   "Connect",
	"Send":                      "Send",
	"Deposit":                   "Deposit",
	"Lock":                      "Lock",
	"New Deposit Address":       "New Deposit Address",
	"Address":                   "Address",
	"Amount":                    "Amount",
	"Authorize the transfer with your app password.": "Authorize the transfer with your app password.",
	"Reconfigure":                 "Reconfigure",
	"pw_change_instructions":      "Changing the password below does not change the password for your wallet software. Use this form to update the DEX client after you have changed your password with the wallet software directly.",
	"New Wallet Password":         "New Wallet Password",
	"pw_change_warn":              "Note: Changing to a different wallet while having active trades might cause funds to be lost.",
	"Show more options":           "Show more options",
	"seed_implore_msg":            "You should carefully write down your application seed and save a copy. Should you lose access to this machine or the critical application files, the seed can be used to restore your DEX accounts and native wallets. Some older accounts cannot be restored from seed, and whether old or new, it's good practice to backup the account keys separately from the seed.",
	"View Application Seed":       "View Application Seed",
	"Remember my password":        "Remember my password",
	"pw_for_seed":                 "Enter your app password to show your seed. Make sure nobody else can see your screen.",
	"Asset":                       "Asset",
	"Balance":                     "Balance",
	"Actions":                     "Actions",
	"Restoration Seed":            "Restoration Seed",
	"Restore from seed":           "Restore from seed",
	"Import Account":              "Import Account",
	"no_wallet":                   "no wallet",
	"create_a_x_wallet":           "Create a {{.Info.Name}} Wallet", // .Info.Name will be an asset symbol like DCR
	"dont_share":                  "Don't share it. Don't lose it.",
	"Show Me":                     "Show Me",
	"Wallet Settings":             "Wallet Settings",
	"add_a_x_wallet":              `Add a <img data-tmpl="assetLogo" class="asset-logo mx-1"> <span data-tmpl="assetName"></span> Wallet`,
	"ready":                       "ready",
	"off":                         "off",
	"Export Trades":               "Export Trades",
	"change the wallet type":      "change the wallet type",
	"confirmations":               "confirmations",
	"how_reg":                     "How will you pay the registration fee?",
	"All markets at":              "All markets at",
	"pick a different asset":      "pick a different asset",
	"Create":                      "Create",
	"Register_loudly":             "Register!",
	"1 Sync the Blockchain":       "1: Sync the Blockchain",
	"Progress":                    "Progress",
	"remaining":                   "remaining",
	"2 Fund the Registration Fee": "2: Fund the Registration Fee",
	"One time anti-spam measure":  "This is a small one-time anti-spam measure to discourage disruptive behavior like backing out on swaps.",
	"Registration fee":            "Registration fee",
	"Your Deposit Address":        "Your Wallet's Deposit Address",
	"Send enough for reg fee":     `Make sure you send enough to also cover network fees. You may deposit as much as you like to your wallet, since only the fee amount will be used in the next step. The deposit must confirm to proceed.`,
	"Send enough with estimate":   `Deposit at least <span data-tmpl="totalFees"></span> <span class="unit">XYZ</span> to also cover network fees. You may deposit as much as you like to your wallet, since only the required amount will be used in the next step. The deposit must confirm to proceed.`,
	"add a different server":      "add a different server",
	"Add a custom server":         "Add a custom server",
	"plus tx fees":                "+ tx fees",
	"Export Seed":                 "Export Seed",
	"Total":                       "Total",
	"Trading":                     "Trading",
	"Receiving Approximately":     "Receiving Approximately",
	"Fee Projection":              "Fee Projection",
	"details":                     "details",
	"to":                          "to",
	"Options":                     "Options",
	"fee_projection_tooltip":      "If network conditions don't change before your order matches, total realized fees (as a percentage of trade) should fall within this range.",
	"unlock_for_details":          "Unlock your wallets to retrieve order details and additional options.",
	"estimate_unavailable":        "Order estimates and options unavailable",
	"Fee Details":                 "Fee Details",
	"estimate_market_conditions":  "Best- and worst-case estimates are based on current network conditions, and might change by the time the order matches.",
	"Best Case Fees":              "Best Case Fees",
	"best_case_conditions":        "The best case fees occur when the entire order is consumed in a single match.",
	"Swap":                        "Swap",
	"Redeem":                      "Redeem",
	"Worst Case Fees":             "Worst Case Fees",
	"worst_case_conditions":       "The worst case can occur if the order matches one lot at a time over the course of many epochs.",
	"Maximum Possible Swap Fees":  "Maximum Possible Swap Fees",
	"max_fee_conditions":          "This is the most you would ever pay in fees on your swap. Fees are normally assessed at a fraction of this rate. The maximum is not subject to changes once your order is placed.",
	"wallet_logs":                 "Wallet Logs",
	"accelerate_order":            "Accelerate Order",
	"acceleration_text":           "If your swap transactions are stuck, you may attempt to accelerate them with an additional transaction. This is helpful when the the fee rate on an existing unconfirmed transaction has become too low to be mined in the next block, but not if blocks are just being mined slowly. When you submit this form, a transaction will be created that sends the change from your own swap initiation transaction to yourself with a higher fee. The effective fee rate of your swap transactions will become the rate you select below. Select a rate that is enough to be included in the next block. Consult a block explorer to be sure.",
	"effective_swap_tx_rate":      "Effective swap tx fee rate",
	"current_fee":                 "Current suggested fee rate",
	"accelerate_success":          `Successfully submitted transaction: <span id="accelerateTxID"></span>`,
	"accelerate":                  "Accelerate",
	"acceleration_transactions":   "Acceleration Transactions",
	"acceleration_cost_msg":       `Increasing the effective fee rate to <span id="feeRateEstimate"></span> will cost <span id="feeEstimate"></span>`,
	"recent_acceleration_msg":     `Your latest acceleration was only <span id="recentAccelerationTime"></span> minutes ago! Are you sure you want to accelerate?`,
	"recent_swap_msg":             `Your oldest unconfirmed swap transaction was submitted only <span id="recentSwapTime"></span> minutes ago! Are you sure you want to accelerate?`,
	"early_acceleration_help_msg": `It will cause no harm to your order, but you might be wasting money. Acceleration is only helpful if the fee rate on an existing unconfirmed transaction has become too low to be mined in the next block, but not if blocks are just being mined slowly. You can confirm this in the block explorer by closing this popup and clicking on your previous transactions.`,
	"recover":                     "Recover",
	"recover_wallet":              "Recover Wallet",
	"recover_warning":             "Recovering your wallet will move all wallet data to a backup folder. You will have to wait until the wallet resyncs with the network, which could potentially take a long time, before you can use the wallet again.",
	"wallet_actively_used":        "Wallet actively used!",
	"confirm_force_message":       "This wallet is actively managing orders. After taking this action, it will take a long time to resync your wallet, potentially causing orders to fail. Only take this action if absolutely necessary!",
	"confirm":                     "Confirm",
	"cancel":                      "Cancel",
	"Update TLS Certificate":      "Update TLS Certificate",
	"registered dexes":            "Registered Dexes:",
	"successful_cert_update":      "Successfully updated certificate!",
	"update dex host":             "Update DEX Host",
	"copied":                      "Copied!",
}

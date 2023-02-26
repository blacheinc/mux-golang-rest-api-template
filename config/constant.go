package config

import (
	"os"

	"github.com/opensaucerer/gotemplate/typing"
)

const (

	// EnvTagName is the tag name for environment variables struct
	envTagName = "env"

	// ShutdownTimeout is the time to wait for the server to shutdown gracefully
	ShutdownTimeout = 5 // seconds

	// TwitterScope is the scope for twitter oauth
	TwitterScope = "like.read%20like.write%20tweet.read%20tweet.write%20users.read%20offline.access%20follows.write%20follows.read"

	// UserCollection is the name of the user collection
	UserCollection = "users"

	// RaidCollection is the name of the raid collection
	RaidCollection = "raids"

	// PricingCollection is the name of the pricing collection
	PricingCollection = "pricing"

	// DefaultAvatar is the default avatar for users
	DefaultAvatar = "https://e7.pngegg.com/pngimages/84/165/png-clipart-united-states-avatar-organization-information-user-avatar-service-computer-wallpaper-thumbnail.png"

	// PaymentMethodSolana is the payment method for solana
	PaymentMethodSolana = "SOLANA"

	// ChannelTwitter is the channel for twitter
	ChannelTwitter = "TWITTER"

	// DefaultPageLimit is the default page limit
	DefaultPageLimit = 10

	//DefaultPageOffset is the default page to skip
	DefaultPageOffset = 0

	// // ActionReply is for comment
	// ActionReply = "REPLY"

	// // ActionLike is for like
	// ActionLike = "LIKE"

	// // ActionRetweet
	// ActionRetweet = "RETWEET"

	// ErrorTransactionFailed is the error message for transaction failure
	ErrTransactionFailed = "transaction failed"

	// ErrTransactionInvalidDestination is the error message for invalid destination
	ErrTransactionInvalidDestination = "transaction made to the wrong destination address"

	// ErrTransactionInvalidAmount is the error message for invalid amount
	ErrTransactionInvalidAmount = "transaction made with an invalid amount"

	// ErrTwitterTooManyRequests is the error message for too many requests
	ErrTwitterTooManyRequests = "Too Many Requests"

	// ErrTwitterNotFound is the error message for not found
	ErrTwitterNotFound = "Not Found Error"

	// CronQueryLimit is the limit for cron queries
	CronQueryLimit = 100

	// RaidActionWaitThreshold is the wait time threshold for raid actions
	RaidActionWaitThreshold = 24 // hours

	// RewardMultiplier is the multiplier for rewards to raiders
	RewardMultiplier = 0.8

	// Lamports is the lamports for solana
	Lamports = 1000000000
)

var (
	// Env is the global environment variable
	Env = new(typing.Env) // global environment variable

	// ShutdownChan is the channel to listen for shutdown signals
	ShutdownChan = make(chan os.Signal, 1)

	// Currency is the currency table
	Currency = map[string]string{
		"USD":    "$",
		"EUR":    "€",
		"GBP":    "£",
		"SOLANA": "SOL",
	}
)

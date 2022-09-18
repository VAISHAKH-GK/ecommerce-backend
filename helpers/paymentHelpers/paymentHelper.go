package paymentHelpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"os"
)

func VerifyOnlinePayment(payment map[string]interface{}) bool {
	var signature = payment["razorpay_signature"].(string)
	var orderId = payment["razorpay_order_id"].(string)
	var paymentId = payment["razorpay_payment_id"].(string)

	var secret = os.Getenv("RP_SECRET")
	var data = orderId + "|" + paymentId

	var h = hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(data))
	helpers.CheckNilErr(err)

	var sha = hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) == 1 {
		return true
	} else {
		return false
	}
}

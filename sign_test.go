package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/qiafan666/quickweb/commons/utils"
	"testing"
)

func TestSign(t *testing.T) {
	apiKey := "1"
	Url := "1"
	Timestamp := "1"
	Parameter := []byte("{\n\t\t\"parameter\":\"1\"\n\t}")
	Rand := "1"
	data := fmt.Sprintf("%s%s%s%s%s", apiKey, Url, Timestamp, Parameter, Rand)
	thirdPartyPrivateKey := "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDAu7lBcj9mIDKk\nrRHdcXfbnSo6e/C++QMhaEKa9J2kyTCKVQ4+2a/WMXjLA3j9Be20nAgHP8WQuiqK\nJ3T90IAoK9X12hTRNOIz8ORb19W27V1fX3KCmol/TNRvF2MtjpWxvBMMnoRvvpQ1\nmmpaQyGDNHuK7G8k2xFFs7ywqZdrSxpxwFzBvNmR8ne6ng1ppkni7KuIbjGxYlcr\n20a+jh2YTBway6k3wfPtf3KKz8Gmf++nTxT4f1cqkOut3L1h+DihIqG6ySOhhP15\niV//RBTwdQYwSeXrFZQEWvwBFOxMBlIehJsGELCA8g2HK4fO3H2VTdfGGwtnJDFX\n3GQBMOEtAgMBAAECggEAHfMN7qWaRHxsYjqitA6V2YKqtTvdRU/ctKxG7V1lwd2h\ntV1SQWICeP5nDuUTP/5T2eUFOlsmkD7drWpEO8zSnWtybCnfYkMdg4TDd6Iqi6qG\n//MNEE6DX0zJFhpERygJCv122FcprVOoJExCipQz+PeG2yeyfb+tB2/OuoUgCoP0\nHek8lLDtNhNfWui24tMd2P1+CMjj4zzsDYrzBdYAPbRbNANKWBvQwFKI3EQkbhPT\nBREzGKj4JHOYH6Hz0jDXrGFJPTAihgID+PeEBshV2y9QGlkJTW0C3gmpIDz1OaKF\nfcZCde2PzmAD65k7aMFNDZmLU7bQ+OthJ2CI4vVO4QKBgQDlb6QyjZZwNUXyRPVO\n57uTqTIaVSF8dgIINUt8TK89Lw98+rGCsfYDBbpsblvlK9W9Hr1IrfJjQHf2DQgD\njIJQuMFluVAa6CglALytU3bH70sXJ0k58Z4y2V/+8rkCTAy9OlGSWBiIdxzmS950\nna/C8hdb5ZYZNynu7dHcMNMplwKBgQDXDDtFznTK66QjxLCKvGLz3CIoeIFHU+PY\ndBKxSZ/7EG+tmGXOJ4bxNBrdHe+FJWRjYNjIOPPlTIreODT3RCY5OMutD+rdIdIJ\ndxpl5M39FFblvio1KcXNFPJcu1peeNeeInIsXrxJHvPdItam81GV3Soxif4XZcNO\nnK0DM4272wKBgFVvYSFCAAcAj29Lpl0fhYXSt0l+8d06xD7yOY2rsIWEBKxxXbBh\nPE6bz3OZFLcdv5WQ4MMzotK6qvEAoT9RDyWn5rxOaTnbwTcmMxwHvG9u9/NDOc1N\n367nqwtwrtvgHc5I7R8llt0aHbTUA55BKbXaGECsGVyCYicKf98Sf085AoGBAKf9\niiASE/Kg+exnLnJyj+poQNbUrEkII6lno2KTXUJHqLY3ou/UuPmb9pBdXkro1u87\nLJ3cv8qUbLcDuXyf5Cw3TgS3toVgci+qtxh6EOBvDyMR1u8I3thCUMJYKVQ7mlSS\nHBbFOtj0MRTCrmRlF4q25sskPTYR7OxwQEeL8mCZAoGAEqv6E0IGV05yqWbFZzkg\nNzj9m9gQUdgd6kLAqx4kwMZZnegHq/DV0HKsgVapPWkCTrDuyNKDq1F82+cRqzdQ\nT2Y1BG04+GgOtK4aflmcKzWZpTPlqDZ9InSLu9y0L9mLHVpB9PO+FWU1sy2DZLmC\nUqMm/hskXy7BYdhABcb3ZWM=\n-----END PRIVATE KEY-----\n"
	bufferString := bytes.NewBufferString(data)
	privateKeyBufferString := bytes.NewBufferString(thirdPartyPrivateKey)

	sign, err := utils.Rsa2Sign(bufferString.Bytes(), privateKeyBufferString.Bytes(), utils.PKCS_8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hex.EncodeToString(sign))
	//4a4a8a564975c8c603f25a9a53652a82389c417a207e37318ef393f858d58b668c9ba85c5b2fb9530b66ac39488c2b42b509b80904a57b1797e95d724feab15445256f8c5513156ec69f6e335b316b8ac6e98944dc5b9f46fa4c2bf77c9eff0fd0d5896d6b3d15b7599028f4062adf3feca3c3ccbadfa67f43ff329764b9344a44f089fea5c4950f8bc43262db630f1c55dc35e5eda2cb3ef3447db1c88fe10c87a84c57b4fb48981c03f1890fe34193659015dea87b5d410aa184df133ad3a1bb799f4727832daf47ff691f4ce01fc718114c012e52faefd19ae0746a06e77d312d7c6f630ab0e0bc09175657e0b36506ba9701546e2f2d4aa65b319d4b3bd2
	/*
		    JJ�VIu���Z�Se*�8�Az ~71����XՋf���\[/�Sf�9H�+B�	�	�{��]rO��TE%o�UnƟn3[1k����D�[�F�L+�|���Չmk=�Y�(�*�?���̺ߦC�2�d�4JD����ĕ��2b�cU�5����>�D}�ȏ���LW��H����A�e�ި{]A
			���:ӡ�y�G'�-�G�iL��L.R��њ�tj�}1-|oc
			���	VW��e��Tn/-J�[1�K;�
	*/
}

func TestVerifySign(t *testing.T) {
	apiKey := "1"
	Url := "1"
	Timestamp := "1"
	Parameter := []byte("{\n\t\t\"parameter\":\"1\"\n\t}")
	Rand := "1"
	decode, err := hex.DecodeString("441c0ca529aefc96fa2e41f89d0182315d60078485eab52c73b9479e74cc8d5f03a87623296a12f48a7407719be565d3418702b695daa573dbe606baf22e4b9c7d65044a07feb9bac06504501471e8e03c7b2d4876b7c8ae08457823fcd8ca46ac3120712fc8c83c369c0fc8a20c8d7111c01156780007a14924ae6c23f23f81fa2e35d27b7a7018635245fcad4323defb880f7ef7d5ada3f6abc481540be721678b2f38966223ccaabc35621a060a5c3e8322df24c23645a52cd4da98d885c8a22dbc2f103d4ad82b3d487b91604d6bfdb8e4391355f0092c29ec6d07a0b8e39ca3382d662dba59aea015e6a4e86acb60cbfa8f279d4dc3bd43a24d6f099fae")
	if err != nil {
		return
	}
	//sign:="JJ�VIu��\u0003�Z�Se*�8�Az ~71����XՋf���\\[/�S\vf�9H�+B�\t�\t\u0004�{\u0017��]rO��TE%o�U\u0013\u0015nƟn3[1k����D�[�F�L+�|��\u000F�Չmk=\u0015�Y�(�\u0006*�?���̺ߦ\u007FC�2�d�4JD����ĕ\u000F��2b�c\u000F\u001CU�5����>�D}�ȏ�\f��LW��H�\u001C\u0003��\u000F�A�e�\u0015ި{]A\n���\u0013:ӡ�y�G'�-�G�i\u001FL�\u001F�\u0018\u0011L\u0001.R��њ�tj\u0006�}1-|oc\n���\t\u0017VW��e\u0006��\u0001Tn/-J�[1�K;�"
	data := fmt.Sprintf("%s%s%s%s%s", apiKey, Url, Timestamp, Parameter, Rand)
	thirdPartyPublicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwLu5QXI/ZiAypK0R3XF3\n250qOnvwvvkDIWhCmvSdpMkwilUOPtmv1jF4ywN4/QXttJwIBz/FkLoqiid0/dCA\nKCvV9doU0TTiM/DkW9fVtu1dX19ygpqJf0zUbxdjLY6VsbwTDJ6Eb76UNZpqWkMh\ngzR7iuxvJNsRRbO8sKmXa0saccBcwbzZkfJ3up4NaaZJ4uyriG4xsWJXK9tGvo4d\nmEwcGsupN8Hz7X9yis/Bpn/vp08U+H9XKpDrrdy9Yfg4oSKhuskjoYT9eYlf/0QU\n8HUGMEnl6xWUBFr8ARTsTAZSHoSbBhCwgPINhyuHztx9lU3XxhsLZyQxV9xkATDh\nLQIDAQAB\n-----END PUBLIC KEY-----\n"
	bufferString := bytes.NewBufferString(data)
	publicKeyBufferString := bytes.NewBufferString(thirdPartyPublicKey)

	block, _ := pem.Decode(publicKeyBufferString.Bytes())
	if block == nil {
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	pub, _ := pubInterface.(*rsa.PublicKey)

	fmt.Println(pub.Size())

	var test response.ThirdPartyLogin
	test.Token = utils.GenerateUUID()
	test.WalletAddress = utils.GenerateUUID()
	test.Uuid = utils.GenerateUUID()
	test.Email = utils.GenerateUUID()
	marshal, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(marshal))
	err = utils.Rsa2VerifySign(sha256.Sum256(bufferString.Bytes()), decode, publicKeyBufferString.Bytes())
	if err != nil {
		fmt.Println(err)
	}

}

/*
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwLu5QXI/ZiAypK0R3XF3
250qOnvwvvkDIWhCmvSdpMkwilUOPtmv1jF4ywN4/QXttJwIBz/FkLoqiid0/dCA
KCvV9doU0TTiM/DkW9fVtu1dX19ygpqJf0zUbxdjLY6VsbwTDJ6Eb76UNZpqWkMh
gzR7iuxvJNsRRbO8sKmXa0saccBcwbzZkfJ3up4NaaZJ4uyriG4xsWJXK9tGvo4d
mEwcGsupN8Hz7X9yis/Bpn/vp08U+H9XKpDrrdy9Yfg4oSKhuskjoYT9eYlf/0QU
8HUGMEnl6xWUBFr8ARTsTAZSHoSbBhCwgPINhyuHztx9lU3XxhsLZyQxV9xkATDh
LQIDAQAB
-----END PUBLIC KEY-----

-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDAu7lBcj9mIDKk
rRHdcXfbnSo6e/C++QMhaEKa9J2kyTCKVQ4+2a/WMXjLA3j9Be20nAgHP8WQuiqK
J3T90IAoK9X12hTRNOIz8ORb19W27V1fX3KCmol/TNRvF2MtjpWxvBMMnoRvvpQ1
mmpaQyGDNHuK7G8k2xFFs7ywqZdrSxpxwFzBvNmR8ne6ng1ppkni7KuIbjGxYlcr
20a+jh2YTBway6k3wfPtf3KKz8Gmf++nTxT4f1cqkOut3L1h+DihIqG6ySOhhP15
iV//RBTwdQYwSeXrFZQEWvwBFOxMBlIehJsGELCA8g2HK4fO3H2VTdfGGwtnJDFX
3GQBMOEtAgMBAAECggEAHfMN7qWaRHxsYjqitA6V2YKqtTvdRU/ctKxG7V1lwd2h
tV1SQWICeP5nDuUTP/5T2eUFOlsmkD7drWpEO8zSnWtybCnfYkMdg4TDd6Iqi6qG
//MNEE6DX0zJFhpERygJCv122FcprVOoJExCipQz+PeG2yeyfb+tB2/OuoUgCoP0
Hek8lLDtNhNfWui24tMd2P1+CMjj4zzsDYrzBdYAPbRbNANKWBvQwFKI3EQkbhPT
BREzGKj4JHOYH6Hz0jDXrGFJPTAihgID+PeEBshV2y9QGlkJTW0C3gmpIDz1OaKF
fcZCde2PzmAD65k7aMFNDZmLU7bQ+OthJ2CI4vVO4QKBgQDlb6QyjZZwNUXyRPVO
57uTqTIaVSF8dgIINUt8TK89Lw98+rGCsfYDBbpsblvlK9W9Hr1IrfJjQHf2DQgD
jIJQuMFluVAa6CglALytU3bH70sXJ0k58Z4y2V/+8rkCTAy9OlGSWBiIdxzmS950
na/C8hdb5ZYZNynu7dHcMNMplwKBgQDXDDtFznTK66QjxLCKvGLz3CIoeIFHU+PY
dBKxSZ/7EG+tmGXOJ4bxNBrdHe+FJWRjYNjIOPPlTIreODT3RCY5OMutD+rdIdIJ
dxpl5M39FFblvio1KcXNFPJcu1peeNeeInIsXrxJHvPdItam81GV3Soxif4XZcNO
nK0DM4272wKBgFVvYSFCAAcAj29Lpl0fhYXSt0l+8d06xD7yOY2rsIWEBKxxXbBh
PE6bz3OZFLcdv5WQ4MMzotK6qvEAoT9RDyWn5rxOaTnbwTcmMxwHvG9u9/NDOc1N
367nqwtwrtvgHc5I7R8llt0aHbTUA55BKbXaGECsGVyCYicKf98Sf085AoGBAKf9
iiASE/Kg+exnLnJyj+poQNbUrEkII6lno2KTXUJHqLY3ou/UuPmb9pBdXkro1u87
LJ3cv8qUbLcDuXyf5Cw3TgS3toVgci+qtxh6EOBvDyMR1u8I3thCUMJYKVQ7mlSS
HBbFOtj0MRTCrmRlF4q25sskPTYR7OxwQEeL8mCZAoGAEqv6E0IGV05yqWbFZzkg
Nzj9m9gQUdgd6kLAqx4kwMZZnegHq/DV0HKsgVapPWkCTrDuyNKDq1F82+cRqzdQ
T2Y1BG04+GgOtK4aflmcKzWZpTPlqDZ9InSLu9y0L9mLHVpB9PO+FWU1sy2DZLmC
UqMm/hskXy7BYdhABcb3ZWM=
-----END PRIVATE KEY-----

*/

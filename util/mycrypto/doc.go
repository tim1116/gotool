/*

	Package mycrypto 实现兼容平台的加密解密.


	加密：

		src := "hello"
		key := "5qHIk7yMbGu29SFA61234567"
		iv := "Xoji5qa9"

		data := TripleDesEncrypt(src, key, iv)

	解密：

		src := "iQv3u6hIj9c="
		key := "5qHIk7yMbGu29SFA61234567"
		iv := "Xoji5qa9"

		data, err := TripleDesDecrypt(src, key, iv)
*/
package mycrypto

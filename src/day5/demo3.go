package main

func main() {
	
}
func authcode(strings,operation,key string, expiry int64) string{		
	key =md5s(key)	
	keya :=md5s(substr(key,0,16))
	keyb := md5s(substr(key, 16, 16));
	var keyc string	 
	if operation =="DECODE"{
		keyc =substr(strings,4,len(strings))
	}else{
		keyc =substr(md5s(strconv.FormatInt(time.Now().UnixNano() / 1000000,10)),-4,0)
	}	
	cryptkey := keya + md5s(keya + keyc)
	key_length := len(cryptkey)	
	if operation =="DECODE"{
		strings,_:=base64.StdEncoding.DecodeString(substr(strings,4,len(strings)))
	}else{
		if expiry !=0 {
			expiry =expiry+time.Now().Unix()
		}
		expirys :=fmt.Sprintf("%10d",expiry)
		expirys =expirys+substr(md5s(strings+keyb),0,16)+strings
	}
	string_length := len(strings);
    var result string;
    box := make(map[int]int)
	rndkey :=make(map[int]string)
	for j:=0;j<=255;j++ {
		box[j]=j
	}
	for i:=0; i<=255; i++ {		
		rndkey[i]=string([]byte(cryptkey)[i%key_length])
	}
	for j,i:=0; i < 256;i++{
		num,_:=strconv.Atoi(rndkey[i])
        j = (j + box[i] +num ) % 256
        tmp := box[i]
        box[i] = box[j]
        box[j] = tmp
    }
	return key
}
//md5加密
func md5s(key string) (keys string){
	h := md5.New()
	h.Write([]byte(key)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	keys = hex.EncodeToString(cipherStr)
	return keys
}
//截取字符串
func substr(str string,start,end int)string{
	rs := []rune(str)
	return string(rs[start:end])
}
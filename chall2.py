from chall1 import conv_hex_base64

def Xor_data(string1, string2):
    # decodes string to hex then XOR's them against each other    
    
    data1 = string1.decode("hex")
    data2 = string2.decode("hex")
    
    return "".join(chr(ord(x) ^ ord(y)) for x, y in zip(data1, data2).encode("hex"))
    return result

if __name__ == "__main__":
    string1= "1c0111001f010100061a024b53535009181c"
    string2= "686974207468652062756c6c277320657965"

    Xor_data(string1, string2)